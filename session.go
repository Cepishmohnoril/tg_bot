package main

type session struct {
	nextStep string
	data     []int
}

func NewSession() session {
	return session{}
}

var storage map[int64]session

func initSessionStorage() {
	storage = make(map[int64]session)
}

func setSession(sesId int64, ses session) {
	storage[sesId] = ses
}

func getSession(sesId int64) session {
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
