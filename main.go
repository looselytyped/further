package main // import "github.com/looselytyped/going-further"

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("%x\n", hash)

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error", err)
		}
	}
}
