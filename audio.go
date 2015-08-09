package main

import (
	"strings"
	"path/filepath"
	"os/exec"
)

type Audio struct {
	Cmd 		*exec.Cmd
	Filename 	string
	Name		string
}

func newAudio(filename string) *Audio{
	var cmd *exec.Cmd
	
	audioType := filepath.Ext(filename)
	
	audioType = strings.ToLower(audioType)
			
	switch(audioType){
		case ".mp3":
			cmd = exec.Command("mpg123", filename, " &")
		case ".wav":
			cmd = exec.Command("aplay", filename)
	}
	
	substr := strings.Split(filename, ".")
	name := substr[0]

	return &Audio{cmd, filename, name}
}

func (a *Audio) Stop() {
	a.Cmd.Process.Kill()
}
