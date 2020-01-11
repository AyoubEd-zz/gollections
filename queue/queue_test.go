package queue

import (
	"testing"
)

func Test(t *testing.T) {
	q := New(7, 4)
	
	if q.Show() != "4 7 " {
		t.Errorf("Not the queue we wanted")	
	}
	
	if q.Pop() != 7 {
		t.Errorf("Not the last element we hoped for")
	}

	if q.Size() != 1 {
		t.Errorf("Size of the queue should be 1")
	}
}
