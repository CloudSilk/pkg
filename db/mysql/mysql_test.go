package mysql

import "testing"

func TestMysql(t *testing.T) {
	dbClient := NewMysql2("test", "test", "lcaolhost", "test", 3307, true)
	dbClient.Close()
}
