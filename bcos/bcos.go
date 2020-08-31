package bcos

import (
	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/client"
	"github.com/yekai1003/gobcos/common"

	"subsys/utils"
	"subsys/config"

	"fmt"
)

func GetControlInstance() (*Control, *client.Client, error) {
	//1. 连接到fisco节点
	cli, err := client.Dial(config.Config.Bcos.Connstr, config.Config.Bcos.GroupId)
	if err != nil {
		fmt.Println("Failed to dial bcos client! ", err)
		return nil, nil, err
	}

	//2. 获取链接合约入口
	instance, err := NewControl(common.HexToAddress(config.Config.Bcos.Contractaddr), cli)
	if err != nil {
		fmt.Println("Failed to new control contract! ", err)
		return nil, nil, err
	}
	return instance, cli, err
}

func GetAdminOpt(password string) (*bind.TransactOpts, error){
	//1. 从ks文件中获取签名
	_, file, err := utils.GetKeyFile(config.Config.Bcos.Keydir)
	if err != nil {
		fmt.Println("failed to get key file! ", err)
		return nil, err
	}
	opt, err := bind.NewTransactor(file, password)
	if err != nil {
		fmt.Println("Failed to new transactor to get opt ! ", err)
		return nil, err
	}
	return opt, nil
}
