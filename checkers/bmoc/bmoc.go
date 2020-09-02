package bmoc

import "github.com/system-pclub/GCatch/analysis/pointer"

func Detect() {
	stPtrResult, vecStOpValue := pointer.AnalyzeAllSyncOp()
	vecChannel := pointer.WithdrawAllChan(stPtrResult, vecStOpValue)
	_ = vecChannel	// TODO: Withdraw all Lockers
}
