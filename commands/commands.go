package commands

import (
	"errors"
	"log"
)

type CommandFunction func([]string) error

type Command struct {
	Cmd    string
	SubCmd map[string]*Command
	CmdFn  CommandFunction
}

type CommandSpec struct {
	CmdFn    CommandFunction
	Sentence []string
}

var commands = []CommandSpec{
	{CmdFn: listAccountsCmd, Sentence: []string{"list", "accounts"}},
	{CmdFn: listAnAccountCmd, Sentence: []string{"list", "account"}},
	{CmdFn: AddAccountCmnd, Sentence: []string{"add", "account"}},
}

func buildCommandTree() (*Command, error) {
	log.Println("buildCommandTree Called")
	root := &Command{Cmd: "Root"}

	for i, c := range commands {
		newCmd := root

		log.Println("Building command :", i)

		for _, s := range c.Sentence {
			log.Println("partial =", s)
			if newCmd.SubCmd == nil {
				newCmd.SubCmd = make(map[string]*Command)
			}

			if newCmd.SubCmd[s] == nil {
				newCmd.SubCmd[s] = &Command{Cmd: s}
			}
			newCmd = newCmd.SubCmd[s]
		}

		if newCmd.CmdFn != nil {
			return nil, errors.New("duplicate command detected")
		}

		newCmd.CmdFn = c.CmdFn

	}

	log.Printf("Returning = %+v\n", *root)
	return root, nil
}

func (c *Command) findSubcommand(s string) *Command {
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
			// did find something with that arg, so consider that it is
			// an actual argument to the command.
			index--
			break
		}

		// we only want next if it actually represents a new command
		cmd = next
	}

	if cmd == nil || cmd.CmdFn == nil {
		return errors.New("no such command found")
	}

	callArgs := args[index:]

	log.Printf("Calling fn for command %s (%v)", cmd.Cmd, callArgs)

	err := cmd.CmdFn(callArgs)

	return err
}
