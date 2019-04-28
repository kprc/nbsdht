package nbsid

import (
	"sync"
	"time"
	"fmt"
	"strconv"
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

	nid:=&NodeID{strid,[]byte(strid)}
	//nid.id = "localid"

	return nid
}

func GetLocalId() INodeID  {
	if localId == nil {
		glock.Lock()
		if localId == nil{
			ts:=strconv.FormatInt(time.Now().UnixNano(),16)
			localId = NewID("nbsid"+ts)
			fmt.Println(localId)
		}
		glock.Unlock()
	}

	return localId
}

