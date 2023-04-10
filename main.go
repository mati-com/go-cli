package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/mati-com/golang-crud-cli/tasks"
)

func main() {

	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []tasks.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []tasks.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
	}

}

func printUsage() {
	fmt.Println("Uso: go-crud-cli [list|add|complete|delete]")
}
