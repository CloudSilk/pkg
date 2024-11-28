package postgres

import (
	"fmt"
	"testing"
	"time"

	"github.com/CloudSilk/pkg/db"
	"gorm.io/gorm"
)

var pg db.DBClientInterface

func init() {
	// pg = NewPostgres("postgresql://postgres:123456@127.0.0.1:5432/test", true)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		"127.0.0.1",
		"admin",
		"quest",
		"station",
		"8812",
	)
	pg = NewPostgres(dsn, true)
}

type Tracker struct {
	Timestamp time.Time `gorm:"type:timestamp" json:"timestamp"`
	VehicleId int       `gorm:"type:int" json:"vehicleId"`
	Latitude  float64   `gorm:"type:double" json:"latitude"`
	Longitude float64   `gorm:"type:double" json:"longitude"`
}

func TestQuestdb(t *testing.T) {
	pg.DB().AutoMigrate(&Tracker{})
	pg.DB().Create(&Tracker{
		Timestamp: time.Now().UTC(),
		VehicleId: 1,
		Latitude:  -7.626923,
		Longitude: 111.5213978,
	})
	var list []Tracker
	records, pages, err := pg.PageQuery(pg.DB().Model(&Tracker{}), 10, 1, "", &list, "*")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(records, pages)
}
func TestAutoMigrate(t *testing.T) {
	pg.DB().AutoMigrate(&Location{}, &LocationTag{}, &LocationFile{})
}

func TestAdd(t *testing.T) {
	l := &Location{
		Name:     "B02",
		Code:     "1001",
		ParentID: 1,
		Tags: []LocationTag{
			{
				Name: "tag1",
			}, {
				Name: "tag2",
			},
		},
	}
	ok, err := pg.CreateWithCheckDuplication(l, "name=?", l.Name)

	if err != nil {
		t.Fatal(err)
	}

	if ok {
		t.Logf("名称(%s)已经存在", l.Name)
		return
	}

	l.Name = "mock_update2"
	l.Tags[0].Name = "tag3"
	l.Tags = append(l.Tags, LocationTag{
		Name: "tag4",
	})

	// 同时更新关联表
	ok, err = pg.UpdateWithCheckDuplication(pg.DB(), l, true, "name=?", l.Name)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Logf("名称(%s)已经存在", l.Name)
	}
}

func TestQueryPage(t *testing.T) {
	var locations []Location
	_, _, err := pg.PageQuery(pg.DB().Model(&Location{}), 10, 1, "name", &locations, "*")
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = pg.PageQueryWithPreload(pg.DB().Model(&Location{}), 10, 1, "name", []string{"Tags", "Files"}, &locations)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = pg.PageQueryWithAssociations(pg.DB().Model(&Location{}), 10, 1, "name", &locations)
	if err != nil {
		t.Fatal(err)
	}
}

type Model struct {
	ID        string         `gorm:"type:varchar"`
	CreatedAt time.Time      `gorm:"type:timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp"`
}
type Location struct {
	Model
	Name     string
	Code     string
	ParentID uint
	Tags     []LocationTag
	Files    []LocationFile
}

type LocationTag struct {
	Model
	LocationID uint
	Name       string
}

type LocationFile struct {
	Model
	LocationID uint
	FileID     string
}
