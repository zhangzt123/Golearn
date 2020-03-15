/*go connect mysql*/
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//初始化连接数据库池
func init() {
	var err error
	//dsn格式 ->   [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//protocol包括
	//另外还包括超文本传输协议（HTTP）和文件传输协议（FTP），以及边界网关协议（BGP）和动态主机配置协议（DHCP），
	//简易邮件传输通讯协议（SMTP），邮局通讯协定（POP3）,网络通讯协议（TCP/IP）,用户数据报协议（UDP）
	dataSourceName := "zzt:Zzt=123456@tcp(192.168.118.130:3306)/nacosdb"
	db, err = sql.Open("mysql", dataSourceName) // db即获得的连接池
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(1000) // 设置做大连接时间
	db.SetMaxIdleConns(10)      //最大闲置连接数
	db.SetMaxOpenConns(100)     //最大连接数
	fmt.Println("db连接success")

}

type users struct {
	username string
	password string
	enabled  uint8
}

func main() {

	fmt.Println(db.Driver())
	fmt.Println(db.Stats())
	//查询
	//sel()
	//增删改
	upd()
	defer db.Close()

}

//查询
func sel() {

	//直接查询

	//无参
	rows, err := db.Query("select t.username,t.password,t.enabled from users t ")
	if err != nil {
		fmt.Println(err)
	}
	Columns, _ := rows.Columns()
	fmt.Println(Columns)
	ColumnType, _ := rows.ColumnTypes()
	fmt.Println(ColumnType)
	// xx := make([]string, 3)
	// for rows.Next() {
	// 	er := rows.Scan(&xx[0], &xx[1], &xx[2])
	// 	if er != nil {
	// 		fmt.Println(er)
	// 	} else {
	// 		fmt.Println(xx)
	// 	}
	// }
	var u users
	for rows.Next() {
		er := rows.Scan(&u.username, &u.password, &u.enabled)
		if er != nil {
			fmt.Println(er)
		} else {
			fmt.Println(u)
		}
	}

	//////////////////////////////////////  带参
	rows, err = db.Query("select t.username,t.password,t.enabled from users t  where t.username=?", "admin") //args 配合?使用 参考jdbc
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		er := rows.Scan(&u.username, &u.password, &u.enabled)
		if er != nil {
			fmt.Println(er)
		} else {
			fmt.Println(u)
		}
	}

	//预处理sql

	prep, _ := db.Prepare("select t.username,t.password,t.enabled from users t  where t.username=?")
	rows, err = prep.Query("admin")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		er := rows.Scan(&u.username, &u.password, &u.enabled)
		if er != nil {
			fmt.Println(er)
		} else {
			fmt.Printf("预处理: %v \n", u)
		}
	}
	defer prep.Close()

	//NOTE rows.Scan 多行查询， row.Scan 单行查询 单行查询不必调用Close（） 会自动调用
	defer rows.Close() //逆向关闭 prepstatement res 参考 java jdbc时候finally的操作

}

//增删改
func upd() {

	u := users{username: "dev", password: "dev", enabled: 1}

	//增
	ins := " insert into users values( ?,?,? ) "
	//删
	del := " delete from users where username =? and enabled=? "
	//改
	upd := " update users t  set t.enabled = ? where t.username=? "

	fmt.Printf("u:%v \n", u)
	prep, _ := db.Prepare(ins)
	res, err := prep.Exec(u.username, u.password, u.enabled)
	if err != nil {
		fmt.Println(err)
	}
	ins_InsertId, _ := res.LastInsertId() //插入的ID
	ins_Affected, _ := res.RowsAffected() //影响行数
	fmt.Printf("新增操作【 LastInsertId : %v RowsAffected : %v 】\n", ins_InsertId, ins_Affected)

	///////////////////////////

	prep, _ = db.Prepare(upd)
	res, err = prep.Exec(0, u.username)
	if err != nil {
		fmt.Println(err)
	}
	ins_InsertId, _ = res.LastInsertId() //插入的ID
	ins_Affected, _ = res.RowsAffected() //影响行数
	fmt.Printf("更新操作【 LastInsertId : %v RowsAffected : %v 】\n", ins_InsertId, ins_Affected)

	//////////////////////////////

	prep, _ = db.Prepare(del)
	res, err = prep.Exec(u.username, 0)
	if err != nil {
		fmt.Println(err)
	}
	ins_InsertId, _ = res.LastInsertId() //插入的ID
	ins_Affected, _ = res.RowsAffected() //影响行数
	fmt.Printf("删除操作【 LastInsertId : %v RowsAffected : %v 】\n", ins_InsertId, ins_Affected)

	///////////////////////////////

	tx, _ := db.Begin() //开启事务
	prep, _ = tx.Prepare(ins)
	res, err = prep.Exec(u.username, u.password, u.enabled)
	if err != nil {
		fmt.Println(err)
	}
	ins_InsertId, _ = res.LastInsertId() //插入的ID
	ins_Affected, _ = res.RowsAffected() //影响行数
	fmt.Printf("事务新增操作【 LastInsertId : %v RowsAffected : %v 】\n", ins_InsertId, ins_Affected)

	time.Sleep(10 * 1e9)
	//tx.Rollback()
	defer prep.Close()
	defer tx.Commit()
}
