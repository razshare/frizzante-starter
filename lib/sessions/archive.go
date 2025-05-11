package sessions

import (
	"encoding/json"
	f "github.com/razshare/frizzante"
	"main/lib"
)

var archiveKey = "session.json"
var archiveNotifier = f.NotifierCreate()
var archiveStores = map[string]lib.State{}
var archiveOperating = map[string]chan int{}
var archive = f.ArchiveCreateOnDisk(".sessions")

func Archive(session *f.Session[lib.State]) {
	f.SessionWithLoader(session, func() {
		state, sessionExists := archiveStores[session.Id]
		if !sessionExists {
			state = lib.InitializeState()
			archiveStores[session.Id] = state
			archiveOperating[session.Id] = make(chan int, 1)
			archiveOperating[session.Id] <- 0

			<-archiveOperating[session.Id]
			if !f.ArchiveHas(archive, session.Id, archiveKey) {
				readBytes, marshalError := json.Marshal(state)
				if nil != marshalError {
					f.NotifierSendError(archiveNotifier, marshalError)
					archiveOperating[session.Id] <- 0
					return
				}
				f.ArchiveSet(archive, session.Id, archiveKey, readBytes)
			}
			archiveOperating[session.Id] <- 0
		}

		<-archiveOperating[session.Id]
		session.Store = archiveStores[session.Id]
		readBytes := f.ArchiveGet(archive, session.Id, archiveKey)
		placeholder := archiveStores[session.Id]
		marshalError := json.Unmarshal(readBytes, &placeholder)
		archiveStores[session.Id] = placeholder
		if nil != marshalError {
			f.NotifierSendError(archiveNotifier, marshalError)
		}
		archiveOperating[session.Id] <- 0
	})

	f.SessionWithValidator(session, func() bool {
		return true
	})

	f.SessionWithSaver(session, func() {
		<-archiveOperating[session.Id]
		readBytes, marshalError := json.Marshal(archiveStores[session.Id])
		if nil != marshalError {
			f.NotifierSendError(archiveNotifier, marshalError)
			archiveOperating[session.Id] <- 0
			return
		}

		f.ArchiveSet(archive, session.Id, archiveKey, readBytes)
		archiveOperating[session.Id] <- 0
	})

	f.SessionWithDestroyer(session, func() {
		<-archiveOperating[session.Id]
		delete(archiveOperating, session.Id)
		delete(archiveStores, session.Id)
		archiveOperating[session.Id] <- 0
	})
}
