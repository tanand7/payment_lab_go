package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func ReadLine() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ReadNonEmptyString(prompt string) string {

	for {
		fmt.Println(prompt)

		line := ReadLine()
		if line != "" {
			return line
		}
		fmt.Println("Value cannot be empty. Please try again")
	}

}
