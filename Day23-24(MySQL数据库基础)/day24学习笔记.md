# day15 课堂资料

## 上节课回顾

mysql数据库

登录：

​	dos窗口：mysql

​		-u 用户名

​		-p 密码

​	mysql自带的命令行：

​		直接输入密码

​	

​	可视化工具：



数据库的基本常识：

​	关系型的数据库：以行和列来组织数据。(二维表)

安装mysql后：

​	创建数据库：(以项目为单位)

​		创建数据表：table

​			行：一条数据，或者叫一条记录

​			列：字段，属性，field



SQL语句：

​	DDL：数据定义语言

​		create database

​		create table

​		alter talble

​		drop table







1.创建数据库：

2.显示所有的数据库：show databases;

3.切换数据库：use 数据库名;

4.创建数据表：

​	id：int，主键(非空+唯一)-->唯一标识，自增：auto_increment

​	数据类型：

​		数值型：int(长度)，float，double。。。

​		字符型：char，varchar

​		日期型：date，time，datetime



5.插入数据：

​	insert into 表名(字段。。。)values (数值。。。); 

6.修改表结构：

​	alter table 表名 add/modify/change/drop 

​	rename to

7.删除数据表：

​	drop table 表名;

8.修改数据：

​	update 表名 set 字段=数值，字段=数值 [where 修改条件]

```
=，!=,<>
>,<,>=,<=
or,and,not
between...and, in()
null	is null/is not null
```



9.删除数据：

​	delete from 表名 [where 删除条件];



## 一、约束

约束：用于限制数据表中某列的数据的存储内容。

​	非空约束，唯一约束

主键约束：非空+唯一

​	用作这个表中，主键所在的字段是该表的唯一标识。

外键约束：保证数据的完整性和有效性。

​	两张表：

​		父表：主表

​			主键

​		子表：从表

​			外键



子表中设置外键的列，是父表中主键。那么子表中外键的列，的数值，就会收到父表中主键的数值的约束。

创建父表：

```Mysql
mysql> create table class(
    -> classno int(4) primary key,
    -> classname varchar(20),
    -> local varchar(30));
```



创建子表：

```mysql
mysql> create table student(
    -> sid int(4) primary key auto_increment,
    -> sname varchar(30),
    -> age int(3),
    -> sex varchar(3),
    -> classno int(4),
    -> constraint fk_stu foreign key (classno) references class(classno));
```



添加数据：

​	父表：1,2,3

​	子表：classno：受到父表的限制



级联操作：父表的数据删除时候，子表中引用的数值？ on delete xxx

​	默认：

​	cascade：子表的数据随着父表一起被删除

​	set null：父表被删除，子表置为null

​	no action：



删除student中原来的外键约束

```mysql
alter table student drop foreign key fk_stu;
```

新增外键约束：

```mysql
mysql> alter table student add constraint fk_stu foreign key(classno) references class(classno) on delete cascade;
```



父表的数据删除，子表的也随之删除。





## 二、查询

### 2.1 简单查询

1.查询指定的列：

```mysql
select 列1，列2，列3.。。 from 表名
```

2.起别名：as 可以省略不写

```mysql
select 列1 as 别名 from 表名
```

3.处理null

使用ifnull(a,b)，如果a列为空，就取b的值。

```mysql
select ifnull(comm,0) from emp;
```

4.去重：distinct

```mysql
select distinct job from emp;
```



### 2.2 条件查询

在检索数据库中的数据时候，需要满足某些条件，才能被检索到，使用where关键字，来限制检索的条件。

```mysql
比较运算符：=，!+=,<>,>,<,>=,<=

逻辑运算符：and ，or， not

范围：between and，in

null：is null，is not null

```



练习1：查询1981年以后入职的员工信息

练习2：查询部门编号为30或者工资大于2000的员工信息。



**模糊查询：like**

%：匹配0-多个任意的字符

_：匹配1个任意字符

```mysql
//名字的第三个字母为a的员工信息
mysql> select * from emp where ename like '__a%';
```



**排序**：orderby

asc：升序，默认

desc：降序

select查询完后，排序要写在整个sql语句的最后。



### 2.3 统计函数

也叫聚合函数，通常用于求整个表中某列的数据的：总和，平均值，最大值，最小值。通常不搭配表中的字段一起查询。

sum(),

avg(),

max()

min(),

count(*/主键)



练习1：求部门20中员工的平均工资，工资总和，工资最大值，最小值，人数。

mysql> select ename,sum(sal),avg(sal),max(sal),min(sal) ,count(empno),count(comm)from emp where deptno=20;



### 2.4 分组

group by 列

按照某列分组，该列有几种取值，就分为几组。

练习1：按照部门来分组，查询每个部门的最高工资，最低工资，平均工资。

练习2：求每种工作的最高薪资，最低薪资，以及人数。

练习3：查询部门人数超过5人的部门。



### 2.5 分页

limit start，count。

练习：按照工资排序，获取前5条数据。





## 三、多表联查

### 3.1 内连接

查询出来的数据一定满足链接的规则。

```mysql
select e.*,d.* from emp e inner join dept d on e.deptno=d.deptno;
```



### 3.2 左外链接

因为内连接的查询结果，并不是所有的数据，而是满足规则的数据。

左外链接，右外连接是为了补充内连接的查询结果的。

左表记录无论是否满足条件都会查询出来，而右表只有满足条件才能查询出来。左表中不满足条件的记录，右表部分都为NULL

```mysql
mysql> select * from emp e left outer join dept d on e.deptno=d.deptno;
```

### 3.3 右外连接

右表记录无论是否满足条件都会查询出来，而左表只有满足条件才能查询出来。右表中不满足条件的记录，左表部分都为NULL

```mysql
mysql> select * from emp e right outer join dept d on e.deptno=d.deptno;
```







练习1：查询所有的部门，以及对应的员工信息。

练习2：查询每个员工的员工信息，工资等级。emp，salgrade

练习3：查询每个员工的员工信息，部门名称，部门位置，工资等级

练习4：查询在部门在纽约的员工信息，部门名称，工资等级。





练习5：查询每个部门的人数，部门名称，部门编号。



自连接：

查询员工的姓名和上级的姓名：

 select e.empno,e.ename,e.mgr,m.empno,m.ename from emp e,emp m where e.mgr=m.empno;





## 四、子查询

1.查询比allen工资高的员工信息。

select * from emp where sal > (select sal from emp where ename='allen');



练习1：查询工资不是最高的，也不是最低的员工信息。

mysql> select * from emp where sal !=(select max(sal) from emp ) and sal !=(select min(sal) from emp);



练习2：不是销售部的员工信息

dname--->deptno

思路一：

select deptno from dept where dname='sales'

mysql> select * from emp where deptno != (select deptno from dept where dname='sales');





思路二：

select deptno from dept where dname !='sales';

mysql> select * from emp where deptno in(select deptno from dept where dname !='sales');







练习3：查询员工信息，要求工资高于部门编号为10的中的任意员工即可

思路一：

select mix(sal) from emp where deptno=10; // 

mysql> select * from emp where sal >(select min(sal) from emp where deptno=10);

思路二：

mysql> select * from emp where sal >any (select sal from emp where deptno=10);



练习4：查询员工信息，工资大于30部门的所有人

思路一：

mysql> select * from emp where sal >(select max(sal) from emp where deptno=30);

思路二：

mysql> select * from emp where sal > all(select sal from emp where deptno=30);



练习5：查询本公司工资最高的员工详细信息

select max(sal) from emp;

select * from emp e,dept d

 where sal=(select max(sal) from emp)  and e.deptno=d.deptno;

mysql> select e.*,d.* from emp e,dept d where sal=(select min(sal) from emp) and e.deptno=d.deptno ;







## 五、总结

主外键约束

​	主键

​	外键：两张表：父表，子表

​	子表中的外键必须是父表中的主键。

​		on delete set null/cascade/no action

查询

​	简单查询：去重，别名，加减乘除运算，ifnull(a, b)

​	条件查询：where

​				=，!=,<>,>,<,>=,<=

​				and, or, not

​				between and , in 

​				is null ,is not null

​			排序：order by 

​				asc，desc

​			模糊查询：like 

​				_,%



​	聚合函数：max(),min(),count(),sum(),avg()

​	分页：limit start，count

​	分组：group by

​		having 



多表联查：

​	内连接：select 。。。 from 表1 别名 inner join 表2 别名 on 连接条件

​	外连接：左外，右外

​		select 。。 from 左表 别名 left outer join 右表 别名 on 连接条件

​		select 。。 from 左表 别名 right outer join 右表 别名 on 连接条件



子查询：

​	select 的嵌套

​	select 的查询结果：





