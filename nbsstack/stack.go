package nbsstack

import "sync"

type PNode struct {
	next *PNode
	Value interface{}
}


type Stack struct {
	top *PNode
	cnt int
	M sync.Mutex
}


func (s *Stack)PopNode() *PNode{
	s.M.Lock()
	defer s.M.Unlock()

	if s.top == nil{
		return nil
	}

	node := s.top
	s.top = s.top.next
	s.cnt --

	return node
}

func (s *Stack)Pop() interface{}  {
	return s.PopNode().Value
}

func (s *Stack) Push(v interface{}) {

	node := &PNode{Value:v}

	s.PushNode(node)
}

func (s *Stack)PushNode(node *PNode)  {
	s.M.Lock()
	defer s.M.Unlock()

	if s.top == nil {
		s.top = node
	}else {
		top := s.top
		s.top = top
		s.top.next = top
	}
	s.cnt ++
}