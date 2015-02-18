go-shell-pipes
==============

For running piped shell commands within go, from a single string (splits on white space),

`s, err := pipes.RunString("ps aux | grep usr")`

or as an array of arguments,

```
tokens := []string{"ps", "aux", "|", "grep", "usr", "|", "awk", "{print $2}"}
s, err := pipes.RunStrings(tokens...)
```

or, go plumbling:

```
cmd1 := exec.Command("ps", "aux")
cmd2 := exec.Command("grep", "usr")
cmd3 := exec.Command("awk", "{print $2}")
cmds := []*exec.Cmd{cmd1, cmd2, cmd3}
pipes.AssemblePipes(cmds, os.Stdin, os.Stdout)
s, err := pipes.RunCmds(cmds)
```

