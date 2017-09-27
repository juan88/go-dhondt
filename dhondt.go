package main

import (
	"flag" //For parsing the command line arguments
	"fmt"
	//"os"
)

type ElectionResults struct {
	Parties uint
	Votes map[string]uint
	Seats uint
}

func main() {
	//filenamePtr := flag.String("file", "", "The input filename with the election results")
	seatsPtr := flag.Uint("seats", 0, "The available seats that are to be allocated.")
	partiesPtr := flag.Uint("parties", 0, "The number of parties in the election.")
	//votesPtr := flag.String("votesDistribution", "", "Vote distribution for parties
	flag.Parse()
	
	//if(len(*filenamePtr) == 0) {
	//	fmt.Println("Error! File is requiered")
	//	os.Exit(1)
	//}

	voteMap := make(map[string]uint)
	voteMap["partido1"] = 30
	voteMap["partido2"] = 190
	voteMap["partido3"] = 329

	hola := &ElectionResults{
	  Parties: *partiesPtr,
	  Seats: *seatsPtr,
	  Votes: voteMap,
	}

	fmt.Println("#Partidos: ", hola.Parties)
	fmt.Println("#Seats: ", hola.Seats)
}

func banner() {
	fmt.Println("d'Hondt")
	fmt.Println("========")
	fmt.Println("Method for allocating proportional seats after an election.")
	fmt.Println("------------------------------------------------------------------------------------")
}