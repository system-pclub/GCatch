package figure2

import "sync"

type T int

type Block struct {
	val T
	_lock *sync.Mutex
}

type BlockChain struct {
	best_block Block
	best_ancient_block Block
}

func (bc BlockChain) chain_info() {
	bc.best_block._lock.Lock()
	bc.best_ancient_block._lock.Lock()
	//Work
	bc.best_block._lock.Unlock()
	bc.best_ancient_block._lock.Unlock()
}

func newBlock1(value T) Block {
	return Block{value, &sync.Mutex{}}
}

func newBlock2(value T) Block {
	return Block{value, &sync.Mutex{}}
}

func (bc BlockChain) commit() {
	bc.best_ancient_block._lock.Lock()
	bc.best_block._lock.Lock()
	//Work
	bc.best_ancient_block._lock.Unlock()
	bc.best_block._lock.Unlock()
}

func main() {
	bc := BlockChain{newBlock1(5), newBlock2(8)}
	go bc.chain_info()
	bc.commit()
}