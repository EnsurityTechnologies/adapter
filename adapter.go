package adapter

import (
	"fmt"

	"github.com/EnsurityTechnologies/config"
	"github.com/EnsurityTechnologies/uuid"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

const (
	postgressDB string = "PostgressSQL"
	sqlDB       string = "SQLServer"
	mysqlDB     string = "MySQL"
)

// TenantIDStr ...
const TenantIDStr string = "TenantId"

// Adapter structer
type Adapter struct {
	db     *gorm.DB
	dbType string
}

// NewAdapter create new adapter
func NewAdapter(cfg *config.Config) (*Adapter, error) {

	var db *gorm.DB
	var err error
	switch cfg.DBType {
	case sqlDB:
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", cfg.DBUserName, cfg.DBPassword, cfg.DBAddress, cfg.DBPort, cfg.DBName)
		db, err = gorm.Open("mssql", dsn)
		break
	default:
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", cfg.DBUserName, cfg.DBPassword, cfg.DBAddress, cfg.DBPort, cfg.DBName)
		db, err = gorm.Open("mssql", dsn)
		break
	}

	if err != nil {
		fmt.Println("Error %v", err)
		return nil, err
	}
	adapter := &Adapter{
		db:     db,
		dbType: cfg.DBType,
	}

	return adapter, err
}

// InitTable Initialize table
func (adapter *Adapter) InitTable(tableName string, item interface{}) error {
	err := adapter.db.Table(tableName).AutoMigrate(item).Error
	return err
}

// InitTable Initialize table
func (adapter *Adapter) InitTwoTable(tableName string, item1 interface{}, item2 interface{}) error {
	err := adapter.db.Table(tableName).AutoMigrate(item1, item2).Error
	return err
}

// DropTable drop the table
func (adapter *Adapter) DropTable(tableName string) error {
	err := adapter.db.DropTable(tableName).Error
	return err
}

// // IsTableExist check whether table exist
// func (adapter *Adapter) IsTableExist(tableName string) bool {
// 	status := adapter.db.Table(tableName).
// 	return status
// }

// DropTable drop the table
func (adapter *Adapter) AddForienKey(tableName string, value interface{}, colStr string, tableStr string) error {
	err := adapter.db.Table(tableName).Model(value).AddForeignKey(colStr, tableStr, "RESTRICT", "RESTRICT").Error
	return err
}

// Delete function delete entry from the table
func (adapter *Adapter) Delete(tenantID interface{}, tableName string, format string, value interface{}, item interface{}) error {
	if tenantID != uuid.Nil {
		formatStr := TenantIDStr + "=? AND " + format
		err := adapter.db.Table(tableName).Where(formatStr, tenantID, value).Delete(item).Error
		return err
	} else {
		err := adapter.db.Table(tableName).Where(format, value).Delete(item).Error
		return err
	}
}

// Create creates and stores the new item in the table
func (adapter *Adapter) Create(tableName string, item interface{}) error {
	err := adapter.db.Table(tableName).Create(item).Error
	return err
}

// Find function finds the value from the table
func (adapter *Adapter) Find(tenantID interface{}, tableName string, format string, value interface{}, item interface{}) error {
	if tenantID != uuid.Nil {
		formatStr := TenantIDStr + "=? AND " + format
		err := adapter.db.Table(tableName).Where(formatStr, tenantID, value).Find(item).Error
		return err
	} else {
		err := adapter.db.Table(tableName).Where(format, value).Find(item).Error
		return err
	}
}

// FindMult function finds the value from the table
func (adapter *Adapter) FindMult(tenantID interface{}, tableName string, format1 string, format2 string, value1 interface{}, value2 interface{}, item interface{}) error {
	if tenantID != uuid.Nil {
		formatStr1 := TenantIDStr + "=? AND " + format1
		formatStr2 := TenantIDStr + "=? AND " + format2
		err := adapter.db.Table(tableName).Where(formatStr1, tenantID, value1).Or(formatStr2, tenantID, value2).Find(item).Error
		return err
	} else {
		err := adapter.db.Table(tableName).Where(format1, value1).Or(format2, value2).Find(item).Error
		return err
	}
}

// FindAnd function finds the value from the table
func (adapter *Adapter) FindAnd(tenantID interface{}, tableName string, format1 string, format2 string, value1 interface{}, value2 interface{}, item interface{}) error {
	if tenantID != uuid.Nil {
		formatStr1 := TenantIDStr + "=? AND " + format1
		formatStr2 := TenantIDStr + "=? AND " + format2
		err := adapter.db.Table(tableName).Where(formatStr1, tenantID, value1).Where(formatStr2, tenantID, value2).Find(item).Error
		return err
	} else {
		err := adapter.db.Table(tableName).Where(format1, value1).Or(format2, value2).Find(item).Error
		return err
	}
}

// FindA function finds the value from the table
func (adapter *Adapter) FindA(tenantID interface{}, tableName string, format string, value interface{}, item interface{}, item1 interface{}) error {
	if tenantID != uuid.Nil {
		formatStr := TenantIDStr + "=? AND " + format
		err := adapter.db.Table(tableName).Where(formatStr, tenantID, value).Find(item).Association("UserId").Find(item1).Error
		return err
	} else {
		err := adapter.db.Table(tableName).Where(format, value).Find(item).Error
		return err
	}
}

// Updates function updates the value in the table
func (adapter *Adapter) Updates(tenantID interface{}, tableName string, format string, value interface{}, item interface{}) error {
	if tenantID != uuid.Nil {
		formatStr := TenantIDStr + "=? AND " + format
		err := adapter.db.Table(tableName).Where(formatStr, tenantID, value).Updates(item).Error
		return err
	} else {
		err := adapter.db.Table(tableName).Where(format, value).Updates(item).Error
		return err
	}
}