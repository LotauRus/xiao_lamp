package digital

type MockDPin struct {
	state bool
}

func (p *MockDPin) Low() {
	p.state = false
}

func (p *MockDPin) High() {
	p.state = true
}

func (p *MockDPin) Get() bool {
	return p.state
}

func (p *MockDPin) Configure(_ Config) {
	return
}

func NewMock() *MockDPin {
	return &MockDPin{}
}
