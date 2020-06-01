Go官方没有提供数据库驱动，而是为开发数据库驱动定义了一些标准接口，开发者可以根据定义的接口来开发相应的数据库驱动。

# prepartion
## start the K8s cluster including the MySQL service.
## check the databse
```bash
$mysql -h <service ip> -u root -p
```
create databse & tables
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
## export the env which are used by PODs of service myapp
```bash
    $ source export.sh
```     
## install mysql driver
```bash 
$ go get -u github.com/go-sql-driver/mysql
$ go get github.com/astaxie/beego  //for ORM
```

# examples
[code](https://github.com/yc-alex-xu/go/tree/master/src/practise/database)
* mysql.go  sql 语句 + rowset 模式
* ORM.go    ORM method


# note
* ORM 即在关系型数据库和业务实体对象之间作一个映射，这样，我们在具体的操作业务对象的时候，就不需要再去和复杂的SQL语句打交道，只要像平时操作对象一样操作它就可以了。常见的ORM框架有：Hibernate、TopLink、Castor JDO、Apache OJB等。
* TBD

