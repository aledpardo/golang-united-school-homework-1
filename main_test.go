package golangunitedlesson0

import (
	"testing"
)

func TestEmojify(t *testing.T) {
	ans := Emojify("Hello :world_map:")
	if ans == "" {
		t.Errorf("Emojify(\"Hello :world_map:\" = %s;", ans)
	}
}
