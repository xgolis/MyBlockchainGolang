package thirdphase

import (
	"testing"
)

func TestAll(t *testing.T) {
	err := TestProcessBlock()
	if err != nil {
		t.Errorf("TestProcessBlock failed: %v", err)
	} else {
		t.Logf("TestProcessBlock OK")
	}

	err = TestWrongValues()
	if err != nil {
		t.Errorf("TestWrongValues failed: %v", err)
	} else {
		t.Logf("TestWrongValues OK")
	}

	err = TestCreateBlock()
	if err != nil {
		t.Errorf("TestCreateBlock failed: %v", err)
	} else {
		t.Logf("TestCreateBlock OK")
	}
}
