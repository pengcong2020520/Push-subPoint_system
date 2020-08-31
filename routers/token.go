package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/common"

	"subsys/bcos"
	"subsys/dbs"
	"subsys/utils"
	"subsys/config"

	"math/big"
	"net/http"
	"time"
	"fmt"
)

//挖矿
func TokenMint(c *gin.Context) {
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp *utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, resp)

	//2. 获取数据 并解析
	ts := &dbs.TokenMB{}
	err := c.ShouldBind(ts)
	if err != nil {
		fmt.Println("Failed to bind transaction")
		return
	}
	ts.ToUserid = c.PostForm("touserid")
	ts.Value = c.GetInt("value")
	adminpassword := c.PostForm("adminpassword")
	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	//4. 获取管理员签名
	opts, err := bcos.GetAdminOpt(adminpassword)
	if err != nil {
		resp.Errno = utils.RECODE_OPTERR
		return
	}
	//5. 调用挖矿合约
	_, err = instance.Mint(opts, ts.ToUserid, big.NewInt(int64(ts.Value)))
	if err != nil {
		resp.Errno = utils.RECODE_GETCONTRAERR
		return
	}
	ts.Timestamp = time.Now().Format("20060102150405")
	resp.Data = ts
	fmt.Println("TokenMint success! ")
}

//销毁token

//func TokenBurn(c *gin.Context) {
//	//1. 响应数据初始化
//	respG := utils.Gin{C: c}
//	var resp *utils.Resp
//	resp.Errno = utils.RECODE_OK
//	defer respG.ResponseData(http.StatusOK, resp)
//
//	//2. 获取数据 并解析
//	ts := &dbs.TokenMB{}
//	err := c.ShouldBind(ts)
//	if err != nil {
//		fmt.Println("Failed to bind transaction")
//		return
//	}
//  ts.ToUserid = c.PostForm("touserid")
//	ts.Value = c.PostForm("value")
//	adminpassword := c.PostForm("adminpassword")
//	//3. 连接网络并获取control合约接口
//	instance, cli, err := bcos.GetControlInstance()
//	if err != nil {
//		resp.Errno = utils.RECODE_BCOSERR
//		return
//	}
//	defer cli.Close()
//	//4. 获取管理员签名
//	opts, err := bcos.GetAdminOpt(adminpassword)
//	if err != nil {
//		resp.Errno = utils.RECODE_OPTERR
//		return
//	}
//	//5. 调用挖矿合约
//	_, err := instance.Burn(opts, ts.ToUserid, big.NewInt(int64(ts.Value)))
//	if err != nil {
//		resp.Errno = utils.RECODE_GETCONTRAERR
//		return
//	}
//	ts.Timestamp = time.Now().Format("20060102150405")
//	resp.Data = ts
//	fmt.Println("TokenBurn success! ")
//}

func Transfer(c *gin.Context) {
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp *utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, resp)

	//2. 获取数据 并解析
	ts := &dbs.Transfer{}
	account := &dbs.Account{} //转账人账户

	ts.FromUserid = c.PostForm("fromuserid")
	ts.ToUserid = c.PostForm("touserid")
	ts.Value = c.GetInt("value")

	adminpassword := c.PostForm("adminpassword")
	fromPass := c.PostForm("frompassword")

	account.UserId = ts.FromUserid
	account.Password = fromPass
	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	//4. 获取from userid的许可
	bcosacc, _, err := utils.GetKeyFile(config.Config.Bcos.Keydir)
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		fmt.Println("Failed to get keyfile ! ", err)
		return
	}
	optsLogin := bind.CallOpts{
		From : common.HexToAddress(bcosacc),
	}
	_, err = instance.Login(&optsLogin, account.UserId, account.Password)
	if err != nil {
		resp.Errno = utils.RECODE_USERERR
		fmt.Println("Fial to identify from userid and password !", err)
		return
	}
	//5. 获取管理员签名
	opts, err := bcos.GetAdminOpt(adminpassword)
	if err != nil {
		resp.Errno = utils.RECODE_OPTERR
		return
	}
	//6. 调用合约
	_, err = instance.Transfer(opts, ts.FromUserid, ts.ToUserid, big.NewInt(int64(ts.Value)))
	if err != nil {
		resp.Errno = utils.RECODE_GETCONTRAERR
		return
	}

	ts.Timestamp = time.Now().Format("20060102150405")

	resp.Data = ts
	fmt.Println("Transfer success! ")
}