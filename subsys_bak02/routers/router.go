package routers

import (
	"fmt"
	"github.com/yekai1003/gobcos/client"
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

const AdminAddr = "0xed11fa12a479480fb7b4a658e42f8ff98f7d1631"

var AdminKey = `{"address":"ed11fa12a479480fb7b4a658e42f8ff98f7d1631","crypto":{"cipher":"aes-128-ctr","ciphertext":"ef584e4cfbb5d1493f03c9451a0621566d7cfc10ba74110ba1bc3c633c6fe622","cipherparams":{"iv":"6702c91f3356df967f62937b869e5b4e"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"2dcce1e262e027d90b9b2357c3a1ffac788e04dde00c676c80bf48a8f58767c0"},"mac":"3d046d146ca179f1e26aff0cfc5ee462a4cbe9b38af1999b274c3d53753c6b74"},"id":"ce5e7062-08df-4d60-976b-a191f7c4374b","version":3}`

func Ping(c *gin.Context) {
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)
	//1. 连接到fisco节点
	cli, err := client.Dial(config.Config.Bcos.Connstr, config.Config.Bcos.GroupId)
	if err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		fmt.Println("Failed to dial bcos client! ", err)
		return
	}
	defer cli.Close()
	//2. 获取链接合约入口
	instance, err := bcos.NewPing(common.HexToAddress("0xaff13ae66f0f6cb4d0fe7ec562c7b7905a49b281"), cli)
	if err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		return
	}
	opts := &bind.CallOpts{
		From : common.HexToAddress(AdminAddr),
	}
	s, _ := instance.GetMsg(opts)
	resp.Data = s
}


func Register(c *gin.Context) {
		/* Test使用
			c.JSON(http.StatusOK, gin.H{"status" : "OK", "function" : "register"})
		*/
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)
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
			//测试时使用 创建bcos用户
			//bcosacc, err := bcos.NewAccount(ADMIN_PASS)
			//if err != nil {
			//	return
			//}
			//fmt.Println(bcosacc)
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
	fmt.Printf("userid %s register success!\n", account.UserId)
}

func Login(c *gin.Context) {
	/* Test使用
	c.JSON(http.StatusOK, gin.H{"status" : "OK", "function" : "register"})
	*/
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)
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
	//bcosacc, _, err := utils.GetKeyFile(config.Config.Bcos.Keydir)
	//if err != nil {
	//	resp.Errno = utils.RECODE_BCOSERR
	//	fmt.Println("Failed to get keyfile ! ", err)
	//	return
	//}
	opts := bind.CallOpts{
		From : common.HexToAddress(AdminAddr),
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
//  模块问题 输入的key不存在或者不对->没有任何反馈 已解决
//  userid不存在时  也会反馈成功

func SetPasswd(c *gin.Context) {
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)

	//2. 获取数据 并解析
	account := &dbs.Account{}
	err := c.ShouldBind(account)
	if err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		fmt.Println("Failed to bind account ! ", err)
		return
	}
	var ok1, ok2, ok3 bool
	account.UserId, ok1 = c.GetPostForm("userid")
	account.Password, ok2 = c.GetPostForm("oldpassword")
	newPass, ok3 := c.GetPostForm("newpassword")
	if !ok1 || !ok2 || !ok3 {
		resp.Errno = utils.RECODE_PARAMERR
		fmt.Println("Get Post Form ERR ", err)
		return
	}
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
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)

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

//已处理
func GetBalance(c *gin.Context) {
	//1. 响应数据初始化
	respG := utils.Gin{C: c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)

	//2.解析数据
	userid, exist := c.GetPostForm("userid")
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
	balance, err := instance.BalanceOf(opts, userid)
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
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)

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






