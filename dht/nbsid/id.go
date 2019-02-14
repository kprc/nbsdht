package nbsid

import (
	"sync"
	"time"
)

type INodeID interface {
	String() string
	Bytes() []byte
}


type NodeID struct {
	id string
	bid []byte
}

var (
	localId INodeID
	glock sync.Mutex
)


func (id *NodeID)String() string  {
	return id.id
}

func (id *NodeID)Bytes() []byte {
	if id.bid==nil {
		//id.bid = make([]byte,)
	}

	return id.bid
}

func NewID(strid string)INodeID {

	nid:=&NodeID{strid,nil}
	//nid.id = "localid"

	return nid
}

func GetLocalId() INodeID  {
	if localId == nil {
		glock.Lock()
		if localId == nil{
			ts := string(time.Now().UnixNano())
			localId = NewID("nbsid"+ts)
		}
		glock.Unlock()
	}

	return localId
}

