package solution

import (
  "github.com/kyokomi/emoji"
)

func GetMessage(text string) string {
  return emoji.Sprint(text)
}

