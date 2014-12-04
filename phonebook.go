package phonebook

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
