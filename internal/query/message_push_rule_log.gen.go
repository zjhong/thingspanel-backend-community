// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"project/internal/model"
)

func newMessagePushRuleLog(db *gorm.DB, opts ...gen.DOOption) messagePushRuleLog {
	_messagePushRuleLog := messagePushRuleLog{}

	_messagePushRuleLog.messagePushRuleLogDo.UseDB(db, opts...)
	_messagePushRuleLog.messagePushRuleLogDo.UseModel(&model.MessagePushRuleLog{})

	tableName := _messagePushRuleLog.messagePushRuleLogDo.TableName()
	_messagePushRuleLog.ALL = field.NewAsterisk(tableName)
	_messagePushRuleLog.ID = field.NewString(tableName, "id")
	_messagePushRuleLog.UserID = field.NewString(tableName, "user_id")
	_messagePushRuleLog.PushID = field.NewString(tableName, "push_id")
	_messagePushRuleLog.Type = field.NewInt16(tableName, "type")

	_messagePushRuleLog.fillFieldMap()

	return _messagePushRuleLog
}

type messagePushRuleLog struct {
	messagePushRuleLogDo

	ALL    field.Asterisk
	ID     field.String
	UserID field.String
	PushID field.String
	Type   field.Int16 // 1 主动失效 2被动失效 3定时任务 4自动清理

	fieldMap map[string]field.Expr
}

func (m messagePushRuleLog) Table(newTableName string) *messagePushRuleLog {
	m.messagePushRuleLogDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m messagePushRuleLog) As(alias string) *messagePushRuleLog {
	m.messagePushRuleLogDo.DO = *(m.messagePushRuleLogDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *messagePushRuleLog) updateTableName(table string) *messagePushRuleLog {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewString(table, "id")
	m.UserID = field.NewString(table, "user_id")
	m.PushID = field.NewString(table, "push_id")
	m.Type = field.NewInt16(table, "type")

	m.fillFieldMap()

	return m
}

func (m *messagePushRuleLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *messagePushRuleLog) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 4)
	m.fieldMap["id"] = m.ID
	m.fieldMap["user_id"] = m.UserID
	m.fieldMap["push_id"] = m.PushID
	m.fieldMap["type"] = m.Type
}

func (m messagePushRuleLog) clone(db *gorm.DB) messagePushRuleLog {
	m.messagePushRuleLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m messagePushRuleLog) replaceDB(db *gorm.DB) messagePushRuleLog {
	m.messagePushRuleLogDo.ReplaceDB(db)
	return m
}

type messagePushRuleLogDo struct{ gen.DO }

type IMessagePushRuleLogDo interface {
	gen.SubQuery
	Debug() IMessagePushRuleLogDo
	WithContext(ctx context.Context) IMessagePushRuleLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMessagePushRuleLogDo
	WriteDB() IMessagePushRuleLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMessagePushRuleLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMessagePushRuleLogDo
	Not(conds ...gen.Condition) IMessagePushRuleLogDo
	Or(conds ...gen.Condition) IMessagePushRuleLogDo
	Select(conds ...field.Expr) IMessagePushRuleLogDo
	Where(conds ...gen.Condition) IMessagePushRuleLogDo
	Order(conds ...field.Expr) IMessagePushRuleLogDo
	Distinct(cols ...field.Expr) IMessagePushRuleLogDo
	Omit(cols ...field.Expr) IMessagePushRuleLogDo
	Join(table schema.Tabler, on ...field.Expr) IMessagePushRuleLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMessagePushRuleLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMessagePushRuleLogDo
	Group(cols ...field.Expr) IMessagePushRuleLogDo
	Having(conds ...gen.Condition) IMessagePushRuleLogDo
	Limit(limit int) IMessagePushRuleLogDo
	Offset(offset int) IMessagePushRuleLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMessagePushRuleLogDo
	Unscoped() IMessagePushRuleLogDo
	Create(values ...*model.MessagePushRuleLog) error
	CreateInBatches(values []*model.MessagePushRuleLog, batchSize int) error
	Save(values ...*model.MessagePushRuleLog) error
	First() (*model.MessagePushRuleLog, error)
	Take() (*model.MessagePushRuleLog, error)
	Last() (*model.MessagePushRuleLog, error)
	Find() ([]*model.MessagePushRuleLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MessagePushRuleLog, err error)
	FindInBatches(result *[]*model.MessagePushRuleLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MessagePushRuleLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMessagePushRuleLogDo
	Assign(attrs ...field.AssignExpr) IMessagePushRuleLogDo
	Joins(fields ...field.RelationField) IMessagePushRuleLogDo
	Preload(fields ...field.RelationField) IMessagePushRuleLogDo
	FirstOrInit() (*model.MessagePushRuleLog, error)
	FirstOrCreate() (*model.MessagePushRuleLog, error)
	FindByPage(offset int, limit int) (result []*model.MessagePushRuleLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMessagePushRuleLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m messagePushRuleLogDo) Debug() IMessagePushRuleLogDo {
	return m.withDO(m.DO.Debug())
}

func (m messagePushRuleLogDo) WithContext(ctx context.Context) IMessagePushRuleLogDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m messagePushRuleLogDo) ReadDB() IMessagePushRuleLogDo {
	return m.Clauses(dbresolver.Read)
}

func (m messagePushRuleLogDo) WriteDB() IMessagePushRuleLogDo {
	return m.Clauses(dbresolver.Write)
}

func (m messagePushRuleLogDo) Session(config *gorm.Session) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Session(config))
}

func (m messagePushRuleLogDo) Clauses(conds ...clause.Expression) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m messagePushRuleLogDo) Returning(value interface{}, columns ...string) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m messagePushRuleLogDo) Not(conds ...gen.Condition) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m messagePushRuleLogDo) Or(conds ...gen.Condition) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m messagePushRuleLogDo) Select(conds ...field.Expr) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m messagePushRuleLogDo) Where(conds ...gen.Condition) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m messagePushRuleLogDo) Order(conds ...field.Expr) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m messagePushRuleLogDo) Distinct(cols ...field.Expr) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m messagePushRuleLogDo) Omit(cols ...field.Expr) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m messagePushRuleLogDo) Join(table schema.Tabler, on ...field.Expr) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m messagePushRuleLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMessagePushRuleLogDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m messagePushRuleLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IMessagePushRuleLogDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m messagePushRuleLogDo) Group(cols ...field.Expr) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m messagePushRuleLogDo) Having(conds ...gen.Condition) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m messagePushRuleLogDo) Limit(limit int) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m messagePushRuleLogDo) Offset(offset int) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m messagePushRuleLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m messagePushRuleLogDo) Unscoped() IMessagePushRuleLogDo {
	return m.withDO(m.DO.Unscoped())
}

func (m messagePushRuleLogDo) Create(values ...*model.MessagePushRuleLog) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m messagePushRuleLogDo) CreateInBatches(values []*model.MessagePushRuleLog, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m messagePushRuleLogDo) Save(values ...*model.MessagePushRuleLog) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m messagePushRuleLogDo) First() (*model.MessagePushRuleLog, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessagePushRuleLog), nil
	}
}

func (m messagePushRuleLogDo) Take() (*model.MessagePushRuleLog, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessagePushRuleLog), nil
	}
}

func (m messagePushRuleLogDo) Last() (*model.MessagePushRuleLog, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessagePushRuleLog), nil
	}
}

func (m messagePushRuleLogDo) Find() ([]*model.MessagePushRuleLog, error) {
	result, err := m.DO.Find()
	return result.([]*model.MessagePushRuleLog), err
}

func (m messagePushRuleLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MessagePushRuleLog, err error) {
	buf := make([]*model.MessagePushRuleLog, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m messagePushRuleLogDo) FindInBatches(result *[]*model.MessagePushRuleLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m messagePushRuleLogDo) Attrs(attrs ...field.AssignExpr) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m messagePushRuleLogDo) Assign(attrs ...field.AssignExpr) IMessagePushRuleLogDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m messagePushRuleLogDo) Joins(fields ...field.RelationField) IMessagePushRuleLogDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m messagePushRuleLogDo) Preload(fields ...field.RelationField) IMessagePushRuleLogDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m messagePushRuleLogDo) FirstOrInit() (*model.MessagePushRuleLog, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessagePushRuleLog), nil
	}
}

func (m messagePushRuleLogDo) FirstOrCreate() (*model.MessagePushRuleLog, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MessagePushRuleLog), nil
	}
}

func (m messagePushRuleLogDo) FindByPage(offset int, limit int) (result []*model.MessagePushRuleLog, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m messagePushRuleLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m messagePushRuleLogDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m messagePushRuleLogDo) Delete(models ...*model.MessagePushRuleLog) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *messagePushRuleLogDo) withDO(do gen.Dao) *messagePushRuleLogDo {
	m.DO = *do.(*gen.DO)
	return m
}
