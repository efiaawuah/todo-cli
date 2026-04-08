package todo

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// creates a hidden .todocli folder for data storage in user's home directory
func getFilePath() string {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".todocli")
	os.MkdirAll(dir, 0755)
	return filepath.Join(dir, "todos.json")
}

// loads todos data from json file into a todos struct
func Load() (Todos, error) {
	path := getFilePath()
	file, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Todos{NEXT_ID: 0, TASKS: make(map[int]Task)}, nil
		}
		return Todos{NEXT_ID: 0, TASKS: make(map[int]Task)}, err
	}
	var todos Todos
	err = json.Unmarshal(file, &todos)
	if err != nil {
		panic(err)
	}
	return todos, nil
}

// puts todos data into json file
func Save(todos Todos) error {
	path := getFilePath()
	data, err := json.MarshalIndent(todos, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// adds new task to json file
func Add(task string) (error, int) {
	todos, err := Load()
	if err != nil {
		return err, 0
	}
	newTask := Task{
		ID:          todos.NEXT_ID,
		DESCRIPTION: task,
		DONE:        false,
	}
	todos.TASKS[newTask.ID] = newTask
	todos.NEXT_ID++
	err = Save(todos)
	if err != nil {
		return err, 0
	}
	return nil, newTask.ID
}

func MarkComplete(task_id int) error {
	return nil
}

func ListTasks(done *bool) []Task {
	return nil
}

func Delete(task_id int) error {
	return nil
}
