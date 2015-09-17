package sleep

import "time"

func sleep(sec int) {
	select {
		case <- time.After(time.Duration(sec) * time.Second):
			return
	}
}
