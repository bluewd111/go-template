package functions

import (
	"fmt"
	"sort"
	"testing"
)

func TestMapValues(t *testing.T) {
	m := map[string]*int{
		"one": new(int),
		"two": new(int),
	}
	*m["one"] = 1
	*m["two"] = 2

	values := MapValues(m)

	if len(values) != 2 {
		t.Errorf("Expected length 2, got %d", len(values))
	}

	// 取得した値をソート
	sort.Slice(values, func(i, j int) bool { return *values[i] < *values[j] })

	// 期待する値を設定
	expected := []*int{new(int), new(int)}
	*expected[0] = 1
	*expected[1] = 2

	for _, v := range values {
		fmt.Println(*v) // 出力して確認
	}

	// 取得した値が期待する値と一致するか確認
	for i, v := range values {
		if *v != *expected[i] {
			t.Errorf("Expected value %d, got %d", *expected[i], *v)
		}
	}
}
