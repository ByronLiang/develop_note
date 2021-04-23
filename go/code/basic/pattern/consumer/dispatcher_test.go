package consumer

import (
	"fmt"
	"testing"
)

type DisHandler struct {
}

func (*DisHandler) Do(data interface{})  {
	switch data.(type) {
	case int:
		fmt.Println("num: ", data.(int))
	case string:
		fmt.Println("str: ", data.(string))
	default:
		fmt.Println("default", data)
	}
}

func TestDispatcher_New(t *testing.T) {
	handle := new(DisHandler)
	dip := NewDispatcher(5, 2, handle)
	for i := 0; i < 10; i++ {
		dip.Add(fmt.Sprintf("test: %d", i))
	}
	dip.Exit()
}
