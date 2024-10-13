/*
package main

import (
	"fmt"
)

func sendData(ch chan string) {
    ch <- "Sample goroutine message!"
}

func main() {
    ch := make(chan string)
    go sendData(ch)
    msg := <-ch
    fmt.Println(msg)
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
*/
/*
package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for b := range ch {
		fmt.Println(b)
	}
}
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	chA := make(chan string)
	chB := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chA <- "Hello from Person A"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chB <- "How's it going from person B"
	}()

	select {
	case msg1 := <-chA:
		fmt.Println(msg1)
	case msg2 := <-chB:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! See You Later!!")
	}
}
