package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	hw := Emojify("Hello :world_map:!")
	fmt.Println(hw)
}

func Emojify(text string) string {
	return emoji.Sprint(text)
}
