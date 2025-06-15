package input

import (
	"bufio"
	"log"
	"os"
)

type inputEngine struct {
}

func (i *inputEngine) Read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}
