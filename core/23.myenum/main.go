package main

import "fmt"

type Season int

const (
	Summer Season = iota
	Autumn
	Winter
	Spring
)

func (s Season) String() string {
	switch s {
	case Summer:
		return "summer"
	case Autumn:
		return "autumn"
	case Winter:
		return "winter"
	case Spring:
		return "spring"
	}
	return "unknown"
}

func printSeason(s Season) {
	fmt.Println("season: ", s)
}

func main() {
	printSeason(Autumn)
}

//package main
//
//import "fmt"
//
//type Card int
//
//const (
//	Spades Card = iota
//	Diamond
//	Heart
//	Club
//)
//
//// returns the corresponding string value to the Card enum
//func (c Card) String() string {
//	switch c {
//	case Spades:
//		return "Spades ♤"
//	case Diamond:
//		return "Diamond ♢"
//	case Heart:
//		return "Heart ♥"
//	case Club:
//		return "Club ♧"
//	}
//	return "Unknown card suit"
//}
//
//func main() {
//	// define the day of type Card (enum)
//	var suit Card
//	//  the day
//	suit = Diamond
//	fmt.Printf("Suit is %s \n", suit)
//	fmt.Printf("Suit is %s \n", suit.String())
//}
