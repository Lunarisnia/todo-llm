package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Read(prompt string) string {
	fmt.Print(prompt, " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}
