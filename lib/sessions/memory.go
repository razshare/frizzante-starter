package sessions

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"time"
)

var memoryOperating = map[string]chan int{}
var memoryStores = map[string]lib.State{}

// Memory builds sessions in memory.
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
		return time.Since(session.State.LastActivity) < 30*time.Minute
	})

	f.SessionWithSaveHandler(session, func() {
		<-memoryOperating[session.Id]
		memoryStores[session.Id] = session.State
		memoryOperating[session.Id] <- 0
	})

	f.SessionWithDestroyHandler(session, func() {
		<-memoryOperating[session.Id]
		delete(memoryStores, session.Id)
		delete(memoryOperating, session.Id)
		memoryOperating[session.Id] <- 0
	})
}
