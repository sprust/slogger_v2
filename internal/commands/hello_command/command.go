package hello_command

import (
	"context"
	"fmt"
)

type Command struct {
}

func (c *Command) Title() string {
	return "Hello!"
}

func (c *Command) Parameters() string {
	return "{no parameters}"
}

func (c *Command) Handle(ctx context.Context, arguments []string) error {
	fmt.Println("Hello, world!")

	return nil
}

func (c *Command) Close() error {
	return nil
}
