package main

import "fmt"

func main() {
	const fixedNum = 50               // 固定数字
	var eventName = "LearningGOEvent" // イベント名
	var numberOfTickets int

	fmt.Printf("This is the start of the %v \nGet your tickets now!", eventName)
	// Printing out the types of my variables
	fmt.Printf("Fixed Num: %T \nEvent Name: %T \nNumber ofTickets: %T\n", fixedNum, eventName, numberOfTickets)

	fmt.Scan(&numberOfTickets) // チケット数を入力
	fmt.Printf("You have purchased %v tickets\nTotal price will be: %v", numberOfTickets, (fixedNum * numberOfTickets))
}
