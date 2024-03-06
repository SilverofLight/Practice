# 什么是资料库？

## 关联式资料库（SQL）

- MySQL
- Oracle
- PostgreSQl
- SQL Server

关联式资料库管理系统（EDBMS）

## 非关联式资料库（noSQL/not just SQL）

- MogoDB
- Redis
- DynamoDB
- Elaticsearch

非关联式资料库管理系统（NRDBMS）

## 什么是SQL？

Structure Query Language 结构性询问语言

![Untitled](2.png)

# 表格和键 tables & keys

table：

| student_id | name | major |
| --- | --- | --- |
| 1 | 小黑 | 生物 |
| 2 | 小白 | 历史 |
| 3 | 小绿 | 英语 |
| 4 | 小白 | 历史 |

可以设定主键（primary key）可以唯一的表示每一笔资料，如student_id

连接到其他表格可以使用外键（foreign key）

![Untitled](3.png)

主键可以设置多个，一点要可以指定唯一的资料

# 创建资料库 create database

创建一个名为sql_turotial 的资料库

一般名称会用 ` ` 包起来

```sql
CREATE DATABASE `sql_tutorial`;
```

展示资料库

```sql
SHOW DATABASE;
```

删除资料库

```sql
DROP DATABASE `sql_tutorial`
```

# 创建表格 create table

以sql_turotial为例

首先指定资料库

```sql
USE `sql_tutorial`;
```

资料库有资料形态：

| INT | 整数 |  |
| --- | --- | --- |
| DECIMAL(m,n) | 有小数点的数 | m表示总位数，n表示小数点后的位数 |
| VARCHAR(n) | 字符串 | n表示最多能存放的字元 |
| BLOB | 图片、影片、档案 | 二进制资料 |
| DATE | ‘YYYY-MM-DD’ 日期 |  |
| TIMESTAMP | ’YYYY-MM-DD   HH : MM : SS’ 记录时间 |  |

如要创建一下表格：

![Untitled](3.png)

```sql
CREATE TABLE `student`(
		`student_id` INT PRIMARY KEY, //属性是int，顺便设定为主键
		`name` VARCHAR(20),
		`major` VARCHAR(20)           //最后一个不用写逗号
		//PEIMARY KEY(`student_id`)     也可以设定为主键
);
```

## 显示表格信息

```sql
DESCRIBE `student`;
```

## 删除表格

```sql
DROP TABLE `student`;
```

## 给表格增加一列gpa

```sql
ALTER TABLE `student` ADD gpa DECIMAL(3,2);
```

## 删掉某一行

```sql
ALTER TABLE `student` DROP COLUMN gpa;
```

# 存储资料 insert

创建完表格之后，根据属性存入资料

## 填入资料的顺序要根据创建时的顺序，即id→名称→学科

```sql
INSERT INTO `student` VALUES(1, "小白", "历史");
```

搜寻表格里的全部资料

```sql
SELECT * FROM `student`;  // * 是全部的意思
```

NULL 表示空

```sql
INSERT INTO `student` VALUES(2, "小黑", NULL);
```

## 或者写入时可以自己决定填入顺序

```sql
INSERT INTO `student`(`name`, `major`, `student_id`) VALUES("小黑", NULL, 2);
```

这个方法也可以只写其中一项或几项

# 限制，约束 constraints

```sql
CREATE TABLE `student`(
		`student_id` INT PRIMARY KEY, 
		`name` VARCHAR(20) NOT NULL,  //NOT NULL 表示这一项不能为空，输入NULL时报错
		`major` VARCHAR(20) UNIQUE    //UNIQUE：不可以重复      
);
```

```sql
`major` VARCHAR(20) DEFAULT `历史`;   //如果没有输入，默认是历史
```

```sql
`student_id` INT AUTO_INCREMENT,   //在新增一行是，不需要再输入这一项，默认是上一行加1
```