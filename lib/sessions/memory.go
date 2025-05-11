package sessions

import (
	f "github.com/razshare/frizzante"
	"main/lib"
)

var memorySessions = map[string]*lib.UserSession{}
var memoryOperating = map[string]chan int{}

// Memory builds sessions in memory.
func Memory(session *f.Session[lib.UserSession]) {
	f.SessionWithLoader(session, func() {
		_, exists := memorySessions[session.Id]
		if !exists {
			memorySessions[session.Id] = lib.InitialState()
			memoryOperating[session.Id] = make(chan int, 1)
			memoryOperating[session.Id] <- 0
		}

		<-memoryOperating[session.Id]
		session.Value = memorySessions[session.Id]
		memoryOperating[session.Id] <- 0
	})

	f.SessionWithValidator(session, func() bool {
		return true
	})

	f.SessionWithSaver(session, func() {
		<-memoryOperating[session.Id]
		memorySessions[session.Id] = session.Value
		memoryOperating[session.Id] <- 0
	})

	f.SessionWithDestroyer(session, func() {
		<-memoryOperating[session.Id]
		delete(memorySessions, session.Id)
		memoryOperating[session.Id] <- 0
	})
}
