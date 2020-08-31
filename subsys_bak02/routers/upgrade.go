package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yekai1003/gobcos/common"

	"subsys/bcos"
	"subsys/utils"
)

//切换管理员  相当于换个ks文件
func UpdateOwner(c *gin.Context) {
	//1. 初始化响应数据
	respG := &utils.Gin{C:c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)

	//2. 参数获取解析
	ownerAddr := c.PostForm("owner")
	password := c.PostForm("password") ///老管理员密码

	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()

	//4. 获取管理员签名
	opts, err := bcos.GetAdminOpt(password)
	if err != nil {
		resp.Errno = utils.RECODE_OPTERR
		return
	}

	//5. 调用更换管理员合约
	_, err = instance.UpdateOwner(opts, common.HexToAddress(ownerAddr))
	if err != nil {
		resp.Errno = utils.RECODE_GETCONTRAERR
		fmt.Println("Failed to updateOwner contract ! ", err)
		return
	}

	//6. 更新管理员相关设置？？？？

	resp.Data = "Update Owner success!"
	fmt.Println("Update Owner success!")


}

//User合约升级
func UpgradeUser(c *gin.Context) {
	//1. 初始化响应数据
	respG := &utils.Gin{C:c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)

	//2. 参数解析
	usercontraAddr := c.PostForm("usercontra")
	password := c.PostForm("password")

	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	// 4. 获取管理员签名
	opts, err := bcos.GetAdminOpt(password)
	if err != nil {
		resp.Errno = utils.RECODE_OPTERR
		return
	}
	// 5. 调用升级User合约
	_, err = instance.UpgradeUser(opts, common.HexToAddress(usercontraAddr))
	if err != nil {
		resp.Errno = utils.RECODE_USERERR
		fmt.Println("Failed to upgrade user contract! ", err)
		return
	}
	resp.Data = "upgrade user contract" + usercontraAddr + "success ! "
	fmt.Println("upgrade user contract %s success ! ", usercontraAddr)

}

//erc200合约升级
func UpgradeErc200(c *gin.Context) {
	//1. 初始化响应数据
	respG := &utils.Gin{C:c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)

	//2. 参数解析
	erc200contraAddr := c.PostForm("erc200contra")
	password := c.PostForm("password")

	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	// 4. 获取管理员签名
	opts, err := bcos.GetAdminOpt(password)
	if err != nil {
		resp.Errno = utils.RECODE_OPTERR
		return
	}
	// 5. 调用升级erc200合约
	_, err = instance.UpgradeErc200(opts, common.HexToAddress(erc200contraAddr))
	if err != nil {
		resp.Errno = utils.RECODE_USERERR
		fmt.Println("Failed to upgrade erc200 contract! ", err)
		return
	}
	resp.Data = "upgrade erc200 contract" + erc200contraAddr + "success!"
	fmt.Println("upgrade erc200 contract %s success !", erc200contraAddr)
}
//erclog 合约升级
func UpgradeErclog(c *gin.Context) {
	//1. 初始化响应数据
	respG := &utils.Gin{C:c}
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer respG.ResponseData(http.StatusOK, &resp)

	//2. 参数解析
	erclogcontraAddr := c.PostForm("erclogcontra")
	password := c.PostForm("password")

	//3. 连接网络并获取control合约接口
	instance, cli, err := bcos.GetControlInstance()
	if err != nil {
		resp.Errno = utils.RECODE_BCOSERR
		return
	}
	defer cli.Close()
	// 5. 通过keystore文件获取签名 以及 对应的bcos addr
	//通过Bcos.Keydir获取私钥文件
	// 4. 获取管理员签名
	opts, err := bcos.GetAdminOpt(password)
	if err != nil {
		resp.Errno = utils.RECODE_OPTERR
		return
	}
	// 5. 调用升级erclog合约
	_, err = instance.UpgradeErclog(opts, common.HexToAddress(erclogcontraAddr))
	if err != nil {
		resp.Errno = utils.RECODE_USERERR
		fmt.Println("Failed to upgrade erclog contract! ", err)
		return
	}
	resp.Data = "upgrade erclog contract" + erclogcontraAddr + "success !"
	fmt.Println("upgrade erclog contract %s success !", erclogcontraAddr)
}

