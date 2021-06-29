package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" //导入mysql包
)

type Doctor struct {
	ID      int64
	Name    string
	Age     int
	Sex     int
	AddTime time.Time
}

func main() {
	//-------1、打开数据库--------
	db, err := sql.Open("mysql", "Toanle:sHf7D3KJkXyXwnM3@tcp(192.168.31.35:3306)/toanle?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("数据库链接错误", err)
		return
	}
	//延迟到函数结束关闭链接
	defer db.Close()

	//-------2、查询单条数据--------
	//定义接收数据的结构
	var doc Doctor
	//执行单条查询
	rows := db.QueryRow("select * from doctor_tb where id = ?", 1)
	rows.Scan(&doc.ID, &doc.Name, &doc.Age, &doc.Sex, &doc.AddTime)
	fmt.Println("单条数据结果：", doc)

	//-------3、查询数据列表--------
	rows2, err := db.Query("select * from doctor_tb where age > ?", 1)
	if err != nil {
		fmt.Println("多条数据查询错误", err)
		return
	}

	//定义对象数组,用于接收数据
	var docList []Doctor
	for rows2.Next() {
		var doc2 Doctor

		rows2.Scan(&doc2.ID, &doc2.Name, &doc2.Age, &doc2.Sex, &doc2.AddTime)
		//加入数组
		docList = append(docList, doc2)
	}
	//fmt.Println("多条数据查询结果", docList)
	//-------4、新增数据--------
	result, err := db.Exec("insert into doctor_tb(name,age,sex,addTime) values(?,?,?,Now())", "花医生", 40, 2)
	if err != nil {
		fmt.Println("新增数据错误", err)
		return
	}
	newID, _ := result.LastInsertId() //新增数据的ID
	i, _ := result.RowsAffected()     //受影响行数
	fmt.Printf("新增的数据ID：%d , 受影响行数：%d \n", newID, i)

	//-------5、修改数据--------
	result2, err := db.Exec("update doctor_tb set age=20 where id = ?", 2)
	if err != nil {
		fmt.Println("修改数据错误", err)
		return
	}
	i2, _ := result2.RowsAffected() //受影响行数
	fmt.Printf("受影响行数：%d \n", i2)
	/*
		//-------6、删除数据--------
		result3, err := db.Exec("delete from doctor_tb where name = ?", "花医生")
		if err != nil {
			fmt.Println("删除数据错误", err)
			return
		}
		i3, _ := result3.RowsAffected()
		fmt.Printf("受影响行数：%d \\n", i3)
	*/
	//-------7、事务--------
	tx, _ := db.Begin()
	result4, _ := tx.Exec("update doctor_tb set age = age + 1 where name = ?", "钟南山")
	result5, _ := tx.Exec("update doctor_tb set age = age + 1 where name = ?", "叶子")

	//影响行数，为0则失败
	i4, _ := result4.RowsAffected()
	i5, _ := result5.RowsAffected()
	if i4 > 0 && i5 > 0 {
		//2条数据都更新成功才提交事务
		err = tx.Commit()
		if err != nil {
			fmt.Println("事务提交失败", err)
			return
		}
		fmt.Println("事务提交成功")
	} else {
		//否则回退事务
		err = tx.Rollback()
		if err != nil {
			fmt.Println("回退事务失败", err)
			return
		}
		fmt.Println("回退事务成功")

	}

}
