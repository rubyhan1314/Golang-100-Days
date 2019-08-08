# 数据表的查询语句

用户对于数据表或视图等最常用的操作就是查询。也叫检索。通过select语句来实现
语法规则：

```mysql
select {columns}
  from {table|view|other select}
  [where 查询条件]
  [group by 分组条件]
  [having 分组后再限定]
  [order by 排序]
```

注意事项：

1. sql的语句对于大小写不敏感。数值也不分。(Oracle中数值区分大小写)
2. 关键字不能被分行

## 一、简单查询

select 的检索语句中，只有select ，以及from子句。也是一个select语句必须的。

### 1.1 指定列名

查询指定的列：

```mysql
 select 列名,列名 from 表名
```

>如果查询表中所有的字段，使用*。

### 1.2 指定别名

 ```mysql
select 列名 as 别名,列名 别名 from 表名
 ```

>as关键字：其实也可以省略。

### 1.3 列运算

可以使用算术表达式

  在select语句中，检索到字段类型如果是number，date可以使用算术表达式
  加：+，减：-，乘：*，除:/,括号

> 如果某一列有null数值，那么计算后的结果也是null。通用函数:ifnull(a,b),a为null那么取b，如果a不为null那么取a的值。

### 1.4 字符串类型可以做连续运算

concat(“我的名字是”, name,'我今年。。');

  把搜索到的列和列进行拼接显示。

### 1.5 去重重复：distinct 字段名

  distinct 列名，表示去除重复
  distinct 列名1，列名2，一行的数据都相同，才会被认为是重复的数据，去除。

## 二、 条件查询

在实际查询的过程中, 我们需要根据需要来过滤出我们想要的数据, 这种查询就叫条件查询。在检索数据信息的时候，使用限定条件。表示满足条件才会被检索到。

使用where子句语法格式：

```mysql
select 检索列 from 表名 where 筛选条件
```

> 条件查询使用`where`子句对表中的数据进行筛选, 结果为`true`的行会出现的结果集中.
>
> 条件的表达式支持多种运算:

### 2.1 比较运算符

- 等于 `=`
- 大于 `>`
- 小于 `<`
- 大于等于 `>=`
- 小于等于 `<=`
- 不等于 `!=` 或 `<>`

---

查找年龄大于30岁的:
```mysql
select * from stus where age>30;
```



查询姓名不是`王二狗`的:

```mysql
select * from stus where name!='王二狗';
或者
select * from stus where name<>'王二狗';
```
### 2.2 逻辑运算符

- and, && (与)
- or, || (或)
- not, !(非)

---

查询年龄大于30并且性别是女的:
```mysql
select * from stus where age>30 and sex='女';
或
select * from stus where age>30 && sex='女';

```
查询年龄小于30或者性别是男的:

```mysql
select * from stus where age<30 or sex='男';
或
select * from stus where age<30 || sex='男';
```

---

### 2.3 模糊查询

比如想查询姓名中含有'张'的这种需求, 就需要用到模糊查询.

模糊查询用到关键字 `like`

```mysql
select * from stus name like '%张%';
李李张张
王张
张
张三
张三三
```



语法:

1. `%`表示任意多个字符（0-多个）

2. `_`(下划线)表示任意一个字符



查询姓名是两个字符的:

```mysql
select * from stus where name like '张__';
```

---

###  2.4 区间查询

#### 2.4.1 in (...)

- `in` 表示在一个离散的非连续的范围中查找

查找id的值等于1或3或4的:

```mysql
select * from stus where id in (1, 3,4);
```


#### 2.4.2 between...and...

- `between...and...` 
- 表示在一个连续的区间中查找.

查找id的值在3到7之间的:
```mysql
select * from stus where id between 3 and 7;
```

###  2.5 NULL判断

`is NULL` 用来判断某列是否为空. 注意: `NULL`用小写也是可以的.



## 三、排序

默认情况下, 我们查到的顺序是按照数据在数据库中的插入顺序排列的.
但是在实际情况下,我们需要根据不同的条件来排序, 比如按照更新日期排序, 或者按照价格从低到高排序, 或者按照人气从高到底排序等等.
**mysql支持对查询的结果进行重新排序.** 
使用order by 列子句可以完成排序的功能. 后面的列表示排序规则.

> order by 子句。语法上位于一条sql语句的末尾。

### 3.1 升序(asc)

默认是升序排序
按照年龄从低到高排序.

```mysql
select * from stus order by age;
```



> 说明:
>
> 1. 默认情况下是按照升序排序.asc
>
>    也可以在列的后面添加asc(ascend 提升)表示升序.
>
> 2. 使用desc表示降序
>
>    desc是单词 descend(下降)的缩写, 来表示降序排序.

### 3.2 降序(desc) 

```mysql
select * from stus order by age asc;
```



### 3.3 多种排序规则 

可以先按照一定规则来排序, 当碰到相同的情况下, 按照第二种规则来排序, ...
先按照年龄升序排序, 如果年龄相等再按照id的降序排序.

```mysql
select * from stus order by age asc,id desc;
```



## 四、聚合函数

在查询中，统计,求和等是很常用的，通过聚合函数完成能十分省事.

如果没有聚合函数，可能需要各种子查询，各种sql嵌套才能完成.

但是有了**聚合函数**，很多问题迎刃而解。

mysql提供了5个常用的聚合函数:`sum(), max(), min(), avg(), count()`

当然配合聚合函数使用的还有分组`group by 和 having子句`, 下一节再细讲.

### 4.1 sum(列)

这个聚合函数用来对那些列不是`null`的值进行求和.

- 对所有的age求和
  ​

- 多所年龄小于100的age求和

  

### 4.2 max(列)

这个聚合函数用来计算那些列不是`null`的值最大值.

- 找到最大的年龄

  

### 4.3 min(列)

这个聚合函数用来计算那些列不是`null`的值最小值.

- 找到最小的年龄


### 4.4 avg(列)

这个聚合函数用来计算那些列不是`null`的值平均值.

- 计算`age`的平均值

  

### 4.5 count(列)

统计函数.

count(*)--->统计所有行的数据,16

count(主键列)-->16

count(非主键列1),是否存在null数值。

count(comm)-->6



- 统计一共有多少行数据

  

- 统计性别不为`null`的数据 行数

  


## 五、分组查询

分组(group by)
group by 列名, 按照指定的列进行分组, 值相同的会分在一组.
语法:

```mysql
select 列名 from 表名 group by 列名;
```


说明:

1. select后面跟的列名必须和group by后的列列名保持一致.
2. 当group by单独使用时, 只显示每组的第一条记录. 所以group by单独使用的意义不大.大多要配合聚合函数。
3. group by后面也可以给多个列进行分组, 表示这些列都相同的时候在一组. 

### 5.1 按照某列分组

```mysql
select sex from  stus group by sex;
```

### 5.2 按照多列分组

```mysql
select name, sex from stus group by name,sex; 
```

### 5.3 分组后使用聚合函数

```mysql
select sex,count(*) from stus group by sex;
```



> 注意：
>
> 1. 如果没有分组，那么组函数(max,min,avg,count....)作用在整张表上。
> 2. 如果有分组，组函数作用在分组后的数据上。
> 3. 在写select 子句中列，如果没有在组函数里，那么就要在group by后边。

```mysql
select a，b，count(c),sum(d) from 表 group by a,b;
```



### 5.4 分组后限定查询：having

二次筛选：二次筛选是指分组后再对数据进行筛选.
需要使用having子句来完成.

```mysql
select 列名 from  表名 group by 列名 having 条件
```

也就是说，分组后再限定，使用having子句

having子句对于group by分组后的结果，进行再次筛选。最后输出的结果就满足having条件的。

> having子句的和where子句的用法类似，都是用于限定条件。
>
>     对比：
>     1. where和having后面都是跟条件
>     2. where是对表中的数据的原始筛选
>     3. having是对group by的结果的二次筛选
>     4. having必须配合group by使用, 一般也会跟着聚合函数一起使用
>     5. 可以先有where, 后面跟着group by和having
>    
>     区别和结论：
>       1.语法上：在having中使用组函数(max,min,avg,count...),而where后不能用组函数。
>       2.执行上：where是先过滤再分组。having是先分组再过滤。



例如：把sex不是null的过滤出来.

例如：把男或女的个数超过6的过滤出来




## 六、分页查询

Limit子句(方言)

> 方言的意思是limit是mysql特有的。

Limit用于限定查询结果的起始行，以及总行数。

```mysql
select * from 表名 limit start,count;
```

例如：查询起始行为第5行，一共查询3行

select * from stus limit 4, 3;

​	其中4表示从第5行开始，其中2表是查询3行。即5,6,7行记录。

## 七、内置函数[扩展]

### 7.1 字符串函数

- 查看字符的ascii码值ascii(str)，str是空串时返回0

```mysql
select ascii('a');
```

- 查看ascii码值对应的字符char(数字)

```mysql
select char(97);

```

- 拼接字符串concat(str1,str2...)

```mysql
select concat(12,34,'ab');

```

- 包含字符个数length(str)

```mysql
select length('abc');

```

- 截取字符串
  - left(str,len)返回字符串str的左端len个字符
  - right(str,len)返回字符串str的右端len个字符
  - substring(str,pos,len)返回字符串str的位置pos起len个字符，从1开始

```mysql
select substring('abc123',2,3);

```

- 去除空格
  - ltrim(str)返回删除了左空格的字符串str
  - rtrim(str)返回删除了右空格的字符串str
  - trim([方向 remstr from str)返回从某侧删除remstr后的字符串str，方向词包括both、leading、trailing，表示两侧、左、右

```mysql
select trim('  bar   ');
select trim(leading 'x' FROM 'xxxbarxxx');
select trim(both 'x' FROM 'xxxbarxxx');
select trim(trailing 'x' FROM 'xxxbarxxx');
```

- 返回由n个空格字符组成的一个字符串space(n)

```mysql
select space(10);

```

- 替换字符串replace(str,from_str,to_str)

```mysql
select replace('abc123','123','def');
```

- 大小写转换，函数如下
  - lower(str)
  - upper(str)

```mysql
select lower('aBcD');

```

### 7.2 数学函数

- 求绝对值abs(n)

```mysql
select abs(-32);
```

- 求m除以n的余数mod(m,n)，同运算符%

```mysql
select mod(10,3);
select 10%3;
```

- 地板floor(n)，表示不大于n的最大整数

```mysql
select floor(2.3);
```

- 天花板ceiling(n)，表示不小于n的最大整数

```mysql
select ceiling(2.3);

```

- 求四舍五入值round(n,d)，n表示原数，d表示小数位置，默认为0

```mysql
select round(1.6);

```

- 求x的y次幂pow(x,y)

```mysql
select pow(2,3);

```

- 获取圆周率PI()

```mysql
select PI();
```

- 随机数rand()，值为0-1.0的浮点数

```mysql
select rand();
```

- 还有其它很多三角函数，使用时可以查询文档

### 7.3 日期时间函数

- 获取子值，语法如下
  - year(date)返回date的年份(范围在1000到9999)
  - month(date)返回date中的月份数值
  - day(date)返回date中的日期数值
  - hour(time)返回time的小时数(范围是0到23)
  - minute(time)返回time的分钟数(范围是0到59)
  - second(time)返回time的秒数(范围是0到59)

```mysql
select year('2016-12-21');
```

- 日期计算，使用+-运算符，数字后面的关键字为year、month、day、hour、minute、second

```mysql
select '2016-12-21'+interval 1 day;
select '2017-12-12' + interval 3 month;
select date_add('2017-12-12', interval 1 day);
```

- 日期格式化date_format(date,format)，format参数可用的值如下

  - 获取年%Y，返回4位的整数

    *　获取年%y，返回2位的整数

    *　获取月%m，值为1-12的整数

  - 获取日%d，返回整数

    *　获取时%H，值为0-23的整数

    *　获取时%h，值为1-12的整数

    *　获取分%i，值为0-59的整数

    *　获取秒%s，值为0-59的整数

```mysql
select date_format('2016-12-21','%Y %m %d');
select str_to_date('12/12/2017','%d/%m/%Y');
```

- 当前日期current_date()

```mysql
select current_date();

```

- 当前时间current_time()

```mysql
select current_time();

```

- 当前日期时间now()

```mysql
select now();
```



## 七、多表查询

连接查询：当需要对有关系的多张表进行查询时，需要使用连接join

**合并结果集**

1. 要求被合并的表中，列的类型和列数相同
2. UNION,去除重复行
3. UNION ALL， 不去除重复行



**分类** 

	1. 内连接
	2. 外连接
		 A：左外连接
		 B：右外连接
		 C：全外连接(Mysql不支持)
	3. 自连接(术语一种简化方式)

### 7.1 笛卡尔积

如果两张表在连接查询时，如果没有连接条件，那么就会产生笛卡尔积。(冗余数据)。

### 7.1 内连接

内连接查询出的所有的记录都满足条件

```mysql
方言：select * from 表1 别名1, 表2 别名2 where 别名1.xx=别名2.xx
	也叫等值链接
标准：select * from 表1 别名1 inner join 表2 别名2 on 别名1.xx=别名2.xx
自然：select * from 表1 别名1 natural join 表2 别名2

```



### 7.2 外连接

内连接所检索出来的结果，都是满足连接条件的。外链接是扩展内连接检索出来的结果集。外链接返回的结果，除了包含满足链接条件的记录，还包括不满足连接条件。

#### 7.2.1 左外连接

左表记录无论是否满足条件都会查询出来，而右表只有满足条件才能查询出来。左表中不满足条件的记录，右表部分都为NULL

语法格式：

```mysql
左外：
	select * from 表1 别名1 left outer join 表2 别名2 on 别名1.xx=别名2.xx
	
左外自然：
	select * from 表1 别名1 natural left outer join 表2 别名2 
```



#### 7.2.2 右外连接

右表记录无论是否满足条件都会查询出来，而左表只有满足条件才能查询出来。右表中不满足条件的记录，左表部分都为NULL

语法格式：

```mysql

右外：
	select * from 表1 别名1 right outer join 表2 别名2 on 别名1.xx=别名2.xx
	

右外自然：
	select * from 表1 别名1 natural right outer join 表2 别名2 

```

#### 7.2.3 全外连接

Mysql不支持全外连接(full outer join)，但是可以使用union来合并左外连接以及右外连接的结果，达到全外连接的效果。

```mysql
左外连接
SELECT e.ename,e.sal, d.dname,e.deptno,d.deptno
FROM emp e LEFT OUTER JOIN dept d
ON e.deptno=d.deptno
UNION
右外连接
SELECT e.ename,e.sal, d.dname,e.deptno,d.deptno
FROM emp e RIGHT OUTER JOIN dept d
ON e.deptno=d.deptno
```



### 7.3 自连接

自连接(术语一种简化方式)。其实就是一张表连接字节。

查询员工的上级名称

```mysql
SELECT e.empno,e.ename,e.mgr,m.empno,m.ename
FROM emp e LEFT OUTER JOIN emp m
ON e.mgr=m.empno
```

## 八、子查询

子查询是指sql语句中包含另一个select语句。内查询或者内select语句

> 一条sql语句中，包含多个select关键字	
>
> 外查询，内查询

子查询可以出现的位置：

​	from后，作为表

​	where后，作为条件

注意事项：
1.子查询必须在()里
2.在子查询中不能使用order by子句。
3.子查询可以再嵌套子查询，最多不能超过255层。



子查询细分为：单行子查询，多行子查询，关联子查询。

1.单行子查询
  子查询的结果是单行数据。
    在where条件后需要配合单行运算符：>,<,>=,<=,!=,=
2.多行子查询
  子查询的结果是多行数据。
    在where条件后需要配合：in，any，all运算符
    in：
    any：匹配上子查询的结果集中任意一个即可。
    all：匹配上子查询的所有的结果。

3.关联子查询
  对于单行子查询和多行子查询，外查询和内查询是分开执行的。
  如果外查询使用到内查询的结果，就使用关联子查询。
  关联子查询，是指在内查询中需要借助于外查询，而外查询离不开内查询的执行。就叫关联子查询。

## 九、总结

完整的select语句书写：

```mysql
select distinct *
from 表名
where ....
group by ... having ...
order by ...
limit start,count
```

执行顺序：

```mysql
from 表名
where ...
group by ...
having ...
select distinct*..
order by ...
limit start,count
```

实际使用中，只是使用语句中某些部分的组合，而不是全部。