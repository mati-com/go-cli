package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

// Enlista las tareas del array con el icono
func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Complete {
			status = "✓"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Name)
	}
}

// añade una tarea nueva al array
func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       GetNextID(tasks),
		Name:     name,
		Complete: false,
	}
	return append(tasks, newTask)
}

// busca la tarea con id y lo marca como completo
func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = true
			break
		}
	}
	return tasks
}

// elimina una tarea con un ID
func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

// obtiene la siguiente ID habilitada, para la siguiente tarea
func GetNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}

// guarda todas las tareas en un json
func SaveTasks(file *os.File, tasks []Task) {
	//convierte a .json
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	//mueve el "pointer" al comienzo
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	//limpiar el archivo seleccionado o elimina todo
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	//escribe el archivo
	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	//asegura que el contenido fue escrito
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
