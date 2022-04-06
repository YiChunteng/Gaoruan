package main

import (
	"fmt"
	"homeworks/src/gaoruan/lab6"
)

//type Node struct {
//	node *LK.LinkTableNode
//	data int
//}

func main() {
	pLinkTable := LK.CreateLinkTable()
	if pLinkTable == nil {
		fmt.Println("Create LinkTable Error")
		return
	}
	for i := 0; i< 10; i++ {
		tNode := &LK.LinkTableNode{}
		tNode.Data = i
		if pLinkTable.AddLinkTableNode(tNode) == -1 {
			fmt.Println("ADD Error")
			return
		}
	}
	pTempNode  := Search(pLinkTable)
	fmt.Printf("%d\n", pTempNode.Data)
	fmt.Println("DelLinkTableNode")
	pLinkTable.DelLinkTableNode(pTempNode)

	if pLinkTable.DeleteLinkTable() == 0 {
		fmt.Println("delete success")
	}
}

func Search(t *LK.LinkTable) *LK.LinkTableNode {
	fmt.Println("Search GetLinkTableHead")
	p := t.GetLinkTableHead()
	for p != nil {
		if p.Data == 5 {
			return p
		}
		fmt.Println("GetNextLinkTableNode")
		p = t.GetNextLinkTableNode(p)
	}
	return nil
}
