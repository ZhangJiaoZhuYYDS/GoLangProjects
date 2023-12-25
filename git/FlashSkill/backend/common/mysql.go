// @Author zhangjiaozhu 2023/10/14 17:33:00
package common

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func NewMysqlConn() (*sql.DB,error) {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/skill?charset=utf8")
	if err != nil {
		return nil,err
	}
	// 设置连接池大小
	db.SetMaxIdleConns(10) // 最大空闲连接数
	db.SetMaxOpenConns(100) // 最大打开连接数
	DB = db
	return db,nil
}

// 获取返回值，获取一条
func GetResultRow(rows *sql.Rows)  map[string]string{
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{},len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	record := make(map[string]string)
	for rows.Next() {
		rows.Scan(scanArgs...)
		for i , v := range values {
			if v != nil {
				record[columns[i]] = string(v.([]byte))
			}
		}
	}
	return record
}
// 多条查询
func GetAllResultRows(rows *sql.Rows) map[int]map[string]string  {
	columns, _ := rows.Columns()
	values := make([][]byte, len(columns))
	scans := make([]interface{},len(columns))
	for i,_ := range values {
		scans[i] = &values[i]
	}
	i := 0
	results := make(map[int]map[string]string)
	for rows.Next() {
		result := make(map[string]string)
		err := rows.Scan(scans...)
		if err != nil {
			log.Println(err)
			return nil
		}
		for i , v := range values {
			if v != nil {
				result[columns[i]] = string(v)
			}
		}
		results[i] = result
		i++
	}
	return results
}