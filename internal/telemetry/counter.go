package telemetry

type Counter struct {
	Name        string
	Description string
}

func (m *Counter) GetName() string {
	return m.Name
}

func (m *Counter) GetDescription() string {
	return m.Description
}
