package sync_check

import "strings"

func Op_to_prim(op_type string) string{
	index_of_ := strings.LastIndex(op_type,"_") // using LastIndex because we want atomic_value from atomic_value_load
	return op_type[:index_of_]
}
