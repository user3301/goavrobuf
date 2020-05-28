package goavrobuf

type void struct{}

type RecordSet map[string]void

func (s RecordSet) Add(recordName string) {
	s[recordName] = struct{}{}
}

func (s RecordSet) Remove(recordName string) {
	delete(s, recordName)
}

func (s RecordSet) Contain(recordName string) bool {
	_, ok := s[recordName]
	return ok
}
