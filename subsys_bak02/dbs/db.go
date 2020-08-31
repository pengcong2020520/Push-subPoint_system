package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"subsys/config"
)
type (
	Account struct {
		UserId string `json:"userid"`
		Password string `json:"password"`
		Point int `json:"point"`
		Log [12]LogInfo `json:"log"`
	}
	Admin struct {
		BcosAcc string `json:"bcosacc"`
		Password string `json:"password"`
	}
	TokenMB struct {
		ToUserid string `json:"toaddr"`
		Value int `json:"value"`
		Timestamp string `timestamp`
	}
	Transfer struct {
		FromUserid string `json:"fromaddr"`
		ToUserid string `json:"toaddr"`
		Value int `json:"value"`
		Timestamp string `timestamp`
	}
	LogInfo struct {
		Type string `json:"type"`
		Subtype string `json:"subtype"`
		Number int `json:"number"`
		Using int `json:"using"`
		IsImport bool `json:"isimport"`
	}
	Logs struct {
		UserId string
		Month int
		Loginfo string
	}
)

//数据库连接的全局变量
var DBConn *sql.DB

func DBSetup() {
	fmt.Println("call dbs.Init", config.Config)
	DBConn = InitDB(config.Config.Db.Connstr, config.Config.Db.Driver)
	//DBConn = InitDB("mysql", "root:root@tcp(127.0.0.1:3306)/copyright?charset=utf8")
}

//初始化数据库连接
func InitDB(connstr, Driver string) *sql.DB {
	db, err := sql.Open(Driver, connstr)
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return db
}

//通用查询，返回map嵌套map
func DBQuery(sql string) ([]map[string]string, int, error) {
	fmt.Println("query is called:", sql)
	rows, err := DBConn.Query(sql)
	if err != nil {
		fmt.Println("query data err", err)
		return nil, 0, err
	}
	//得到列名数组
	cols, err := rows.Columns()
	//获取列的个数
	colCount := len(cols)
	values := make([]string, colCount)
	oneRows := make([]interface{}, colCount)
	for k, _ := range values {
		oneRows[k] = &values[k] //将查询结果的返回地址绑定，这样才能变参获取数据
	}
	//存储最终结果
	results := make([]map[string]string, 1)
	idx := 0
	//循环处理结果集
	for rows.Next() {
		rows.Scan(oneRows...)
		rowmap := make(map[string]string)
		for k, v := range values {
			rowmap[cols[k]] = v

		}
		if idx > 0 {
			results = append(results, rowmap)
		} else {
			results[0] = rowmap
		}
		idx++
		//fmt.Println(values)
	}
	//fmt.Println("---------------------------------------")
	fmt.Println("query..idx===", idx)
	return results, idx, nil

}
func Create(sql string) (int64, error) {
	res, err := DBConn.Exec(sql)
	if err != nil {
		fmt.Println("exec sql err,", err, "sql is ", sql)
		return -1, err
	}
	return res.LastInsertId()
}
