package lamp

import (
	"context"
	"testing"
	"time"

	"github.com/LotauRus/xiao_lamp/internal/pio/digital"
)

func TestLamp(t *testing.T) {
	pinLamp := digital.NewMock()
	testLamp := New(pinLamp)

	testLamp.Start(context.Background())
	defer testLamp.Stop()

	t.Run("включение на 1 секунду", func(t *testing.T) {
		testLamp.SwitchOn(time.Second * 1)
		if !testLamp.IsSwitchedOn() {
			t.Errorf("Isn't switchedOn! state = %v \n", testLamp.IsSwitchedOn())
		}
		time.Sleep(time.Second * 3)
		if testLamp.IsSwitchedOn() {
			t.Errorf("Isn't switchedOff! state = %v \n", testLamp.IsSwitchedOn())
		}
	})

	t.Run("повторно не должно включиться при отправке команды", func(t *testing.T) {
		testLamp.SwitchOn(time.Second * 1)
		if testLamp.IsSwitchedOn() {
			t.Errorf("Isn't switchedOff! state = %v \n", testLamp.IsSwitchedOn())
		}
	})

	t.Run("после выключения не должно включится", func(t *testing.T) {
		testLamp.SwitchOff()
		testLamp.SwitchOn(time.Second * 5)
		if testLamp.IsSwitchedOn() {
			t.Errorf("Isn't switchedOff! state = %v \n", testLamp.IsSwitchedOn())
		}
	})

	t.Run("после остановки и нового запуска всё должно работать", func(t *testing.T) {
		testLamp.Stop()
		testLamp.Start(context.Background())

		testLamp.SwitchOn(time.Second * 1)
		if !testLamp.IsSwitchedOn() {
			t.Errorf("Isn't switchedOn after stop! state = %v \n", testLamp.IsSwitchedOn())
		}
		time.Sleep(time.Second * 3)
		if testLamp.IsSwitchedOn() {
			t.Errorf("Isn't switchedOff after stop! state = %v \n", testLamp.IsSwitchedOn())
		}
	})
}
