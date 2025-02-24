package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type LocalDate time.Time

// String returns the string representation of LocalDate in YYYY-MM-DD format
func (d LocalDate) String() string {
	return time.Time(d).Format(time.DateOnly)
}

// MarshalJSON implements json.Marshaler interface
func (d LocalDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON implements json.Unmarshaler interface
func (d *LocalDate) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	t, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return err
	}

	*d = LocalDate(t)
	return nil
}

// Value implements driver.Valuer interface
func (d LocalDate) Value() (driver.Value, error) {
	return time.Time(d), nil
}

// Scan implements sql.Scanner interface
func (d *LocalDate) Scan(value any) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*d = LocalDate(v)
		return nil
	case string:
		t, err := time.Parse(time.DateOnly, v)
		if err != nil {
			return err
		}
		*d = LocalDate(t)
		return nil
	default:
		return fmt.Errorf("cannot scan %T into LocalDate", value)
	}
}

// SysUser represents a system user entity
type SysUser struct {
	UserID         int32     `json:"userId" gorm:"column:user_id;primaryKey;autoIncrement"`
	LoginName      string    `json:"loginName" gorm:"column:login_name"`
	UserPwd        *string   `json:"userPwd" gorm:"column:user_pwd"`
	PwdExpiredDays *int32    `json:"pwdExpiredDays" gorm:"column:pwd_expired_days"`
	CreatedTime    LocalDate `json:"createdTime" gorm:"column:created_time"`
}
