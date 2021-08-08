package main

import ("testing"
        "os"
)

func TestNewDeck(t *testing.T){
	d:=newDeck()
	if(len(d)!=16){
		t.Errorf("expected size of 16 but was %v",len(d))
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T){
	os.Remove("_deckTesting")
	deck:=newDeck()
	deck.saveToFile("_deckTesting")
	loadedDeck:=newDecFromFile("_deckTesting")
	if(len(loadedDeck)!=16){
		t.Errorf("Expected 16 but was %v",len(loadedDeck))
	}
	os.Remove("_deckTesting")
}