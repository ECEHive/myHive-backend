package entity

const (
	HiveUserType_EndUser = iota
	HiveUserType_PI
)

type HiveUser struct {
	BaseModel
	// Hash Derived from gtid
	UniqueIdentifier string `gorm:"unique;"`

	// Basic Information
	FirstName  string
	MiddleName string
	LastName   string
	Alias      *string

	// Contacts
	GTEmail       string `gorm:"unique;"`
	PersonalEmail string `gorm:"unique;"`
	Phone         string `gorm:"unique;"`
	GTUsername    string `gorm:"unique;"`

	HiveUserType int // Described in the enum (const group) above
}
