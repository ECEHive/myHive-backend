package entity

type Sequence struct {
	BaseModel
	SequenceName  *string `gorm:"unique; NOT NULL"`
	SequenceValue int64
}
