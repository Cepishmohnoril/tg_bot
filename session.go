package main

import (
	"time"
)

type session struct {
	data      []int
	nextStep  string
	waiting   bool
	createdAt time.Time
}

func NewSession(sesId int64) *session {
	storage[sesId] = &session{
		createdAt: time.Now(),
		waiting:   false,
	}
	return storage[sesId]
}

var storage map[int64]*session

func initSessionStorage() {
	storage = make(map[int64]*session)

	go cleanerLoop()
}

func cleanerLoop() {
	for {
		time.Sleep(10 * time.Minute)
		cleanStorage()
	}
}

func cleanStorage() {

	for id, session := range storage {
		diff := time.Since(session.createdAt)

		if diff.Minutes() >= 10 {
			delete(storage, id)
		}
	}
}

func getSession(sesId int64) *session {
	return storage[sesId]
}

func terminateSession(sesId int64) {
	delete(storage, sesId)
}

func sessionExists(sesId int64) bool {
	_, ok := storage[sesId]

	return ok
}

func (s *session) setNextStep(nextStep string) {
	s.nextStep = nextStep
}

func (s *session) addData(record int) {
	s.data = append(s.data, record)
}
