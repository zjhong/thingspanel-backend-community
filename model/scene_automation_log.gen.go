// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSceneAutomationLog = "scene_automation_log"

// SceneAutomationLog mapped from table <scene_automation_log>
type SceneAutomationLog struct {
	SceneAutomationID string    `gorm:"column:scene_automation_id;not null;comment:场景联动ID（外键-关联删除）" json:"scene_automation_id"` // 场景联动ID（外键-关联删除）
	ExecutedAt        time.Time `gorm:"column:executed_at;not null;comment:执行时间" json:"executed_at"`                            // 执行时间
	Detail            string    `gorm:"column:detail;not null;comment:执行说明：详细的执行过程" json:"detail"`                              // 执行说明：详细的执行过程
	ExecutionResult   string    `gorm:"column:execution_result;not null;comment:执行状态S：成功F：失败 全部执行成功才算" json:"execution_result"` // 执行状态S：成功F：失败 全部执行成功才算
	TenantID          string    `gorm:"column:tenant_id;not null" json:"tenant_id"`
	Remark            *string   `gorm:"column:remark" json:"remark"`
}

// TableName SceneAutomationLog's table name
func (*SceneAutomationLog) TableName() string {
	return TableNameSceneAutomationLog
}