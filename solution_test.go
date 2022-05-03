package solution

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	ans := GetMessage()
	if ans == "" {
		t.Errorf("GetMessage = %s; want \"Hello :world_map:\"", ans)
	}
}
