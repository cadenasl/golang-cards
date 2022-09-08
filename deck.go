package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string
func newDeck()deck{
	cards:= deck{}
	cardSuits:=[]string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues:=[]string{"Ace","Two","Three","Four"}
	for _,cardSuit:=range cardSuits{
		for _,cardValue:=range cardValues{
			cards=append(cards, cardValue+" of "+cardSuit)
		} 
	}
	return cards
}
func(d deck)print(){
	for i,card :=range d{
		fmt.Println(i,card)
	}
}
func deal(d deck,handsize int)(deck,deck){
return d[:handsize],d[handsize:]
}
func (d deck) toString() string{
return strings.Join([]string(d),",")
}

func(d deck) saveToFile(filename string)error{
	return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}
func newDeckFromFile(filename string)(deck){
bs,err:=ioutil.ReadFile(filename)
if err!=nil{
	fmt.Println("Error:",err)
	 os.Exit(1)
}
s := strings.Split(string(bs),",")
return deck(s)
}
func(d deck) shuffleDeck(){
for i:=range d{
	newPosition:= rand.Intn(len(d)-1)
	d[i],d[newPosition]=d[newPosition],d[i]
}
}