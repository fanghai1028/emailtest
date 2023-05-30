package config

import (
	"fmt"
	"net/mail"
)

// 初始化配置文件
func InitConfigFile() {

	// 需要监控的文件
	fileName1 := ""
	fileName2 := ""

	// 写配置信息
	ci := ConfigInfo{
		configFileType: "toml",
		configFileName: "config.toml",
		Config: &TailConfigCollectionEntity{
			MailServer: SmtpMailServerEntity{
				ServerAddress:     "",
				ServerAddressPort: 465,
				NeedLogin:         true,
				LoginUser:         "",
				LoginPassword:     "*******",
				SendMailUserMail:  mail.Address{Name: "", Address: ""},
			},
			Stat: StatConfig{
				Enable:     true,
				ServerName: "10.162.222.210",
			},
			ConfigArr: []TailConfigEntity{
				{
					FileName:            fileName1,
					FileNameUseTemplate: false,
					Subject:             "异常监控报告，服务器：61.235",
					Remark:              "",
					ToMailArr: []mail.Address{{Name: "fanghao", Address: ""},
						{Name: "", Address: ""}}},
				{
					FileName:            fileName2,
					FileNameUseTemplate: true,
					Subject:             "测试邮件标题",
					Remark:              "",
					ToMailArr:           []mail.Address{{Name: "", Address: ""}}}}},
	}
	err := ci.WriteConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("初始化配置文件完成！")
	}

}
