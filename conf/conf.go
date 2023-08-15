package conf

// ftp服务器相关配置
const FtpHost = "39.101.72.240:210" // 这里的控制端口映射是210 默认映射端口是21
const FtpUsername = "user"
const FtpPassword = "123456"
const LiveTime = 120

// HostSSH SSH配置
const HostSSH = "39.101.72.240"
const UserSSH = "7m5GUXUK09tu5jQz.knr4RfKIQA+oo0tR92/xfnIP6JU="
const PasswordSSH = "tFJ7I5SD5fQQV1Ao.Q/VwOeIhPxg78Dc9hNtd/mPHz1QrYuKlFuczQvHuRw=="
const TypeSSH = "password"
const PortSSH = 22
const SSHLiveTime = 10 * 60

// VideoCount 每次获取视频流的数量
const VideoCount = 3

// UrlPre 存储的图片和视频的链接的前缀信息
const PlayUrlPre = "http://39.101.72.240:9002/videos/"
const CoverUrlPre = "http://39.101.72.240:9002/images/"

// 数据库配置 可根据生产环境改动
const User = "guest"         // 用户名
const Pass = "123456"        // 密码
const Adrr = "39.101.72.240" // 地址
const Port = "3306"          // 端口
const Dbname = "tiktok"      // 数据库名称

// 加密算法密钥
const KeyString = "mysecretkey123456789012345600000"

// jwt密钥
const JwtKey = "123456"
