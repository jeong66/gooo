package main

import(
	"fmt"
	"log"
	"../datafile"
)

func main(){
	lines, err := datafile.GetStrings("votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	var counts map[string]int 
	counts = make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}
	
	for name, count := range counts {
		fmt.Printf("votes for %s : %d\n", name, count)
	}
}
