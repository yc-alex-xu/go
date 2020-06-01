package main

import (
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

type Userinfo struct { //表名，field名字，在数据库中实际都是小写
	Uid        int `orm:"PK"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username   string
	Departname string
	Created    time.Time
}
type User struct {
	Uid     int `orm:"PK"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"` //设置一对多关系
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

/* after init
mysql> show tables;
+----------------+
| Tables_in_test |
+----------------+
| post           |
| post_tags      |
| profile        |
| tag            |
| user           |
| userdetail     |
| userinfo       |
+----------------+
*/

func init() {
	fmt.Println("ORM init start")
	orm.Debug = true
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:"+os.Getenv("MySQL_PASSWORD")+"@tcp("+os.Getenv("MySQL_HOST")+":3306)/test?charset=utf8", 30)

	orm.RegisterModel(new(Userinfo), new(User), new(Profile), new(Post), new(Tag))
	orm.RunSyncdb("default", false, true) //create table
	fmt.Println("ORM init end")
}

func main() {
	o := orm.NewOrm()
	var Uid int
	var u Userinfo
	for _, s := range []string{"smith", "john", "kate", "mary"} {
		u.Username = s
		u.Departname = "tbd"
		fmt.Println(s)
		id, err := o.Insert(&u)
		if err != nil {
			panic(err)
		} else {
			Uid = int(id)
			fmt.Printf("inserted: (%d, %v)\n", id, s)
		}

	}

	// 更新表
	u.Username = "new"
	num, err := o.Update(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u = Userinfo{Uid: Uid}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// 删除表
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
