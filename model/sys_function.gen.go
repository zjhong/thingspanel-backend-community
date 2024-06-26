// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSysFunction = "sys_function"

// SysFunction mapped from table <sys_function>
type SysFunction struct {
	ID          string  `gorm:"column:id;primaryKey;comment:id" json:"id"`                                        // id
	Name        string  `gorm:"column:name;not null;comment:功能名称" json:"name"`                                    // 功能名称
	EnableFlag  string  `gorm:"column:enable_flag;not null;comment:启用标志 enable-启用 disable-禁用" json:"enable_flag"` // 启用标志 enable-启用 disable-禁用
	Description *string `gorm:"column:description;comment:描述" json:"description"`                                 // 描述
	Remark      *string `gorm:"column:remark;comment:备注" json:"remark"`                                           // 备注
}

// TableName SysFunction's table name
func (*SysFunction) TableName() string {
	return TableNameSysFunction
}
