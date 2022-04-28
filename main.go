package golangunitedlesson0

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	hw := emoji.Sprint("Hello :world_map:!")
	fmt.Println(hw)
}

func Emojify(text string) string {
	return emoji.Sprint(text)
}
