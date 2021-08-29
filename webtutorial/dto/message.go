package dto

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Todo struct {
	Author string `json:"author"`
	Title  string `json:"title,omitempty"`
	Done   bool   `json:"done,omitempty"`
}

func (t *Todo) Marshal() ([]byte, error) {
	if t == nil {
		return nil, errors.New("Todo is nil")
	}
	return json.Marshal(t)
}

func (t *Todo) Unmarshal(payload []byte) error {
	object := Todo{}
	err := json.Unmarshal(payload, &object)
	fmt.Println("Object is ", object, err, string(payload))
	*t = object
	return err
}
