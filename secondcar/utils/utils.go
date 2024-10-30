package utils

import (
	"math/rand"
	"time"
)

//RandomStr 随机生成字符串
func RandomStr(length int) string {
	str := "0123456789abcdefg"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//func FindApartmentInfoById(bodyBytes [][]byte) map[string]interface{} {
//	result, _ := service.ServiceSetup.FindProInfoByEntityID(bodyBytes)
//	data := make(map[string]interface{}, 0)
//
//	err := json.Unmarshal(bytes.NewBuffer(result.Payload).Bytes(), &data)
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	} else {
//		return data
//	}
//
//}
