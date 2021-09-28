package plugin

import (
	"io/ioutil"
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
