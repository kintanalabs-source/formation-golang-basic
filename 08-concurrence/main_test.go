package concurrence

import (
	"reflect"
	"testing"
	"time"
)

func TestSendValue(t *testing.T) {
	ch := make(chan int, 1)
	SendValue(ch, 42)
	select {
	case val := <-ch:
		if val != 42 {
			t.Errorf("SendValue failed, got %d; want 42", val)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("SendValue timed out")
	}
}

func TestSumParallel(t *testing.T) {
	got := SumParallel(10, 20)
	want := 30
	if got != want {
		t.Errorf("SumParallel(10, 20) = %d; want %d", got, want)
	}
}

func TestPingPong(t *testing.T) {
	count := 3
	pings, pongs := PingPong(count)

	wantPings := []string{"ping", "ping", "ping"}
	wantPongs := []string{"pong", "pong", "pong"}

	if !reflect.DeepEqual(pings, wantPings) {
		t.Errorf("PingPong pings = %v; want %v", pings, wantPings)
	}
	if !reflect.DeepEqual(pongs, wantPongs) {
		t.Errorf("PingPong pongs = %v; want %v", pongs, wantPongs)
	}
}
