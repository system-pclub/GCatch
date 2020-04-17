package instinfo

import (
	"fmt"
	"github.com/system-pclub/gochecker/config"
	"github.com/system-pclub/gochecker/tools/go/ssa"
	"github.com/system-pclub/gochecker/util"
	"go/types"
	"strings"
)

/*
func GetMutexName1(inputInst ssa.Instruction) string {
	instLocation := (config.Prog.Fset).Position(inputInst.Pos())
	strFileName := instLocation.Filename
	numLine := instLocation.Line

	if numLine <  1 {
		return ""
	}

	strCodeLine, err := util.ReadFileLine(strFileName, numLine)

	if err != nil {
		fmt.Println("Error: during read file:", strFileName,"\tline:", numLine,"\tfor inst:", inputInst)
		return ""
	}

	if numCommentStart := strings.Index(strCodeLine,"//"); numCommentStart > -1 {
		strCodeLine = strCodeLine[:numCommentStart]
	}

	strMutexName := ""

	if numIndexLock := strings.Index(strCodeLine,"Lock"); numIndexLock > -1 {
		strMutexName = strCodeLine[:numIndexLock]
	} else if numIndexUnlock := strings.Index(strCodeLine,"Unlock"); numIndexUnlock > -1 {
		strMutexName = strCodeLine[:numIndexUnlock]
	} else if numIndexRUnlock := strings.Index(strCodeLine,"RUnlock"); numIndexRUnlock > -1 {
		strMutexName = strCodeLine[:numIndexRUnlock]
	} else if numIndexLastDot := strings.LastIndex(strCodeLine,"."); numIndexLastDot > -1 {
		strMutexName = strCodeLine[:numIndexLastDot]
	}

	if strings.Contains(strMutexName,"defer") {
		splits := strings.Split(strMutexName," ")
		strMutexName = splits[len(splits) - 1]
	}
	strMutexName = strings.TrimSpace(strMutexName)

	return strMutexName
}

 */


func GetMutexName(inputInst ssa.Instruction) string {
	instLoc := (config.Prog.Fset).Position(inputInst.Pos())
	strFileName := instLoc.Filename
	numLine := instLoc.Line
	if numLine < 1 {
		return ""
	}

	strCodeLine, err := util.ReadFileLine(strFileName, numLine)

	if err != nil {
		fmt.Println("Error: during read file:", strFileName,"\tline:", numLine,"\tfor inst:", inputInst)
		return ""
	}

	if numCommentStart := strings.Index(strCodeLine,"//"); numCommentStart > -1 {
		strCodeLine = strCodeLine[:numCommentStart]
	}

	numDotIndex := -1

	numLockIndex := strings.LastIndex(strCodeLine,".Lock")
	numRLockIndex := strings.LastIndex(strCodeLine,".RLock")
	numUnlockIndex := strings.LastIndex(strCodeLine,".Unlock")
	numRUnlockIndex := strings.LastIndex(strCodeLine,".RUnlock")
	switch {
	case numLockIndex > 0:
		numDotIndex = numLockIndex
	case numRLockIndex > 0:
		numDotIndex = numRLockIndex
	case numUnlockIndex > 0:
		numDotIndex = numUnlockIndex
	case numRUnlockIndex > 0:
		numDotIndex = numRUnlockIndex
	}
	if numDotIndex < 1 {
		//fmt.Println("Error: calculating last dot. str_same_line:",str_same_line,"\tfor inst:",inst)
		return ""
	}

	strMutexName := strCodeLine[:numDotIndex]
	splits := strings.Split(strMutexName," ")
	strMutexName = splits[len(splits) - 1]
	strMutexName = strings.TrimSpace(strMutexName)
	strMutexName = strings.ReplaceAll(strMutexName,"case ","")
	strMutexName = strings.ReplaceAll(strMutexName,"\t","")
	strMutexName = strings.TrimSpace(strMutexName)


	return strMutexName
}


func IsMutexMake(inputInst ssa.Instruction) bool {
	instAlloc, ok := inputInst.(*ssa.Alloc)

	if ok {
		typeInst := instAlloc.Type().Underlying().(*types.Pointer).Elem()

		if typeInst.String() == "sync.Mutex" {
			return true
		}
	}

	return false
}

func IsMutexLock(inputInst ssa.Instruction) bool {
	var fnCall * ssa.CallCommon
	instCall, ok := inputInst.(*ssa.Call)

	if ok {
		fnCall = instCall.Common()
	}

	instDefer, ok := inputInst.(*ssa.Defer)
	if ok {
		fnCall = instDefer.Common()
	}

	if fnCall != nil {
		strFnName := GetCallName(fnCall)
		if strFnName == "(*sync.Mutex).Lock" {
			return true
		}
		if fnCall.IsInvoke() {
			strFnName1 := fnCall.Method.Name()
			if strings.ToLower(strFnName1) == "lock" {
				return true
			}

		}
	}

	return false
}

func IsMutexUnlock(inputInst ssa.Instruction) bool {
	var fnCall *ssa.CallCommon

	instCall, ok := inputInst.(*ssa.Call)
	if ok {
		fnCall = instCall.Common()
	}

	instDefer, ok := inputInst.(*ssa.Defer)

	if ok {
		fnCall = instDefer.Common()
	}

	if fnCall != nil {
		strFnName := GetCallName(fnCall)

		if strFnName == "(*sync.Mutex).Unlock" {
			return true
		}
		if fnCall.IsInvoke() == true {
			strFnName2 := fnCall.Method.Name()

			if strings.ToLower(strFnName2) == "unlock" {
				return true
			}

		}
	}

	return false
}

func IsMutexOperation(inputInst ssa.Instruction) bool {
	return IsMutexLock(inputInst) || IsMutexUnlock(inputInst)
}

