package scheduler

import (
	"testing"
	"time"
)

func TestShouldRunTaskEverySecond(t *testing.T) {
	resultChan := make(chan int, 2)
	funcToRun := func() {
		resultChan <- 1
	}

	RunTaskAtInterval(funcToRun, time.Second*1, time.Second*0)
	if len(resultChan) != 0 {
		t.Error()
	}
	time.Sleep(time.Millisecond * 2500)
	if len(resultChan) != 2 {
		t.Error()
	}
}

func TestShouldRunTaskEverySecondWithStartTimeShift(t *testing.T) {
	resultChan := make(chan int, 2)
	funcToRun := func() {
		resultChan <- 1
	}

	RunTaskAtInterval(funcToRun, time.Second*1, time.Second*1)
	if len(resultChan) != 0 {
		t.Error()
	}
	time.Sleep(time.Millisecond * 2500)
	if len(resultChan) != 1 {
		t.Error()
	}
	time.Sleep(time.Second * 1)
	if len(resultChan) != 2 {
		t.Error()
	}
}
