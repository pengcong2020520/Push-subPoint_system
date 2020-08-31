package eths

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"os"
	"strings"
	"subsys/config"
	"subsys/utils"
)

const keyDir = "./keystore"

//新建eth账户  这个方法创建的账户需要手动从geth文件中的keystore文件中拷贝出来
func NewEthAcc(pass string) (string, error) {
	//1. 连接geth
	client, err := rpc.Dial(config.Config.Eth.Connstr)
	if err != nil {
		log.Panic("Failed to rpc Dial client!", err)
	}
	defer client.Close()
	//2. 操作geth来创建账户
	var account string
	if err = client.Call(&account, "personal_newAccount", pass); err != nil {
		log.Println("Failed to new account by geth !", err)
		return "", err
	}
	fmt.Println(account)
	return account, err
}

//可以指定文件目录创建账户的keystore文件, 并改名为账户地址
func NewKeystore(pass string) (string, error) {
	ks := keystore.NewKeyStore(keyDir, keystore.LightScryptN, keystore.LightScryptP)
	acc, err := ks.NewAccount(pass)
	if err != nil {
		fmt.Println("Failed to new account keystore file ! ", err)
		return "", err
	}
	src, _ := utils.GetFileName(acc.Address.Hex(), keyDir)
	os.Rename("./keystore"+"/"+src, "./keystore"+"/"+strings.ToLower(acc.Address.Hex()))
	return acc.Address.Hex(), nil
}






