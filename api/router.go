package api

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine){
	//注册路由组
	r.POST("/register",RegisterUser)//路由必须和get路由保持一致
	r.GET("/register",GoRegister)
	//点击首页路由
	r.GET("/",AdmainPost)
	//根据id，跳转到详情页
	r.GET("/listpost/:id",PostDetail)
	//登录路由组
	r.GET("/login",GoLogin)
	r.POST("/login",Login)
	//添加博客列表
	r.POST("indexpost",AddPost)
	r.GET("/addpost",AdmainPost)
	r.GET("/listpost",ListPost)
	//
    r.GET("/indexpost",Gopost)
	r.GET("/adminpost",Admin)
	//删除操作
	r.GET("/delete/:id",Delete)
//跳转到文章更新界面
	r.GET("/update/:id",ToUpdate)
		//更新数据库操作
	r.POST("/updatepost/:id",UpdatePost)
	r.GET("/updatepost",GetPostList)


	
	
	
}