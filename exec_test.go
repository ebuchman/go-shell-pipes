package pipes

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestAssemblePipes(t *testing.T){
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "usr")
	cmd3 := exec.Command("awk", "{print $2}")
	buf := bytes.NewBuffer([]byte{})
	cmds := []*exec.Cmd{cmd1, cmd2, cmd3}
	AssemblePipes(cmds, os.Stdin, os.Stdout)
	if err := RunCmds(cmds); err != nil{
		t.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))
}

func TestString(t *testing.T){
	s, err := RunString("ps aux | grep usr")
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(s)
}

func TestStrings(t *testing.T) {
	tokens := []string{"ps", "aux", "|", "grep", "usr", "|", "awk", "{print $2}"}
	s, err := RunStrings(tokens...)
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(s)
}

func TestLittle(t *testing.T){
	tokens := []string{"ps", "aux"}
	s, err := RunStrings(tokens...)
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(s)
}

