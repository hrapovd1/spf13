package todo_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/hrapovd1/spf13/tri/todo"
	"github.com/stretchr/testify/require"
)

func TestSaveItems(t *testing.T) {
	testFile := "/tmp/todo_test.json"
	items := []todo.Item{
		{Text: "one", Priority: 0, Done: true},
		{Text: "second todo", Priority: 0, Done: true},
	}

	err := todo.SaveItems(testFile, items)
	require.NoError(t, err)
}

func TestReadItems(t *testing.T) {
	// create file with json data
	fileName := "/tmp/read-test-todo.json"
	items := []todo.Item{
		{Text: "one", Priority: 1, Done: false},
		{Text: "two", Priority: 2, Done: true},
	}
	b, err := json.Marshal(items)
	require.Nil(t, err)
	require.Nil(t, ioutil.WriteFile(fileName, b, 0644))

	// run func ReadItems(filename string) ([]Item, error)
	_, err = todo.ReadItems(fileName)
	// check result
	require.NoError(t, err)
}

func TestSetPriority(t *testing.T) {
	priorTests := []struct {
		pri      int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 2},
		{0, 2},
		{-1, 2},
	}
	item := todo.Item{Text: "one", Priority: 1, Done: false}

	for _, pri := range priorTests {
		item.SetPriority(pri.pri)
		require.Equal(t, pri.expected, item.Priority)
	}
}

func TestPrettyP(t *testing.T) {
	tests := []struct {
		pri    int
		expect string
	}{
		{0, " "},
		{1, "(1)"},
		{2, "(2)"},
		{3, "(3)"},
		{4, " "},
	}
	items := []todo.Item{
		{Text: "text", Priority: 0, Done: false},
		{Text: "text", Priority: 1, Done: false},
		{Text: "text", Priority: 2, Done: false},
		{Text: "text", Priority: 3, Done: false},
		{Text: "text", Priority: -1, Done: false},
	}
	for _, test := range tests {
		require.Equal(t, test.expect, items[test.pri].PrettyP())
	}
}

func TestPrettyDone(t *testing.T) {
	tests := []struct {
		done   bool
		expect string
	}{
		{true, "X"},
		{false, ""},
	}
	item := todo.Item{Text: "text", Priority: 0, Done: false}
	for _, test := range tests {
		item.Done = test.done
		require.Equal(t, test.expect, item.PrettyDone())
	}
}

/*
func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func TestLabel(t *testing.T) {
	tests := []string{
		"1.",
		"2.",
		"3.",
	}
	items := []todo.Item{
		{Text: "text", Priority: 0, Done: false},
		{Text: "text", Priority: 1, Done: false},
		{Text: "text", Priority: 2, Done: false},
		{Text: "text", Priority: 3, Done: false},
		{Text: "text", Priority: -1, Done: false},
	}
	for n, item := range items {
		require.Equal(t, tests[n], item.Label())
	}
}
*/

func TestLen(t *testing.T) {
	count := 3
	items := make(todo.ByPri, count)
	require.Equal(t, count, items.Len())
}

func TestSwap(t *testing.T) {
	items := todo.ByPri{
		{Text: "1", Priority: 1, Done: false},
		{Text: "2", Priority: 2, Done: false},
		{Text: "3", Priority: 3, Done: false},
		{Text: "4", Priority: 4, Done: false},
	}
	tests := []struct {
		i int
		j int
	}{
		{0, 2},
		{3, 0},
		{1, 2},
	}
	for _, test := range tests {
		copyItems := make(todo.ByPri, 4)
		require.Equal(t, len(items), copy(copyItems, items))
		copyItems.Swap(test.i, test.j)
		require.Equal(t, items[test.i].Text, copyItems[test.j].Text)
		require.Equal(t, items[test.j].Text, copyItems[test.i].Text)
	}
}

func TestLess(t *testing.T) {
	items := todo.ByPri{
		{Text: "1", Priority: 1, Done: false},
		{Text: "2", Priority: 2, Done: false},
		{Text: "3", Priority: 3, Done: false},
		{Text: "4", Priority: 4, Done: false},
	}
	tests := []struct {
		i int
		j int
	}{
		{0, 2},
		{3, 0},
		{1, 2},
	}
	for _, test := range tests {
		require.Equal(t, test.i < test.j, items.Less(test.i, test.j))
	}
}
