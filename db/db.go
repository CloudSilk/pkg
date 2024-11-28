package db

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DBClientInterface interface {
	Close()
	DB() *gorm.DB
	PageQuery(db *gorm.DB, pageSize, pageIndex int64, order string, result, selects interface{}) (records int64, pages int64, err error)
	PageQueryWithPreload(db *gorm.DB, pageSize, pageIndex int64, order string, preload []string, result interface{}) (records int64, pages int64, err error)
	PageQueryWithAssociations(db *gorm.DB, pageSize, pageIndex int64, order string, result interface{}) (records int64, pages int64, err error)
	CheckDuplication(db *gorm.DB, query interface{}, args ...interface{}) (bool, error)
	CheckDuplicationByTableName(db *gorm.DB, tableName string, query string, args ...interface{}) (bool, error)
	CreateWithCheckDuplication(info, query interface{}, args ...interface{}) (bool, error)
	CreateWithCheckDuplicationWithDB(db *gorm.DB, info, query interface{}, args ...interface{}) (bool, error)
	CreateWithCheckDuplicationByTableName(tableName string, info, query interface{}, args ...interface{}) (bool, error)
	UpdateWithCheckDuplication(
		db *gorm.DB,
		info interface{},
		fullSaveAssociations bool,
		checkDuplicationQuery interface{},
		checkDuplicationParams ...interface{}) (bool, error)
	UpdateWithCheckDuplicationAndOmit(
		db *gorm.DB, info interface{},
		fullSaveAssociations bool,
		omit []string,
		checkDuplicationQuery interface{},
		checkDuplicationParams ...interface{}) (bool, error)
	UpdateWithCheckDuplicationByTableName(db *gorm.DB, tableName string, info, query interface{}, args ...interface{}) (bool, error)
}

// DBClient DBClient
type DBClient struct {
	db    *gorm.DB
	debug bool
}

// NewDBClient New DBClient
func NewDBClient(db *gorm.DB, debug bool) *DBClient {

	return &DBClient{
		db:    db,
		debug: debug,
	}
}

// Close 关闭
func (m *DBClient) Close() {
}

// DB DB
func (m *DBClient) DB() *gorm.DB {
	if m.debug {
		return m.db.Session(&gorm.Session{}).Debug()
	}
	return m.db.Session(&gorm.Session{})
}

// PageQuery 分页查询
func (m *DBClient) PageQuery(db *gorm.DB, pageSize, pageIndex int64, order string, result, selects interface{}) (records int64, pages int64, err error) {
	err = db.Count(&records).Error
	if err != nil {
		return
	}
	if records == 0 {
		return
	}
	if pageSize == 0 {
		pageSize = 10
	}
	if pageIndex <= 0 {
		pageIndex = 1
	}
	pages = records / pageSize
	if records%pageSize > 0 {
		pages++
	}

	offset := pageSize * (pageIndex - 1)
	db = db.Order(order).Offset(int(offset)).Limit(int(pageSize))
	if selects != nil {
		db = db.Select(selects)
	}
	db = db.Find(result)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return
	}
	err = db.Error
	return
}

// PageQueryWithPreload 分页查询
func (m *DBClient) PageQueryWithPreload(db *gorm.DB, pageSize, pageIndex int64, order string, preload []string, result interface{}) (records int64, pages int64, err error) {
	err = db.Count(&records).Error
	if err != nil {
		return
	}
	if records == 0 {
		return
	}
	if pageSize == 0 {
		pageSize = 10
	}
	if pageIndex <= 0 {
		pageIndex = 1
	}
	pages = records / pageSize
	if records%pageSize > 0 {
		pages++
	}

	offset := pageSize * (pageIndex - 1)
	db = db.Offset(int(offset)).Limit(int(pageSize))
	for _, s := range preload {
		db = db.Preload(s)
	}
	db = db.Order(order)
	db = db.Find(result)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return
	}
	err = db.Error
	return
}

// PageQueryWithAssociations 分页查询
func (m *DBClient) PageQueryWithAssociations(db *gorm.DB, pageSize, pageIndex int64, order string, result interface{}) (records int64, pages int64, err error) {
	err = db.Count(&records).Error
	if err != nil {
		return
	}
	if records == 0 {
		return
	}
	if pageSize == 0 {
		pageSize = 10
	}
	if pageIndex <= 0 {
		pageIndex = 1
	}
	pages = records / pageSize
	if records%pageSize > 0 {
		pages++
	}

	offset := pageSize * (pageIndex - 1)
	db = db.Order(order)
	db = db.Offset(int(offset)).Limit(int(pageSize))

	db = db.Preload(clause.Associations).Find(result)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return
	}
	err = db.Error
	return
}

func (m *DBClient) CheckDuplication(db *gorm.DB, query interface{}, args ...interface{}) (bool, error) {
	var count int64
	err := db.Where(query, args...).Count(&count).Error
	if err != nil {
		return true, err
	}
	fmt.Println("count=", count)
	return count > 0, nil
}

func (m *DBClient) CheckDuplicationByTableName(db *gorm.DB, tableName string, query string, args ...interface{}) (bool, error) {
	var result = make(map[string]int64)
	err := db.Raw(fmt.Sprintf("select count(1) as count from %s where %s", tableName, query), args...).Scan(&result).Error
	if err != nil {
		return true, err
	}
	return result["count"] > 0, nil
}

func (m *DBClient) CreateWithCheckDuplication(info, query interface{}, args ...interface{}) (bool, error) {
	var count int64
	err := m.DB().Model(info).Where(query, args...).Count(&count).Error
	if err != nil {
		return true, err
	}

	if count > 0 {
		return true, nil
	}
	err = m.DB().Create(info).Error
	return false, err
}

func (m *DBClient) CreateWithCheckDuplicationWithDB(db *gorm.DB, info, query interface{}, args ...interface{}) (bool, error) {
	var count int64
	err := db.Model(info).Where(query, args...).Count(&count).Error
	if err != nil {
		return true, err
	}

	if count > 0 {
		return true, nil
	}
	err = db.Create(info).Error
	return false, err
}

func (m *DBClient) CreateWithCheckDuplicationByTableName(tableName string, info, query interface{}, args ...interface{}) (bool, error) {
	var count int64
	err := m.DB().Table(tableName).Where(query, args...).Count(&count).Error
	if err != nil {
		return true, err
	}
	if count > 0 {
		return true, nil
	}
	err = m.DB().Create(info).Error
	return false, err
}

func (m *DBClient) UpdateWithCheckDuplication(
	db *gorm.DB,
	info interface{},
	fullSaveAssociations bool,
	checkDuplicationQuery interface{},
	checkDuplicationParams ...interface{}) (bool, error) {
	return m.UpdateWithCheckDuplicationAndOmit(db, info, fullSaveAssociations, []string{}, checkDuplicationQuery, checkDuplicationParams...)
}

func (m *DBClient) UpdateWithCheckDuplicationAndOmit(
	db *gorm.DB, info interface{},
	fullSaveAssociations bool,
	omit []string,
	checkDuplicationQuery interface{},
	checkDuplicationParams ...interface{}) (bool, error) {
	var count int64
	err := db.Model(info).Where(checkDuplicationQuery, checkDuplicationParams...).Count(&count).Error
	if err != nil {
		return true, err
	}
	if count > 0 {
		return true, nil
	}
	db = db.Omit(omit...)
	if fullSaveAssociations {
		err = db.Session(&gorm.Session{FullSaveAssociations: true}).Save(info).Error
	} else {
		err = db.Save(info).Error
	}

	return false, err
}

func (m *DBClient) UpdateWithCheckDuplicationByTableName(db *gorm.DB, tableName string, info, query interface{}, args ...interface{}) (bool, error) {
	var count int64
	err := db.Table(tableName).Where(query, args...).Count(&count).Error
	if err != nil {
		return true, err
	}
	if count > 0 {
		return true, nil
	}
	err = db.Session(&gorm.Session{FullSaveAssociations: true}).Save(info).Error
	return false, err
}
