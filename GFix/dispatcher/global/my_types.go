package global

import (
	"go/token"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
)

type Struct struct {
	Name string
	Field map[string]string
}

type Inst_value struct {
	Inst ssa.Instruction
	Value ssa.Value
	Comment string
}

type Temp_struct struct {
	Name string
	Field string
}

type Sync_struct struct {
	Type string
	ContainSync bool
	Field map[string]string
	Potential bool
}

type Parent_path struct {
	Parent string
	Children []string
	Total_num_lock int
	Total_num_send int
}

type Operation struct {
	Position token.Position
	Instruction ssa.Instruction
	Type string
}

type C3_struct struct {
	Name string
	Field map[string](map[ssa.Instruction][]string) //map[field_name](map[inst_used_field][]str_alive_mutexs)
}

type Path_stat struct {
	Path string
	Self_num_lock int
	Self_num_send int
}
