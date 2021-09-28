package plugin

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"strings"
)

// getFiles recursively reads all files in a
// given directory while omitting directories
// and returning a list of dir/filename
func getFiles(dir string) ([]string, error) {
	results, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, file := range results {
		if file.IsDir() {
			continue
		}
		files = append(files, dir+"/"+file.Name())
	}
	return files, nil
}

// isAllowed is a curry function that handles
// predication and is to be passed as a callback
func isAllowed(denyList []string) func(string) bool {
	return func(file string) bool {
		for _, deny := range denyList {
			if strings.Contains(file, deny) {
				return false
			}
		}
		return true
	}
}

// filter takes a list and callback and filters
// out all values that are falsey predicates
func filter(list []string, predicate func(string) bool) []string {
	var filtered []string

	for _, v := range list {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// findAndReplace loops through a list of files
// a replaces all given instances of a 'find' with
// a base64 corresponding 'value' and overwriting
// the existing file
func findAndReplace(files []string, find, value string) error {
	for _, fileName := range files {
		replace := base64.
			StdEncoding.
			EncodeToString([]byte(value))

		file, _ := ioutil.ReadFile(fileName)
		data := bytes.Replace(file, []byte(find), []byte(replace), -1)

		if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
			return err
		}
	}
	return nil
}
