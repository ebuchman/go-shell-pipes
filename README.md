go-shell-pipes
==============

For running piped shell commands within go:

`pipes.RunString("ps aux | grep usr")`

or, 

```
cmd1 := exec.Command("ps", "aux")
cmd2 := exec.Command("grep", "usr")
cmd3 := exec.Command("awk", "{print $2}")
cmds := []*exec.Cmd{cmd1, cmd2, cmd3}
pipes.AssemblePipes(cmds, os.Stdin, os.Stdout)
pipes.RunCmds(cmds)
```

