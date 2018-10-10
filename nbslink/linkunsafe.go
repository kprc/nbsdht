package nbslink

type LinkNode struct {
	prev,next *LinkNode
	Value interface{}
}


func NewLink(v interface{}) *LinkNode{
	return &LinkNode{Value:v}
}

func (ln *LinkNode) Init() *LinkNode {
	ln.next = ln
	ln.prev = ln

	return ln
}

//insert into after ln
func (ln *LinkNode) Add(node *LinkNode) {
	node.next = ln.next
	ln.next.prev = node
	node.prev = ln
	ln.next = node
}
//insert into befor ln
func (ln *LinkNode) Insert(node *LinkNode){
	node.next = ln
	node.prev = ln.prev
	ln.prev.next = node
	ln.prev = node
}

func  (ln *LinkNode)Remove() {
	ln.prev.next = ln.next
	ln.next.prev = ln.prev
}

func (ln *LinkNode) Prev() *LinkNode {
	prev := ln.prev

	return prev
}

func (ln *LinkNode) Next() *LinkNode {
	nxt := ln.next

	return nxt
}