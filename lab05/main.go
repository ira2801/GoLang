/* package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int
	mtx     sync.Mutex
	evenCh  = make(chan int)
	oddCh   = make(chan int)
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	close(evenCh)
	close(oddCh)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go increment(&wg)

	go func() {
		defer wg.Done()
		eCh, oCh := evenCh, oddCh

		for {
			select {
			case val, ok := <-eCh:
				if !ok {
					eCh = nil
				} else if val%3 == 0 {
					mtx.Lock()
					counter++
					mtx.Unlock()
				}
			case val, ok := <-oCh:
				if !ok {
					oCh = nil
				} else if val%33 == 0 {
					mtx.Lock()
					counter--
					mtx.Unlock()
				}
			}

			if eCh == nil && oCh == nil {
				break
			}
		}
	}()

	wg.Wait()
	fmt.Printf("Фінальне значення counter: %d\n", counter)
}
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	evenCh := make(chan int)
	oddCh := make(chan int)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 1000; i++ {
			if i%2 == 0 {
				evenCh <- i
			} else {
				oddCh <- i
			}
		}
		close(evenCh)
		close(oddCh)
	}()

	go func() {
		defer wg.Done()
		eCh, oCh := evenCh, oddCh
		for {
			select {
			case val, ok := <-eCh:
				if !ok {
					eCh = nil
				} else if val%3 == 0 {
					atomic.AddInt64(&counter, 1)
				}
			case val, ok := <-oCh:
				if !ok {
					oCh = nil
				} else if val%33 == 0 {
					atomic.AddInt64(&counter, -1)
				}
			}
			if eCh == nil && oCh == nil {
				break
			}
		}
	}()

	wg.Wait()
	fmt.Printf("Фінальне значення counter: %d\n", atomic.LoadInt64(&counter))
}
