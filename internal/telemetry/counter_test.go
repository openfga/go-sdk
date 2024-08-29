package telemetry

import (
	"testing"
)

func TestCounterCreation(t *testing.T) {
	counterName := "test-counter"
	counterDescription := "This is a test counter."

	counter := &Counter{
		Name:        counterName,
		Description: counterDescription,
	}

	if counter.GetName() != counterName {
		t.Errorf("Expected Counter Name to be '%s', but got '%s'", counterName, counter.GetName())
	}

	if counter.GetDescription() != counterDescription {
		t.Errorf("Expected Counter Description to be '%s', but got '%s'", counterDescription, counter.GetDescription())
	}
}

func TestEmptyCounterCreation(t *testing.T) {
	counter := &Counter{}

	if counter.GetName() != "" {
		t.Errorf("Expected Counter Name to be empty, but got '%s'", counter.GetName())
	}

	if counter.GetDescription() != "" {
		t.Errorf("Expected Counter Description to be empty, but got '%s'", counter.GetDescription())
	}
}

func TestCounterWithWhitespaceName(t *testing.T) {
	counterName := " "
	counterDescription := "Counter with whitespace name."

	counter := &Counter{
		Name:        counterName,
		Description: counterDescription,
	}

	if counter.GetName() != counterName {
		t.Errorf("Expected Counter Name to be '%s', but got '%s'", counterName, counter.GetName())
	}

	if counter.GetDescription() != counterDescription {
		t.Errorf("Expected Counter Description to be '%s', but got '%s'", counterDescription, counter.GetDescription())
	}
}

func TestCounterWithSpecialCharacters(t *testing.T) {
	counterName := "!@#$%^&*()_+{}|:\"<>?"
	counterDescription := "Description with special characters: !@#$%^&*()_+{}|:\"<>?"

	counter := &Counter{
		Name:        counterName,
		Description: counterDescription,
	}

	if counter.GetName() != counterName {
		t.Errorf("Expected Counter Name to be '%s', but got '%s'", counterName, counter.GetName())
	}

	if counter.GetDescription() != counterDescription {
		t.Errorf("Expected Counter Description to be '%s', but got '%s'", counterDescription, counter.GetDescription())
	}
}

func TestCounterWithLongNameAndDescription(t *testing.T) {
	counterName := "ThisIsAVeryLongCounterNameToTestEdgeCasesInTheTelemetryModule"
	counterDescription := "This is a very long description to test how the Counter struct handles long strings."

	counter := &Counter{
		Name:        counterName,
		Description: counterDescription,
	}

	if counter.GetName() != counterName {
		t.Errorf("Expected Counter Name to be '%s', but got '%s'", counterName, counter.GetName())
	}

	if counter.GetDescription() != counterDescription {
		t.Errorf("Expected Counter Description to be '%s', but got '%s'", counterDescription, counter.GetDescription())
	}
}
