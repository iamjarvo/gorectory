package phonebook

type Entry struct {
}

type PhoneBook struct {
	entries []Entry
}

func NewPhoneBook() PhoneBook {
	return PhoneBook{}
}
