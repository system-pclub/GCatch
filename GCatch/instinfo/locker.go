package instinfo

import "github.com/system-pclub/GCatch/GCatch/tools/go/ssa"

// This file defines primitive Locker and its operations
// Locker has similar definitions to Channel in channel.go
// They are used by the BMOC checker, not traditional checkers
// Locker includes Mutex and RWMutex, and doesn't handle RLock and RUnlock of RWMutex

// Define Locker
type Locker struct{
	Name string
	Type string
	Locks []*LockOp
	Unlocks []*UnlockOp
	Pkg string // Don't use *ssa.Package here! It's not reliable
	Value ssa.Value

	Status string
}

func (l *Locker) AllOps() []LockerOp {
	vecOp := []LockerOp{}
	for _, lock := range l.Locks {
		vecOp = append(vecOp, lock)
	}
	for _, unlock := range l.Unlocks {
		vecOp = append(vecOp, unlock)
	}
	return vecOp
}

func (l *Locker) ModifyStatus(str string) {
	l.Status = str
}

// Define interface LockerOp and its two implementations
//		Define LockerOp
type LockerOp interface {
	Instr() ssa.Instruction
	Prim() *Locker
}

//		Define LockOp
type LockOp struct {
	Name    string
	Inst    ssa.Instruction
	IsRLock bool
	IsDefer bool

	Parent *Locker
}

func (l *LockOp) Instr() ssa.Instruction {
	return l.Inst
}

func (l *LockOp) Prim() *Locker {
	return l.Parent
}

//		Define UnlockOp
type UnlockOp struct {
	Name string
	Inst ssa.Instruction
	IsRUnlock bool
	IsDefer bool

	Parent *Locker
}

func (u *UnlockOp) Instr() ssa.Instruction {
	return u.Inst
}


func (u *UnlockOp) Prim() *Locker {
	return u.Parent
}

// A map from inst to its corresponding LockerOp
var MapInst2LockerOp map[ssa.Instruction]LockerOp

func ClearLockerOpMap() {
	MapInst2LockerOp = make(map[ssa.Instruction]LockerOp)
}

func init() {
	MapInst2LockerOp = make(map[ssa.Instruction]LockerOp)
}
