package sessions

import (
	f "github.com/razshare/frizzante"
	"main/lib"
)

var memoryOperating = map[string]chan int{}
var memoryStores = map[string]lib.State{}

func Memory(session *f.Session[lib.State]) {
	f.SessionWithLoader(session, func() {
		state, sessionExists := memoryStores[session.Id]
		if !sessionExists {
			state = lib.InitializeState()
			memoryStores[session.Id] = state
			memoryOperating[session.Id] = make(chan int, 1)
			memoryOperating[session.Id] <- 0
		}

		<-memoryOperating[session.Id]
		session.Store = memoryStores[session.Id]
		memoryOperating[session.Id] <- 0
	})

	f.SessionWithValidator(session, func() bool {
		return true
	})

	f.SessionWithSaver(session, func() {
		// Noop.
	})

	f.SessionWithDestroyer(session, func() {
		<-memoryOperating[session.Id]
		delete(memoryStores, session.Id)
		delete(memoryOperating, session.Id)
		memoryOperating[session.Id] <- 0
	})
}
