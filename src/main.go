package main

import (
	"flag"
	"fmt"

	gitlocalstats "github.com/KennyZ69/gitlocalstats"
)

var (
	email_flag *string = flag.String("email", "email@email.com", "email to scan")
	dir_flag   *string = flag.String("add", "", "new directory to add for git scan")
)

func main() {
	flag.Parse()

	if *dir_flag != "" {
		gitlocalstats.Scan(*dir_flag)
		return
	}

	stats(*email_flag)
}

func stats(email string) {
	fmt.Println("stats")
}
