package dhttable

import (
	"github.com/kprc/nbsdht/dht/nbsid"
	"sync"
	"github.com/kprc/nbsdht/nbslink"
	"github.com/NBSChain/go-nbs/thirdParty/account"
	"github.com/NBSChain/go-nbs/utils"
	"math/big"
	"sync/atomic"
)

type DHTHashTable struct {
	selfID nbsid.INodeID      //self id
	rwlock sync.RWMutex       //lock
	maxLatency uint            //
	buckets *nbslink.HashLink
	candBuckets *nbslink.HashLink    //candidate node for update
	bucketSize uint
	totalCount uint32
	candTotalCount uint32
}

type NBSNode interface {
	nbslink.HLNode
	//Less(node NBSNode) bool   //for sort
	Duration() uint      //now time divide into the last access time
	Update(node NBSNode)  //update connections, access time, etc.
	Clone() NBSNode    //clone a node
	Replace()   //replace value
	Distance(node NBSNode) big.Int  //caculate and store
	StoreDistance(bgi big.Int)
}


var (
	maxLatency uint = 500    //500ms
 	bucketSize uint = 160    //160 bits
 	instance *DHTHashTable
	pingChannel chan NBSNode = make(chan NBSNode,COUNT_PING_CHANNEL)
 	once sync.Once
	logger          = utils.GetLogInstance()
)

func NewDHTHashTable() *DHTHashTable {
	once.Do(func() {
		dht := newDHTHashTable()

		instance = dht
	})

	return instance

}

func newDHTHashTable() *DHTHashTable {
	selfid := nbsid.NewID(account.GetAccountInstance().GetPeerID())

	hlink := nbslink.NewHashLink(bucketSize)

	return &DHTHashTable{selfID:selfid,maxLatency:maxLatency,buckets:hlink,bucketSize:bucketSize}
}

func (dht *DHTHashTable)Store(node NBSNode)  {
	pingNode := dht.add(node)

	if pingNode != nil {
		//put to candbuckets
		dht.putCandNode(node)
		dht.Ping(pingNode)
	}

}

func (dht *DHTHashTable)Ping(node NBSNode)  {
	pingChannel <- node    //if channel full, the process will block
}



func removeOldNode(hb *nbslink.HashBucket){
	hn := hb.Node
	max := hn

	for  {
		hn = hn.Next()
		if hn == hb.Node {
			break
		}
		if hn.Value.(NBSNode).Duration() > max.Value.(NBSNode).Duration()  {
			max = hn
		}

	}

	if max == hb.Node {
		nxt:=hb.Node.Next()
		if nxt == hb.Node {
			hb.Node = nil
		}else {
			hb.Node = hb.Node.Next()
			max.Remove()
		}
	}else {
		max.Remove()
	}
	hb.Cnt --
}

func (dht *DHTHashTable)add(node NBSNode) NBSNode{

	nhash := node.Hash()

	rootlink := dht.buckets

	hb := rootlink.Hbs[nhash]

	hb.M.Lock()
	defer hb.M.Unlock()

	cnt,n := update(node,&hb)

	atomic.AddUint32(&dht.totalCount,uint32(cnt))

	return n
}

func (dht *DHTHashTable)FindNode(node NBSNode) ([]NBSNode,error){
	nhash := node.Hash()

	rootlink := dht.buckets

	hb := rootlink.Hbs[nhash]

	hb.M.Lock()
	defer hb.M.Unlock()

	ns := make([]NBSNode,0)

	//ns = append(ns, )

	return ns,nil
}

func (dht *DHTHashTable)FindValue(node NBSNode) (v interface{}, nodes []NBSNode,err error) {

	return nil,nil,nil
}


func update(node interface{},  hb *nbslink.HashBucket) (uint,NBSNode)  {
	if hb.Node == nil {
		ln := nbslink.NewLink(node)
		ln.Init()
		hb.Node = ln
		hb.Cnt ++
		return 1,nil
	}

	hn := hb.Node

	for  {

		if hn.Value.(NBSNode).Equals(node) {
			//update time
			hn.Value.(NBSNode).Update(node.(NBSNode))
			return 0,nil
		}
		hn = hn.Next()

		if hn == hb.Node {
			break
		}

	}

	if hb.Cnt >= COUNT_PER_BUCKET {
		return 0,findLastNode(hb)
	}else {
		ln := nbslink.NewLink(node)
		hb.Node.Add(ln)
		hb.Cnt ++
	}

	return 1,nil

}

func findLastNode(hb *nbslink.HashBucket) NBSNode {
	hn := hb.Node
	max := hn

	for  {
		hn = hn.Next()
		if hn == hb.Node {
			break
		}
		if hn.Value.(NBSNode).Duration() > max.Value.(NBSNode).Duration()  {
			max = hn
		}

	}

	return max.Value.(NBSNode).Clone()

}


func (dht *DHTHashTable)putCandNode(node NBSNode){
	nhash := node.Hash()

	rootlink := dht.candBuckets

	hb := rootlink.Hbs[nhash]

	hb.M.Lock()
	defer hb.M.Unlock()

	cnt := updateCandNode(node,&hb)

	atomic.AddUint32(&dht.candTotalCount,uint32(cnt))

}

func updateCandNode(node interface{},hb *nbslink.HashBucket) uint {

	var cnt uint = 0

	if hb.Node == nil {
		ln := nbslink.NewLink(node)
		ln.Init()
		hb.Node = ln
		hb.Cnt ++
		return 1
	}

	hn := hb.Node

	for  {

		if hn.Value.(NBSNode).Equals(node) {
			//update time
			hn.Value.(NBSNode).Update(node.(NBSNode))
			return 0
		}
		hn = hn.Next()

		if hn == hb.Node {
			break
		}

	}

	if hb.Cnt >= COUNT_PER_CANDBUCKET {
		//find last and remove
		removeOldNode(hb)
	}else {
		hb.Cnt ++
		cnt = 1
	}
	ln := nbslink.NewLink(node)
	hb.Node.Add(ln)

	return cnt
}

func send(node NBSNode)  {
	return
}

func PingChannelProcess(){
	for node := range pingChannel {
		send(node)
	}
}


