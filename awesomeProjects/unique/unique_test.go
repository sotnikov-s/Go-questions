package main

import (
	"testing"
	"io/ioutil"
)

func TestUnique(t *testing.T) {

	expectedValue, err := ioutil.ReadFile("test/testIn.md")
	if err != nil {
		println("TestUnique failed:", err)
	}

	unique("test/testIn.md", "test/testOut.md")

	ObtainedValue, err := ioutil.ReadFile("test/testOut.md")
	if err != nil {
		println("TestUnique failed:", err)
	}

	if expectedValue == nil && ObtainedValue == nil {
		t.Errorf("TestUnique failed in stage of comparison: %d", err)
	}

	if expectedValue == nil || ObtainedValue == nil {
		t.Errorf("TestUnique failed in stage of comparison: %d", err)
	}

	if len(expectedValue) != len(ObtainedValue) {
		t.Errorf("TestUnique failed in stage of comparison: %d", err)
	}

	for i := range expectedValue {
		if expectedValue[i] != ObtainedValue[i] {
			t.Errorf("TestUnique failed in stage of comparison: %d", err)
		}
	}
}


