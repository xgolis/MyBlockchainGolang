package firstPhase

import "testing"

// call tests from Tests.go here
func TestAll(t *testing.T) {
	ok := TestHandler()
	if ok != true {
		t.Errorf("TestHandler failed: got %v expected true", ok)
	} else {
		t.Logf("TestHandler OK")
	}

}
