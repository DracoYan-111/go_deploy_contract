package try_test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestDeploy(t *testing.T) {
	//stream :=
	//	deploy.Structure{
	//		Name:           "TianYun",
	//		Symbol:         "TianYun",
	//		Minter:         common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
	//		TokenURIPrefix: "test",
	//	}
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
	status := make(chan bool)

	// 在一个新的 goroutine 中执行后续代码
	go processTask(status)

	// 立即返回状态信息给前端
	fmt.Println("正在处理，请稍候...")

	// 等待从 channel 中接收到状态信息
	if <-status {
		fmt.Println("处理成功！")
	} else {
		fmt.Println("处理失败！")
	}
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
