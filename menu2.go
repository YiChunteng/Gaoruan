package main

import (
	"fmt"
	"os"
)

type help interface {
	Help()
}

type DataNode struct {
	cmd string
	desc string
	handler func()
	next *DataNode
}

var head = []*DataNode{
	{"help", "this is a help cmd", nil, nil},
	{"version", "menu version 1.0", nil, nil},
	{"quit", "quit from menu", Quit, nil},
}

func init() {
	head[0].handler = Help
	head[0].next = head[1]
	head[1].next = head[2]
}

func Help() {
	fmt.Println("menu list:")
	for _, h := range head {
		fmt.Printf("%s - %s\n", h.cmd, h.desc)
	}
}

func Quit() {
	os.Exit(0)
}

func main() {
	fmt.Println("this is a menu program:")
	for {
		fmt.Println("please input your cmd:")
		var input string
		fmt.Scanln(&input)
		p := head[0]
		for p != nil {
			if p.cmd == input {
				fmt.Printf("%s - %s\n", p.cmd, p.desc)
				if p.handler != nil {
					p.handler()
				}
				break
			}
			p = p.next
		}
		if p == nil {
			fmt.Println("this is a wrong cmd")
		}
		fmt.Println()
	}
}
