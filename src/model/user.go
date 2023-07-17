package model

import (
	"auth-cert/src/db/mysql"
	"fmt"
	"log"
)

// 引入gorm 与mysql  go get -u gorm.io/gorm   go get -u gorm.io/driver/mysql
type User struct { //定义实体
	Id         int     `json:"id" gorm:"column:id"`
	UName      *string `json:"uname" gorm:"column:uname"`
	PWD        *string `json:"pwd" gorm:"pwd"`
	CreateTime *string `json:"createtime" gorm:"createtime"`
}

// 映射表名
func (u *User) TableNam() string {
	return "user"
}

// 新增方法
func (u *User) Insert() int {
	//插入语法
	rs, err := mysql.SqlDB.Exec("insert into User(username)value(?,?)", u.UName)
	//错误打印
	if err != nil {
		log.Fatal(err)
	}
	//更新条数
	res, err := rs.RowsAffected()
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}
	//返回新增id
	return u.Id
}

// 查询其中某一条记录
func (u *User) First() (user User, err error) {
	user = User{}
	err = mysql.SqlDB.QueryRow("select * from User where id = ?", u.Id).Scan(
		&user.Id, &user.UName, &user.PWD, &user.CreateTime)
	return
}

// 查询全部记录
func (u *User) All() (users []User, err error) {
	rows, err := mysql.SqlDB.Query("select * from User")
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.UName, &u.PWD, &u.CreateTime)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	rows.Close()
	return
}

// 修改更新
func (u *User) Update() int {
	rs, err := mysql.SqlDB.Exec("update User set uname=?  where id=?", u.UName, u.Id)
	if err != nil {
		log.Fatal(err)
	}
	row, err := rs.RowsAffected()
	fmt.Println(row)
	if err != nil {
		log.Fatal(err)
	}
	return u.Id
}

// 删除
func (u *User) Deleted() int64 {
	rs, err := mysql.SqlDB.Exec("delete from User where id=?", u.Id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return rows

}
