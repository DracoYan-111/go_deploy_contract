package try_test

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
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

	comparePassword("Tianyun", "$2a$10$yuKr.CVq0o9K4a8QUDDmbumc6E6rn7L7jme8RP26MR92p3jsty2E6")
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
