package lamp

import (
	"context"
	"time"

	"github.com/LotauRus/xiao_lamp/internal/pio/digital"
)

type Lamp struct {
	cancelFunc   context.CancelFunc
	pinLight     digital.Pin
	jobChannel   chan time.Duration
	resetChannel chan bool
	switchTimer  *time.Timer

	isLocked bool
}

func New(pinLight digital.Pin) *Lamp {
	lamp := &Lamp{
		pinLight:     pinLight,
		cancelFunc:   func() {},
		switchTimer:  time.NewTimer(time.Minute),
		jobChannel:   make(chan time.Duration),
		resetChannel: make(chan bool),
	}
	lamp.switchTimer.Stop()
	return lamp
}

func (l *Lamp) Start(ctx context.Context) {
	l.cancelFunc()

	ctx, cancel := context.WithCancel(ctx)
	l.cancelFunc = func() {
		cancel()
		l.switchTimer.Stop()
	}
	go l.run(ctx)
}

func (l *Lamp) Stop() {
	l.isLocked = false
	l.cancelFunc()
}

func (l *Lamp) Reset() {
	l.resetChannel <- true
}

func (l *Lamp) SwitchOn(period time.Duration) {
	if l.isLocked {
		return
	}
	l.isLocked = true
	l.jobChannel <- period
}

func (l *Lamp) SwitchOff() {
	if l.IsSwitchedOn() {
		return
	}
	l.switchTimer.Stop()
	l.pinLight.Low()
	l.isLocked = false
}

func (l *Lamp) IsSwitchedOn() bool {
	return l.pinLight.Get()
}

func (l *Lamp) run(ctx context.Context) {
	defer func() {
		l.pinLight.Low()
		l.isLocked = false
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case period, ok := <-l.jobChannel:
			if !ok || period == 0 {
				continue
			}
			l.pinLight.High()
			l.switchTimer.Reset(period)
		case _, ok := <-l.switchTimer.C:
			l.switchTimer.Stop()
			if !ok {
				return
			}
			l.pinLight.Low()
		case reset, ok := <-l.resetChannel:
			if !ok {
				continue
			}
			if reset {
				switch l.IsSwitchedOn() {
				case true:
					l.switchTimer.Stop()
					l.pinLight.Low()
				case false:
					if l.isLocked {
						l.isLocked = false
					}
				}
			}
		}
	}
}
