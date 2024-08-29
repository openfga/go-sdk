package telemetry

type Histogram struct {
	Name        string
	Unit        string
	Description string
}

func (m *Histogram) GetName() string {
	return m.Name
}

func (m *Histogram) GetDescription() string {
	return m.Description
}

func (m *Histogram) GetUnit() string {
	return m.Unit
}
