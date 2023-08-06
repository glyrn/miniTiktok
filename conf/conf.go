package conf

// ftp服务器相关配置
const FtpHost = "39.101.72.240:210" // 这里的控制端口映射是210 默认映射端口是21
const FtpUsername = "user"
const FtpPassword = "123456"
const LiveTime = 120

// HostSSH SSH配置
const HostSSH = "39.101.72.240"
const UserSSH = "haha"
const PasswordSSH = "123456"
const TypeSSH = "password"
const PortSSH = 22
const MaxMsgCount = 100
const SSHLiveTime = 10 * 60

// VideoCount 每次获取视频流的数量
const VideoCount = 5

// UrlPre 存储的图片和视频的链接的前缀信息
const PlayUrlPre = "http://39.101.72.240:9002/videos/"
const CoverUrlPre = "http://39.101.72.240:9002/images/"
