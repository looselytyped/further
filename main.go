package main // import "github.com/looselytyped/going-further"

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	inputs := make(chan string)

	defer close(inputs)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputs <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for {
		in, ok := <-inputs
		if !ok {
			break
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
		fmt.Printf("%x\n", hash)
	}
}
