package model

type HiveUserModel struct {
	UniqueIdentifier string

	// Basic Information
	FirstName  string
	MiddleName string
	LastName   string
	Alias      *string

	// Contacts
	GTEmail       string
	PersonalEmail string
	Phone         string
	GTUsername    string
}

type HiveUserQuery struct {
	UniqueIdentifier string

	// Basic Information
	FirstName  string
	MiddleName string
	LastName   string
	Alias      *string

	// Contacts
	GTEmail       string
	PersonalEmail string
	Phone         string
	GTUsername    string
}
