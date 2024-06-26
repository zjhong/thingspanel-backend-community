// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDeviceTemplate = "device_templates"

// DeviceTemplate mapped from table <device_templates>
type DeviceTemplate struct {
	ID             string    `gorm:"column:id;primaryKey;comment:Id" json:"id"`        // Id
	Name           string    `gorm:"column:name;not null;comment:模板名称" json:"name"`    // 模板名称
	Author         *string   `gorm:"column:author;comment:作者" json:"author"`           // 作者
	Version        *string   `gorm:"column:version;comment:版本号" json:"version"`        // 版本号
	Description    *string   `gorm:"column:description;comment:描述" json:"description"` // 描述
	TenantID       string    `gorm:"column:tenant_id;not null" json:"tenant_id"`
	CreatedAt      time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
	Flag           *int16    `gorm:"column:flag;default:1;comment:标志 默认1" json:"flag"`                // 标志 默认1
	Label          *string   `gorm:"column:label;comment:标签" json:"label"`                            // 标签
	WebChartConfig *string   `gorm:"column:web_chart_config;comment:web图表配置" json:"web_chart_config"` // web图表配置
	AppChartConfig *string   `gorm:"column:app_chart_config;comment:app图表配置" json:"app_chart_config"` // app图表配置
	Remark         *string   `gorm:"column:remark;comment:备注" json:"remark"`                          // 备注
	Path           *string   `gorm:"column:path;comment:图片路径" json:"path"`                            // 图片路径
}

// TableName DeviceTemplate's table name
func (*DeviceTemplate) TableName() string {
	return TableNameDeviceTemplate
}
