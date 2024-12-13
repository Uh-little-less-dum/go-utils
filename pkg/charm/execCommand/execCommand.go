package charm_exec_command

import (
	"io"

	tea "github.com/charmbracelet/bubbletea"
)

type charmExecCommand struct {
	f func() error
}

func (c charmExecCommand) Run() error {
	return c.f()
}

func (c charmExecCommand) SetStdin(io.Reader)  {}
func (c charmExecCommand) SetStdout(io.Writer) {}
func (c charmExecCommand) SetStderr(io.Writer) {}

func NewCharmExecCommand(f func() error) tea.ExecCommand {
	c := charmExecCommand{
		f: func() error {
			err := f()
			return err
		},
	}
	return c
}
