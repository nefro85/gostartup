package command

import (
	"fmt"
	"strings"
)

const CMD_QUIT = "quit"

type Command struct {
	id   uint64
	text string
}

func (c *Command) Fill(id uint64, text string) {
	c.id = id
	c.text = text
}

func (c *Command) Execute() {

	fmt.Printf("executing, commandID: %d, content: %s", c.id, c.text)
}

func (c *Command) IsQuit() bool {
	return strings.Contains(c.text, CMD_QUIT)
}
