package entity

type Workbench struct {
	BaseModel
	BenchName      string
	CheckedOut     bool
	WorkbenchItems []*WorkbenchItem         `gorm:"many2many:workbench_item_relations;"`
	RentalRecords  []*WorkbenchRentalRecord `gorm:"foreign_key:workbench_ref_id;"json:"-"`
}

type WorkbenchItem struct {
	BaseModel
	Name          string
	InventoryLink *InventoryItemClass
}

const (
	WorkbenchRecordType_Checkout = iota
	WorkbenchRecordType_Return
)

var WorkbenchRecordTypes = map[int]string{
	WorkbenchRecordType_Checkout: "CheckOut",
	WorkbenchRecordType_Return:   "Return",
}

type WorkbenchRentalRecord struct {
	BaseModel
	RecordType     int
	WorkbenchRefId EntityIDType
	WorkbenchRef   *Workbench
	UserRef        *HiveUser
	UserRefId      EntityIDType
}
