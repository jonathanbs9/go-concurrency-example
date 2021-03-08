package main

import (
	"fmt"
	"time"
)

func main() {
	// Creamos 2 channels
	out1 := make(chan string)
	out2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			out1 <- "Orden procesada!"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			out2 <- "Reembolso procesado!"
		}
	}()

	for {
		select {
		// Cuando out1 tiene lista info, la imprime!
		case msg := <-out1:
			fmt.Println(msg)
		// Cuando out2 tiene lista info, la imprime!
		// Si out2 no esta lista, rechequea out1 para ver si tiene info lista
		case msg := <-out2:
			fmt.Println(msg)
		}
	}

}

func process(item string, out chan string) {
	defer close(out)
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		// Le agrego data al channel
		out <- item
	}

}
