package gitlocalstats

import (
	"bufio"
	"log"
	"os"
	"os/user"
	"strings"
)

func getReposFile() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	f := usr.HomeDir + "/.localgitstats"

	return f
}

func addRepos(path string, news []string) {
	existing := parseRepos(path)
	repos := joinSlice(existing, news)
	dumpToFile(repos, path)
}

func openFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(path)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
		// log.Fatal(err)
	}

	return f
}

// parse each line of file into a slice of strings
func parseRepos(path string) []string {
	f := openFile(path)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func sliceContains[T comparable](slice []T, x T) bool {
	for _, v := range slice {
		if v == x {
			return true
		}
	}
	return false
}

// Generic function to join two slices or arrays
func joinSlice[T comparable](a, b []T) []T {
	for _, v := range b {
		if !sliceContains(a, v) {
			a = append(a, v)
		}
	}
	return a
}

func dumpToFile[T any](data T, path string) {
	var content string

	switch t := any(data).(type) {
	case string:
		content = t
		break
	case []string:
		content = strings.Join(t, "\n")
		break
	default:
		break
	}

	os.WriteFile(path, []byte(content), 0755)
}
