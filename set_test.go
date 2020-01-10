package set

import (
	"testing"
)
                    
func Test(t *testing.T) {
	st := New(7, 4)
	st.Insert(3)
	if st.Size() != 3 {
		t.Errorf("Set should contain three elements")
	}
	if st.Show(0) != "7 4 3 " {
		t.Errorf("Set does not show as intended")
	}
	st.Remove(3)
	if st.Size() != 2 {
		t.Errorf("Set should contain two elements")
	}

	if !st.Contains(4) {
		t.Errorf("Set should contain 4")
	}
	if st.Contains(3) {
		t.Errorf("Set shouldn't contain 3 since it was removed")
	}
}
