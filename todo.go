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

	_, sortedKey:= getAllTasks()
	highestTaskKey := sortedKey[len(sortedKey)-1] + 1
	writeString :=strconv.Itoa(highestTaskKey) + "|" +task

	f, err := os.OpenFile(dbFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	_,err1 :=w.WriteString(writeString)
	if err1 != nil {
		panic(err1)
	}
	w.Flush()

}

func showTask() {
	tasks, sortedKey:= getAllTasks()
	printTasks(tasks,sortedKey)
}

func printTasks(tasks map[int]string, sortedKey []int){
	for _,key:= range sortedKey {
		fmt.Println(strconv.Itoa(key)+"> " + tasks[key])
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
		id := toInt(s[0])
		sortedKey = append(sortedKey,id)
		tasks[id] = s[1]
	}
	sort.Ints(sortedKey)
	return
}

func reOrderTask(oldNumber int, newNumber int) {
	fmt.Println("REORDER TASKS : file : " + dbFile + ". Old :" + strconv.Itoa(oldNumber) + ".New :" + strconv.Itoa(newNumber))
	_, sortedKey:= getAllTasks()
	if (oldNumber >= len(sortedKey) || newNumber >= len(sortedKey)){
		fmt.Println("Wrong order number. Max order number is "+strconv.Itoa(len(sortedKey)))
	}

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
		reOrderTask(toInt(args[1]), toInt(args[2]))
	}

}
