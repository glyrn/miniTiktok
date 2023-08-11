package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"miniTiktok/dao"
	"strconv"
	"time"
)

type UserServiceImpl struct {
	// 关注模块服务
	//FollowService
	// 待添加 点赞模块等
}

var fsi = FollowServiceImpl{}

func (UserServiceImpl *UserServiceImpl) GetUserByName(name string) (dao.User_dao, error) {
	user_dao, err := dao.GetUserByName(name)
	if err != nil {
		fmt.Println("用户不存在与数据库")
		fmt.Println(err)
		return user_dao, err
	}
	fmt.Println("用户已经找到")
	return user_dao, nil
}

func (UserServiceImpl *UserServiceImpl) Insert2User(user *dao.User_dao) bool {

	if dao.Insert2User(user) == false {
		fmt.Println("数据插入失败")
		return false
	}
	fmt.Println("数据插入成功")
	return true

}

// 未登录状态 获取 根据userID 获取到user组装后得到对象
func (UserServiceImpl *UserServiceImpl) GetUser_serviceById(userId int64) (User_service_final, error) {
	user := User_service_final{
		Id:            0,
		Name:          "",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
		Avatar:        "",
	}
	user_dao, err := dao.GetUserById(userId)
	if err != nil {
		fmt.Println("获取dao层usr失败")
		return user, err
	}
	fmt.Println("获取dao层usr成功")

	// 获取关注人数, 被关注人数
	//followcount, followercount := UserServiceImpl.GetFansDndAttention(userId)
	followcount, followercount := fsi.GetFansDndAttention(userId)

	//fmt.Println(followcount, followercount)
	//followcount := int64(12)
	//followercount := int64(10)

	// 随机头像地址
	avatarAPI := "https://api.multiavatar.com/" + string(userId) + ".png"

	// 用户信息获取成功
	// 组装信息
	user = User_service_final{
		Id:            userId,
		Name:          user_dao.Name,
		FollowCount:   followcount,
		FollowerCount: followercount,
		IsFollow:      false,
		Avatar:        avatarAPI,
	}
	fmt.Println(user)

	return user, err

}

func CreateTokenByUserName(userName string) string {
	user, err := UserService.GetUserByName(new(UserServiceImpl), userName) // 先把接口赋值 然后传参
	if err != nil {
		fmt.Println("CreateTokenByUserName 中出现错误")
	}
	token := CreateTokenByUser_dao(user)
	fmt.Println("生成的token", token)
	return token
}

// 这里是用于注入payload部分
// 根据user信息创建token
func CreateTokenByUser_dao(user dao.User_dao) string {

	fmt.Println("开始合成token")

	// 标准接口类型
	claims := jwt.StandardClaims{

		Audience: user.Name,
		// token 过期时间 这里存的是时间戳 表示经过多少时间过期 这里暂时写死数据 后期放配置文件中
		ExpiresAt: time.Now().Unix() + int64(60*60*24),
		// 这里把id 转化为10进制的字符串
		Id: strconv.FormatInt(user.Id, 10),
		// 签发时间
		IssuedAt: time.Now().Unix(),
		// 签发者
		Issuer: "tiktok",
		// 生效时间
		NotBefore: time.Now().Unix(),
		// jwt字段的主题 用于什么字段
		Subject: "token",
	}

	var jwtSecret = []byte("123456") // 这里使加密算法的私钥  token需要同时有公钥和私钥才能解析

	//用加密算法生成标准JWT结构体
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 对payload进行数字签名  --用header和payload以及密钥进行数字签名 得到可信token
	token, err := tokenClaims.SignedString(jwtSecret)

	if err != nil {
		fmt.Println("token 生成失败", err)
		return "token 生成失败"
	} else {
		token = "Bearer " + token // 这个是jwt的规范 用bearer前缀
		fmt.Println("token 生成成功", token)
		return token
	}

}

// parseToken 解析给定的 JWT 令牌并返回 JWT 负载 (StandardClaims) 及任何遇到的错误。
// 该函数接收 JWT 令牌作为输入，并使用指定的密钥 来验证令牌的签名。
// 如果令牌有效并成功解析，它将返回 JWT 负载 (StandardClaims) 的指针以及 nil 错误。
// 如果在解析过程中遇到错误或者令牌无效，则返回 nil 负载和相应的错误。
func ParseToken(token string) (*jwt.StandardClaims, error) {
	// 使用 jwt.ParseWithClaims 解析 JWT 令牌并提取 StandardClaims 负载。
	// 使用提供的字节切片  作为签名验证所需的密钥。
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		fmt.Println([]byte("123456"))
		return []byte("123456"), nil
	})

	// 检查解析和验证是否成功 (err == nil)，以及 jwtToken 是否非空。
	if err == nil && jwtToken != nil {
		// 类型断言检查 jwtToken.Claims 是否为 *jwt.StandardClaims 类型。
		// 同时，检查令牌是否有效，调用 jwtToken.Valid 进行验证。
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			// 如果一切成功，返回解析的 JWT 负载 (StandardClaims) 和 nil 错误。
			return claim, nil
		}
	}
	// 如果在解析过程中遇到错误或者令牌无效，则返回 nil 负载和相应的错误。
	return nil, err
}
