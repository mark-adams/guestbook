package main

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"os/exec"
)

func runCommand(cmdPath string) {

	// ruleid:dangerous-exec-cmd
	cmd := &exec.Cmd{
		Path:   cmdPath,
		Args:   []string{"foo", "bar"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	cmd.Start()
}

func hashPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	sumBytes := h.Sum(nil)
	return hex.EncodeToString(sumBytes)
}
