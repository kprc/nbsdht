package nbslink

import (
	"sync"
	"nbserr"
)

type HashBucket struct {
	Node *LinkNode
	M sync.RWMutex
	Cnt int
}

type HLNode interface {
	Hash() uint64
	Equals(node interface{}) bool
}

type HashLink struct {
	Hbs []HashBucket
	HashMax uint64
}


func NewHashLink(max uint) *HashLink{
	if max<1 || max>31{
		return nil
	}
	return &HashLink{Hbs:make([]HashBucket,1 << max),HashMax:1<<max}
}


func (hl *HashLink)Add(node interface{}) error{

	hcnt := node.(HLNode).Hash()

	if hcnt > hl.HashMax-1 {
		return nbserr.NbsErr{"Hash Code calculate error ",nbserr.ERROR_DEFAULT}
	}

	hb := &hl.Hbs[hcnt]

	ln := NewLink(node)

	hb.M.Lock()
	defer hb.M.Unlock()
	if hb.Node == nil {
		ln.Init()
		hb.Node = ln
	}else {
		hb.Node.Add(ln)
	}
	hb.Cnt ++

	return nil

}

func (hl *HashLink) AddNode(node *LinkNode) error{
	hln := node.Value.(HLNode)
	hcnt := hln.Hash()

	if hcnt > hl.HashMax-1 {
		return nbserr.NbsErr{"Hash Code calculate error ",nbserr.ERROR_DEFAULT}
	}
	hb := &hl.Hbs[hcnt]
	hb.M.Lock()
	defer hb.M.Unlock()
	if hb.Node == nil {
		node.Init()
		hb.Node = node
	}else {
		hb.Node.Add(node)
	}
	hb.Cnt ++
	return nil
}

func (hl *HashLink)RemoveNode(node *LinkNode) error {
	hln := node.Value.(HLNode)
	hcnt := hln.Hash()
	if hcnt > hl.HashMax-1 {
		return nbserr.NbsErr{"Hash Code calculate error ",nbserr.ERROR_DEFAULT}
	}
	hb := &hl.Hbs[hcnt]
	hb.M.Lock()
	defer hb.M.Unlock()

	if hb.Node == node {
		nxt:=hb.Node.Next()
		if nxt == hb.Node {
			hb.Node = nil
		}else {
			hb.Node = hb.Node.Next()
		}
	}else {
		node.Remove()
	}
	hb.Cnt --

	return nil
}

func (hl *HashLink)Remove(node interface{}) error{
	hcnt := node.(HLNode).Hash()

	if hcnt > hl.HashMax-1 {
		return nbserr.NbsErr{"Hash Code calculate error ",nbserr.ERROR_DEFAULT}
	}
	hb := &hl.Hbs[hcnt]
	hb.M.Lock()
	defer hb.M.Unlock()
	hn := hb.Node
	if hn == nil {
		return nil
	}

	for {
		if hn.Value.(HLNode).Equals(node){
			if hb.Node == hn {

				nxt:=hb.Node.Next()

				if nxt == hb.Node {

					hb.Node = nil
				}else {
					hb.Node = hb.Node.Next()
				}
			}else {
				hn.Remove()
			}
			break
		}
		hn = hn.Next()
		if hn == hb.Node {
			break
		}
	}
	hb.Cnt--

	return nil
}

func (hl *HashLink)Find(node interface{}) (HLNode,error){
	hcnt:=node.(HLNode).Hash()

	if hcnt > hl.HashMax-1 {
		return nil,nbserr.NbsErr{"Hash Code calculate error ",nbserr.ERROR_DEFAULT}
	}
	hb := &hl.Hbs[hcnt]
	hb.M.RLock()
	defer hb.M.RUnlock()

	hn:=hb.Node
	if hn == nil {
		return nil,nil
	}
	for {
		if hn.Value.(HLNode).Equals(node){
			return hn.Value.(HLNode),nil
		}
		hn = hn.Next()
		if hn == hb.Node {
			break
		}
	}

	return nil,nil

}

