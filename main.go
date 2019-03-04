package main // import "github.com/looselytyped/going-further"

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var wg sync.WaitGroup
	inputs := make(chan string)

	wg.Add(1)
	go func() {
		defer close(inputs)
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputs <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for in := range inputs {
				hash, _ := bcrypt.GenerateFromPassword([]byte(in), bcrypt.DefaultCost)
				fmt.Printf("%x\n", hash)
			}
		}()
	}

	wg.Wait()
}
