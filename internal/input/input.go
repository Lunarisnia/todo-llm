package input

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
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

func ReadKeyboard(prompt string) (rune, keyboard.Key) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		r, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			os.Exit(0)
		}
		return r, key
	}

}
