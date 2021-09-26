package controller

import (
	"gintest1/common"
	"gintest1/model"
	"gintest1/response"
	"gintest1/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//post新建用户
func Register(c *gin.Context) {
	//得到数据库链接
	DB := common.GetDB()
	//注册参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//验证手机位数
	if len(telephone) != 11 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}
	//验证密码不得少于6位
	if len(password) < 6 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}
	//缺省name时自动补齐10位
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	//服务端打印日志
	log.Println(name, telephone, password)
	//查询数据库手机号是否被注册
	if isTelephoneExist(DB, telephone) {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号已经存在")
		return
	}
	//创建用户
	//加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c,http.StatusInternalServerError,500,nil,"加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	//数据写入数据库
	DB.Create(&newUser)
	response.Success(c,nil,"成功注册")
}

//用户登陆
func Login(c *gin.Context) {
	//得到数据库链接
	DB := common.GetDB()
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//验证数据
	//判断手机号格式是否合法
	if len(telephone) != 11 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}
	//判断密码格式是否合法
	if len(password) < 6 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号不存在")
		return
	}
	//密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c,http.StatusBadRequest,400,nil,"密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c,http.StatusInternalServerError,500,nil,"系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	//发放结果
	response.Success(c,gin.H{"token": token},"登陆成功")
}

//得到数据库中的用户数据 返回上下文中的user数据
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": user},
	})
}

//判断手机号是否已经被注册
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
