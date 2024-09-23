package telemetry

type Metric struct {
	Name        string
	Description string
	MetricInterface
}

type MetricInterface interface {
	GetName() string
	GetDescription() string
}

func (m *Metric) GetName() string {
	return m.Name
}

func (m *Metric) GetDescription() string {
	return m.Description
}
