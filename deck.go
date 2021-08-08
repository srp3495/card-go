package main
 
import ("fmt"
		"strings"
		"io/ioutil"
		"os"
		"math/rand"
		"time"		
 )
//create a type of deck
//which is actually a slice of string

type deck []string

func newDeck() deck{
	cards:=deck{}

	cardSuites :=[]string{"spades","diamonds","hearts","clever"}
	cardValues:=[]string{"Ace","two","three","four"}

	for _,suit:=range cardSuites{
		for _,value:=range cardValues{
			cards=append(cards,value+" of"+suit)
		}
	}
	return cards
}

func deal(d deck,handSize int) (deck,deck){

	return d[:handSize],d[handSize:]
}

func(d deck) toString() string{	
return strings.Join([]string(d),",")
}

func(d deck) saveToFile(filename string) error{
 return ioutil.WriteFile(filename,[]byte(d.toString()),0666) //0666 is the permission
}

func newDecFromFile(filename string) deck{
 bs,err:=ioutil.ReadFile(filename)
 if err!=nil{
	 fmt.Println("Error :",err)
	 os.Exit(1)
 }
 s:=strings.Split(string(bs),",") //will give  slice of strings
 return deck(s)

}

func(d deck) shuffle(){
	source:=rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)
	for i:=range d{
		newPosition:=r.Intn(len(d)-1)
		d[i],d[newPosition]=d[newPosition],d[i]
	}
}
func(d deck) print(){
   for i,card:=range d{
	   fmt.Println(i,card)
   }
}

