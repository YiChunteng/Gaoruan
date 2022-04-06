package LK

const (
	SUCCESS = 0
	FAILURE = -1
)

/*
	LinkTable Node Type
*/
type LinkTableNode struct {
	pNext *LinkTableNode
	Data int
}

/*
	LinkTable Type
*/
type LinkTable struct {
	Head *LinkTableNode
	Tail *LinkTableNode
	SumOfNode int
}

/*
	Create a LinkTable
*/
func CreateLinkTable() *LinkTable {
	linkTable := &LinkTable{}
	if linkTable == nil {
		return nil
	}
	return linkTable
}

/*
	Delete a LinkTable
*/
func (lk *LinkTable) DeleteLinkTable() int {
	if lk == nil {
		return FAILURE
	}
	lk.Head = nil
	lk.Tail = nil
	lk.SumOfNode = 0
	return SUCCESS
}

/*
	Add a LinkTableNode to LinkTable
*/
func (lk *LinkTable) AddLinkTableNode(pNode *LinkTableNode) int {
	if lk == nil || pNode == nil {
		return FAILURE
	}
	pNode.pNext = nil
	if lk.Head == nil {
		lk.Head = pNode
	}
	if lk.Tail == nil {
		lk.Tail = pNode
	} else {
		lk.Tail.pNext = pNode
		lk.Tail = pNode
	}
	lk.SumOfNode++
	return SUCCESS
}

/*
	Delete a LinkTableNode from LinkTable
*/
func (lk *LinkTable) DelLinkTableNode(pNode *LinkTableNode) int {
	if lk == nil || pNode == nil {
		return FAILURE
	}
	if lk.Head == pNode {
		lk.Head = lk.Head.pNext
		lk.SumOfNode--
		if lk.SumOfNode == 0 {
			lk.Tail = nil
		}
		return SUCCESS
	}
	pTempNode := lk.Head
	for pTempNode != nil {
		if pTempNode.pNext == pNode {
			pTempNode.pNext = pTempNode.pNext.pNext
			lk.SumOfNode--
			if lk.SumOfNode == 0 {
				lk.Tail = nil
			}
			return SUCCESS
		}
		pTempNode = pTempNode.pNext
	}
	return FAILURE
}

/*
  get LinkTableHead
*/
func (lk *LinkTable) GetLinkTableHead() *LinkTableNode {
	if lk == nil {
		return nil
	}
	return lk.Head
}

/*
	get next LinkTableNode
*/
func (lk *LinkTable) GetNextLinkTableNode(pNode *LinkTableNode) *LinkTableNode {
	if lk == nil || pNode == nil {
		return nil
	}
	pTempNode := lk.Head
	for pTempNode != nil {
		if pTempNode == pNode {
			return pTempNode.pNext
		}
		pTempNode = pTempNode.pNext
	}
	return nil
}




















