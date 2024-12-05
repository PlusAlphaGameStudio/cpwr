package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*

cpwr: Copy with replace

[original directory]에 있는 모든 *.h, *.cpp 파일을

[original directory]/../[new name] 위치에 다 복사하되,

파일명과 파일 내용에 있는 [original directory의 마지막 디렉토리명] 문자열을 [new name]으로 모두 치환하여 복사한다.

 */
func main() {
	if len(os.Args) != 3 {
		fmt.Print("cpwr [original directory] [new name]\n")
		os.Exit(-1)
	}

	dirPath := os.Args[1]
	_, dirName := filepath.Split(dirPath)
	newDirName := os.Args[2]

	var pathList []string

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		pathList = append(pathList, path)
		return nil
	})
	if err != nil {
		panic(err)
		return 
	}

	for _, path := range pathList {
		if strings.HasSuffix(path, ".cpp") || strings.HasSuffix(path, ".h") {
			log.Print("       " + path)
			newPath := strings.ReplaceAll(path, dirName, newDirName)
			log.Print("    -> " + newPath)

			newBasePath, _ := filepath.Split(newPath)
			err := os.MkdirAll(newBasePath, os.ModeDir | 0644)
			if err != nil && os.IsExist(err) == false {
				panic(err)
			}

			fileContent, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}

			newFileContent := strings.ReplaceAll(string(fileContent), dirName, newDirName)

			err = os.WriteFile(newPath, []byte(newFileContent), 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}
