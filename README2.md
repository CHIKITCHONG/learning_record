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

