package commands

import (
	"errors"
	"log"
)

type CommandFunction func([]string) error

type CommandSpec struct {
	Fn       CommandFunction
	Args     []string
	HelpText string
}

type CommandNode struct {
	Cmd     string
	SubCmd  map[string]*CommandNode
	CmdSpec CommandSpec
}

func (c *CommandNode) findSubcommand(s string) *CommandNode {
	if c == nil || c.SubCmd == nil {
		return nil
	}

	retval, exists := c.SubCmd[s]
	if !exists {
		return nil
	}

	return retval
}

func (c *CommandNode) runCommand(args []string) error {
	if len(args) != len(c.CmdSpec.Args) {
		return errors.New("wrong number of arguments to command")
	}

	log.Printf("Calling fn for command %s (%v)", c.Cmd, args)

	err := c.CmdSpec.Fn(args)

	return err
}
