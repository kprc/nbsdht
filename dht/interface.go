package dht

import "github.com/kprc/nbsdht/dht/dhttable"

type NBSKDhter interface {
	Ping(node dhttable.NBSNode)
	Store(node dhttable.NBSNode)
	FindNode(node dhttable.NBSNode) ([]dhttable.NBSNode,error)
	FindValue(node dhttable.NBSNode) (v interface{}, nodes []dhttable.NBSNode,err error)
}
