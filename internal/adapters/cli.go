package adapters

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	Reader *bufio.Reader
}

func NewCLI() *CLI {
	return &CLI{
		Reader: bufio.NewReader(os.Stdin),
	}
}

func (cli *CLI) ReadInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := cli.Reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (cli *CLI) PrintOutput(output string) {
	fmt.Println(output)
}
