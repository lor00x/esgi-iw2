package main

import "fmt"

func info(s string) {
	fmt.Printf("Char: '%s', runes:  %d, bytes: %d\n", s, len([]rune(s)), len(s))
}

func main() {
	heart := "❤️"
	lol := '😆'
	msg := "Happy birthday ! 🎉🥳🎂"
	info(heart)
	info(string(lol))
	info(msg)

	for i, run := range msg {
		fmt.Println(i, string(run))
	}

}
