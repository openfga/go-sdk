package telemetry

import (
	"testing"
)

func TestHistogramCreation(t *testing.T) {
	histogramName := "request_duration"
	histogramUnit := "milliseconds"
	histogramDescription := "The duration of client requests."

	histogram := &Histogram{
		Name:        histogramName,
		Unit:        histogramUnit,
		Description: histogramDescription,
	}

	if histogram.GetName() != histogramName {
		t.Errorf("Expected Histogram Name to be '%s', but got '%s'", histogramName, histogram.GetName())
	}

	if histogram.GetUnit() != histogramUnit {
		t.Errorf("Expected Histogram Unit to be '%s', but got '%s'", histogramUnit, histogram.GetUnit())
	}

	if histogram.GetDescription() != histogramDescription {
		t.Errorf("Expected Histogram Description to be '%s', but got '%s'", histogramDescription, histogram.GetDescription())
	}
}

func TestEmptyHistogramCreation(t *testing.T) {
	histogram := &Histogram{}

	if histogram.GetName() != "" {
		t.Errorf("Expected Histogram Name to be empty, but got '%s'", histogram.GetName())
	}

	if histogram.GetUnit() != "" {
		t.Errorf("Expected Histogram Unit to be empty, but got '%s'", histogram.GetUnit())
	}

	if histogram.GetDescription() != "" {
		t.Errorf("Expected Histogram Description to be empty, but got '%s'", histogram.GetDescription())
	}
}

func TestHistogramWithSpecialCharacters(t *testing.T) {
	histogramName := "request_duration!@#$%"
	histogramUnit := "ms!@#$%"
	histogramDescription := "The duration of client requests!@#$%."

	histogram := &Histogram{
		Name:        histogramName,
		Unit:        histogramUnit,
		Description: histogramDescription,
	}

	if histogram.GetName() != histogramName {
		t.Errorf("Expected Histogram Name to be '%s', but got '%s'", histogramName, histogram.GetName())
	}

	if histogram.GetUnit() != histogramUnit {
		t.Errorf("Expected Histogram Unit to be '%s', but got '%s'", histogramUnit, histogram.GetUnit())
	}

	if histogram.GetDescription() != histogramDescription {
		t.Errorf("Expected Histogram Description to be '%s', but got '%s'", histogramDescription, histogram.GetDescription())
	}
}

func TestHistogramWithLongStrings(t *testing.T) {
	histogramName := "this_is_a_very_long_histogram_name_to_test_edge_cases_in_the_telemetry_module"
	histogramUnit := "milliseconds_with_long_unit_name"
	histogramDescription := "This is a very long description to test how the Histogram struct handles long strings in the telemetry module."

	histogram := &Histogram{
		Name:        histogramName,
		Unit:        histogramUnit,
		Description: histogramDescription,
	}

	if histogram.GetName() != histogramName {
		t.Errorf("Expected Histogram Name to be '%s', but got '%s'", histogramName, histogram.GetName())
	}

	if histogram.GetUnit() != histogramUnit {
		t.Errorf("Expected Histogram Unit to be '%s', but got '%s'", histogramUnit, histogram.GetUnit())
	}

	if histogram.GetDescription() != histogramDescription {
		t.Errorf("Expected Histogram Description to be '%s', but got '%s'", histogramDescription, histogram.GetDescription())
	}
}
