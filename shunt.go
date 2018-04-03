/*	Shunting Arm Algorithm
Author: Adrian McNulty
Based on instructional video "Shunting yard algorithm in Go"
https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e?channelId=f9970e30-b336-4145-8af3-a2bbe2938f5e&channelName=Graph%20Theory
*/
package main

import (
	"fmt"
)

/*	Create a function to convert
expressions from infix notation to postfix notation. */
func intopost(infix string) string {
	//	Create a map of the special characters and assign values to keep track of the precedence of the characters(i.e. * has a higher precedence than . which has precedence over |).
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	//	Initialise an empty array of runes
	pofix := []rune{}
	//	Create a stack to temporarily hold operators asthey are read from the string
	s := []rune{}
	//	Cast pofix to a string
	return string(pofix)
}

func main() {
	// Answer: ab.c*.
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Pofix: ", intopost("a.b.c*"))
	// Answer: abd|.*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Pofix: ", intopost("(a.(b|d))*"))
	// Answer: abd|.c*.
	fmt.Println("Infix: ", "a.(b|d).c*")
	fmt.Println("Pofix: ", intopost("a.(b|d).c*"))
	// Answer: abb.+.c
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Pofix: ", intopost("a.(b.b)+.c"))
}
