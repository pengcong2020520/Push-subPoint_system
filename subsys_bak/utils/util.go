package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"subsys/dbs"
	"github.com/gin-gonic/gin"

)

const (
	RECODE_OK         = "0"
	RECODE_DBERR      = "4001"
	RECODE_NODATA     = "4002"
	RECODE_DATAEXIST  = "4003"
	RECODE_DATAERR    = "4004"
	RECODE_SESSIONERR = "4101"
	RECODE_LOGINERR   = "4102"
	RECODE_PARAMERR   = "4103"
	RECODE_USERERR    = "4104"
	RECODE_HASHERR    = "4105"
	RECODE_PWDERR     = "4106"
	RECODE_BCOSERR    = "4201"
	
	RECODE_UNKNOWERR  = "4500"
	
)

var recodeText = map[string]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库操作错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或密码错误",
	RECODE_PWDERR:     "密码错误",
	RECODE_BCOSERR:    "BCOS连接错误",
	
	RECODE_UNKNOWERR:   "未知错误",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}

type Gin struct {
	C *gin.Context
}

type Resp struct {
	Errno  string      `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

//resp数据响应
func (g *Gin) ResponseData(httpCode int, resp *Resp) {
	resp.ErrMsg = RecodeText(resp.Errno)
	g.C.JSON(httpCode, *resp)
	return
}

//加密算法 MD5
func EncodeMD5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

func CheckUserid(id string) ([]map[string]string, bool) {
	sql := fmt.Sprintf("select * from user where userid='%s'", id)
	fmt.Println(sql);
	m, n, err := dbs.DBQuery(sql)
	if err != nil || n <= 0 {
		return make([]map[string]string, 0), false
	}
	return m, true
}