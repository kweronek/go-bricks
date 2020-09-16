package main

import (
"fmt"
"unicode/utf8"
)

func printBytes(s string) {
	fmt.Println("Printing Bytes")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printCharsUsingFor(s string) {
	fmt.Println("Printing runes using for")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytesUsingForRange(s string) {
	fmt.Println("Printing characters and bytes using for Range")
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutate(s []rune, i int, insChar []rune ) string {
	if len(insChar) > 1 || i > len(s) {
		fmt.Println("too long, not mutated")
	} else {
		s[i] = insChar[0]
	}
	return string(s)
}

func mutates(s []byte, i int, insChar []byte ) string {
	if len(insChar) > 1 || i > len(s) {
		fmt.Println("too long, not mutated")
	} else {
		s[i] = insChar[0]
	}
	return string(s)
}



func main() {
	fmt.Println("Demo for strings")

	name := "Señor"
	printBytes(name)
	fmt.Println("\n")

	printCharsUsingFor(name)
	fmt.Println("\n")

	printCharsAndBytesUsingForRange(name)
	fmt.Println("\n")

	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println("String constructed from hexadecimal byte slice")
	fmt.Println(str)
	fmt.Println()

	byteSlice = []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str = string(byteSlice)
	fmt.Println("String constructed from decimal byte slice")
	fmt.Println(str + "\n")

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println("String constructed from rune slice")
	fmt.Println(str)
	fmt.Println()

	length(str)

	fmt.Println("Mutate:")
	fmt.Println(mutate([]rune("hello"),3, []rune("k")))
	fmt.Println(mutates([]byte("hello"),3, []byte("k")))

	var string1, string2 []byte
	string1 = []byte("qwertz")
	string2 = []byte("ha")
	string1[4] = string2[0]
	fmt.Println(string(string2))

}
