package bcos
import(
	"fmt"
	"github.com/yekai1003/gobcos/accounts/keystore"
)

const keyDir = "./keystore"

func NewAccount(password string) (string, error){
	ks := keystore.NewKeyStore(keyDir, keystore.LightScryptN, keystore.LightScryptP)
	acc, err := ks.NewAccount(password)
	if err != nil {
		fmt.Println("Failed to new ACCOUNT! ", err)
		return "", err
	}
	return acc.Address.Hex(), nil
}


