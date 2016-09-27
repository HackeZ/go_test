package BufferUages

import (
	"container/list"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var makes int
var frees int

func newBuffer() []byte {
	makes++
	return make([]byte, rand.Intn(5000000)+5000000)
}

type queued struct {
	When  time.Time
	Slice []byte
}

func makeRecycler() (get, give chan []byte) {
	get = make(chan []byte)
	give = make(chan []byte)

	go func() {
		q := new(list.List)

		for {
			if q.Len() == 0 {
				q.PushFront(queued{When: time.Now(), Slice: newBuffer()})
			}

			ele := q.Front()

			timeout := time.NewTimer(time.Minute)
			select {
			case b := <-give:
				timeout.Stop()
				q.PushFront(queued{When: time.Now(), Slice: b})

			case get <- ele.Value.(queued).Slice:
				timeout.Stop()
				q.Remove(ele)

			case <-timeout.C:
				e := q.Front()
				for e != nil {
					n := e.Next()
					if time.Since(e.Value.(queued).When) > time.Minute {
						q.Remove(e)
						e.Value = nil
					}
					e = n
				}
			}
		}

	}()
	return
}

// BufferInAction it is a simple demo in real project.
func BufferInAction() {
	pool := make([][]byte, 20)

	get, give := makeRecycler()

	var m runtime.MemStats
	for {
		b := <-get

		i := rand.Intn(len(pool))
		if pool[i] != nil {
			give <- pool[i]
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
		fmt.Printf("%d,%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc, m.HeapIdle, m.HeapReleased, makes, frees)
	}
}
