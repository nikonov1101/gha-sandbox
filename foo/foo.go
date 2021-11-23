package foo

import (
	"fmt"
)

type foo struct {
	id int
}

func (f *foo) String() string {
	return fmt.Sprintf("<*foo: %d>", f.id)
}

func New(id int) (*foo, error) {
	if id < 1 {
		return nil, fmt.Errorf("invalid id")
	}

	return &foo{id: id}, nil
}
