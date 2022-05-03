package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	// ans := GetMessage("Hello World")
	ans := GetMessage("Hello World")
	if ans == "" {
		t.Errorf("Hello World = %s; want \"Hello World\"", ans)
	}
}
