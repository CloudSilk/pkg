package postgres

import (
	"testing"

	"github.com/CloudSilk/pkg/db"
	"gorm.io/gorm"
)

var pg db.DBClientInterface

func init() {
	pg = NewPostgres("postgresql://postgres:123456@127.0.0.1:5432/test", true)
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
	_, _, err := pg.PageQuery(pg.DB().Model(&Location{}), 10, 1, "name", &locations)
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

type Location struct {
	gorm.Model
	Name     string
	Code     string
	ParentID uint
	Tags     []LocationTag
	Files    []LocationFile
}

type LocationTag struct {
	gorm.Model
	LocationID uint
	Name       string
}

type LocationFile struct {
	gorm.Model
	LocationID uint
	FileID     string
}
