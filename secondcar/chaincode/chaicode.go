package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
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

type Chaincode struct {
}
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Car
}

func (t *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println(" ==== Init ====")

	return shim.Success(nil)
}
func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 获取用户意图
	fun, args := stub.GetFunctionAndParameters()
	if fun == "addCar" {
		return t.addCar(stub, args) // 添加信息
	} else if fun == "addOrder" {
		return t.addOrder(stub, args) // 添加信息
	} else if fun == "queryCarByCarID" {
		return t.queryByID(stub, args)
	}
	return shim.Error("指定的函数名称错误")

}

const DOC_TYPE = "medicineObj"
const DOC_TYPE1 = "prescriptionObj"

func (t *Chaincode) addCar(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {
		return shim.Error("给定的参数个数不符合要求")
	}
	err := stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(args[2], []byte(args[0]))
	if err != nil {
		return peer.Response{}
	}
	return shim.Success([]byte("信息添加成功"))
}

func (t *Chaincode) addOrder(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {
		return shim.Error("给定的参数个数不符合要求")
	}
	err := stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(args[2], []byte(args[0]))
	if err != nil {
		return peer.Response{}
	}
	return shim.Success([]byte("信息添加成功"))
}

func (t *Chaincode) queryByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	b, _ := stub.GetState(args[0])
	return shim.Success(b)
}

//func (t *Chaincode) QueryAllPro(stub shim.ChaincodeStubInterface, args []string) peer.Response {
//	startKey := ""
//	endKey := ""
//
//	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
//
//	if err != nil {
//		return shim.Error("错了")
//	}
//	defer resultsIterator.Close()
//
//	var results []QueryResult
//
//	for resultsIterator.HasNext() {
//		queryResponse, err := resultsIterator.Next()
//		//CoughDB  是一个键值对数据库，支持富查询   值是JSON形式
//		if err != nil {
//			return shim.Error("错了")
//		}
//
//		product := new(Medicine)
//		_ = json.Unmarshal(queryResponse.Value, product)
//
//		queryResult := QueryResult{Key: queryResponse.Key, Record: product}
//		results = append(results, queryResult)
//	}
//	jsontex, _ := json.Marshal(results)
//	return shim.Success(jsontex)
//}

//
//func (t *Chaincode) UpPrescription(stub shim.ChaincodeStubInterface, args []string) peer.Response {
//
//	if len(args) != 3 {
//		return shim.Error("给定的参数个数不符合要求")
//	}
//	err := stub.PutState(args[2], []byte(args[0]))
//	if err != nil {
//		return peer.Response{}
//	}
//	err = stub.SetEvent(args[1], []byte{})
//	if err != nil {
//		return shim.Error(err.Error())
//	}
//	return shim.Success([]byte("信息添加成功"))
//}
//
//func (t *Chaincode) GetPrescription(stub shim.ChaincodeStubInterface, args []string) peer.Response {
//	b, _ := stub.GetState(args[0])
//	return shim.Success(b)
//}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("启动MedicineChaincode时发生错误: %s", err)
	}
}
