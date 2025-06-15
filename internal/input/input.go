package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type InputEngine interface {
	Read(prompt string) string
}

type inputEngine struct {
}

func NewInputEngine() InputEngine {
	return &inputEngine{}
}

func (i *inputEngine) Read(prompt string) string {
	fmt.Print(prompt, " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}
