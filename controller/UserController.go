package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"xietong.me/ginessential/common"
	"xietong.me/ginessential/dto"
	"xietong.me/ginessential/model"
	"xietong.me/ginessential/response"
	"xietong.me/ginessential/util"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	err := ctx.ShouldBind(&requestUser)
	if err != nil {
		log.Println("Bind Error:" + err.Error())
	}
	//获取参数
	name := requestUser.Name
	//telephone := requestUser.Telephone
	password := requestUser.Password
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//如果名称为空给一个随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:     name,
		Password: string(hasePassword),
	}
	DB.Create(&newUser)
	//发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error:%v", err)
		return
	}
	//返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

func Login(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.UserId == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发送token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
	}

	//返回结果
	response.Success(c, gin.H{"token": token}, "登陆成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func Remove(ctx *gin.Context) {
	//db := common.GetDB()
	//user, _ := ctx.Get("user")
	//tel := user.(model.User).Telephone
	//err := db.Where("Telephone = ?", tel).Delete(&model.User{})
	//log.Print(err.Error)
}

func UpdateByUser(ctx *gin.Context) {
	var requestUser = model.User{}
	//originUser, _ := ctx.Get("user")
	err := ctx.ShouldBind(&requestUser)
	if err != nil {
		log.Println(err.Error())
		response.Fail(ctx, gin.H{"name": "error"}, "参数错误")
	}
	name := requestUser.Name
	password := requestUser.Password
	updateUser := model.User{
		Name:     name,
		Password: password,
	}
	db := common.GetDB()
	db.Save(&updateUser)
}
