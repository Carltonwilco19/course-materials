// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=B0WFcVTmIGluwigF6ys0un4jnA1Vqsab ./main <search term> <pagenumber>

package main

import (
	"fmt"
	"log"
	"os"
//	"encoding/json"
	"shodan/shodan"
	"strconv"
)

func main() {
	if len(os.Args) != 3  {
		log.Fatalln("Usage: main <searchterm> <startingpage>") // main is the first argument and the serach term is the second
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	i,_ := strconv.Atoi(os.Args[2]) // parse string arg to int
	hostSearch, err := s.HostSearch(os.Args[1], i)// here I would add in the page number, extending main.go
	if err != nil {
		log.Panicln(err)
	}

	// fmt.Printf("Host Data Dump\n")
	// for _, host := range hostSearch.Matches {
	// 	fmt.Println("==== start ",host.IPString,"====")
	// 	h,_ := json.Marshal(host)
	// 	fmt.Println(string(h))
	// 	fmt.Println("==== end ",host.IPString,"====")
	// 	//fmt.Println("Press the Enter Key to continue.")
	// 	//fmt.Scanln()
	// }

	



	fmt.Printf("IP, Port, City\n")

	for _, host := range hostSearch.Matches {
		fmt.Printf("IP: %s, Port: %d, City: %s\n", host.IPString, host.Port, host.Location.City)
	}


}