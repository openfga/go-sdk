package telemetry

import (
	"testing"
)

type MockMetric struct {
	Name        string
	Description string
}

func (m *MockMetric) GetName() string {
	return m.Name
}

func (m *MockMetric) GetDescription() string {
	return m.Description
}

func TestMetric_GetName(t *testing.T) {
	metricName := "test-metric"
	metricDescription := "This is a test metric."

	metric := &Metric{
		Name:        metricName,
		Description: metricDescription,
		MetricInterface: &MockMetric{
			Name:        metricName,
			Description: metricDescription,
		},
	}

	if metric.GetName() != metricName {
		t.Errorf("Expected Metric Name to be '%s', but got '%s'", metricName, metric.GetName())
	}

	if metric.GetDescription() != metricDescription {
		t.Errorf("Expected Metric Description to be '%s', but got '%s'", metricDescription, metric.GetDescription())
	}
}

func TestMetricInterfaceImplementation(t *testing.T) {
	metricName := "interface-metric"
	metricDescription := "Metric implemented using MetricInterface."

	mockMetric := &MockMetric{
		Name:        metricName,
		Description: metricDescription,
	}

	metric := &Metric{
		Name:            metricName,
		Description:     metricDescription,
		MetricInterface: mockMetric,
	}

	if metric.MetricInterface.GetName() != metricName {
		t.Errorf("Expected MetricInterface Name to be '%s', but got '%s'", metricName, metric.MetricInterface.GetName())
	}

	if metric.MetricInterface.GetDescription() != metricDescription {
		t.Errorf("Expected MetricInterface Description to be '%s', but got '%s'", metricDescription, metric.MetricInterface.GetDescription())
	}
}

func TestEmptyMetric(t *testing.T) {
	metric := &Metric{}

	if metric.GetName() != "" {
		t.Errorf("Expected Metric Name to be empty, but got '%s'", metric.GetName())
	}

	if metric.GetDescription() != "" {
		t.Errorf("Expected Metric Description to be empty, but got '%s'", metric.GetDescription())
	}
}
