package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math"
	"miniTiktok/conf"
	"miniTiktok/midddleWare/redis"
	"miniTiktok/response"
	"net/http"
	"strconv"
	"time"
)

type Claims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

var JwtSecret = []byte(conf.JwtKey) // 这里使加密算法的私钥
// token 鉴权中间件
func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 取出token auth
		token := context.Query("token")
		if token == "" {
			token = context.Request.PostFormValue("token")
		}
		Auth(context, token)
	}
}

// 验证
func Auth(context *gin.Context, tokenStr string) {

	// 如果 用户不合法
	if len(tokenStr) == 0 {
		// 拦截请求
		context.Abort()
		// 传回json
		context.JSON(http.StatusOK, response.Response{
			StatusCode: http.StatusUnauthorized,
			StatusMsg:  "未获取到token",
		})
	}

	// 解析token
	token, ok := ParseToken(tokenStr)
	if !ok {
		// 鉴权后不合法
		context.Abort()
		context.JSON(http.StatusUnauthorized, response.Response{
			StatusCode: http.StatusForbidden,
			StatusMsg:  "token 错误",
		})
	} else if GetJWTFromID(strconv.FormatInt(token.UserId, 10)) != tokenStr {
		fmt.Println("该token已经作废")
		context.Abort()
		context.JSON(http.StatusUnauthorized, response.Response{
			StatusCode: http.StatusForbidden,
			StatusMsg:  "该token已经作废，请重新登录",
		})
	} else {
		fmt.Println("token正确 身份验证通过")
	}

	context.Set("userId", strconv.FormatInt(token.UserId, 10))
	context.Set("userName", token.UserName)
	// 身份验证通过 放行
	context.Next()
}

// 插入JWT或者修改
func SetJWT2Redis(id string, jwt string) error {
	err := redis.Rdb15.Set(redis.Ctx, "userID:"+id, jwt, 60*60*24*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

// 查询JWT
func GetJWTFromID(id string) string {
	JWT, err := redis.Rdb15.Get(redis.Ctx, "userID:"+id).Result()
	if err != nil {
		return ""
	}
	// 如果存在 延长jwt令牌过期时间
	redis.Rdb15.Set(redis.Ctx, "userID:"+id, JWT, 60*60*24*time.Second)
	return JWT
}

/*
这里是用于注入payload部分
生成 token
*/
func CreateToken(userId int64, userName string) string {

	//fmt.Println("开始合成token")

	claims := Claims{
		UserId:   userId,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			Subject:   "token",           // 主题
			Issuer:    "tiktok",          // 签发者
			IssuedAt:  time.Now().Unix(), // 签发时间
			ExpiresAt: math.MaxInt64,     // 过期时间 永不过期 -> 通过redis主动控制过期时间
			NotBefore: time.Now().Unix(), // 生效时间
		},
	}

	//用加密算法生成标准JWT结构体
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 对payload进行数字签名  --用header和payload以及密钥进行数字签名 得到可信token
	token, err := tokenClaims.SignedString(JwtSecret)

	if err != nil {
		fmt.Println("token 生成失败", err)
		return "token 生成失败"
	} else {
		fmt.Println("token 生成成功", token)
		return token
	}

}

func ParseToken(token string) (*Claims, bool) {

	jwtToken, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if key, _ := jwtToken.Claims.(*Claims); jwtToken.Valid { // 检查令牌是否有效，调用 jwtToken.Valid 进行验证
		return key, true
	} else {
		return nil, false
	}
}
