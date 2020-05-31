package main

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

type User struct {
	ID   int
	Name string `orm:"size(100)"`
}

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:"+os.Getenv("MySQL_PASSWORD")+"@tcp("+os.Getenv("MySQL_HOST")+":3306)/test?charset=utf8", 30)

	// 注册定义的 model
	orm.RegisterModel(new(User))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

	/* 创建 table user
	mysql> show tables;
	+----------------+
	| Tables_in_test |
	+----------------+
	| user           |

	mysql> drop table user;

	*/
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// 插入 table user
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 更新表
	user.Name = "John"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := User{ID: user.ID}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// 删除表
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
