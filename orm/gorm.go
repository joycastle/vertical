package orm

import (
	"database/sql"

	"github.com/joycastle/vertical/connector"
	"gorm.io/gorm"
)

type GormSwitch struct {
	Node   string
	Master *gorm.DB
	Slave  *gorm.DB
}

func NewGormSwitch(node string) *GormSwitch {
	gs := &GormSwitch{
		Node:   node,
		Master: connector.GetMysqlMaster(node),
		Slave:  connector.GetMysqlSlave(node),
	}
	return gs
}

func (gs *GormSwitch) GetMasterDB() *gorm.DB {
	return gs.Master
}
func (gs *GormSwitch) GetSlaveDB() *gorm.DB {
	return gs.Slave
}

func (gs *GormSwitch) Create(value interface{}) *gorm.DB {
	return gs.Master.Create(value)
}
func (gs *GormSwitch) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return gs.Master.CreateInBatches(value, batchSize)
}
func (gs *GormSwitch) Save(value interface{}) *gorm.DB {
	return gs.Master.Save(value)
}
func (gs *GormSwitch) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return gs.Slave.First(dest, conds...)
}
func (gs *GormSwitch) Take(dest interface{}, conds ...interface{}) *gorm.DB {
	return gs.Slave.Take(dest, conds...)
}
func (gs *GormSwitch) Last(dest interface{}, conds ...interface{}) *gorm.DB {
	return gs.Slave.Last(dest, conds...)
}
func (gs *GormSwitch) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return gs.Slave.Find(dest, conds...)
}
func (gs *GormSwitch) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return gs.Slave.FindInBatches(dest, batchSize, fc)
}
func (gs *GormSwitch) FirstOrInit(dest interface{}, conds ...interface{}) *gorm.DB {
	return gs.Master.FirstOrInit(dest, conds...)
}
func (gs *GormSwitch) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return gs.Master.FirstOrCreate(dest, conds...)
}
func (gs *GormSwitch) Update(column string, value interface{}) *gorm.DB {
	return gs.Master.Update(column, value)
}
func (gs *GormSwitch) Updates(values interface{}) *gorm.DB {
	return gs.Master.Updates(values)
}
func (gs *GormSwitch) UpdateColumn(column string, value interface{}) *gorm.DB {
	return gs.Master.UpdateColumn(column, value)
}
func (gs *GormSwitch) UpdateColumns(values interface{}) *gorm.DB {
	return gs.Master.UpdateColumns(values)
}
func (gs *GormSwitch) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return gs.Master.Delete(value, conds...)
}
func (gs *GormSwitch) Count(count *int64) *gorm.DB {
	return gs.Slave.Count(count)
}
func (gs *GormSwitch) Row() *sql.Row {
	return gs.Slave.Row()
}
func (gs *GormSwitch) Rows() (*sql.Rows, error) {
	return gs.Slave.Rows()
}
func (gs *GormSwitch) Connection(fc func(*gorm.DB) error) (err error) {
	return gs.Master.Connection(fc)
}
func (gs *GormSwitch) Transaction(fc func(*gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return gs.Master.Transaction(fc, opts...)
}
func (gs *GormSwitch) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return gs.Master.Begin(opts...)
}
func (gs *GormSwitch) Commit() *gorm.DB {
	return gs.Master.Commit()
}
func (gs *GormSwitch) Rollback() *gorm.DB {
	return gs.Master.Rollback()
}
func (gs *GormSwitch) SavePoint(name string) *gorm.DB {
	return gs.Master.SavePoint(name)
}
func (gs *GormSwitch) RollbackTo(name string) *gorm.DB {
	return gs.Master.RollbackTo(name)
}
func (gs *GormSwitch) Exec(sql string, values ...interface{}) *gorm.DB {
	return gs.Master.Exec(sql, values)
}
