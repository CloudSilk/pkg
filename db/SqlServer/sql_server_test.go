package sqlserver

import (
	"encoding/json"
	"testing"
)

func TestMysql(t *testing.T) {
	dbClient := NewSqlServer2("sa", "ABC123def", "127.0.0.1", "tes", 1433, true)
	var result []map[string]interface{}
	err := dbClient.DB().Raw(`SELECT d.name as table_name,a.name AS column_name,a.length AS character_maximum_length
    , CASE
        WHEN (
            SELECT COUNT(*)
            FROM sysobjects
            WHERE name IN (
                    SELECT name
                    FROM sysindexes
                    WHERE id = a.id
                        AND indid IN (
                            SELECT indid
                            FROM sysindexkeys
                            WHERE id = a.id
                                AND colid IN (
                                    SELECT colid
                                    FROM syscolumns
                                    WHERE id = a.id
                                        AND name = a.name
                                )
                        )
                )
                AND xtype = 'PK'
        ) > 0 THEN 'YES'
        ELSE 'NO'
    END AS column_key, b.name AS data_type
    , CASE
        WHEN a.isnullable = 0 THEN 'YES'
        ELSE 'NO'
    END AS is_nullable
    , isnull(g.[value], '') AS column_comment,e.text as column_default
FROM syscolumns a
    LEFT JOIN systypes b ON a.xtype = b.xusertype
    INNER JOIN sysobjects d
    ON a.id = d.id
        AND d.xtype = 'U'
        AND d.name <> 'dtproperties'
    LEFT JOIN syscomments e on a.cdefault=e.id 
 
 left join sys.extended_properties g 
 on a.id=g.major_id AND a.colid= g.minor_id  
 where d.name=?
 order by a.id,a.colorder `, "Role").Find(&result).Error
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(result)
	t.Log(string(data))
	dbClient.Close()
}
