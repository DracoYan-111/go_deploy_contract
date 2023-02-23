package try_test

import (
	phMysql "GoContractDeployment/handler/http"
	"GoContractDeployment/internal"
	"GoContractDeployment/internal/deploy"
	"GoContractDeployment/navigation"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/big"
	"testing"
	"time"
)

func TestDeploy(t *testing.T) {
	_ =
		deploy.Structure{
			Name:           "TianYun",
			Symbol:         "TianYun",
			Minter:         common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
			TokenURIPrefix: "test",
		}
	connection, _ := navigation.CreateData()
	phHandler := phMysql.NewJobHandler(connection)
	one, _ := phHandler.Repo.GetOne()
	a := internal.GetBnbToUsdt(big.NewInt(one.GasUsed))

	log.Panicln(a, "++++++++++++++++++++")

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
