package main

import "fmt"

import "github.com/kyokomi/emoji"

func main() {
  hw := emoji.Sprint("Hello :world_map:!")
  fmt.Println(hw)
}