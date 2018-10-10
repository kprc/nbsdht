package nbslink

import (
	"sync"
)


type Element interface {
	//Large(e Element) bool
	Equals(e Element) bool
	//Less(e Element) bool
	//String() string
	//Value() interface{}
}

type entry struct {
	prev *entry
	next *entry
	e *Element
}

type DLinkRoot struct {
	m sync.RWMutex
	cnt int
	l *entry
}


type DLink interface {
	Add(e Element)
	Remove(e Element)
	IsIn(e Element) bool
	Count() int
	WalkAll(f func(e Element,args ...interface{}),args ...interface{})
	Prev(e Element, jump int) *Element
	Next(e Element, jump int) *Element
	Get(e Element) []*Element
	Put(e []*Element)
}

func (dlr *DLinkRoot)WalkAll(f func(e Element,args ...interface{}),args ...interface{}){
	dlr.m.RLock()
	defer dlr.m.RUnlock()
	lloop:=dlr.l
	if dlr.l == nil{
		return
	}

	for {
		f(*lloop.e,args)
		lloop = lloop.next
		if dlr.l == lloop {
			break
		}
	}
}

func NewDLink() *DLinkRoot{
	return &DLinkRoot{}
}


func (dlr *DLinkRoot) Add(e Element){
	l :=&entry{prev:nil,next:nil,e:&e}
	dlr.m.Lock()
	defer dlr.m.Unlock()
	if dlr == nil {
		dlr.cnt ++
		dlr.l = l
	}else {
		lhead := dlr.l
		l.next = lhead.next
		lhead.next.prev = l
		l.prev = lhead
		lhead.next = l
	}
}

func (dlr *DLinkRoot) Remove(e Element) {
	dlr.m.Lock()
	defer dlr.m.Unlock()
	lloop:=dlr.l
	if dlr.l == nil{
		return
	}
	for {
		if e.Equals(*lloop.e) {
			lloop.prev.next = lloop.next
			lloop.next.prev = lloop.prev
			if lloop == dlr.l {
				dlr.l = lloop.next
			}
			dlr.cnt --

			break
		}
		lloop = lloop.next
		if dlr.l == lloop {
			break
		}
	}

}

func (dlr *DLinkRoot) IsIn(e Element) bool{
	dlr.m.RLock()
	defer dlr.m.RUnlock()
	lloop:=dlr.l
	if dlr.l == nil{
		return false
	}
	for {
		if e.Equals(*lloop.e) {
			return true
		}
		lloop = lloop.next
		if dlr.l == lloop {
			break
		}
	}

	return false
}

func (dlr *DLinkRoot) Count() int{
	return dlr.cnt
}


func (dlr *DLinkRoot) Prev(e Element, jump int) *Element{
	dlr.m.RLock()
	defer dlr.m.RUnlock()

	if dlr.l == nil {
		return nil
	}
	var cnt int
	lloop := dlr.l
	for {
		if e.Equals(*lloop.e) {
			if cnt >= jump {
				return lloop.e
			}
			cnt ++
		}
		lloop = lloop.prev
		if dlr.l == lloop {
			break
		}
	}

	return nil

}


func (dlr *DLinkRoot) Next(e Element, jump int) *Element  {
	dlr.m.RLock()
	defer dlr.m.RUnlock()

	if dlr.l == nil {
		return nil
	}
	var cnt int
	lloop := dlr.l
	for {
		if e.Equals(*lloop.e) {
			if cnt >= jump {
				return lloop.e
			}
			cnt ++
		}
		lloop = lloop.next
		if dlr.l == lloop {
			break
		}
	}

	return nil
}

func (dlr *DLinkRoot) Get(e Element) []*Element {
	var es []*Element

	dlr.m.RLock()
	defer dlr.m.RUnlock()

	if dlr.l == nil {
		return nil
	}
	lloop := dlr.l
	for{
		if e.Equals(*lloop.e) {
			es = append(es,lloop.e)
		}
		lloop = lloop.next
		if dlr.l == lloop {
			break
		}
	}

	return es

}

func (dlr *DLinkRoot) Put(es []*Element) {
	if len(es)==0 {
		return
	}
	for _,e:=range es{
		dlr.Add(*e)
	}
}






