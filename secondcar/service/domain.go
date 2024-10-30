package service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"reflect"
	"tce/sdkInit"
	"time"
)

type User struct {
	int64    `json:"id"`
	Account  string `gorm:"column:account" json:"account" `
	Password string `gorm:"column:password" json:"password"`
	Username string `gorm:"column:name" json:"name"`
	Identity string `gorm:"column:Identity" json:"Identity"`
	Phone    string `gorm:"column:phone" json:"phone"`
	Sex      string `gorm:"column:sex" json:"sex"`
	Address  string `gorm:"column:address" json:"address"`
	Mail     string `gorm:"column:mail" json:"mail"`
	UserId   string `gorm:"column:user_id" json:"user_id"`
	UserType string `gorm:"column:user_type" json:"usertype"`
	Amount   int    `json:"amount"`
}

type Car struct {
	int64           `json:"id"`
	CarBrand        string `gorm:"column:car_brand" json:"CarBrand" `
	CarName         string `gorm:"column:car_name" json:"CarName"`
	CarClass        string `gorm:"column:car_class" json:"CarClass"`
	Domestic        string `gorm:"column:domestic" json:"Domestic"`
	Engin           string `gorm:"column:engin" json:"Engin"`    //引擎号
	ExDate          string `gorm:"column:ex_date" json:"ExDate"` //出厂日期
	GuidePrice      string `gorm:"column:guide_price" json:"GuidePrice"`
	SecondHandLevel string `gorm:"column:secondhandlevel" json:"SecondHandLevel"`
	CarAges         string `gorm:"column:car_ages" json:"CarAges"`
	LossAmount      string `gorm:"column:loss_amount" json:"LossAmount"`
	CarId           string `gorm:"column:car_id" json:"CarId"`
	Manufacture     string `gorm:"column:manufacture" json:"manufacture"`
	CarOwner        string `gorm:"column:car_owner" json:"carowner" `
	CarState        string `gorm:"column:state" json:"car_state" `
}

type Order struct {
	int64      `json:"id"`
	BuyerName  string `gorm:"column:buyername" json:"buyername"`
	SellerName string `gorm:"column:sellername" json:"sellername"`
	Time       string `gorm:"column:time" json:"time"` //购买日期
	Price      string `gorm:"column:price" json:"price"`
	CarId      string `gorm:"column:car_id" json:"CarId"`
	TxHash     string `gorm:"column:tx_hash" json:"TxHash"`
}

type Repair struct {
	int64    `json:"id"`
	CarId    string `gorm:"column:car_id" json:"CarId"`
	CarOwner string `gorm:"column:car_owner" json:"carowner" `
	M_amount string `gorm:"column:m_amount" json:"m_amount" `
	R_amount string `gorm:"column:r_amount" json:"r_amount" `
	Describe string `gorm:"column:describe" json:"describe" `
	Time     string `gorm:"column:time" json:"time"`
	Detail   string `gorm:"column:detail" json:"detail"`
}

type PriceRange struct {
	StartKey string `json:"StartKey" `
	EndKey   string `json:"EndKey"`
}

func (r Repair) IsEmpty() bool {
	return reflect.DeepEqual(r, Repair{})
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(u, User{})
}

func (c Car) IsEmpty() bool {
	return reflect.DeepEqual(c, Car{})
}

func (c Order) IsEmpty() bool {
	return reflect.DeepEqual(c, Order{})
}

func (r Repair) TableName() string {
	return "t_repair"
}
func (o Order) TableName() string {
	return "t_order"
}

func (c Car) TableName() string {
	return "t_car"
}

func (u User) TableName() string {
	return "t_user"
}

//嵌套获取历史信息

type TraceId struct {
	TxId string
}

//ServiceSetup sdk
type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}

func InitService(chaincodeID, channelID string, org *sdkInit.OrgInfo, sdk *fabsdk.FabricSDK) (*ServiceSetup, error) {
	handler := &ServiceSetup{
		ChaincodeID: chaincodeID,
	}
	//prepare channel client context using client context
	clientChannelContext := sdk.ChannelContext(channelID, fabsdk.WithUser(org.OrgUser), fabsdk.WithOrg(org.OrgName))
	// Channel client is used to query and execute transactions (Org1 is default org)
	client, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new channel client: %s", err)
	}
	handler.Client = client

	return handler, nil
}
