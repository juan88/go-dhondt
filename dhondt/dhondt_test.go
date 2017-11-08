package dhondt 

import (
	"testing"
)


func TestValidateThrowsError(t *testing.T) {
	voteMap := make(map[string]uint64)
	voteMap["partido1"] = 30
	e := ElectionResults{Parties: 1, Votes: voteMap, Seats: 0}

	err := e.Validate()
	if err == nil {
		t.Errorf("Should not validate!")
	}
}

func TestValidateStructOk(t *testing.T) {
	voteMap := make(map[string]uint64)
	voteMap["partido1"] = 30
	e := ElectionResults{Parties: 100, Votes: voteMap, Seats: 13}

	err := e.Validate()
	if err != nil {
		t.Errorf("Struct is valid. Should not be here! Error: %s", err)
	}
}