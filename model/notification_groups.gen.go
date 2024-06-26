// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameNotificationGroup = "notification_groups"

// NotificationGroup mapped from table <notification_groups>
type NotificationGroup struct {
	ID                 string    `gorm:"column:id;primaryKey" json:"id"`
	Name               string    `gorm:"column:name;not null;comment:名称" json:"name"`                                                                                       // 名称
	NotificationType   string    `gorm:"column:notification_type;not null;comment:通知类型MEMBER-成员通知 EMAIL-邮箱通知 SME-短信通知 VOICE-语音通知 WEBHOOK-webhook" json:"notification_type"` // 通知类型MEMBER-成员通知 EMAIL-邮箱通知 SME-短信通知 VOICE-语音通知 WEBHOOK-webhook
	Status             string    `gorm:"column:status;not null;comment:通知状态ON-启用 OFF-停用" json:"status"`                                                                     // 通知状态ON-启用 OFF-停用
	NotificationConfig *string   `gorm:"column:notification_config;comment:通知配置" json:"notification_config"`                                                                // 通知配置
	Description        *string   `gorm:"column:description;comment:描述" json:"description"`                                                                                  // 描述
	TenantID           string    `gorm:"column:tenant_id;not null;comment:租户id" json:"tenant_id"`                                                                           // 租户id
	CreatedAt          time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                                                                         // 创建时间
	UpdatedAt          time.Time `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`                                                                         // 更新时间
	Remark             *string   `gorm:"column:remark;comment:备注" json:"remark"`                                                                                            // 备注
}

// TableName NotificationGroup's table name
func (*NotificationGroup) TableName() string {
	return TableNameNotificationGroup
}
