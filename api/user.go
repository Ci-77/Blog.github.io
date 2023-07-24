package api

import (
	"Blog/model"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

// post获取数据方法
func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := model.User{
		Username: username,
		Password: password,
	}
	model.RegisterRouter(&user)
	c.Redirect(301, "/")
}
func ListUser(c *gin.Context) {
	c.HTML(200, "first.html", nil)
}

// 首页路由方法
// func Index(c *gin.Context) {
// 	c.HTML(200, "Addpost.html", nil)
// }

// 注册路由方法
func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

// 登录路由方法
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	u := model.Login(username)
	if u.Username == " " {
		c.HTML(200, "login.html","用户不存在")
		fmt.Println("用户不存在！")

	} else {
		if u.Password != password {
			c.HTML(200, "login.html", "密码错误")
			fmt.Println("密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}

	}
}

// 博客添加数据操作
func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	desc := c.PostForm("desc")
	post := model.Post{
		Title:   title,
		Content: content,
		Desc:    desc,
	}
	model.AddPost(&post)
	c.Redirect(301, "/addpost")
}
func Gopost(c *gin.Context) {
	c.HTML(200, "indexpost.html", nil)
}

// 返回博客列表


// 获取到所有的数据以后，跳转到博客列表页面
func AdmainPost(c *gin.Context) {
	posts := model.GetPost()
	c.HTML(200, "listpost.html", posts)
}
//文章详情页面
func PostDetail(c *gin.Context) {
	id:= c.Param("id")
    var post model.Post
	model.DB.First(&post,id)
	if post.ID==0 {
		c.JSON(404,gin.H{
			"error":"id is not found",
		
		})
		return	
		}
		content := blackfriday.Run([]byte(post.Content))
	c.HTML(200, "readpost.html", gin.H{
		"Title":   post.Title,
		"Content": template.HTML(content),
	})
	}
	//设置文章列表函数
	func ListPost(c *gin.Context)  {
		//返回存放列表的切片
		err,post:=model.GetAll()
		if err!=nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.HTML(200,"listpost.html",gin.H{
			"posts":post,
		})
	}
	//后台操作-获取到博客数据后，跳转到管理后台界面
	func Admin(c *gin.Context)  {
		posts := model.GetPost()
		c.HTML(200,"adminpost.html",posts)
	}
	//删除文章操作
	func Delete(c *gin.Context)  {
		id:=c.Param("id")
		var post model.Post
		model.DB.First(&post,id)
		if post.ID==0 {
			c.JSON(404,"id is not found")
			return
		}
		model.DeletePost(&post)
		c.Redirect(301,"/adminpost")
	}
	//绑定参数,更新文章操作
	func UpdatePost(c *gin.Context)  {
		id:=c.Param("id")
		var post model.Post
		
		//post操作，拿到
		model.DB.First(&post,id)
		//绑定表单数据

		//更新文章数据
		post.Desc=c.PostForm("desc")
		post.Title=c.PostForm("title")
		post.Content=c.PostForm("content")
		//行post操作获取post的值
		model.DB.Save(&post)
		c.Redirect(301,"/updatepost")

	}
	//跳转到编辑页操作
	func ToUpdate(c *gin.Context)  {
		id:=c.Param("id")
		var post model.Post
		model.DB.First(&post,id)
		if post.ID==0 {
			c.JSON(404,"id is not found")
			return
		}
		c.HTML(200,"update.html",gin.H{
			"Title":post.Title,
			"Desc":post.Desc,
			"Context":post.Content,
		})
	}
//返回最新文章列表页
	func GetPostList(c *gin.Context) {
		posts:=model.NewPost()
		c.HTML(200,"listpost.html",posts)
	}
	

