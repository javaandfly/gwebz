package manager

import (
	"sync/atomic"
)

type GoStatus uint8
type GoFunction func()

const (
	Createing GoStatus = iota + 1
	Running
	Waiting
	Stop
)

// 协程状态树
type StateManagerNode struct {
	parentNode   *StateManagerNode
	nextNode     []*StateManagerNode
	runningCount atomic.Int64
	status       GoStatus
	signChan     chan struct{}
	doFunc       GoFunction
}

func NewNode(parentNode *StateManagerNode, doFunc GoFunction) *StateManagerNode {
	return &StateManagerNode{
		nextNode:     make([]*StateManagerNode, 0),
		parentNode:   parentNode,
		runningCount: atomic.Int64{},
		status:       Createing,
		signChan:     make(chan struct{}),
		doFunc:       doFunc,
	}
}

func (node *StateManagerNode) Do() {
	defer func() {
		if node.parentNode != nil {
			node.parentNode.runningCount.Add(-1)
			if node.parentNode.runningCount.Load() == 0 {
				node.parentNode.signChan <- struct{}{}
			}
		}
	}()
	for _, next := range node.nextNode {
		go func(newNode *StateManagerNode) {
			newNode.Do()
		}(next)
	}
	if node.doFunc != nil {
		node.doFunc()
	}

	if node != nil && len(node.nextNode) != 0 {
		<-node.signChan
	}

}

func (node *StateManagerNode) RegisterNode(newNode ...*StateManagerNode) {

	node.nextNode = append(node.nextNode, newNode...)

	node.runningCount.Add(1)

}
