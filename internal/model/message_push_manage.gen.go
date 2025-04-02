// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMessagePushManage = "message_push_manage"

// MessagePushManage mapped from table <message_push_manage>
type MessagePushManage struct {
	ID           string     `gorm:"column:id;primaryKey" json:"id"`
	UserID       string     `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`               // 用户id
	PushID       string     `gorm:"column:push_id;not null;comment:推送id" json:"push_id"`               // 推送id
	DeviceType   string     `gorm:"column:device_type;not null;comment:设备类型" json:"device_type"`       // 设备类型
	Status       int16      `gorm:"column:status;not null;default:1;comment:类型 1正常 2注销" json:"status"` // 类型 1正常 2注销
	CreateTime   time.Time  `gorm:"column:create_time;not null;comment:创建类型" json:"create_time"`       // 创建类型
	UpdateTime   *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                // 更新时间
	DeleteTime   *time.Time `gorm:"column:delete_time;comment:删除时间" json:"delete_time"`                // 删除时间
	LastPushTime *time.Time `gorm:"column:last_push_time;comment:最后一次推送时间" json:"last_push_time"`      // 最后一次推送时间
	ErrCount     *int32     `gorm:"column:err_count;default:1;comment:联系推送错误次数" json:"err_count"`      // 联系推送错误次数
	InactiveTime *time.Time `gorm:"column:inactive_time;comment:标记不活跃时间" json:"inactive_time"`	//标记不活跃时间
}

// TableName MessagePushManage's table name
func (*MessagePushManage) TableName() string {
	return TableNameMessagePushManage
}
