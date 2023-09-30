package main

import (
	"context"
	"machine"
	"time"

	"github.com/LotauRus/xiao_lamp/internal/lamp"
	"github.com/LotauRus/xiao_lamp/internal/pio"
	"github.com/LotauRus/xiao_lamp/internal/pio/analog"
	"github.com/LotauRus/xiao_lamp/internal/pio/digital"
	"github.com/LotauRus/xiao_lamp/internal/resolver"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	machine.InitADC()

	lightSensor := pio.NewAnalog(machine.A0, analog.Config{Range: analog.Range{Min: 0, Max: 65535}})
	pinReset := pio.NewDigital(machine.D4, digital.Config{Mode: digital.Mode(machine.PinInputPullup)})
	switchOnTime := pio.NewAnalog(machine.A6, analog.Config{Range: analog.Range{Min: 0, Max: 65535}})
	pinLamp := pio.NewDigital(machine.D7, digital.Config{Mode: digital.Mode(machine.PinOutput)})
	darkLevel := pio.NewAnalog(machine.A8, analog.Config{Range: analog.Range{Min: 0, Max: 65535}})

	lmp := lamp.New(pinLamp)
	lmp.Start(ctx)

	w := resolver.New(lmp, lightSensor, switchOnTime, pinReset, darkLevel)
	w.Start(ctx)

	for {
		time.Sleep(time.Millisecond * 3000)
	}
}
