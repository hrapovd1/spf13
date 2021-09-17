package todo_test

import (
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
