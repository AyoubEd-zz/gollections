package stack

import (
	"testing"
)

func Test(t *testing.T) {
	st := New(7)
	st.Insert(4)

	if st.Show() != "7 4" {
		t.Errorf("Not the right stack")
	}

	st.Remove()

	if st.Show() != "7" {
		t.Errorf("Element wasn't removed properly")
	}

	if st.Size() != 1 {
		t.Errorf("Stack should have a size of 1")
	}
}
