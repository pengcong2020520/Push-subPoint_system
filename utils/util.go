package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"subsys/dbs"
	"github.com/gin-gonic/gin"

)

const (
	RECODE_OK           = "0"
	RECODE_DBERR        = "4001"
	RECODE_NODATA       = "4002"
	RECODE_DATAEXIST    = "4003"
	RECODE_DATAERR      = "4004"
	RECODE_SESSIONERR   = "4101"
	RECODE_LOGINERR     = "4102"
	RECODE_PARAMERR     = "4103"
	RECODE_USERERR      = "4104"
	RECODE_HASHERR      = "4105"
	RECODE_PWDERR       = "4106"
	RECODE_SETPWDERR    = "4107"
	RECODE_BCOSERR      = "4201"
	RECODE_OPTERR       = "4202"
	RECODE_GETCONTRAERR = "4203"

	RECODE_LOGDATAERR   = "4301"
	RECODE_RDFILEERR    = "4302"

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
	RECODE_SETPWDERR:   "修改密码错误",
	RECODE_BCOSERR:    "BCOS连接错误",
	RECODE_OPTERR:     "认证失败",
	RECODE_GETCONTRAERR: "调用合约失败",
	RECODE_LOGDATAERR:   "日志数据错误",
	RECODE_RDFILEERR:    "读取日志文件错误",
	
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

func GetKeyFile(keyDir string) (string, *os.File, error) {
	data, err := ioutil.ReadDir(keyDir)
	if err != nil {
		fmt.Println("Failed to read key dir! ", err)
		return "", nil, err
	}
	accAddr := data[0].Name()
	file, err := os.Open(keyDir + "/" + accAddr)
	if err != nil {
		fmt.Println("Failed to read accAddr key !")
		return accAddr, nil, err
	}
	return accAddr, file, nil
}

func ParseFileData(head *multipart.FileHeader) (string, error) {
	src, err := head.Open() //打开源文件
	if err != nil {
		fmt.Println("Failed to open file")
		return "", err
	}
	cData := make([]byte, head.Size)
	n, err := src.Read(cData)
	if err != nil || int64(n) != head.Size {
		fmt.Println("Failed to read file ! ", err)
		return "", err
	}
	jsonData, err := json.Marshal(cData)
	if err != nil {
		fmt.Println("Failed to json marshal data ! ")
		return "", err
	}
	return string(jsonData), nil
}