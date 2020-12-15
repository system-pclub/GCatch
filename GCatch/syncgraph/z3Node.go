package syncgraph

func (s *ZNodeNbSend) hasPairWith(r *ZNodeNbRecv) bool {
	for _, pair := range s.Pairs {
		if pair.Recv == r {
			return true
		}
	}
	return false
}

func (s *ZNodeNbSend) findOtherThreadRecvs() (result []*ZNodeNbRecv) {
	for _, zthread := range s.ZGoroutine.Z3Sys.vecZGoroutines {
		if zthread == s.ZGoroutine {
			continue
		}
		for _, znode := range zthread.Nodes {
			zrecv, ok := znode.(*ZNodeNbRecv)
			if !ok {
				continue
			}
			if zrecv.boolIsBlocking {
				continue
			}
			if zrecv.PtrPNode.Node.(SyncOp).MapSyncOp()[s.PtrPNode.Node.(SyncOp)] {
				result = append(result, zrecv)
			}
		}
	}
	return
}

func (r *ZNodeNbRecv) hasPairWith(s *ZNodeNbSend) bool {
	for _, pair := range r.Pairs {
		if pair.Send == s {
			return true
		}
	}
	return false
}

func (r *ZNodeNbRecv) findOtherThreadSends() (result []*ZNodeNbSend) {
	for _, zthread := range r.ZGoroutine.Z3Sys.vecZGoroutines {
		if zthread == r.ZGoroutine {
			continue
		}
		for _, znode := range zthread.Nodes {
			zsend, ok := znode.(*ZNodeNbSend)
			if !ok {
				continue
			}
			if zsend.boolIsBlocking {
				continue
			}
			if zsend.PtrPNode.Node.(SyncOp).MapSyncOp()[r.PtrPNode.Node.(SyncOp)] {
				result = append(result, zsend)
			}
		}
	}
	return
}

func (r *ZNodeNbRecv) findAllThreadCloses() (result []*ZNodeClose) {
	for _, zthread := range r.ZGoroutine.Z3Sys.vecZGoroutines {
		for _, znode := range zthread.Nodes {
			zclose, ok := znode.(*ZNodeClose)
			if !ok {
				continue
			}
			if zclose.PtrPNode.Node.(SyncOp).MapSyncOp()[r.PtrPNode.Node.(SyncOp)] {
				result = append(result, zclose)
			}
		}
	}
	return
}

func (s *ZNodeBSend) findAllThreadOtherSendRecv() (result []ZNode) {
	for _, zthread := range s.ZGoroutine.Z3Sys.vecZGoroutines {
		for _, znode := range zthread.Nodes {
			if znode == s {
				continue
			}
			_, is_B_send := znode.(*ZNodeBSend)
			_, is_B_recv := znode.(*ZNodeBRecv)
			if is_B_send == false && is_B_recv == false {
				continue
			}
			if s.PtrPNode.Node.(SyncOp).MapSyncOp()[znode.PNode().Node.(SyncOp)] {
				result = append(result, znode)
			}
		}
	}
	return
}

func (r *ZNodeBRecv) findAllThreadOtherSendRecv() (result []ZNode) {
	for _, zthread := range r.ZGoroutine.Z3Sys.vecZGoroutines {
		for _, znode := range zthread.Nodes {
			if znode == r {
				continue
			}
			_, is_B_send := znode.(*ZNodeBSend)
			_, is_B_recv := znode.(*ZNodeBRecv)
			if is_B_send == false && is_B_recv == false {
				continue
			}
			if r.PtrPNode.Node.(SyncOp).MapSyncOp()[znode.PNode().Node.(SyncOp)] {
				result = append(result, znode)
			}
		}
	}
	return
}

func (r *ZNodeBRecv) findAllThreadCloses() (result []*ZNodeClose) {
	for _, zthread := range r.ZGoroutine.Z3Sys.vecZGoroutines {
		for _, znode := range zthread.Nodes {
			zclose, ok := znode.(*ZNodeClose)
			if !ok {
				continue
			}
			if zclose.PtrPNode.Node.(SyncOp).MapSyncOp()[r.PtrPNode.Node.(SyncOp)] {
				result = append(result, zclose)
			}
		}
	}
	return
}

