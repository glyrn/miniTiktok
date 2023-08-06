package ffmpeg

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"miniTiktok/conf"
	"time"
)

/**
这里是在服务器安装了ffmpeg，然后通过ssh远程连接服务器
调用shell命令，完成视频截图，并储存在指定的位置
*/

type Ffmsg struct {
	VideoName string
	ImageName string
}

var ClientSSH *ssh.Client

func InitSSH() {
	var err error
	//创建sshp登陆配置
	SSHconfig := &ssh.ClientConfig{
		Timeout:         5 * time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            conf.UserSSH,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //通过密码方式连接服务器，但是不够安全

	}

	// 这里用密码方式连接
	if conf.TypeSSH == "password" {
		SSHconfig.Auth = []ssh.AuthMethod{ssh.Password(conf.PasswordSSH)}
	}
	// 登录，创建会话
	addr := fmt.Sprintf("%s:%d", conf.HostSSH, conf.PortSSH)
	ClientSSH, err = ssh.Dial("tcp", addr, SSHconfig)
	if err != nil {
		fmt.Println("创建会话失败", err)
	}
	fmt.Println("创建会话成功")

	go keepAlive()

}

// SaveVideoCover 通过远程调用ffmpeg命令来获取视频截图，并保存
func SaveVideoCover(videoName string, imageName string) error {
	session, err := ClientSSH.NewSession()
	if err != nil {
		fmt.Println("创建会话失败")
	}
	defer session.Close()

	//执行远程shell命令 /root/ffmpeg/bin/ffmpeg -ss 00:00:01 -i /root/ftp/Ftpfile/user/videos/test2.mp4 -vframes 1 /root/ftp/Ftpfile/user/images/test3.jpg
	combo, err := session.CombinedOutput("/root/ffmpeg/bin/ffmpeg -ss 00:00:01 -i /root/ftp/Ftpfile/user/videos/" + videoName + ".mp4 -vframes 1 /root/ftp/Ftpfile/user/images/" + imageName + ".jpg")
	if err != nil {
		fmt.Println("远程执行cmd 失败", err)
		return err
	}
	fmt.Println("远程执行命令:", string(combo))
	return nil
}

// 维持会话长连接
func keepAlive() {
	time.Sleep(time.Duration(conf.SSHLiveTime) * time.Second)
	session, _ := ClientSSH.NewSession()
	session.Close()
}
