package analog

type MockAPin struct {
	value float32
}

func (p *MockAPin) SetMockValue(value float32) {
	p.value = value
}

func (p *MockAPin) Get(_, _ float32) float32 {
	return p.GetRaw()
}

func (p *MockAPin) GetRaw() float32 {
	return p.value
}

func (p *MockAPin) Configure(_ Config) {
	return
}

func NewMock() *MockAPin {
	return &MockAPin{}
}
