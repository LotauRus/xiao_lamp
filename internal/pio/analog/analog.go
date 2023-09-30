package analog

type Range struct {
	Min float32
	Max float32
}

type Config struct {
	Reference  uint32
	Resolution uint32
	Samples    uint32
	Range      Range
}

type Pin interface {
	Configure(config Config)
	GetRaw() uint16
	Get(from, to float32) float32
}
