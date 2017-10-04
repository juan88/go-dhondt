package main

import (
	"flag" //For parsing the command line arguments
	"fmt"
    "os"
    "./dhondt"
)

func main() {
	filenamePtr := flag.String("file", "", "The input filename with the election results")
	seatsPtr := flag.Uint("seats", 0, "The available seats that are to be allocated.")
	partiesPtr := flag.Uint("parties", 0, "The number of parties in the election.")
	//votesPtr := flag.String("votesDistribution", "", "Vote distribution for parties
	flag.Parse()
	
	if(len(*filenamePtr) == 0) {
		fmt.Println("Error! File is requiered")
		os.Exit(1)
	}

	voteMap := make(map[string]uint64)
	voteMap["partido1"] = 30
	voteMap["partido2"] = 190
	voteMap["partido3"] = 329

	hola := &dhondt.ElectionResults{
	  Parties: *partiesPtr,
	  Seats: *seatsPtr,
	  Votes: voteMap,
	}

	fmt.Println("#Partidos: ", hola.Parties)
	fmt.Println("#Seats: ", hola.Seats)

    // Uncomment the following lines (and comment those above) to
    // test the behaviour of `LoadResults'.

    // var res *dhondt.ElectionResults = new(dhondt.ElectionResults)

    // res.LoadResults("file.txt", ",")

    // for k, v := range res.Votes {
    //     fmt.Printf("Party %s | Votes %d\n", k, v)
    // }
}

func banner() {
	fmt.Println("d'Hondt")
	fmt.Println("========")
	fmt.Println("Method for allocating proportional seats after an election.")
	fmt.Println("------------------------------------------------------------------------------------")
}
