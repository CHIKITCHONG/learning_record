### golang 循环引用

#### 接口方法的循环引用解耦,先看图

```
结构目录
- main_project(dic)
	- article(dic)
		- article.go
	- article-favor(dic)
		-article-favor.go
	- main.go
```

- main.go

  ```go
  article_svc := article.New(article.NewArticleDB(),article_favor.NewArticleFavorDB())
  article_favor_svc := article_favor.New(article_favor.NewArticleFavorDB(),article.NewArticleDB())
  ```

- article

  - article.go

    ```go
    func New(idb IArticleDB, fs article_favor.IArticleLikeDB) *Service {
    	return &Service{idb: idb, fs: fs}
    }
    
    type Service struct {
        idb IArticleDB
    	fs article_favor.IArticleLikeDB
    }
    
    // 通过一个函数调用将 IArticleDB 作为一个实例返回
    func NewArticleDB() *NewArticleDB_ {
        return &NewArticleDB_{}
    }
    
    type IArticleDB interface {
    	List(fn query.SQFunc) (err error)
    }
    
    // 启用一个结构体实现 IArticleDB 的接口, 成为接收者
    type NewArticleDB_ struct {}
    
    func (s *NewArticleDB_) List(fn query.SQFunc) (err error) {}
    ```

- article-favor

  - article-favor.go

    ```go
    func New(idb IArticleDB) *Service {
    	return &Service{idb: idb}
    }
    
    // Service represents user application service
    type Service struct {
        idb IArticleDB
    }
    
    type NewArticleFavorDB_ struct {}
    
    // 通过一个函数调用将 IArticleDB 作为一个实例返回
    func NewArticleFavorDB() *NewArticleFavorDB_ {
        return &NewArticleFavorDB_{}
    }
    
    type IArticleFavorDB interface {
    	List(fn query.SQFunc) (err error)
    }
    
    // 启用一个结构体实现 IArticleFavorDB 的接口, 成为接收者
    func (s *NewArticleFavorDB_) List(fn query.SQFunc) (err error) {}
    
    // 在这里我想用 article.go 的 IArticleDB 接口的 List 接口
    ```

  

####  但运行main.go， 会提示循环引用

```
article_svc := article.New(article.NewArticleDB(),article_favor.NewArticleFavorDB())
article_favor_svc := article_favor.New(article_favor.NewArticleFavorDB(),article.NewArticleDB())

// 若在 article_favor 中想用 article 的 List 方法, 需要把需要把接口单独写到另一个服务
```

#### 在新的文件里，将 article 中的 articleDB 单独抽离

```
结构目录
- main_project(dic)
	- article(dic)
		- article.go
	- article-favor(dic)
		-article-favor.go
	- main.go
	- tools(dic)
		- cycle_use.go
```

#### 将接口抽离

```
func ContractDB() *Cycle_contract {
	return &Cycle_contract{}
}

type Cycle_contract struct {
}

type Cycle_db_contract interface {
	List(fn query.SQFunc) (err error)
}

func (this *Cycle_contract) List(fn query.SQFunc) (err error) {}


// main.go 现调用如下,就不会报错
article_svc := article.New(article.NewArticleDB(),article_favor.NewArticleFavorDB())
article_favor_svc := article_favor.New(article_favor.NewArticleFavorDB(),tools.ContractDB())
```

