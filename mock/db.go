package mock

import (
	"database/sql"
	"fmt"
)

// 数据库连接测试

// OrderDBI 定义了一个订单接口，有一个获取名称的方法
type OrderDBI interface {
	GetName(conStr string, userId int) string
}

// OrderInfo 定义结构体
type OrderInfo struct {
	orderid int
}

// GetName 实现接口的方法
func (order OrderInfo) GetName(conStr string, userId int) string {
	fmt.Println("connect string is", conStr)
	fmt.Println("user id is", userId)
	db, err := sql.Open("mysql", conStr)
	if err != nil {
		fmt.Errorf("db connect error %+v", err)
		return ""
	}
	_, err = db.Query("select * from user where id=" + fmt.Sprintf("%d", userId))
	if err != nil {
		fmt.Errorf("select error %+v", err)
		return ""
	}

	return fmt.Sprintf("%d", userId)
}
