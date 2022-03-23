package main

import "fmt"

var cmds = map[string]string{
	"hello":   "hello",
	"help":    "this is a help cmd",
	"version": "menu version 1.0",
	"quit":    "quit!",
}

func main() {
	fmt.Println("this is a menu program:")
	for {
		fmt.Println("please input your cmd:")
		var input string
		fmt.Scanln(&input)
		if output, ok := cmds[input]; ok {
			fmt.Println(output)
		} else {
			fmt.Println("no cmd!")
		}
		if input == "quit" {
			return
		}
	}
}
