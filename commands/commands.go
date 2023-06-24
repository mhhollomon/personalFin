package commands

import (
	"errors"
	"fmt"
	"log"
)

type CommandLineSpec struct {
	Spec     CommandSpec
	Sentence []string
}

var commands = []CommandLineSpec{
	{Spec: listAccountsSpec, Sentence: []string{"list", "accounts"}},
	{Spec: listAnAccountSpec, Sentence: []string{"list", "account"}},
	{Spec: addAccountSpec, Sentence: []string{"add", "account"}},
}

var commandTree *CommandNode = nil

func buildCommandTree() (*CommandNode, error) {

	if commandTree != nil {
		return commandTree, nil
	}

	// Done this way to break a loop in the initialization since
	// helpConnad references commands
	commands = append(commands,
		CommandLineSpec{Spec: helpCommandSpec, Sentence: []string{"help"}},
	)

	log.Println("buildCommandTree Called")
	root := &CommandNode{Cmd: "Root"}

	for _, c := range commands {
		newCmd := root

		//log.Println("Building command :", i)

		for _, s := range c.Sentence {
			//log.Println("partial =", s)
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

	commandTree = root
	return root, nil
}

// ------------------------------------------------------------
// Find and run a command
// ------------------------------------------------------------
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
			// did not find something with that arg, so assume that it is
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

	return cmd.runCommand(callArgs)
}

//-------------------------------------
// Help command
//-------------------------------------

var helpCommandSpec = CommandSpec{
	Fn:       helpCommand,
	HelpText: `Get list of command`,
}

func helpCommand([]string) error {

	for _, c := range commands {
		for _, s := range c.Sentence {
			fmt.Print(s, " ")
		}
		for _, a := range c.Spec.Args {
			fmt.Printf("<%s> ", a)
		}
		fmt.Print("\n")
		fmt.Println("   ", c.Spec.HelpText)
	}
	return nil
}
