package worker

import "testing"

func TestConsumerLimit(t *testing.T) {
	ConsumerLimit()
}

func TestCalChanLen(t *testing.T) {
	res := CalChanLen()
	t.Log(res)
}
