package pio

import (
	"machine"

	"github.com/LotauRus/xiao_lamp/internal/pio/analog"
	"github.com/LotauRus/xiao_lamp/internal/pio/digital"
	"github.com/LotauRus/xiao_lamp/internal/support"
)

type Digital struct {
	pin machine.Pin
}

func NewDigital(pin machine.Pin, config digital.Config) digital.Pin {
	dPin := &Digital{
		pin: pin,
	}
	dPin.Configure(config)
	return dPin
}

func (p *Digital) Low() {
	p.pin.Low()
}

func (p *Digital) High() {
	p.pin.High()
}

func (p *Digital) Get() bool {
	return p.pin.Get()
}

func (p *Digital) Configure(config digital.Config) {
	p.pin.Configure(machine.PinConfig{Mode: machine.PinMode(config.Mode)})
}

type Analog struct {
	pin    machine.ADC
	config analog.Config
}

func NewAnalog(pin machine.Pin, config analog.Config) analog.Pin {
	analogPin := &Analog{
		pin: machine.ADC{Pin: pin},
	}
	analogPin.Configure(config)
	return analogPin
}

func (p *Analog) GetRaw() uint16 {
	return p.pin.Get()
}

func (p *Analog) Get(from, to float32) float32 {
	return support.MapValue(p.GetRaw(), p.config.Range.Min, p.config.Range.Max, from, to)
}

func (p *Analog) Configure(config analog.Config) {
	p.config = config
	p.pin.Configure(machine.ADCConfig{
		Reference:  config.Reference,
		Resolution: config.Resolution,
		Samples:    config.Samples,
	})
}
