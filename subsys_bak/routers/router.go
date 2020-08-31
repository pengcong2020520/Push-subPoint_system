package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/client"
	"github.com/yekai1003/gobcos/common"
	"log"
	"net/http"
	"subsys/bcos"
	"subsys/config"
	"subsys/dbs"
	"subsys/utils"
)

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
	account.UserName = c.PostForm("username")
	account.Organization = c.PostForm("organization")
	account.Password = c.PostForm("password")

	//3. 生成userid,并通过id验证账号是否存在  userid是用户名称和组织名称组合起来的字符串的MD5值
	userid := utils.EncodeMD5(account.UserName + account.Organization)
	account.UserId = userid
	_, exist := utils.CheckUserid(account.UserId)
	if exist {
		resp.Errno = utils.RECODE_DATAEXIST
		return
	}
	//4. 通过bcos创建账户
	bcosacc, err := bcos.NewAccount(account.Password)
	if err != nil {
		fmt.Println("failed to new bcos account! ", err)
		return 
	}
	account.BcosAcc = bcosacc
	//5. 将用户信息加入数据库
	sql := fmt.Sprintf("insert into accounts(username, organization, userid, password, bcosacc) value('%s', '%s', '%s', '%s', '%s')",
		utils.EncodeMD5(account.UserName),
		utils.EncodeMD5(account.Organization),
		account.UserId,
		utils.EncodeMD5(account.Password),
		bcosacc,
	)

	fmt.Println(sql)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("Failed to insert account to database", err)
		resp.Errno = utils.RECODE_PARAMERR
		return
	}
	//6. 连接到fisco节点
	cli, err := client.Dial(config.Config.Bcos.Connstr, config.Config.Bcos.GroupId)
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		log.Panic("Failed to dial bcos client! ", err)
	}
	defer cli.Close()
	//设置签名
	opts := bind.CallOpts{
		From : common.HexToAddress(account.BcosAcc),
	}
	//7. 链接合约入口
	instance, err := bcos.NewControl(common.HexToAddress(config.Config.Bcos.Contractaddr), cli)
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		log.Panic("Failed to new control contract! ", err)
	}
	_, err = instance.Register(&opts, userid, account.Password)
	if err != nil {
		resp.Errno = utils.RECODE_USERERR
		log.Panic("Failed to register")
	}
	fmt.Println("userid %s register success!", userid)
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
	account.UserName = c.PostForm("username")
	account.Password = c.PostForm("password")

	//3. 验证用户名与密码  如果成功则获取userid
	sql := fmt.Sprintf("select userid,bcosacc from accounts where username='%s' and password='%s'",
		utils.EncodeMD5(account.UserName),
		utils.EncodeMD5(account.Password),
	)
	m, n, err := dbs.DBQuery(sql)
	if err != nil || n <= 0 {
		resp.Errno = utils.RECODE_NODATA
		log.Fatal("Failed to db query! ", err)
	}
	fmt.Println(sql)
	account.UserId = m[0]["userid"]
	account.BcosAcc = m[0]["bcosacc"]

	//4. 连接到fisco节点
	cli, err := client.Dial(config.Config.Bcos.Connstr, config.Config.Bcos.GroupId)
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		log.Panic("Failed to dial bcos client! ", err)
	}
	defer cli.Close()
	//设置签名
	opts := bind.CallOpts{
		From : common.HexToAddress(account.BcosAcc),
	}
	//5. 链接合约入口
	instance, err := bcos.NewControl(common.HexToAddress(config.Config.Bcos.Contractaddr), cli)
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		log.Panic("Failed to new control contract! ", err)
	}
	ok, err := instance.Login(&opts, account.UserId, account.Password)
	if err != nil {
		resp.Errno = utils.RECODE_LOGINERR
		log.Panic("Failed to login")
	}

	if !ok {
		resp.Errno = utils.RECODE_LOGINERR
		return
	}
	fmt.Println("userid %s login success!", account.UserId)
}