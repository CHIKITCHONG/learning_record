###  2019-3-14

#### 再议 golang 指针

`前段时间中使用 pg 遇到一个问题：`

```go
func (this *ArticleBackstageDB) Create(o orm.DB, req model.Article) (*model.Article, error) {

	var (
		err error
		res *model.Article
		result orm.Result
	)

	//res = new(model.Article)

	// 检查此分类 ID 下是否已经存在此文章
	if result, err = o.Query(res, "select id from articles where (title = ? and category_id = ?) and deleted_at is null",
		req.Title, req.CategoryId); err != nil {
		fmt.Println("***debug here 1 ***")
		return nil, err
	}
	return nil, err
}
# 这样创建一个东西,pg会报错：pg: Scan(nonsettable *model.Article)
# 就是说不是一个可扫描的类型,说明 Query 接受指针类, *xxx 是一个实体类
```

- 正确的做法是

```go
func (this *ArticleBackstageDB) Create(o orm.DB, req model.Article) (*model.Article, error) {

	var (
		err error
		res *model.Article
		result orm.Result
	)
    # new 一个类型(创建指针引用)
	res = new(model.Article)

	// 检查此分类 ID 下是否已经存在此文章
    if result, err = o.Query(res, "select id from articles where (title = ? and category_id = ?) and deleted_at is null",
		req.Title, req.CategoryId); err != nil {
		fmt.Println("***debug here 1 ***")
		return nil, err
	}
	return nil, err
}
```



#### 补充一些 stackoverflow 的资料

- 1.Why would I make() or new()?

```go
// Things you can do with make that you can't do any other way:

# 这里主要是说 new、make 都可以建造一个新的指针引用

- Create a channel
- Create a map with space preallocated
- Create a slice with space preallocated or with len != cap

// It's a little harder to justify new. The main thing it makes easier is creating pointers to non-composite types. The two functions below are equivalent. One's just a little more concise:

func newInt1() *int { return new(int) }

func newInt2() *int {
    var i int
    return &i
}
```

- 2.Golang - Difference between &Struct{} vs Struct{}

```go
Well, they will have different behavior. Essentially if you want to modify state using a method on a struct, then you will need a pointer, otherwise a value will be fine. Maybe an example will be better:

package main
import "fmt"

# 这里主要是说你要修改一个结构体，在golang里必须拿到他的引用(指针)
# 直接修改是没有用的

type test_struct struct {
  Message string
}

func (t test_struct)Say (){
   fmt.Println(t.Message)
}

func (t test_struct)Update(m string){
  t.Message = m; 
}

func (t * test_struct) SayP(){
   fmt.Println(t.Message)
}

func (t* test_struct) UpdateP(m string)  {
  t.Message = m;
}

func main(){
  ts := test_struct{}
  ts.Message = "test";
  ts.Say()
  ts.Update("test2")
  ts.Say() // will still output test

  tsp := &test_struct{}
  tsp.Message = "test"
  tsp.SayP();
  tsp.UpdateP("test2")
  tsp.SayP() // will output test2

}
```
### 2019-3-18

`若编译时没报什么错,但自动停了,是端口被占用`


### 2019-3-22

`当收到一些不规则的json回复类似：`

```markdown
filename="goodcs-dir/113b152de2-7a16-402a-982f-87e1709a594eImage.png"&size=312038&mimeType="image/png"&height=949&width=1856

# 并不是标准的：
{
    "xxx": x,
    "xxxx": "xxxx",
}
```

##### 这时需要分割字符(&),构造字典

```go
//e.g.
/*[filename="goodcs-dir/113b152de2-7a16-402a-982f-87e1709a594eImage.png"
size=312038 mimeType="image/png" height=949 width=1856]
*/
func Handle(req string) (resp map[string]string) {
	condi := make(map[string]string)
	result := strings.Split(req, "&")
	for _, v := range result {
		v := strings.Split(v, "=")
		condi[v[0]] = strings.Replace(v[1], `"`, "", -1)
	}
	return condi
}
```

### 2019-3-25

`golang 读取字典`

```go
func main() {
	test := make(map[string]string)
	test["hello"] = "nihk"
	test["heelo2"] = "nihk2"
	for k, v := range test {
		fmt.Println("k", k, "v", v)
	}
}
# 结果是
k hello v nihk
k heelo2 v nihk2
```



### 2019-3-27

##### 关于 sql 联结查询(左联结是以左表为基准,右联结是以右表为基准)

```go
// 演示联结查询, 以 go-pg 为例(github.com/go-pg)

type Profile struct {
	Base
	UserId        int64  ``
	Name          string `sql:"type:varchar(64)"`
	Signature     string `sql:"type:varchar(64)"`
	Birth         string `sql:"type:varchar(64)"`
	Job           string `sql:"type:varchar(64)"`
	Constellation string `sql:"type:varchar(64)"`
	Phone 		  string `sql:"type:varchar(16)"`
}

type Address struct {
	Base
	UserID    int64  ``
	Name      string `sql:"type:varchar(64)"`
	Phone     string `sql:"type:varchar(64)"`
	Addr      string `sql:"type:varchar(64)"`
	FirstAddr bool   `` // default 是默认地址
}

func joinsql(db *pg.DB) error {

	var (
		err	error
		res []*pgmodel.Address
	)
	// 若不加入ORDER BY 两表中一个唯一的字段，就会重复打印两表中共同的值
	_, err = db.Query(&res,
		`SELECT s.id, s.created_at, s.updated_at, s.deleted_at,
				s.user_id, s.name,s.phone, s.addr 
				From addresses s
				FULL OUTER JOIN profiles c 
				ON s.user_id = c.user_id
				GROUP BY s.id
				ORDER BY s.id ASC
				`)
	if err != nil {
		return err
	}
	time.Sleep(2 * time.Second)
	for _, v := range res {
		fmt.Println("查找的结果是:", v)
	}

	return nil
}

2019/04/04 18:00:43 2.388556ms | SELECT s.id, s.created_at, s.updated_at, s.deleted_at,
				s.user_id, s.name,s.phone, s.addr 
				From addresses s
				LEFT OUTER JOIN profiles c 
				ON s.user_id = c.user_id
				GROUP BY s.id
				ORDER BY s.id ASC
// ------------------ 下面是结果			
查找的结果是: &{{1 2019-03-22 15:30:32 +0800 CST 2019-03-22 15:30:37 +0800 CST <nil>} 1 zzj 13559775273 深圳益田路 false}
查找的结果是: &{{2 2019-03-22 15:31:08 +0800 CST 2019-03-22 15:31:11 +0800 CST <nil>} 2 zzj-7 123 深圳南山 false}
```



