package main // import "github.com/looselytyped/going-further"

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
	wg sync.WaitGroup
)

func read(inputs chan<- string) {
	defer close(inputs)
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputs <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func operate(inputs <-chan string, outputs chan<- string) {
	defer wg.Done()
	for in := range inputs {
		hash, _ := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
		outputs <- string(hash)
	}
}

func main() {
	inputs := make(chan string)
	outputs := make(chan string)
	// read in
	wg.Add(1)
	go read(inputs)

	// process
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go operate(inputs, outputs)
	}

	go func() {
		wg.Wait()
		close(outputs)
	}()

	// output
	for v := range outputs {
		fmt.Printf("the output is %x\n", v)
	}
}
