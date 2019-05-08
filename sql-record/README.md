# sql 记录

### 记录 sql 遇到的一些问题

### - 1. unique、primarykey的区别

#####  **UNIQUED** 可空，可以在一个表里的一个或多个字段定义；**PRIMARY KEY** 不可空不可重复，在一个表里可以定义联合主键；

简单的说，**primary key = unique +  not null** 

一个表只能有一个主键，但是可以有好多个UNIQUE，而且UNIQUE可以为NULL值，如员工的电话号码一般就用UNIQUE，因为电话号码肯定是唯一的，但是有的员工可能没有电话。

主键肯定是唯一的，但唯一的不一定是主键；

不要总把UNIQUE索引和UNIQUE约束混为一谈

1、primary key = unique + not null

2、唯一约束和主键一样都是约束的范畴，而且都可以作为外键的参考，不同的是，一张表只能有一个主键

3、主键和唯一约束的创建需要依靠索引，如果在创建主键或唯一约束的时候没有已经建好的索引可以使用的话，Oracle会自动建立一个唯一的索引。

### - 2. 索引的两大类型hash与btree

```go
索引类型，分两类
hash类型的索引：查询单条快，范围查询慢
btree类型的索引：b+树，层数越多，数据量指数级增长（我们就用它，因为innodb默认支持它）
```

### -3. 联合索引举例

```go
## E.g.1.
type test struct {
  ID  		int64
  GoodsId int64				`sql:",unique:user_id_goods_id"`
  UserId  int64				`sql:",unique:user_id_goods_id"`
}

## E.g.2.
type test2 struct {
  ID     			int64		
  ParcentCid  int64   `json:"parcent_cid" sql:",unique:parent_cid_name"`
  Name        string	`json:"name" sql:",unique:parent_cid_name"`
}
```

### -4. 联合查询索引的范围

- 建立联合索引

  ```sql
  CREATE TABLE `user2` (
    `userid` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(20) NOT NULL DEFAULT '',
    `password` varchar(20) NOT NULL DEFAULT '',
    `usertype` varchar(20) NOT NULL DEFAULT '',
    PRIMARY KEY (`userid`),
    KEY `a_b_c_index` (`username`,`password`,`usertype`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
  ```

- 规则

  ```sql
  // 上表中有一个联合索引，下面开始验证最左匹配原则。
  // 当存在username时会使用索引查询：
  
  explain select * from user2 where username = '1' and password = '1';
  
  // 当没有username时，不会使用索引查询：
  
  explain select * from user2 where password = '1';
  
  // 当有username，但顺序乱序时也可以使用索引：
  
  explain select * from user2 where password = '1' and username = '1';
  ```

### -5.使用例子

- 主键索引

![image-20190508175547926](/Users/chikitchong/Develop/learning_record/sql-record/README.assets/image-20190508175547926.png)

- 其他索引

![image-20190508175633125](/Users/chikitchong/Develop/learning_record/sql-record/README.assets/image-20190508175633125.png)

- 没有索引

![image-20190508175650900](/Users/chikitchong/Develop/learning_record/sql-record/README.assets/image-20190508175650900.png)

