package BufferUages

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// GoodBufferUsage a good buffer usage.
func GoodBufferUsage() {
	pool := make([][]byte, 20)

	buffer := make(chan []byte, 5)

	var m runtime.MemStats
	makes := 0

	for {
		var b []byte

		select {
		case b = <-buffer:
		default:
			makes++
			b = makeBuffer()
		}

		i := rand.Intn(len(pool))
		if pool[i] != nil {
			select {
			case buffer <- pool[i]:
				pool[i] = nil
			default:
			}
		}
		pool[i] = b

		time.Sleep(time.Second)

		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
			m.HeapIdle, m.HeapReleased, makes)
	}
}
