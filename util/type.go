package util

import (
	"fmt"
	"github.com/system-pclub/gochecker/config"
	"go/types"
	"reflect"
	"strings"
)

var mapStruct2Pointer map[* types.Struct] * types.Pointer

func GetStructPointerMapping()  {
	mapStruct2Pointer = make(map[* types.Struct] * types.Pointer)
	for _, T := range config.Prog.RuntimeTypes() {
		if pPointer, ok := T.(* types.Pointer); ok {
			if pNamed, ok := pPointer.Elem().(* types.Named); ok {
				if pStruct, ok := pNamed.Underlying().(* types.Struct); ok {
					mapStruct2Pointer[pStruct] = pPointer
				}
			}
		}
	}
	//fmt.Println("mapping size", len(mapStruct2Pointer))
}


/*
func GetFieldTypes(t  types.Type, m map[types.Type] bool) {
	if pPointer, ok := t.(* types.Pointer); ok {
		GetFieldTypes(pPointer.Elem(), m)
	} else if pNamed, ok := t.(* types.Named); ok {
		GetFieldTypes(pNamed.Underlying(), m)
	} else if pStruct, ok := t.(* types.Struct); ok {
		m[pStruct] = true
		for i := 0; i < pStruct.NumFields(); i ++ {
			GetFieldTypes(pStruct.Field(i).Type(), m)
		}
	} else if pInterface, ok := t.(* types.Interface); ok {
		m[pInterface] = true
		for i := 0; i < pInterface.NumEmbeddeds(); i ++ {
			GetFieldTypes(pInterface.EmbeddedType(i), m)
		}

	} else if pMap, ok := t.(*types.Map); ok {
		m[pMap] = true
	} else if pBasic, ok := t.(*types.Basic); ok {
		m[pBasic] = true
	} else {
		fmt.Println("not handle", reflect.TypeOf(t), t)
		panic("in GetFieldTypes()")
	}
}

func PrintTypes(m map[types.Type] bool) {
	for t, _ := range m {
		fmt.Println(reflect.TypeOf(t), t)
	}
}
*/

func GetTypeMethods(t types.Type, m map[string] bool, mVisited map[types.Type] bool)  {

	if _, ok := mVisited[t]; ok {
		return
	}

	mVisited[t] = true

	if pPointer, ok := t.(* types.Pointer); ok {
		GetTypeMethods(pPointer.Elem(), m, mVisited)
	} else if pNamed, ok := t.(* types.Named); ok {
		GetTypeMethods(pNamed.Underlying(), m, mVisited)
	} else if pStruct, ok := t.(* types.Struct); ok {
		if pPointer, ok := mapStruct2Pointer[pStruct]; ok {
			mset := config.Prog.MethodSets.MethodSet(pPointer)
			if mset != nil {
				for i := 0; i < mset.Len(); i++ {
					if pFunc, ok := mset.At(i).Obj().(*types.Func); ok {
						m[pFunc.FullName()] = true
					}
				}
			}
		}
		for i := 0; i < pStruct.NumFields(); i ++ {
			GetTypeMethods(pStruct.Field(i).Type(), m, mVisited)
		}
	} else if pInterface, ok := t.(* types.Interface); ok {
		for i := 0; i < pInterface.NumMethods(); i ++ {
			m[pInterface.Method(i).FullName()] = true
		}
		for i := 0; i < pInterface.NumEmbeddeds(); i ++ {
			GetTypeMethods(pInterface.EmbeddedType(i), m, mVisited)
		}

	} else if _, ok := t.(*types.Map); ok {
	} else if _, ok := t.(*types.Basic); ok {
	} else if _, ok := t.(*types.Slice); ok {
	} else if _, ok := t.(*types.Chan); ok {
	} else if _, ok := t.(*types.Array); ok {
	} else if _, ok := t.(*types.Signature); ok {
	}else {
		fmt.Println("not handle", reflect.TypeOf(t), t)
		panic("in GetTypeMethods()")
	}
}

func DecoupleTypeMethods(m map[string] bool) map[string] map[string] bool {
	mapResult := make(map[string] map[string] bool)
	for strFunName, _ := range m {
		var strStructName string
		var strFName string
		if strings.LastIndex(strFunName, ".") < 0 {
			strStructName = ""
			strFName = strFunName
		} else {
			strStructName = strFunName[:strings.LastIndex(strFunName, ".")]
			strFName = strFunName[strings.LastIndex(strFunName, ".") + 1:]
		}

		if _, ok := mapResult[strStructName]; !ok {
			mapResult[strStructName] = make(map[string] bool)
		}

		mapResult[strStructName][strFName] = true
	}

	return mapResult
}

func PrintTypeMethods(m map[string] map[string] bool) {
	for sPackage, mapMethods := range m {
		fmt.Println(sPackage)
		for sFuncName, _ := range mapMethods {
			fmt.Println(sFuncName)
		}
	}

}





