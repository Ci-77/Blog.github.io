package model

import (
	"log"

)
type User struct {
	Username string 
	Password string
}

func (User) TableName() string {
	return "blog_user"
}

type Post struct {
	//创建博客模型
    ID            uint   
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
}


func (Post) TableName() string {
	return "blog_article"
}
func RegisterRouter(user *User) {
	//添加数据库
	err := DB.Create(user).Error
	if err != nil {
		log.Println("add data err...", err)
	}
}
//验证登录，查询数据库对比
func Login(username string) User {
	var user User
	//查询数据库
	err := DB.Where("username=?", username).Find(&user)
	if err != nil {
		log.Println("find err...", err)
	}
	return user
}

// 添加博客数据库操作
func AddPost(post *Post) {
	err := DB.Create(post).Error
	if err != nil {
		log.Println("Add post err..", err)
	}
}

// 创建一个数据库切片，存放所有的数据库数据
//查找博客数据库操作(全部查询)
func GetAll() (posts []Post,err error) {
	//声明存放所有博客列表的切片
	 err=DB.Find(&posts).Error
	return 
}
func GetPost() []Post {
	//声明存放所有博客列表的切片
	 var posts=make([]Post,20)
	 DB.Find(&posts)
	 return posts
	
}

// 查询数据库的单个操作
// func GetPost(pid int)(*Post) {
// 	var post Post
// 	DB.First(&post,pid)
// 	return &post
// }
//更新博客文章操作
func Update(post *Post)  {
	 err:=DB.Save(&post).Error
	 if err!=nil {
		log.Println("更新数据库失败")
	 }
}
func NewPost() []Post {
	var posts =make([]Post,100)
if err := DB.Find(&posts).Error; err != nil {
	log.Printf("返回出错%s\n",err)
	
}
return posts
}







//删除博客文章操作
func DeletePost(post *Post) (err error)  {
	
	return DB.Delete(post).Error
}
