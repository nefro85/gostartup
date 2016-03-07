package command

import "fmt"

type Command struct {
	id uint64
	text string
}

func(c *Command) Fill(id uint64,text string ){
	c.id = id
	c.text = text
}

func (c *Command) Execute()  {

	fmt.Printf("executing, commandID: %d, content: %s", c.id, c.text)
}