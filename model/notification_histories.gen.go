// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameNotificationHistory = "notification_histories"

// NotificationHistory mapped from table <notification_histories>
type NotificationHistory struct {
	ID               string    `gorm:"column:id;primaryKey" json:"id"`
	SendTime         time.Time `gorm:"column:send_time;not null;comment:发送时间" json:"send_time"`                                                                           // 发送时间
	SendContent      *string   `gorm:"column:send_content;comment:发送内容" json:"send_content"`                                                                              // 发送内容
	SendTarget       string    `gorm:"column:send_target;not null;comment:发送目标" json:"send_target"`                                                                       // 发送目标
	SendResult       *string   `gorm:"column:send_result;comment:发送结果SUCCESS-成功FAILURE-失败" json:"send_result"`                                                            // 发送结果SUCCESS-成功FAILURE-失败
	NotificationType string    `gorm:"column:notification_type;not null;comment:通知类型MEMBER-成员通知 EMAIL-邮箱通知 SME-短信通知 VOICE-语音通知 WEBHOOK-webhook" json:"notification_type"` // 通知类型MEMBER-成员通知 EMAIL-邮箱通知 SME-短信通知 VOICE-语音通知 WEBHOOK-webhook
	TenantID         string    `gorm:"column:tenant_id;not null;comment:租户id" json:"tenant_id"`                                                                           // 租户id
	Remark           *string   `gorm:"column:remark;comment:备注" json:"remark"`                                                                                            // 备注
}

// TableName NotificationHistory's table name
func (*NotificationHistory) TableName() string {
	return TableNameNotificationHistory
}
