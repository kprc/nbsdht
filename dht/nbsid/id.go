package nbsid


type INodeID interface {
	String() string
	Bytes() []byte
}


type NodeID struct {
	id string
	bid []byte
}

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
	return &NodeID{strid,nil}
}


