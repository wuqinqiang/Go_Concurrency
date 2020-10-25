package main

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue(100)
	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
	if len(q.data) != 100 {
		t.Errorf("queue len is not 100")
	}
}
