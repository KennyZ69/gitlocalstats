package gitlocalstats

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Scan(dir string) {
	log.Println("Scanning dir ...", dir)
	repos := scanDir(dir)
	log.Println("Got the repos:", repos)
	f := getReposFile()
	addRepos(f, repos)
}

func scanDir(dir string) []string {
	return scanGitFolders(make([]string, 0), dir)
}

// returns list of subdirectories of the provided dir ending with .git
func scanGitFolders(dirs []string, dir string) []string {
	dir = strings.TrimSuffix(dir, "/")

	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = dir + "/" + file.Name()
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println("Found:", path)
				dirs = append(dirs, path)
				continue
			}
			dirs = scanGitFolders(dirs, path)
		}
	}

	return dirs
}
