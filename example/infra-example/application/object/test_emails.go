// Code generated by 'freedom new-crud'
package object

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TestEmails struct {
	changes    map[string]interface{}
	Id         int       `gorm:"primary_key" column:"id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at"`
	TypeId     int       `gorm:"column:type_id"`
	Subscribed int       `gorm:"column:subscribed"`
	TestUserId int       `gorm:"column:test_user_id"`
}

func (obj *TestEmails) TableName() string {
	return "test_emails"
}

// TakeChanges .
func (obj *TestEmails) TakeChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// SetCreatedAt .
func (obj *TestEmails) SetCreatedAt(createdAt time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.CreatedAt = createdAt
	obj.changes["created_at"] = createdAt
}

// SetUpdatedAt .
func (obj *TestEmails) SetUpdatedAt(updatedAt time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.UpdatedAt = updatedAt
	obj.changes["updated_at"] = updatedAt
}

// SetDeletedAt .
func (obj *TestEmails) SetDeletedAt(deletedAt time.Time) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.DeletedAt = deletedAt
	obj.changes["deleted_at"] = deletedAt
}

// SetTypeId .
func (obj *TestEmails) SetTypeId(typeId int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.TypeId = typeId
	obj.changes["type_id"] = typeId
}

// SetSubscribed .
func (obj *TestEmails) SetSubscribed(subscribed int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Subscribed = subscribed
	obj.changes["subscribed"] = subscribed
}

// SetTestUserId .
func (obj *TestEmails) SetTestUserId(testUserId int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.TestUserId = testUserId
	obj.changes["test_user_id"] = testUserId
}

// AddTypeId .
func (obj *TestEmails) AddTypeId(typeId int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.TypeId += typeId
	obj.changes["type_id"] = gorm.Expr("type_id + ?", typeId)
}

// AddSubscribed .
func (obj *TestEmails) AddSubscribed(subscribed int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.Subscribed += subscribed
	obj.changes["subscribed"] = gorm.Expr("subscribed + ?", subscribed)
}

// AddTestUserId .
func (obj *TestEmails) AddTestUserId(testUserId int) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.TestUserId += testUserId
	obj.changes["test_user_id"] = gorm.Expr("test_user_id + ?", testUserId)
}
