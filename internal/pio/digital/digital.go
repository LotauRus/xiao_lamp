package digital

type Mode uint8

type Config struct {
	Mode Mode
}

type Pin interface {
	Low()
	High()
	Get() bool
	Configure(config Config)
}
