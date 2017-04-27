// package conf 为配置相关.
package conf

import (
	"github.com/astaxie/beego"
	"strings"
)

// 登录用户的Session名
const LoginSessionName = "LoginSessionName"

const RegexpEmail  = `^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`

const RegexpAccount = `^[a-zA-Z][a-zA-z0-9]{2,50}$`

// PageSize 默认分页条数.
const PageSize  = 15

// 用户权限
const (
	// 超级管理员.
	MemberSuperRole = 0
	//普通管理员.
	MemberAdminRole = 1
	//普通用户.
	MemberGeneralRole = 2
)

const (
	// 创始人.
	BookFounder = 0
	//管理者
	BookAdmin = 1
	//编辑者.
	BookEditor = 2
	//观察者
	BookObserver = 3
)

// app_key
func GetAppKey()  (string) {
	return beego.AppConfig.DefaultString("app_key","go-git-webhook")
}

func GetDatabasePrefix() string  {
	return beego.AppConfig.DefaultString("db_prefix","md_")
}
//获取默认头像
func GetDefaultAvatar() string {
	return beego.AppConfig.DefaultString("avatar","/static/images/headimgurl.jpg")
}

func GetTokenSize() int {
	return beego.AppConfig.DefaultInt("token_size",12)
}

func GetDefaultCover() string {
	return beego.AppConfig.DefaultString("cover","/static/images/book.jpg")
}

func GetUploadFileExt()  []string {
	ext := beego.AppConfig.DefaultString("upload_file_ext","png|jpg|jpeg|gif|txt|doc|docx|pdf")
	
	temp := strings.Split(ext,"|")
	
	exts := make([]string,len(temp))
	
	i := 0
	for _,item := range temp {
		if item != "" {
			exts[i] = item
			i++
		}
	}
	return exts
}

func IsAllowUploadFileExt(ext string) bool  {

	if strings.HasPrefix(ext,".") {
		ext = string(ext[1:])
	}
	exts := GetUploadFileExt()

	for _,item := range exts {
		if strings.EqualFold(item,ext) {
			return  true
		}
	}
	return false
}