package pipes

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
)

// Convert a shell command with a series of pipes into
// correspondingly piped list of *exec.Cmd
// If an arg has spaces, this will fail
func RunString(s string) string {
	buf := bytes.NewBuffer([]byte{})
	sp := strings.Split(s, "|")
	cmds := make([]*exec.Cmd, len(sp))
	// create the commands
	for i, c := range sp {
		cs := strings.Split(strings.TrimSpace(c), " ")
		cmd := cmdFromStrings(cs)
		cmds[i] = cmd
	}

	cmds = AssemblePipes(cmds, nil, buf)
	RunCmds(cmds)

	b := buf.Bytes()
	return string(b)
}

func cmdFromStrings(cs []string) *exec.Cmd {
	if len(cs) == 1 {
		return exec.Command(cs[0])
	}
	return exec.Command(cs[0], cs[1:]...)
}

// Convert sequence of tokens into commands,
// using "|" as a delimiter
func RunStrings(tokens ...string) string {
	buf := bytes.NewBuffer([]byte{})
	cmds := []*exec.Cmd{}
	args := []string{}
	// accumulate tokens until a |
	for _, t := range tokens {
		if t != "|" {
			args = append(args, t)
		} else {
			cmds = append(cmds, cmdFromStrings(args))
			args = []string{}
		}
	}
	cmds = append(cmds, cmdFromStrings(args))
	cmds = AssemblePipes(cmds, nil, buf)
	RunCmds(cmds)

	b := buf.Bytes()
	return string(b)
}

// Pipe stdout of each command to into stdin of next
func AssemblePipes(cmds []*exec.Cmd, stdin io.Reader, stdout io.Writer) []*exec.Cmd {
	cmds[0].Stdin = stdin
	// assemble pipes
	for i, c := range cmds {
		if i < len(cmds)-1 {
			cmds[i+1].Stdin, _ = c.StdoutPipe()
		} else {
			c.Stdout = stdout
		}
	}
	return cmds
}

// Run series of piped commands
func RunCmds(cmds []*exec.Cmd) {
	// start processes in descending order
	for i := len(cmds) - 1; i > 0; i-- {
		cmds[i].Start()
	}
	// run the first process
	cmds[0].Run()
	// wait on processes in ascending order
	for i := 1; i < len(cmds); i++ {
		cmds[i].Wait()
	}
}
