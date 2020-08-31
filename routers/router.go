package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/common"

	"subsys/bcos"
	"subsys/config"
	"subsys/dbs"
	"subsys/utils"
)

const ADMIN_PASS = "admin_password"

func Register(c *gin.Context) {
		/* Test使用
			c.JSON(http.StatusOK, gin.H{"status" : "OK", "function" : "register"})
		*/
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp *utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, resp)
	//2. 解析数据
	account := &dbs.Account{}
	err := c.ShouldBind(account)
	if err != nil {
		fmt.Println(account)
		resp.Errno = utils.RECODE_PARAMERR
	}
	account.UserId = c.PostForm("userid")
	account.Password = c.PostForm("password")

	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()

	// 4. 获取管理员签名
	opts, err := bcos.GetAdminOpt(ADMIN_PASS)
	if err != nil {
		resp.Errno = utils.RECODE_OPTERR
		return
	}
	//5. 调用注册合约函数
	_, err = instance.Register(opts, account.UserId, account.Password)
	if err != nil {
		resp.Errno = utils.RECODE_USERERR
		fmt.Println("Failed to register")
		return
	}
	fmt.Println("userid %s register success!", account.UserId)
}

func Login(c *gin.Context) {
	/* Test使用
	c.JSON(http.StatusOK, gin.H{"status" : "OK", "function" : "register"})
	*/
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp *utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, resp)
	//2. 解析数据
	account := &dbs.Account{}
	err := c.ShouldBind(account)
	if err != nil {
		fmt.Println(account)
		resp.Errno = utils.RECODE_PARAMERR
	}
	account.UserId = c.PostForm("userid")
	account.Password = c.PostForm("password")

	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	//4. 获取签名
	bcosacc, _, err := utils.GetKeyFile(config.Config.Bcos.Keydir)
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		fmt.Println("Failed to get keyfile ! ", err)
		return
	}
	opts := bind.CallOpts{
		From : common.HexToAddress(bcosacc),
	}
	//5. 调用合约 进行登录请求
	ok, err := instance.Login(&opts, account.UserId, account.Password)
	if err != nil {
		resp.Errno = utils.RECODE_LOGINERR
		fmt.Println("Failed to login")
		return
	}
	if !ok {
		resp.Errno = utils.RECODE_LOGINERR
		fmt.Println("Failed to login")
		return
	}
	fmt.Println("userid %s login success!", account.UserId)
}

//  修改密码
func SetPasswd(c *gin.Context) {
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp *utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, resp)

	//2. 获取数据 并解析
	account := &dbs.Account{}
	err := c.ShouldBind(account)
	if err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		fmt.Println("Failed to bind account ! ", err)
		return
	}
	account.UserId = c.PostForm("userid")
	account.Password = c.PostForm("oldpassword")
	newPass := c.PostForm("newpassword")
	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	//4. 获取管理员签名
	opts, err := bcos.GetAdminOpt(ADMIN_PASS)
	if err != nil {
		resp.Errno = utils.RECODE_OPTERR
		return
	}
	//5.调用合约函数 修改密码
	_, err = instance.SetPasswd(opts, account.UserId, account.Password, newPass)
	if err != nil {
		resp.Errno = utils.RECODE_SETPWDERR
		fmt.Println("Failed to set password !")
		return
	}
	account.Password = newPass
	resp.Data = "set password success!"
	fmt.Println("userid %s set password success! ", account.UserId)
}

func GetTotalSupply(c *gin.Context) {
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp *utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, resp)

	//2. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	//3. 获取管理员签名
	bcosacc, _, err := utils.GetKeyFile(config.Config.Bcos.Keydir)
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		fmt.Println("Failed to get keyfile ! ", err)
		return
	}
	opts := bind.CallOpts{
		From : common.HexToAddress(bcosacc),
	}
	//4. 调用合约
	totalSupply, err := instance.TotalSupply(&opts)
	if err != nil {
		resp.Errno = utils.RECODE_GETCONTRAERR
		fmt.Println("Failed to get TotalSupply by instance !", err)
		return
	}
	resp.Data = totalSupply.Int64()
	fmt.Println("get total supply success!")
}

func GetBalance(c *gin.Context) {
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp *utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, resp)

	//2.解析数据
	userid, exist := c.Get("userid")
	if !exist {
		fmt.Println("Failed to get userid , userid err ")
		resp.Errno = utils.RECODE_NODATA
		return
	}
	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	//3. 获取管理员签名
	bcosacc, _, err := utils.GetKeyFile(config.Config.Bcos.Keydir)
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		fmt.Println("Failed to get keyfile ! ", err)
		return
	}
	opts := &bind.CallOpts{
		From : common.HexToAddress(bcosacc),
	}
	//4. 调用合约
	balance, err := instance.BalanceOf(opts, userid.(string))
	if err != nil {
		resp.Errno = utils.RECODE_GETCONTRAERR
		fmt.Println("Failed to get balance by instance !", err)
		return
	}
	resp.Data = balance.Int64()
	fmt.Println("get %s balance success!", userid)
}

func GetContraAddr(c *gin.Context) {
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp *utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, resp)

	//2. 解析数据
	itype := c.GetInt("itype")
	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	//4. 获取管理员签名
	bcosacc, _, err := utils.GetKeyFile(config.Config.Bcos.Keydir)
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		fmt.Println("Failed to get keyfile ! ", err)
		return
	}
	opts := &bind.CallOpts{
		From : common.HexToAddress(bcosacc),
	}

	//5. 调用合约
	addr, err := instance.GetAddr(opts, uint8(itype))
	if err != nil {
		resp.Errno = utils.RECODE_GETCONTRAERR
		fmt.Println("Failed to get addr by instance !", err)
		return
	}
	resp.Data = addr.Hex()
	fmt.Println("get total supply success!")
}






