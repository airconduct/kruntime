package api

import "encoding/json"

type Task struct {
	Num int
}

func (t *Task) Unmarshal() []byte {
	raw, _ := json.Marshal(t)
	return raw
}
