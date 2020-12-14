package prepare

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"sync"
)


var found = 1
var wg sync.WaitGroup
var total int

func readFile(wg *sync.WaitGroup, path string, query string) {
	defer wg.Done()

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		if strings.Contains(scanner.Text(), query) {
			found = 0
			total ++
		}
	}
}

func Grep_count_current(query string, root string) int {
	total = 0

	err := filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() && second_last_index(path,"/") == strings.LastIndex(root,"/") {
			wg.Add(1)
			go readFile(&wg, path, query)
		}
		return nil
	})

	wg.Wait()

	if err != nil {
		return 0
	}

	return total
}

func Grep_count_recursive(query string, root string) int {
	total = 0

	err := filepath.Walk(root, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			wg.Add(1)
			go readFile(&wg, path, query)
		}
		return nil
	})

	wg.Wait()

	if err != nil {
		return 0
	}

	return total
}

func second_last_index(str string, sub_str string) int {
	str_cut := str[:strings.LastIndex(str,sub_str)]
	return strings.LastIndex(str_cut,sub_str)
}
