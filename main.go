package main
//import "fmt"

func main(){
	//var cards string="ace of spades"
	cards:= newDeck()
	//hand,remainingDeck:=deal(cards,5)
	//cards.print()
	//hand.print()
	//remainingDeck.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	//cards2:=newDecFromFile("my_cards")
	//cards2.print()
	cards.shuffle();
	cards.print()
}