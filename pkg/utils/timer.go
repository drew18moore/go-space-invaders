package utils

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	currentTicks int
	targetTicks  int
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		currentTicks: 0,
		targetTicks:  int(d.Milliseconds()) * ebiten.TPS() / 1000,
	}
}

func (t *Timer) Update() {
	if t.currentTicks < t.targetTicks {
		t.currentTicks++
	}
}

func (t *Timer) IsReady() bool {
	return t.currentTicks >= t.targetTicks
}

func (t *Timer) Reset() {
	t.currentTicks = 0
}

func (t *Timer) SetDuration(d time.Duration) {
	t.targetTicks = int(d.Milliseconds()) * ebiten.TPS() / 1000
}

func (t *Timer) DecreaseTimer(d time.Duration) {
	newDuration := time.Duration(t.targetTicks*1000/ebiten.TPS()) * time.Millisecond - d
	t.SetDuration(newDuration)
}

func (t *Timer) CurrentTarget() int {
	return t.targetTicks
}
