package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func Root_Path() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicln("worong", err)
	}
	i := strings.LastIndex(s, "//")
	path := s[:i+1]
	return path
}
