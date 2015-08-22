package main

import (
	"strings"
	"os/exec"
	"syscall" 
	"encoding/binary"
)

type Audio struct {
	Cmd 		*exec.Cmd
	Filename 	string
	Name		string
	Duration 	int64
}

func newAudio(filename string) *Audio{
	var cmd, dur *exec.Cmd
	
	file := "./audio/" + filename
	
	substr := strings.Split(filename, ".")
	name := substr[0]
	
	dur = exec.Command("soxi", "-D", file)
	
	b_dur,_ := dur.Output()
	
	i_dur,_ := binary.Varint(b_dur)

	return &Audio{cmd, file, name, i_dur}
}

func (a *Audio) Play() error {
	
	a.Cmd = exec.Command("play", "-q", a.Filename)
	
	return a.Cmd.Run()
}

func (a *Audio) Stop() error {
	return a.Cmd.Process.Kill()
}

func (a *Audio) Pause() error {
	return a.Cmd.Process.Signal(syscall.SIGSTOP)
}

func (a *Audio) Resume() error {
	return a.Cmd.Process.Signal(syscall.SIGCONT)
}
