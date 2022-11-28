package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/Nelwhix/todo"
	"time"
	"strconv"
)

const (
	FileName = "todos.json"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
		"Jarvis. Developed by Nelson Isioma \n")
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright " + strconv.Itoa(time.Now().Local().Year()) + "\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}
	task := flag.String("add", "", "Add a new task")
	listItems := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("finish", 0, "Complete a task")
	delete := flag.Int("del", 0, "Delete a task")

	flag.Parse()

	list := &todo.List{}

	if err := list.Get(FileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *listItems:
		fmt.Print(list)
	case *delete > 0:
		if err := list.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete > 0:
		if err := list.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := list.Save(FileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		list.Add(*task)

		if err := list.Save(FileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}