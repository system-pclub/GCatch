package prepare

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"strings"
)

func List_all_sync_global(all_sync_struct []*global.Sync_struct, prog *ssa.Program) []*ssa.Global {//TODO: This function is proved to be sound by checking with etcd, but not proved to be complete

	all_sync_global := *new([]*ssa.Global)
	for _, pkg := range prog.AllPackages() { //loop all packages
		if pkg == nil {
			continue
		}

		//Skip builtin packages, vendor packages. Test functions are automatically skipped. Include packages in "include"
		if Is_path_include(pkg.Pkg.Path()) {
		} else {
			continue
		}

		all_sync_struct_type := *new([]string)
		for _,sync_struct := range all_sync_struct {
			all_sync_struct_type = append(all_sync_struct_type,sync_struct.Type)
		}

		for _,mem := range pkg.Members { //loop through all members; the member may be a func or a type; if it is type, loop through all its methods
			mem_as_global,ok := mem.(*ssa.Global)
			if ok {
				//This member is a global

				if is_type_sync((*mem_as_global).Type().String()) {//This global is of sync type
					all_sync_global = append(all_sync_global,mem_as_global)
					continue
				}

				if is_type_sync_struct((*mem_as_global).Type().String(),all_sync_struct_type) {
					all_sync_global = append(all_sync_global,mem_as_global)
					continue
				}

			}

		} // end of member loop
	} //end of package loop

	return all_sync_global
}

func is_type_sync_struct(str string, all_sync_struct_type []string) bool {
	str = strings.ReplaceAll(str,"*","")
	str = strings.ReplaceAll(str,"[]","")
	for _,sync_struct_type := range all_sync_struct_type {
		if strings.HasPrefix(str,sync_struct_type){
			return true
		}
	}

	if strings.Contains(str,"map[") {
		for _,sync_struct_type := range all_sync_struct_type {
			if strings.Contains(str,sync_struct_type){
				return true
			}
		}
	}
	return false
}