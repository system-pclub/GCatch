package figure2_translate
import "sync"

type BlockᐸDummyᐸᐳᐳ struct { val Dummyᐸᐳ; _lock *sync.Mutex };
type BlockChainᐸDummyᐸᐳᐳ struct { best_block BlockᐸDummyᐸᐳᐳ; best_ancient_block BlockᐸDummyᐸᐳᐳ };
type Dummyᐸᐳ struct {};
func (bc BlockChainᐸDummyᐸᐳᐳ) chain_infoᐸᐳ() Dummyᐸᐳ { bc.best_block._lock.Lock();
	bc.best_ancient_block._lock.Lock();
	bc.best_block._lock.Unlock();
	bc.best_ancient_block._lock.Unlock();
	return Dummyᐸᐳ{};
};
func (bc BlockChainᐸDummyᐸᐳᐳ) chain_info___Dummy__() Top { return bc;
};
func (d Dummyᐸᐳ) newBlock1ᐸDummyᐸᐳᐳ(value Dummyᐸᐳ) BlockᐸDummyᐸᐳᐳ { return BlockᐸDummyᐸᐳᐳ{value, &sync.Mutex{}};
};
func (d Dummyᐸᐳ) newBlock2ᐸDummyᐸᐳᐳ(value Dummyᐸᐳ) BlockᐸDummyᐸᐳᐳ { return BlockᐸDummyᐸᐳᐳ{value, &sync.Mutex{}};
};
func (d Dummyᐸᐳ) newBlock__β1_Any____β1_Block_β1_() Top { return d;
};
func (bc BlockChainᐸDummyᐸᐳᐳ) commitᐸᐳ() Dummyᐸᐳ { bc.best_ancient_block._lock.Lock();
	bc.best_block._lock.Lock();
	bc.best_ancient_block._lock.Unlock();
	bc.best_block._lock.Unlock();
	return Dummyᐸᐳ{};
};
func (bc BlockChainᐸDummyᐸᐳᐳ) commit___Dummy__() Top { return bc;
};
func (d Dummyᐸᐳ) Mainᐸᐳ() Dummyᐸᐳ { bc := BlockChainᐸDummyᐸᐳᐳ{Dummyᐸᐳ{}.newBlock1ᐸDummyᐸᐳᐳ(Dummyᐸᐳ{}), Dummyᐸᐳ{}.newBlock2ᐸDummyᐸᐳᐳ(Dummyᐸᐳ{})};
	go bc.chain_infoᐸᐳ();
	_ = bc.commitᐸᐳ();
	return Dummyᐸᐳ{};
};
func (d Dummyᐸᐳ) Main___Dummy__() Top { return d;
};
type Top interface {};
func main() { _ = Dummyᐸᐳ{}.Mainᐸᐳ() }