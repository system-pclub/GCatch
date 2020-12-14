package prepare

import (
	"github.com/system-pclub/GCatch/GFix/dispatcher/global"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func Exclude_path(names_str string) []string {
	splits := strings.Split(names_str,":")
	result := []string{}
	for _,split := range splits {
		if split != "" {
			result = append(result,"/"+split+"/")
		}
	}
	return result
}


func ContainsAnyStr(str string,subs []string) bool {
	flag_any_contains := false
	for _,sub := range subs {
		if strings.Contains(str,sub) {
			flag_any_contains = true
		}
	}
	return flag_any_contains
}

func List_all_pkg_paths() []string {
	root := global.Absolute_root + global.Root
	var all_pkg_paths []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if ContainsAnyStr(path,global.Exclude) || info.IsDir() == false {
			return nil
		}
		file_infos,err := ioutil.ReadDir(path)
		if err != nil {
			return nil
		}
		flag_has_go := false
		for _,file_info := range file_infos {
			if strings.HasSuffix(file_info.Name(),".go") {
				flag_has_go = true
			}
		}
		if flag_has_go == false {
			return nil
		}

		all_pkg_paths = append(all_pkg_paths,path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return all_pkg_paths
}

func List_worthy_paths() (worthy_paths []global.Parent_path) {

	worthy_paths = *new([]global.Parent_path)
	all_path_stats := *new([]global.Path_stat)

	all_pkg_paths := List_all_pkg_paths()

	for _,path := range all_pkg_paths {
		self_num_lock := Grep_count_current(".Lock()",path)
		self_num_send := Grep_count_current("<-",path)
		path_stat := global.Path_stat{
			path,
			self_num_lock,
			self_num_send,
		}
		all_path_stats = append(all_path_stats,path_stat)
	}

	sort.Slice(all_path_stats, func(i, j int) bool {
			return ( 1 * all_path_stats[i].Self_num_lock + 1 * all_path_stats[i].Self_num_send) > ( 1 * all_path_stats[j].Self_num_lock + 1 * all_path_stats[j].Self_num_send)
	})

	global.All_pkg_paths = all_path_stats

	outer:
	for _,all_path_stat := range all_path_stats {
		//if i > 100 || len(worthy_paths) > 50 {
		//	break
		//}
		for index,_ := range worthy_paths {
			if strings.HasPrefix(all_path_stat.Path,worthy_paths[index].Parent) { //this path is a child of an existing path in worthy_paths
				worthy_paths[index].Children = append(worthy_paths[index].Children,all_path_stat.Path)
				continue outer
			}

			if strings.HasPrefix(worthy_paths[index].Parent,all_path_stat.Path) { //this path is a parent of an existing path in worthy_paths
				new_pair := global.Parent_path{
					all_path_stat.Path,
					[]string{worthy_paths[index].Parent},
					0,
					0,
				}
				worthy_paths = append(worthy_paths,new_pair)
				continue outer
			}
		}

		new_pair := global.Parent_path{
			all_path_stat.Path,
			[]string{},
			0,
			0,
		}
		worthy_paths = append(worthy_paths,new_pair)

	}

	for index,_ := range worthy_paths {
		worthy_paths[index].Total_num_lock = Grep_count_recursive(".Lock()",worthy_paths[index].Parent)
		worthy_paths[index].Total_num_send = Grep_count_recursive("<-",worthy_paths[index].Parent)
		worthy_paths[index].Parent = strings.ReplaceAll(worthy_paths[index].Parent,global.Absolute_root,"")

		for j,_ := range worthy_paths[index].Children {
			worthy_paths[index].Children[j] = strings.ReplaceAll(worthy_paths[index].Children[j],global.Absolute_root,"")
		}

	}


	return
}




