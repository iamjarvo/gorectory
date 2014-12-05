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

func Test_searching_for_an_entry(t *testing.T) {
	phoneBook := NewPhoneBook()
	dave := Entry{name: "Dave Lumberg"}
	harold := Entry{name: "Harold Peterson"}
	phoneBook.add(dave)
	phoneBook.add(harold)
	found := phoneBook.search("Dave")
	assert.Equal(t, found.name, dave.name)
}
