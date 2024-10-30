package service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *ServiceSetup) SaveCar(car Car) (string, error) {
	eventID := "eventAddPro1"
	//注册事件
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	//事件defer
	defer t.Client.UnregisterChaincodeEvent(reg)
	// 将edu对象序列化成为字节数组
	value, err := json.Marshal(car)
	key := car.CarId
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}
	fmt.Println([]byte(key))
	//rep 是执行调用链码需要的参数
	req := channel.Request{
		ChaincodeID: t.ChaincodeID,                                 //通道名字
		Fcn:         "addCar",                                      //函数名
		Args:        [][]byte{value, []byte(eventID), []byte(key)}, //函数参数
	}

	//t.Client.Execute(req) 是在后端执行该函数
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	//事件结果
	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) AddOrder(order Order) (string, error) {
	eventID := "eventAddPro1"
	//注册事件
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	//事件defer
	defer t.Client.UnregisterChaincodeEvent(reg)
	// 将edu对象序列化成为字节数组
	value, err := json.Marshal(order)
	key := order.CarId
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}
	fmt.Println([]byte(key))
	//rep 是执行调用链码需要的参数
	req := channel.Request{
		ChaincodeID: t.ChaincodeID,                                 //通道名字
		Fcn:         "addOrder",                                    //函数名
		Args:        [][]byte{value, []byte(eventID), []byte(key)}, //函数参数
	}

	//t.Client.Execute(req) 是在后端执行该函数
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	//事件结果
	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) FindProInfoByEntityID(entityID [][]byte) (channel.Response, error) {

	req := channel.Request{
		ChaincodeID: t.ChaincodeID,
		Fcn:         "queryProInfoByEntityID",
		Args:        entityID,
	}
	resp, _ := t.Client.Query(req)

	return resp, nil
}

func (t *ServiceSetup) QueryAllPro() (channel.Response, error) {

	eventID := "QueryAllPro"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{
		ChaincodeID: t.ChaincodeID,
		Fcn:         "QueryAllPro",
		Args:        [][]byte{[]byte(eventID)},
	}
	respone, _ := t.Client.Query(req)

	_ = eventResult(notifier, eventID)

	return respone, nil
}

//func (t *ServiceSetup) UpPrescription(prescription Prescription, key []byte) (string, error) {
//
//	eventID := "eventUpPrescription"
//	//注册事件
//	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
//
//	//事件defer
//	defer t.Client.UnregisterChaincodeEvent(reg)
//
//	src, _ := json.Marshal(prescription)
//
//	cipherText, err := sm4.Sm4Cbc(key, src, true)
//	fmt.Println("sm4加密后，密文如下\n", cipherText)
//	fmt.Println("key: ", prescription.Id)
//	//rep 是执行调用链码需要的参数
//	req := channel.Request{
//		ChaincodeID: t.ChaincodeID,                                                  //通道名字
//		Fcn:         "UpPrescription",                                               //函数名
//		Args:        [][]byte{cipherText, []byte(eventID), []byte(prescription.Id)}, //函数参数
//	}
//
//	//t.Client.Execute(req) 是在后端执行该函数
//	respone, err := t.Client.Execute(req)
//	if err != nil {
//		return "", err
//	}
//
//	//事件结果
//	err = eventResult(notifier, eventID)
//	if err != nil {
//		return "", err
//	}
//
//	return string(respone.TransactionID), nil
//}
//
//func (t *ServiceSetup) GetPrescription(Id string, key []byte) (Prescription, error) {
//
//	eventID := "eventGetPrescription"
//	//注册事件
//	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
//
//	//事件defer
//	defer t.Client.UnregisterChaincodeEvent(reg)
//
//	fmt.Println(Id, string(key))
//	//rep 是执行调用链码需要的参数
//	req := channel.Request{
//		ChaincodeID: t.ChaincodeID,        //通道名字
//		Fcn:         "GetPrescription",    //函数名
//		Args:        [][]byte{[]byte(Id)}, //函数参数
//	}
//
//	resp, _ := t.Client.Execute(req)
//	fmt.Println("=========================")
//
//	fmt.Println(resp.Payload)
//
//	plainText, _ := sm4.Sm4Cbc(key, bytes.NewBuffer(resp.Payload).Bytes(), false)
//
//	p := Prescription{}
//	_ = json.Unmarshal(plainText, &p)
//
//	//事件结果
//	_ = eventResult(notifier, eventID)
//	return p, nil
//}
