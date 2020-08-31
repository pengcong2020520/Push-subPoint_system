package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/common"
	"math/big"
	"net/http"
	"strconv"
	"subsys/bcos"
	"subsys/config"
	"subsys/dbs"
	"subsys/utils"
)

const (
	MAX_PAGE = 50
)

func PushLog(c *gin.Context) {
	respG := utils.Gin{C: c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)
	//解析数据
	userid := c.PostForm("userid")
	smonth := c.PostForm("month")
	month, err := strconv.Atoi(smonth)
	if err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		return
	}

	logJsonData := c.PostForm("log")
	adminpassword := c.PostForm("adminpassword")

	logs := &dbs.Logs{
		UserId:  userid,
		Month:   month,
		Loginfo: logJsonData,
	}
	//loginfo, err := c.FormFile("log")
	//if err != nil {
	//	fmt.Println("Failed to get loginfo", err)
	//	resp.Errno = utils.RECODE_LOGDATAERR
	//	return
	//}
	////处理日志文件
	//logJsonData, err = utils.ParseFileData(loginfo)
	//if err != nil || logJsonData == "" {
	//	resp.Errno = utils.RECODE_RDFILEERR
	//	fmt.Println("Failed to parse log file! ", err)
	//	return
	//}

	//连接节点并获得合约入口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	//获取管理员签名
	opts, err := bcos.GetAdminOpt(adminpassword)
	if err != nil {
		resp.Errno = utils.RECODE_OPTERR
		fmt.Println("Failed to get admin opt ! ", err)
		return
	}
	//调用合约
	_, err = instance.PushLog(opts, userid, logJsonData, big.NewInt(int64(month)))
	if err != nil {
		fmt.Println("Failed to pushlog")
		resp.Errno = utils.RECODE_GETCONTRAERR
		return
	}
	resp.Data = logs
	fmt.Println("push userid"+ userid + "log success")
}

func QueryLog(c *gin.Context) {
	respG := utils.Gin{C: c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)

	userid := c.PostForm("userid")
	sbegin := c.PostForm("begin")
	begin, _ := strconv.Atoi(sbegin)
	send := c.PostForm("end")
	end, _ := strconv.Atoi(send)

	//连接节点并获得合约入口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	//获取管理员bcos账户
	bcosacc, _, err := utils.GetKeyFile(config.Config.Bcos.Keydir)
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	opts := &bind.CallOpts{
		From : common.HexToAddress(bcosacc),
	}
	//调用合约

	logInfos, err := instance.QueryLog(opts, userid, big.NewInt(int64(begin)), big.NewInt(int64(end)))
	if err != nil {
		resp.Errno =utils.RECODE_GETCONTRAERR
		fmt.Println("Failed to query logs")
		return
	}


	//处理响应信息
	total_page := int(len(logInfos))/MAX_PAGE + 1
	current_page := 1
	mapResp := make(map[string]interface{})
	mapResp["total_page"] = total_page
	mapResp["current_page"] = current_page
	mapResp["loginfos"] = logInfos
	resp.Data = mapResp
}



