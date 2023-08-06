package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniTiktok/dao"
	"miniTiktok/service"
	"net/http"
)

// 用于用户登录注册的返回
type LoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(context *gin.Context) {

	userName := context.Query("username")
	passWord := context.Query("password")

	fmt.Println("注册信息请求中得到了" + userName)
	fmt.Println("注册信息请求中得到了" + passWord)

	// 获取业务层实例
	UserService_Impl := service.UserServiceImpl{}

	user, err := UserService_Impl.GetUserByName(userName)
	if user.Name == userName && err == nil {
		fmt.Println("用户已存在")
		// 用户存在
		context.JSON(http.StatusOK, LoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "用户已存在",
			},
		})
	} else {
		// 进入注册流程
		fmt.Println("开始注册")
		User := dao.User_dao{
			Name:     userName,
			Password: passWord,
		}
		// 插入数据 新增用户信息
		flag := UserService_Impl.Insert2User(&User)
		if flag == false {
			fmt.Println("新增用户信息失败")
		}
		// 用户信息插入成功
		fmt.Println(userName + " 注册成功")
		// 从表中得到新增的user对象 防止注册失败 而使用未入库的对象, 从而得到user的id
		userInTable, err := UserService_Impl.GetUserByName(userName)
		if err != nil {
			fmt.Println("从表中获取对象错误")
		}
		token := service.CreateTokenByUserName(userInTable.Name)
		fmt.Println("用户的id是", userInTable.Id)

		context.JSON(http.StatusOK, LoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   int64(userInTable.Id),
			Token:    token,
		})
	}
}

func Login(context *gin.Context) {
	userName := context.Query("username")
	passWord := context.Query("password")

	UserService_Impl := service.UserServiceImpl{}

	user, err := UserService_Impl.GetUserByName(userName)

	if err != nil {
		fmt.Println("GetUserByName 出现错误")
	}

	if passWord == user.Password {
		// 通过登录验证
		token := service.CreateTokenByUserName(userName)
		context.JSON(http.StatusOK, LoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   int64(user.Id),
			Token:    token,
		})
	} else {
		context.JSON(http.StatusOK, LoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "用户密码错误",
			},
		})
	}

}
