package commands

import (
	"errors"
	"log"
)

type CommandFunction func([]string) error

type CommandSpec struct {
	Fn   CommandFunction
	Args []string
}

type CommandNode struct {
	Cmd     string
	SubCmd  map[string]*CommandNode
	CmdSpec CommandSpec
}

type CommandLineSpec struct {
	Spec     CommandSpec
	Sentence []string
}

var commands = []CommandLineSpec{
	{Spec: listAccountsSpec, Sentence: []string{"list", "accounts"}},
	{Spec: listAnAccountSpec, Sentence: []string{"list", "account"}},
	{Spec: addAccountSpec, Sentence: []string{"add", "account"}},
}

func buildCommandTree() (*CommandNode, error) {
	log.Println("buildCommandTree Called")
	root := &CommandNode{Cmd: "Root"}

	for i, c := range commands {
		newCmd := root

		log.Println("Building command :", i)

		for _, s := range c.Sentence {
			log.Println("partial =", s)
			if newCmd.SubCmd == nil {
				newCmd.SubCmd = make(map[string]*CommandNode)
			}

			if newCmd.SubCmd[s] == nil {
				newCmd.SubCmd[s] = &CommandNode{Cmd: s}
			}
			newCmd = newCmd.SubCmd[s]
		}

		if newCmd.CmdSpec.Fn != nil {
			return nil, errors.New("duplicate command detected")
		}

		newCmd.CmdSpec = c.Spec

	}

	log.Printf("Returning = %+v\n", *root)
	return root, nil
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

func Execute(args []string) error {

	log.Println("Execute: called with args:", args)
	index := 0
	cmd, cerr := buildCommandTree()
	if cerr != nil {
		return cerr
	}
	log.Println("walking command tree")
	for {
		if index >= len(args) {
			break
		}
		next := cmd.findSubcommand(args[index])
		index++

		if next == nil {
			// did not find something with that arg, so consider that it is
			// an actual argument to the command.
			index--
			break
		}

		// we only want next if it actually represents a new command
		cmd = next
	}

	if cmd == nil || cmd.CmdSpec.Fn == nil {
		return errors.New("no such command found")
	}

	callArgs := args[index:]

	if len(callArgs) != len(cmd.CmdSpec.Args) {
		return errors.New("wrong number of arguments to command")
	}

	log.Printf("Calling fn for command %s (%v)", cmd.Cmd, callArgs)

	err := cmd.CmdSpec.Fn(callArgs)

	return err
}
