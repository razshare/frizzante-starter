package sessions

import (
	"encoding/json"
	f "github.com/razshare/frizzante"
	"main/lib"
)

var archive = f.ArchiveCreateLocal(archiveNotifier, ".sessions")
var archiveNotifier = f.NotifierCreate()
var archiveOperating = map[string]chan int{}
var archiveSessions = map[string]*lib.UserSession{}

// Archive builds session in an archive that is backed by the file system.
//
// Sessions will be persisted in ".session/{sessionId}/session.json",
// where "sessionId" is the session id of the user.
func Archive(session *f.Session[lib.UserSession]) {
	f.SessionWithLoader(session, func() {
		_, exists := archiveSessions[session.Id]
		if !exists {
			archiveSessions[session.Id] = lib.InitialState()
			archiveOperating[session.Id] = make(chan int, 1)
			archiveOperating[session.Id] <- 0

			<-archiveOperating[session.Id]
			if !f.ArchiveHas(archive, session.Id, "session.json") {
				readBytes, marshalError := json.Marshal(archiveSessions[session.Id])
				if nil != marshalError {
					f.NotifierSendError(archiveNotifier, marshalError)
					archiveOperating[session.Id] <- 0
					return
				}
				f.ArchiveSet(archive, session.Id, "session.json", readBytes)
			}
			archiveOperating[session.Id] <- 0
		}

		<-archiveOperating[session.Id]
		session.Value = archiveSessions[session.Id]
		readBytes := f.ArchiveGet(archive, session.Id, "session.json")
		marshalError := json.Unmarshal(readBytes, archiveSessions[session.Id])
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
		readBytes, marshalError := json.Marshal(archiveSessions[session.Id])
		if nil != marshalError {
			f.NotifierSendError(archiveNotifier, marshalError)
			archiveOperating[session.Id] <- 0
			return
		}
		f.ArchiveSet(archive, session.Id, "session.json", readBytes)
		archiveOperating[session.Id] <- 0
	})

	f.SessionWithDestroyer(session, func() {
		<-archiveOperating[session.Id]
		delete(archiveOperating, session.Id)
		archiveOperating[session.Id] <- 0
	})
}
