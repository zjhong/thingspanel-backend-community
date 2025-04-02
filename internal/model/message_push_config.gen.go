// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMessagePushConfig = "message_push_config"

// MessagePushConfig mapped from table <message_push_config>
type MessagePushConfig struct {
	ID         string     `gorm:"column:id;primaryKey" json:"id"`
	URL        string     `gorm:"column:url;not null;comment:推送地址" json:"url"`                                  // 推送地址
	ConfigType int16      `gorm:"column:config_type;not null;default:1;comment:配置类型 1 推送地址" json:"config_type"` // 配置类型 1 推送地址
	CreateTime time.Time  `gorm:"column:create_time;not null;comment:创建时间" json:"create_time"`                  // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                           // 更新时间
}

// TableName MessagePushConfig's table name
func (*MessagePushConfig) TableName() string {
	return TableNameMessagePushConfig
}
