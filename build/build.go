package main

import (
	"github.com/pellared/taskflow"
)

func main() {
	flow := taskflow.New()

	hi := flow.MustRegister(taskHello())

	lint := flow.MustRegister(taskLint())

	test := flow.MustRegister(taskflow.Task{
		Name:        "test",
		Description: "go test",
		Command:     taskflow.Exec("go", "test", "./..."),
	})

	flow.MustRegister(taskflow.Task{
		Name:         "all",
		Description:  "workflow",
		Dependencies: taskflow.Deps{hi, lint, test},
	})

	flow.Main()
}

func taskHello() taskflow.Task {
	return taskflow.Task{
		Name:        "hello",
		Description: "my hello world task",
		Command: func(tf *taskflow.TF) {
			tf.Log("Hello world")
			if tf.Params().Bool("ci") {
				tf.Log("I am the CI server!")
			}
		},
	}
}

func taskLint() taskflow.Task {
	return taskflow.Task{
		Name:        "lint",
		Description: "golangci-lint",
		Command: func(tf *taskflow.TF) {
			if err := tf.Cmd("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.38.0").Run(); err != nil {
				tf.Fatalf("golangci-lint install failed: %v", err)
			}
			if err := tf.Cmd("golangci-lint", "run").Run(); err != nil {
				tf.Fatalf("golangci-lint run failed: %v", err)
			}
		},
	}
}
