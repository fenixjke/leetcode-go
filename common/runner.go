package common

import (
	"fmt"
	"testing"
)

type TaskRunner interface {
	Run(option string)
}

type Info struct {
	Runner   TaskRunner
	Options  []string
	TaskName string
}

func (i Info) Run(option string) {
	f := func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			i.Runner.Run(option)
		}
	}
	for _, v := range i.Options {
		option = v
		br := testing.Benchmark(f)
		fmt.Printf("Task: \"%s\", Option: %v -> %s\n", i.TaskName, v, br)
	}
}
