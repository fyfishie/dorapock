/*
 * @Author: fyfishie
 * @Date: 2023-03-27:18
 * @LastEditors: fyfishie
 * @LastEditTime: 2023-03-27:18
 * @Description: :)
 * @email: fyfishie@outlook.com
 */
// 每个Generator管理自己的一套id,这些id互相不重复
package idgen

import "math/rand"

type Generator struct {
	idReqChan   chan newIdReq
	recycleChan chan recycleReq
	usedIdList  map[int]bool
}

type recycleReq struct {
	done chan bool
	id   int
}
type newIdReq struct {
	idOut chan int
}

func (g *Generator) Run() {
	go func() {
		for {
			select {
			case req := <-g.idReqChan:
				req.idOut <- g.newID()
			case req := <-g.recycleChan:
				g.recycle(req.id)
				req.done <- true
			}
		}
	}()
}

func (g *Generator) newID() int {
	id := rand.Int()
	for {
		if _, exist := g.usedIdList[id]; exist {
			id++
		} else {
			g.usedIdList[id] = true
			return id
		}
	}
}

func (g *Generator) recycle(id int) {
	delete(g.usedIdList, id)
}

func NewIdGen() Generator {
	g := Generator{
		idReqChan:   make(chan newIdReq),
		recycleChan: make(chan recycleReq),
		usedIdList:  map[int]bool{},
	}
	return g
}

// 供用户使用的接口，并且提供线程安全
func (g *Generator) NewID() int {
	idOut := make(chan int)
	req := newIdReq{
		idOut: idOut,
	}
	g.idReqChan <- req
	return <-req.idOut
}

// 供用户使用的接口，并且提供线程安全
func (g *Generator) RecycleID(id int) {
	done := make(chan bool)
	req := recycleReq{
		id:   id,
		done: done,
	}
	g.recycleChan <- req
	<-done
}
