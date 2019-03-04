package main // import "github.com/looselytyped/going-further"

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type tasker interface {
	process()
	print()
}

type hashed struct {
	input     string
	encrypted string
}

func (h *hashed) process() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(h.input), bcrypt.DefaultCost)
	h.encrypted = string(hash)
}

func (h *hashed) print() {
	fmt.Printf("the input is %s, the output is %x\n", h.input, h.encrypted)

	// err := bcrypt.CompareHashAndPassword([]byte(h.encrypted), []byte(h.input))
	// if err != nil {
	// 	log.Println(err)
	// }
}

type maker interface {
	create(s string) tasker
}

type hashedMaker struct{}

func (m *hashedMaker) create(s string) tasker {
	return &hashed{s, ""}
}

var (
	wg sync.WaitGroup
)

func read(f maker, inputs chan<- tasker) {
	defer close(inputs)
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputs <- f.create(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func operate(inputs <-chan tasker, outputs chan<- tasker) {
	defer wg.Done()
	for in := range inputs {
		in.process()
		outputs <- in
	}
}

func run(f maker) {
	inputs := make(chan tasker)
	outputs := make(chan tasker)
	// read in
	wg.Add(1)
	go read(f, inputs)

	// process
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go operate(inputs, outputs)
	}

	go func() {
		wg.Wait()
		close(outputs)
	}()

	// output
	for v := range outputs {
		v.print()
	}
}

func main() {
	factory := hashedMaker{}
	run(&factory)
}
