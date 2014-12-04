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

func Test_adding_entries(t *testing.T) {
	phonebook := NewPhoneBook()
	entry := Entry{name: "Jearvon"}
	phonebook.add(entry)
	assert.Equal(t, len(phonebook.entries), 1)
}
