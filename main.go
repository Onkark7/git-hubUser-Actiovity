package main

import (
	"flag"
	opra "user-activitiy-github/opration"
)

func main() {
	username := flag.String("username", "", "give username for show Activity")

	flag.Parse()
	opra.FetchUserActivity(*username)

}
