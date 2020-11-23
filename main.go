package main

import (
	"errors"
	"fmt"
	"github.com/caochong01/gin-one/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
)

func main() {
	router := gin.Default()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// TODO 数据库练习
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:caochong0521@tcp(127.0.0.1:3306)/gogogo?charset=utf8mb4&parseTime=True&loc=Local"
	//_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:caochong0521@tcp(127.0.0.1:3306)" +
			"/gogogo?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         1024,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   "",   // 表名前缀`t_`，则`User` 的表名应该是 `t_users`
		SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
	}})
	if err != nil {
		fmt.Println("数据库连接异常错误：")
		panic(err)
	}

	var diary entity.Diary
	var count int64

	// 获取全部记录
	result := db.Find(&diary).Count(&count)
	fmt.Println("获取全部记录: ", count, result.Error, result.RowsAffected)

	result = db.Create(&entity.Diary{
		DiaryTitle:  "nihao",
		DiaryDetail: "nihaocontent",
	})
	fmt.Println(result.Error, result.RowsAffected)

	// 获取第一条记录（主键升序）
	// SELECT * FROM users ORDER BY id LIMIT 1;
	result = db.First(&entity.Diary{})
	fmt.Println(result.Error, result.RowsAffected)
	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)
	// 获取一条记录，没有指定排序字段
	// SELECT * FROM users LIMIT 1;
	result = db.Take(&entity.Diary{})
	fmt.Println(result.Error, result.RowsAffected)
	// 获取最后一条记录（主键降序）
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	result = db.Last(&entity.Diary{})
	fmt.Println(result.Error, result.RowsAffected)
	// 获取全部记录
	rows, err := db.Find(&entity.Diary{}).Rows()
	for rows.Next() {
		columns, errs := rows.Columns()
		if errs != nil {
			panic(errs)
		}
		fmt.Println(columns)
	}
	fmt.Println("获取全部记录: ", result.Error, result.RowsAffected)

	routerUrl(router)
	//runErr := router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//if runErr != nil {
	//	fmt.Println("服务启动异常错误：")
	//	panic(runErr)
	//}
}

func routerUrl(router *gin.Engine) {
	// Per route middleware, you can add as many as you desire.
	router.GET("/benchmark", func(c *gin.Context) {}, func(c *gin.Context) {})

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := router.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(func(c *gin.Context) {})
	{
		authorized.POST("/login", func(c *gin.Context) {})
		authorized.POST("/submit", func(c *gin.Context) {})
		authorized.POST("/read", func(c *gin.Context) {})

		// nested group
		testing := authorized.Group("testing")
		{
			testing.GET("/analytics", func(c *gin.Context) {})
		}
	}

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {})
		v1.POST("/submit", func(c *gin.Context) {})
		v1.POST("/read", func(c *gin.Context) {})
	}

	router.GET("/ping", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"message": name,
		})
	})
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		fmt.Println(c.FullPath() == "/user/:name/*action") // true
	})
}
