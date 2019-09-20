package entity

import (
	"database/sql/driver"
	"fmt"
	"github.com/ECEHive/myHive-backend/util"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
)

var logger = util.GetLogger("entity")

type UnixTime time.Time
type EntityIDType int64

func (t UnixTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%d", time.Time(t).UTC().Unix())
	return []byte(stamp), nil
}

func (t *UnixTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	zone := time.Now().Location()
	*t = UnixTime(time.Unix(q, 0).In(zone))
	return
}

func (t UnixTime) String() string {
	ts := time.Time(t)
	return ts.Format(time.RFC3339)
}

func (t UnixTime) Value() (driver.Value, error) {
	return time.Time(t).UTC(), nil
}

func (t *UnixTime) Scan(src interface{}) error {
	if val, ok := src.(time.Time); ok {
		*t = UnixTime(val.In(time.Now().Location()))
	}
	return nil
}

type BaseModel struct {
	Id        EntityIDType `gorm:"primary_key"json:"id"`
	CreatedAt *UnixTime    `gorm:"type:timestamp"json:"createdAt"`
	UpdatedAt *UnixTime    `gorm:"type:timestamp"json:"updatedAt"`
}

func (m *BaseModel) BeforeCreate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("CreatedAt", UnixTime(time.Now()))
	if err != nil {
		return
	}
	err = scope.SetColumn("UpdatedAt", UnixTime(time.Now()))
	return
}

func (m *BaseModel) BeforeUpdate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("UpdatedAt", UnixTime(time.Now()))
	return
}
