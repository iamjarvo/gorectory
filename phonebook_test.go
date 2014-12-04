package phonebook

import (
	"testing"

	"github.com/bmizerany/assert"
)

var emptyPhoneBook = NewPhoneBook()

func Test_entries_are_initially_empty(t *testing.T) {
	entries := emptyPhoneBook.entries
	assert.Equal(t, len(entries), 0)
}
