package entity

const (
	HiveUserType_EndUser = iota
	HiveUserType_PI
)

type HiveUser struct {
	BaseModel
	UniqueIdentifier string `gorm:"unique;"`
	FirstName        string
	LastName         string
	Alias            *string
	Email            string
	Phone            string

	Type int
}
