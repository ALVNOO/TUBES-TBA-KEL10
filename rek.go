package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
) dfaTransisi

// dfaTransisi mengembalikan state berikutnya berdasarkan state saat ini dan simbol input
func(state int, symbol rune) int {
	switch state {
	case 0:
		if symbol == '<' {
			return 1
		}
	case 1:
		switch symbol {
		case 'h', 'H':
			return 2
		case 't', 'T':
			return 8
		case 'b', 'B':
			return 14
		case 'p', 'P':
			return 18
		case '/':
			return 22
		}
	case 2:
		switch symbol {
		case 't', 'T':
			return 3
		case '1':
			return 7
		case 'e', 'E':
			return 27
		}
	case 3:
		if symbol == 'm' || symbol == 'M' {
			return 4
		}
	case 4:
		if symbol == 'l' || symbol == 'L' {
			return 5
		}
	case 5, 7:
		if symbol == '>' {
			return 6
		}
	case 8:
		if symbol == 'i' || symbol == 'I' {
			return 9
		}
	case 9:
		if symbol == 't' || symbol == 'T' {
			return 10
		}
	case 10:
		if symbol == 'l' || symbol == 'L' {
			return 11
		}
	case 11:
		if symbol == 'e' || symbol == 'E' {
			return 12
		}
	case 12:
		if symbol == '>' {
			return 6
		}
	case 14:
		if symbol == 'o' || symbol == 'O' {
			return 15
		}
	case 15:
		if symbol == 'd' || symbol == 'D' {
			return 16
		}
	case 16:
		if symbol == 'y' || symbol == 'Y' {
			return 17
		}
	case 17, 18:
		if symbol == '>' {
			return 6
		}
	case 22:
		switch symbol {
		case 'h', 'H':
			return 23
		case 't', 'T':
			return 24
		case 'b', 'B':
			return 30
		case 'p', 'P':
			return 34
		}
	case 23:
		switch symbol {
		case 't', 'T':
			return 24
		case '1':
			return 7
		case 'e', 'E':
			return 27
		}
	case 24:
		if symbol == 'm' || symbol == 'M' {
			return 25
		}
	case 25:
		if symbol == 'l' || symbol == 'L' {
			return 26
		}
	case 26:
		if symbol == '>' {
			return 6
		}
	case 27:
		if symbol == 'a' || symbol == 'A' {
			return 28
		}
	case 28:
		if symbol == 'd' || symbol == 'D' {
			return 29
		}
	case 29:
		if symbol == '>' {
			return 6
		}
	case 30:
		if symbol == 'o' || symbol == 'O' {
			return 31
		}
	case 31:
		if symbol == 'd' || symbol == 'D' {
			return 32
		}
	case 32:
		if symbol == 'y' || symbol == 'Y' {
			return 33
		}
	case 33:
		if symbol == '>' {
			return 6
		}
	}
	return -1
}

// terimadfa memeriksa apakah inputString diterima oleh DFA
func terimadfa(inputString string) bool {
	state := 0
	for _, symbol := range inputString {
		state = dfaTransisi(state, symbol)
		if state == -1 {
			return false
		}
	}
	return state == 6
}

func main() {
	recognizedTags := []string{"<html>", "<head>", "<title>", "<body>", "<h1>", "<p>", "<h1>", "</html>", "</head>", "</body>", "</h1>"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan tag HTML: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	accepted := false
	for _, tag := range recognizedTags {
		if terimadfa(userInput) && strings.EqualFold(userInput, tag) {
			accepted = true
			break
		}
	}

	if accepted {
		fmt.Printf("Tag '%s' : Accepted\n", userInput)
	} else {
		fmt.Printf("Tag '%s' : Rejected\n", userInput)
	}
}
