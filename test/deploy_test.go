package try_test

import (
	"GoContractDeployment/internal"
	"GoContractDeployment/internal/deploy"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestDeploy(t *testing.T) {
	stream :=
		deploy.Structure{
			Name:           "TianYun",
			Symbol:         "TianYun",
			Minter:         common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
			TokenURIPrefix: "test",
		}
	//fmt.Print(stream)
	////
	////amount := big.NewInt(1e18)
	////usdtAmount := internal.GetBnbToUsdt(amount, t)
	////
	////log.Printf(usdtAmount)
	//
	//a := deploy.GoContractDeployment(stream, t)
	//
	//t.Log(a)
	//password, err := hashPassword("Tianyun")
	//if err != nil {
	//	return
	//}
	//fmt.Println(password)
	//a := internal.GetBnbToUsdt(big.NewInt(1e10))
	//fmt.Print(a)
	//comparePassword("Tianyun", "$2a$10$yuKr.CVq0o9K4a8QUDDmbumc6E6rn7L7jme8RP26MR92p3jsty2E6")
	// 创建用于接收状态的 channel

	addressHex, txDataHashHex, gasUsed := deploy.GoContractDeployment(stream)
	fmt.Println(addressHex)
	fmt.Println(txDataHashHex)
	fmt.Println("========================")
	//
	////gasUsed := deploy.GoTransactionNews(client, txDataHashHex)
	//fmt.Println(gasUsed)
	//fmt.Println("========================")
	//
	gasToUsdt := internal.GetBnbToUsdt(gasUsed)
	fmt.Println(gasToUsdt)

}
func processTask(status chan<- bool) {
	// 执行一些耗时的操作
	time.Sleep(5 * time.Second)

	// 处理完后向 channel 中发送状态信息
	status <- true
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
