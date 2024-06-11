// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDeviceModelTelemetry = "device_model_telemetry"

// DeviceModelTelemetry mapped from table <device_model_telemetry>
type DeviceModelTelemetry struct {
	ID               string    `gorm:"column:id;primaryKey;comment:id" json:"id"`                                   // id
	DeviceTemplateID string    `gorm:"column:device_template_id;not null;comment:设备模板id" json:"device_template_id"` // 设备模板id
	DataName         *string   `gorm:"column:data_name;comment:数据名称" json:"data_name"`                              // 数据名称
	DataIdentifier   string    `gorm:"column:data_identifier;not null;comment:数据标识符" json:"data_identifier"`        // 数据标识符
	ReadWriteFlag    *string   `gorm:"column:read_write_flag;comment:读写标志R-读 W-写 RW-读写" json:"read_write_flag"`     // 读写标志R-读 W-写 RW-读写
	DataType         *string   `gorm:"column:data_type;comment:数据类型String Number Boolean" json:"data_type"`         // 数据类型String Number Boolean
	Unit             *string   `gorm:"column:unit;comment:单位" json:"unit"`                                          // 单位
	Description      *string   `gorm:"column:description;comment:描述" json:"description"`                            // 描述
	AdditionalInfo   *string   `gorm:"column:additional_info;comment:附加信息" json:"additional_info"`                  // 附加信息
	CreatedAt        time.Time `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                   // 创建时间
	UpdatedAt        time.Time `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`                   // 更新时间
	Remark           *string   `gorm:"column:remark;comment:备注" json:"remark"`                                      // 备注
	TenantID         string    `gorm:"column:tenant_id;not null" json:"tenant_id"`
}

// TableName DeviceModelTelemetry's table name
func (*DeviceModelTelemetry) TableName() string {
	return TableNameDeviceModelTelemetry
}