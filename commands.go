package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

)

type Commands struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCommandFlags() *Commands {
	Comflags := Commands{}
	flag.StringVar(&Comflags.Add ,"add", "", "Add a new Todo specify title") //go run ./ -add "This is a Project"
	flag.StringVar(&Comflags.Edit ,"edit", "", "Edit the Todo using index and add new title id:id.new_title") // go run ./ -edit 2:Dickiao
	flag.IntVar(&Comflags.Delete ,"delete", -1, "Specific the todo by id to delete that") //go run ./ -delete 1
	flag.IntVar(&Comflags.Toggle ,"toggle", -1, "Specific the tod0 by its id and Change the Completed tag in the file") //go run ./ -toggle id(number) //go run ./ -toggle 2
	flag.BoolVar(&Comflags.List ,"list", false, "list all todos") //go run ./ -list


	flag.Parse()
	return &Comflags
}

func (cf *Commands) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Delete != -1:
		todos.delete(cf.Delete)

	default:
		fmt.Println("Invalid command")
	}
}