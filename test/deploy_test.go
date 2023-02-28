package try_test

import (
	utils "GoContractDeployment/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
	"time"
)

func TestDeploy(t *testing.T) {
	//_ =
	//	deploy.Structure{
	//		Name:           "TianYun",
	//		Symbol:         "TianYun",
	//		Minter:         common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
	//		TokenURIPrefix: "test",
	//	}
	//connection, _ := navigation.CreateData()
	//phHandler := phMysql.NewJobHandler(connection)
	//one, _ := phHandler.Repo.GetOne()
	//a := internal.GetBnbToUsdt(big.NewInt(one.GasUsed))
	//
	//log.Panicln(a, "++++++++++++++++++++")

	//cfg, err := ini.Load("/Users/zhumaomao.eth/Desktop/myCode/golang/GoContractDeployment/config.ini")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//s := cfg.Section("database").Key("password").String()
	//log.Println(s)
	//
	//password, _ := hashPassword(s)
	//log.Println(password)
	//comparePassword(s, password)
	//key := []byte("0123456789abcdef")
	//
	//// 需要加密的数据
	//plaintext := []byte("Hello, World!")
	//
	//encrypt, err := aescrypt.Encrypt(plaintext, key)
	//if err != nil {
	//}
	//log.Println(encrypt)
	//
	//dncrypt, err := aescrypt.Dncrypt("vk1/8UY6+ZzU2058Uot6Iw==", key)
	//if err != nil {
	//}
	//log.Println(dncrypt)
	// 128位密钥
	// 密钥
	// 128位密钥
	// 128位密钥

	//生成一个16字节的随机数
	//randomBytes := make([]byte, 8)
	//_, err := rand.Read(randomBytes)
	//if err != nil {
	//	panic(err)
	//}
	//// 将随机数转换为16进制字符串
	//randomString := hex.EncodeToString(randomBytes)
	//fmt.Println(randomString)
	//
	////("0123456789abcdef")
	//key := []byte("ca5b20230224b5ac")
	//
	////// 需要加密的数据
	//plaintext := "[{\"chainId\":2,\"id\":\"1605133286670282753\",\"name\":\"创世纪念勋章\"},{\"chainId\":2,\"id\":\"1605133287098101762\",\"name\":\"石氏星经\"}]"
	//
	//encrypt, err := utils.Encrypt(plaintext, key)
	//if err != nil {
	//	return
	//}
	//
	//log.Println(encrypt)
	//
	//decrypt, err := utils.AesDecrypt(encrypt, key)
	//if err != nil {
	//	return
	//}
	//log.Println(decrypt)

	var configIni = []string{"username", "host", "port", "password", "database"}
	loading, _ := utils.ConfigurationLoading("database", configIni)
	log.Panicln(loading)

}

// processTask 暂停
func processTask(status chan<- bool) {
	// 执行一些耗时的操作
	time.Sleep(5 * time.Second)

	// 处理完后向 channel 中发送状态信息
	status <- true
}

// hashPassword 更改密码
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// comparePassword 比较密码
func comparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
