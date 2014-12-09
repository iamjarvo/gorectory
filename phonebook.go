package phonebook

import (
	"regexp"
)

type Entry struct {
	name string
}

type PhoneBook struct {
	entries []Entry
}

func NewPhoneBook() PhoneBook {
	return PhoneBook{}
}

func (phoneBook *PhoneBook) add(entry Entry) {
	phoneBook.entries = append(phoneBook.entries, entry)
}

func (phoneBook PhoneBook) search(name string) Entry {
	r, _ := regexp.Compile(name)
	entries := phoneBook.entries
	entry := Entry{}
	for _, current_entry := range entries {
		if r.MatchString(current_entry.name) {
			entry.name = current_entry.name
			break
		}
	}
	return entry
}
