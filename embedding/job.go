package main

import (
	"fmt"
	"log"
	"os"
)

type Job struct {
	Command string
	*log.Logger
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func (job *Job) Printf(format string, args ...interface{}) {
	job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}

func main() {
	command := "ls"
	user := os.Getenv("USER")
	job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Printf("/home/%s", user)
}
