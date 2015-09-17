package sleep

import "time"
import "testing"

func TestSleep(t *testing.T) {
	start := time.Now()
	sleep(2)
	end := time.Now()

	if end.Sub(start) < 2 {
		t.Error("Sleep exited before 2 sec")
	}
}
