Go官方没有提供数据库驱动，而是为开发数据库驱动定义了一些标准接口，开发者可以根据定义的接口来开发相应的数据库驱动，这样做有一个好处，只要是按照标准接口开发的代码， 以后需要迁移数据库时，不需要任何修改。

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
$go get -u github.com/go-sql-driver/mysql
```

# examples

