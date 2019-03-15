package codegen

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CodeGen = NewCodeGenAction()

const (
	version = "0.1"
)

type CodeGenAction struct {
	Version bool
	Cmd     *cobra.Command
}

func NewCodeGenAction() *CodeGenAction {

	var action = &CodeGenAction{}
	command := &cobra.Command{
		Use:   "codegen",
		Short: "code generator",
		Long:  "tool for code generator",
		Run: func(cmd *cobra.Command, args []string) {
			if action.Version {
				fmt.Printf("codegen version %s\n", version)
			} else {
				cmd.Usage()
			}
		},
	}
	command.Flags().BoolVarP(&action.Version, "version", "v", false, "print version")
	action.Cmd = command

	return action
}
