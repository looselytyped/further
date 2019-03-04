package main // import "github.com/looselytyped/going-further"

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type hashed struct {
	input     string
	encrypted string
}

func create(s string) *hashed {
	return &hashed{s, ""}
}

var (
	wg sync.WaitGroup
)

func read(inputs chan<- *hashed) {
	defer close(inputs)
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputs <- create(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func operate(inputs <-chan *hashed, outputs chan<- *hashed) {
	defer wg.Done()
	for in := range inputs {
		hash, _ := bcrypt.GenerateFromPassword([]byte(in.input), bcrypt.DefaultCost)
		in.encrypted = string(hash)
		outputs <- in
	}
}

func main() {
	inputs := make(chan *hashed)
	outputs := make(chan *hashed)
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
		fmt.Printf("the output is %x\n", v.encrypted)
	}
}
