package prepare

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"github.com/system-pclub/GCatch/GFix/dispatcher/tools/go/ssa"
	"strings"
)



func Is_path_include(str string) bool {
	if strings.Contains(str, global.Root)   &&
		!ContainsAnyStr(str, global.Exclude) {
		return true
	} else {
		return false
	}
}

func List_all_struct(prog *ssa.Program) []*global.Struct {
	all_structs := *new([]*global.Struct)

	for _, pkg := range prog.AllPackages() { //loop all packages
		if pkg == nil {
			continue
		}

		if Is_path_include(pkg.Pkg.Path()) == false {
			continue
		}

		for mem_name, _ := range pkg.Members { //loop through all members; the member may be a func, a type, etc

			//check if this member is a type
			mem_as_type := pkg.Type(mem_name)
			if mem_as_type != nil {
				//this member is a type

				//check if this member is in our interested path
				if Is_path_include(mem_as_type.String()) == false {
					continue
				}

				struct_name := mem_as_type.String()
				fields_str := mem_as_type.Object().Type().Underlying().String()


				if !strings.HasPrefix(fields_str,"struct{") || fields_str == "struct{}" {
					continue
				}
				fields_str = strings.Replace(fields_str,"struct{","",1)
				fields_str = fields_str[:len(fields_str) - 1]//delete the last char, which is "}"
				fields := strings.Split(fields_str,"; ")
				if len(fields) == 0  {
					continue
				} else {
					str := strings.ReplaceAll(fields[0]," ","")
					if len(str) == 0 {
						continue
					}
				}

				struct_field := make(map[string]string)
				for _,field := range fields {
					field_element := strings.Split(field," ")
					var field_name,field_type string
					if len(field_element) == 1 { //this is an anonymous field
						field_name = field_element[0]
						last_dot_index := strings.LastIndex(field_name,".") // from "*github.com/coreos/etcd/mvcc.store", we only want "store"
						field_name = field_name[last_dot_index+1:]
						//fmt.Println("Anonymous field:",field_element[0],"\trefined:",field_name,"\tstruct.Name:",struct_name)

						field_type = field_element[0]
						if field_element[0] == "chan" && len(field_element) > 1 {
							field_type = "chan " + field_element[1]
						}
					} else {
						field_name = field_element[0]
						field_type = field_element[1]
						if field_element[1] == "chan" && len(field_element) > 2 {
							field_type = "chan " + field_element[2]
						}
					}
					struct_field[field_name] = field_type
				}
				new_struct_ptr := &global.Struct{
					Name: struct_name,
					Field: struct_field,
				}

				all_structs = append(all_structs,new_struct_ptr)

			}
		}
	}

	return all_structs
}

func List_all_sync_struct(prog *ssa.Program) []*global.Sync_struct { //TODO: This function is proved to be sound by checking with etcd, but not proved to be complete
	all_sync_struct := *new([]*global.Sync_struct)

	all_potential_struct := *new([]*global.Sync_struct)
	for _, pkg := range prog.AllPackages() { //loop all packages
		if pkg == nil {
			continue
		}

		if Is_path_include(pkg.Pkg.Path()) == false {
			continue
		}

		for mem_name, _ := range pkg.Members { //loop through all members; the member may be a func or a type; if it is type, loop through all its methods

			//check if this member is a type
			mem_as_type := pkg.Type(mem_name)
			if mem_as_type != nil {
				//this member is a type

				//check if this member is in our interested path
				if Is_path_include(mem_as_type.String()) == false {
					continue
				}

				struct_type := mem_as_type.String()
				struct_containsync := false
				struct_potential := false

				fields_str := mem_as_type.Object().Type().Underlying().String()
				if !strings.HasPrefix(fields_str,"struct{") || fields_str == "struct{}" {
					continue
				}
				fields_str = strings.Replace(fields_str,"struct{","",1)
				fields_str = fields_str[:len(fields_str) - 1]//delete the last char, which is "}"
				fields := strings.Split(fields_str,"; ")
				if len(fields) == 0  {
					continue
				} else {
					str := strings.ReplaceAll(fields[0]," ","")
					if len(str) == 0 {
						continue
					}
				}

				struct_field := make(map[string]string)
				for index,field := range fields {
					field_element := strings.Split(field," ")
					var field_name,field_type string
					if len(field_element) == 1 { //anonymous field
						field_name = string(index)
						field_type = field_element[0]
						if field_element[0] == "chan" && len(field_element) > 1 {
							field_type = "chan " + field_element[1]
						}
					} else {
						field_name = field_element[0]
						field_type = field_element[1]
						if field_element[1] == "chan" && len(field_element) > 2 {
							field_type = "chan " + field_element[2]
						}
					}
					struct_field[field_name] = field_type
					if is_type_sync(field_type) {
						struct_containsync = true
						struct_potential = true
					}
					if strings.Contains(field_type,"/") {
						struct_potential = true
					}
				}
				new_struct_ptr := &global.Sync_struct{
					Type: struct_type,
					ContainSync: struct_containsync,
					Field: struct_field,
					Potential: struct_potential,
				}

				if struct_potential == true {
					all_potential_struct = append(all_potential_struct,new_struct_ptr)
				}
			}
		}
	}



	//now all potential structs are listed
	//at this point, each struct in all_potential_struct has struct.potential == true, meaning it may be moved to all_sync_struct
	//append structs that ContainSync = true, to all_sync_struct

	for _,struct_potential_ptr := range all_potential_struct {
		if (*struct_potential_ptr).ContainSync {
			all_sync_struct = append(all_sync_struct,struct_potential_ptr)
			(*struct_potential_ptr).Potential = false
		}
	}


	for {
		flag_break := true

		potential:
		for _,struct_potential := range all_potential_struct {
			if struct_potential.Potential == false {//this struct has already been moved to all_sync_struct
				continue
			}

			for _,field_type := range struct_potential.Field {
				for _,struct_sync := range all_sync_struct {
					if struct_sync.Type == field_type {// This struct contains a field, which is in all_sync_struct
						all_sync_struct = append(all_sync_struct,struct_potential)
						struct_potential.Potential = false
						flag_break = false // don't break the infinite loop since there is still at least one struct moving from all_potential_struct to all_sync_struct
						continue potential
					}
				}
			}
		}

		if flag_break == true {
			break
		}

	}

	return all_sync_struct
}


func is_type_sync(str string) bool {
	str = strings.ReplaceAll(str,"*","")
	str = strings.ReplaceAll(str,"[]","")
	if strings.HasPrefix(str,"sync.") || strings.Contains(str,"<-") || strings.HasPrefix(str,"chan ") {
		return true
	} else if strings.Contains(str,"map[") {
		if strings.Contains(str,"sync.") || strings.Contains(str,"<-") || strings.Contains(str,"chan ") {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

