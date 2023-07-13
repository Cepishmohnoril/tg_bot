package main

type session struct {
	sesType  string
	nextStep string
	data     []int
}

var storage map[int64]session

func createSession(sesId int64) session {
	storage[sesId] = session{}

	return storage[sesId]
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

func (s *session) setType(sesType string) {
	s.sesType = sesType
}

func (s *session) setNextStep(nextStep string) {
	s.nextStep = nextStep
}

func (s *session) addData(record int) {
	s.data = append(s.data, record)
}
