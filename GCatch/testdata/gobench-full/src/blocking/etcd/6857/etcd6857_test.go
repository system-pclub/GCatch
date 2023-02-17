/*
 * Project: etcd
 * Issue or PR  : https://github.com/etcd-io/etcd/pull/6857
 * Buggy version: 7c8f13aed7fe251e7066ed6fc1a090699c2cae0e
 * fix commit-id: 7afc490c95789c408fbc256d8e790273d331c984
 * Flaky: 19/100
 */
package etcd6857

import (
	"testing"
)

type Status struct{}

type node struct {
	status chan chan Status
	stop   chan struct{}
	done   chan struct{}
}

func (n *node) Status() Status {
	c := make(chan Status)
	n.status <- c
	return <-c
}

func (n *node) run() {
	for {
		select {
		case c := <-n.status:
			c <- Status{}
		case <-n.stop:
			close(n.done)
			return
		}
	}
}

func (n *node) Stop() {
	select {
	case n.stop <- struct{}{}:
	case <-n.done:
		return
	}
	<-n.done
}

func NewNode() *node {
	return &node{
		status: make(chan chan Status),
		stop:   make(chan struct{}),
		done:   make(chan struct{}),
	}
}

///
/// G1				G2				G3
/// n.run()
///									n.Stop()
///									n.stop<-
/// <-n.stop
///									<-n.done
/// close(n.done)
///	return
///									return
///					n.Status()
///					n.status<-
///----------------G2 leak-------------------
///

func TestEtcd6857(t *testing.T) {
	n := NewNode()
	go n.run()    // G1
	go n.Status() // G2
	go n.Stop()   // G3
}
