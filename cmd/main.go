package main

import (
	"bytes"
	"fmt"
	"github.com/ebuchman/go-shell-pipes"
	"os"
	"os/exec"
)

func main() {
	//s := pipes.RunString("ps aux | grep usr")
	//fmt.Println(s)

	//s := pipes.Run("ps aux | grep usr | awk '{print $2}'")
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "usr")
	cmd3 := exec.Command("awk", "{print $2}")
	buf := bytes.NewBuffer([]byte{})
	cmds := []*exec.Cmd{cmd1, cmd2, cmd3}
	pipes.AssemblePipes(cmds, os.Stdin, os.Stdout)
	pipes.RunCmds(cmds)
	fmt.Println(string(buf.Bytes()))

}
