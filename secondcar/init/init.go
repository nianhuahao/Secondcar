package Init

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"tce/middlewares"
	"tce/sdkInit"
	"tce/service"
)

const (
	username   = "root"           //账号
	password   = "123456"         //密码
	host       = "118.25.195.230" //数据库地址，可以是Ip或者域名
	port       = 3306             //数据库端口
	Dbname     = "SecondHandCar"  //数据库名
	cc_name    = "TCE_cc"
	cc_version = "1.0.0"
)

func InitGinGrom() (*gin.Engine, *gorm.DB) {

	r := gin.Default()
	r.Use(middlewares.Cors())
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil
	}
	return r, db
}

func InitFabric() *service.ServiceSetup {
	orgs := []*sdkInit.OrgInfo{ //SDK 中封装好的组织结构体
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: os.Getenv("GOPATH") + "/src/tce/fixtures/channel-artifacts/Org1MSPanchors.tx",
		}, //  os.Getenv  作用是返回参数内的环境变量
	}
	// sdk 环境信息
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    os.Getenv("GOPATH") + "/src/tce/fixtures/channel-artifacts/channel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    os.Getenv("GOPATH") + "/src/tce/chaincode/",
		ChaincodeVersion: cc_version,
	}
	// sdk setup   调用sdkinit 内的 Setup 函数，将config.yaml  和上面 建立好的sdk环境信息传入 ，返回一个完整的SDK
	sdk, err := sdkInit.Setup("config.yaml", &info)
	fmt.Println("-------------------")
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}
	// create channel and join
	if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	fmt.Println("-------------------")

	// create chaincode lifecycle
	if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		fmt.Println(">> create chaincode lifecycle error: %v", err)
		os.Exit(-1)
	}
	// invoke chaincode set status
	fmt.Println(">> 通过链码外部服务设置链码状态......")
	//初始化服务
	serviceSetup, err := service.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk)
	if err != nil {
		fmt.Println()
		os.Exit(-1)
	}

	return serviceSetup
}
