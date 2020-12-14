package output

import (
	"encoding/json"
	"fmt"
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"go/token"
	"strconv"
)


func Write_to_JSON(all_syn_op map[token.Pos]global.Operation) {
	op_outer_map := make(map[string]interface{})
	for op_Pos,op := range all_syn_op {
		op_inner_map := make(map[string]interface{})
		op_inner_map["Filename"] = op.Position.Filename
		op_inner_map["Line"] = op.Position.Line
		op_inner_map["Type"] = op.Type
		op_inner_map["Instruction"] = op.Instruction.String()
		json_str, err := json.Marshal(op_inner_map)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Instruction ID:",op_Pos)
		fmt.Println("Instruction Information(JSON):\t",string(json_str))

		op_outer_map[strconv.Itoa(int(op_Pos))] = string(json_str)
	}
	json_str, err := json.Marshal(op_outer_map)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("----------JSON of all operations---------")
	fmt.Println(string(json_str))
}