go-shell-pipes
==============

For running piped shell commands within go, from a single string (splits on white space),

`fmt.Println(pipes.RunString("ps aux | grep usr"))`

or as an array of arguments,

```
tokens := []string{"ps", "aux", "|", "grep", "usr", "|", "awk", "{print $2}"}
fmt.Println(pipes.RunStrings(tokens...)
```

or, go plumbling:

```
cmd1 := exec.Command("ps", "aux")
cmd2 := exec.Command("grep", "usr")
cmd3 := exec.Command("awk", "{print $2}")
cmds := []*exec.Cmd{cmd1, cmd2, cmd3}
pipes.AssemblePipes(cmds, os.Stdin, os.Stdout)
fmt.Println(pipes.RunCmds(cmds))
```

