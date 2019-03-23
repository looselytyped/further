# Go-ing Further
This repository contains the code for my "Go-ing Further" talk.
It attempts to solve a simple problem, incrementally refactoring the code to make it faster or better organized.

## Usage

- You will need Go [installed](https://golang.org/doc/install)
- Clone this project
- Check out the `encrypt-passwords` branch
- Run `go run main.go < passwords.txt`

## Details

This repository contains individual commits that reveal our attempts at refactoring.
Feel free to look over the Git history and ideally, start at the beginning and checkout each commit to see how we changed our code.
Certain commits are prefixed with `[bad]` that introduces a delta that **will not** work.
See if you can figure out why.

## Credits

This talk was heavily influenced by a talk called ["I came for the easy concurrency I stayed for the easy composition"](https://www.youtube.com/watch?v=woCg2zaIVzQ) presented by the awesome [John Graham-Cumming](https://twitter.com/jgrahamc) at [The European Go conference](https://www.dotgo.eu/).

## License

The code in this repository is released under the [MIT License](https://opensource.org/licenses/MIT).
