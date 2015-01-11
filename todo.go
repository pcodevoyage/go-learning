package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sort"
)

func getFilenameFromConfig() {

	f, err := os.Open("config")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	bf := bufio.NewReader(f)
	line, _, err := bf.ReadLine()

	if err != nil {
		panic(err)
	}

	//dbFile is a global variable.
	dbFile = string(line)

	//TODO : validate that dbFile is a valid file path.
}

func addTask(task string) {
	fmt.Println("ADD TASKS: file is " + dbFile + ". Task :" + task)
}

func showTask() {
	tasks, sortedKey:= getAllTasks()
	for _,key:= range sortedKey {
		fmt.Print(key)
		fmt.Println("> " + tasks[key])
	}
}

func getAllTasks() (tasks map[int]string, sortedKey []int) {
	f, err := os.Open(dbFile)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	tasks = make(map[int]string)

	bf := bufio.NewReader(f)
	for {
		line, isPrefix, err := bf.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil || isPrefix {
			panic(err)
		}

		s := strings.Split(string(line), "|")
		i, _ := strconv.ParseInt(s[0], 0, 64)
		id := int(i)
		sortedKey = append(sortedKey,id)
		tasks[id] = s[1]
	}
	sort.Ints(sortedKey)
	return
}

func reOrderTask(oldNumber string, newNumber string) {
	fmt.Println("REORDER TASKS : file : " + dbFile + ". Old :" + oldNumber + ".New :" + newNumber)
}

func removeTask(taskNum string) {
	fmt.Println("REMOVE TASKS : file : " + dbFile + " taskNum:" + taskNum)
}

var dbFile = ""

func main() {

	getFilenameFromConfig()
	// fmt.Println(dbFile)

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Arguments\n-a Add a task \n-s Show all tasks\n-r remove task")
		return
	}

	switch args[0] {
	case "-a":
		addTask(strings.Join(args[1:], " "))
	case "-s":
		showTask()
	case "-r":
		removeTask(args[1])
	case "-c":
		reOrderTask(args[1], args[2])
	}

}
