package sessions

import (
	f "github.com/razshare/frizzante"
	"main/lib"
)

var memoryOperating = map[string]chan int{}
var memoryStores = map[string]lib.State{}

func Memory(session *f.Session[lib.State]) {
	f.SessionWithLoadHandler(session, func() {
		state, sessionExists := memoryStores[session.Id]
		if !sessionExists {
			state = lib.InitializeState()
			memoryStores[session.Id] = state
			memoryOperating[session.Id] = make(chan int, 1)
			memoryOperating[session.Id] <- 0
		}

		<-memoryOperating[session.Id]
		session.State = memoryStores[session.Id]
		memoryOperating[session.Id] <- 0
	})

	f.SessionWithValidateHandler(session, func() bool {
		return true
	})

	f.SessionWithSaveHandler(session, func() {
		// Noop.
	})

	f.SessionWithDestroyHandler(session, func() {
		<-memoryOperating[session.Id]
		delete(memoryStores, session.Id)
		delete(memoryOperating, session.Id)
		memoryOperating[session.Id] <- 0
	})
}
