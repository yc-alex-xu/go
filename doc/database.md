Go官方没有提供数据库驱动，而是为开发数据库驱动定义了一些标准接口，开发者可以根据定义的接口来开发相应的数据库驱动。

# prepartion
## start the K8s cluster including the MySQL/Redis service.
## setup in MySQL
```bash
$mysql -h <service ip> -u root -p
```
sql
```sql
use test;
mysql> create database test;
mysql> use test;
Database changed
mysql>

CREATE TABLE `userinfo` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NULL DEFAULT NULL,
	`department` VARCHAR(64) NULL DEFAULT NULL,
	`created` DATE NULL DEFAULT NULL,
	PRIMARY KEY (`uid`)
);
mysql>
CREATE TABLE `userdetail` (
	`uid` INT(10) NOT NULL DEFAULT '0',
	`intro` TEXT NULL,
	`profile` TEXT NULL,
	PRIMARY KEY (`uid`)
);

```  
## setup in redis
```bash
$ src/redis-cli -h 10.152.183.202
10.152.183.202:6379> set name abc
OK
```

## emulate the env of POD
```bash
    $ source export.sh
```     
## install libs
```bash 
$ go get -u github.com/go-sql-driver/mysql　// for mysql
$ go get -u github.com/astaxie/beego  //for ORM
$ go get -u github.com/gomodule/redigo/redis
```

# examples
[code](https://github.com/yc-alex-xu/go/tree/master/src/practise/database)
* mysql.go  sql 语句 + rowset 模式
* ORM.go    MySQL/ORM 
* redis.go    


# note
* ORM 即在关系型数据库和业务实体对象之间作一个映射，这样，我们在具体的操作业务对象的时候，就不需要再去和复杂的SQL语句打交道，只要像平时操作对象一样操作它就可以了。常见的ORM框架有：Hibernate、TopLink、Castor JDO、Apache OJB等。
* 至少beego的 ORM　是比较原始的，其他的ORM框架没有用过，但只要sql->代码的mapping 不流畅，可用性就不会很高。

# referenc
* https://github.com/go-sql-driver/mysql
* https://github.com/gomodule/redigo