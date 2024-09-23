package telemetry

import (
	"testing"
)

func TestAttributeCreation(t *testing.T) {
	attrName := "test-attribute"
	attr := &Attribute{Name: attrName}

	if attr.Name != attrName {
		t.Errorf("Expected Attribute Name to be '%s', but got '%s'", attrName, attr.Name)
	}
}

func TestEmptyAttributeCreation(t *testing.T) {
	attr := &Attribute{}

	if attr.Name != "" {
		t.Errorf("Expected Attribute Name to be empty, but got '%s'", attr.Name)
	}
}
