package resolver

import (
	"context"
	"time"

	"github.com/LotauRus/xiao_lamp/internal/lamp"
	"github.com/LotauRus/xiao_lamp/internal/pio/analog"
	"github.com/LotauRus/xiao_lamp/internal/pio/digital"
)

const diffLevel = 15

type Worker struct {
	cancelFunc   context.CancelFunc
	lamp         *lamp.Lamp
	lightSensor  analog.Pin
	switchOnTime analog.Pin
	darkLevel    analog.Pin
	pinReset     digital.Pin
}

func New(lamp *lamp.Lamp, lightSensor analog.Pin, switchOnTime analog.Pin, pinReset digital.Pin, lowLevel analog.Pin) *Worker {
	return &Worker{
		lamp:         lamp,
		lightSensor:  lightSensor,
		switchOnTime: switchOnTime,
		pinReset:     pinReset,
		darkLevel:    lowLevel,
	}
}

func (w *Worker) Start(ctx context.Context) {
	go w.run(ctx)
	go w.runCheckReset(ctx)
}

func (w *Worker) run(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			w.lamp.Stop()
			return
		}

		switchTime := w.switchOnTime.Get(0, 300)
		lightSensor := w.lightSensor.Get(0, 100)
		lowLevel := w.darkLevel.Get(0, 100-diffLevel)
		highLevel := lowLevel + diffLevel

		if lightSensor < lowLevel {
			w.lamp.SwitchOn(time.Second * time.Duration(switchTime))
		}

		if lightSensor > highLevel {
			w.lamp.SwitchOff()
		}

		time.Sleep(time.Millisecond * 500)
	}
}

func (w *Worker) runCheckReset(ctx context.Context) {
	state := true
	for {
		if ctx.Err() != nil {
			return
		}

		time.Sleep(time.Millisecond * 50)
		value := w.pinReset.Get()

		if value != state {
			state = value
			if state == false {
				w.lamp.Reset()
				time.Sleep(time.Millisecond * 1000)
			}
		}
	}
}
