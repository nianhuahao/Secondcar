package CollectRoute

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"tce/service"
	"tce/utils"
	"time"
)

func CollectRoute(r *gin.Engine, db *gorm.DB) *gin.Engine {
	//路由
	//注册路由
	r.POST("/Register", func(ctx *gin.Context) {
		var requestUser = service.User{}
		var getUser = service.User{}
		err := ctx.ShouldBind(&requestUser)
		db.First(&getUser, "account = ?", requestUser.Account)

		fmt.Println(getUser)
		if !getUser.IsEmpty() {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "已经存在该用户",
			})
			return
		}

		fmt.Println(requestUser)
		if err != nil || requestUser.IsEmpty() {
			ctx.String(http.StatusNotFound, "绑定form失败", err)
			return
		}
		requestUser.UserId = utils.RandomStr(8)
		if err := db.Create(&requestUser).Error; err != nil {
			ctx.String(http.StatusNotFound, "插入数据库失败")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "数据库插入成功",
		})
	})
	//登录路由
	r.POST("/Login", func(ctx *gin.Context) {
		var requestUser = service.User{}
		var temp = service.User{}
		err := ctx.ShouldBind(&requestUser)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "绑定错误",
			})
			return
		}
		db.First(&temp, "account = ?", requestUser.Account) //
		if temp.Password == requestUser.Password && temp.UserType == requestUser.UserType {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "登录成功",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "密码或角色错误",
			})
		}

		fmt.Println(requestUser)
	})
	//修改信息
	r.POST("/Modify", func(ctx *gin.Context) {
		var requestUser = service.User{}
		var getUser = service.User{}
		err := ctx.ShouldBind(&requestUser)
		if err != nil {
			return
		}

		db.First(&getUser, "account = ?", requestUser.Account)
		fmt.Println(getUser)
		if getUser.IsEmpty() {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "不存在该用户,无法修改",
			})
			return
		}
		//删除原来记录
		db.Delete(&getUser)
		//插入新数据

		requestUser.Password = getUser.Password
		requestUser.UserId = getUser.UserId
		requestUser.Sex = getUser.Sex

		if err := db.Create(&requestUser).Error; err != nil {
			ctx.String(http.StatusNotFound, "插入数据库失败")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "修改成功",
		})
	})
	//获取我的信息
	r.POST("/GetMyInf", func(ctx *gin.Context) {
		var requestUser = service.User{}
		var getUser = service.User{}
		err := ctx.ShouldBind(&requestUser)
		if err != nil {
			return
		}
		db.First(&getUser, "account = ?", requestUser.Account)

		fmt.Println(getUser)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "get my inf success",
			"data": getUser,
		})
	})
	//获取所有用户信息
	r.POST("/FindAllUser", func(ctx *gin.Context) {
		var Result []service.User
		var GetCar []service.Car
		db.Find(&Result)
		db.Find(&GetCar)

		fmt.Println(GetCar)

		for i, v := range Result {
			account := 0
			for _, v1 := range GetCar {
				if v1.CarOwner == v.Identity {
					account++
				}
			}
			Result[i].Amount = account
		}

		fmt.Println(Result)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "查询成功",
			"data": Result,
		})
	})
	//生产商上传新车信息
	r.POST("/NewCar", func(ctx *gin.Context) {

		var NewCar = service.Car{}
		var GetCar = service.Car{}
		err := ctx.ShouldBind(&NewCar)
		db.First(&GetCar, "car_id = ?", NewCar.CarId)

		fmt.Println(GetCar)
		if !GetCar.IsEmpty() {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "该车已在库",
			})
			return
		}

		fmt.Println(NewCar)
		NewCar.SecondHandLevel = "0"
		NewCar.CarAges = "0"
		NewCar.LossAmount = "0"
		NewCar.CarOwner = "0"
		NewCar.CarState = "0"

		if err != nil || NewCar.IsEmpty() {
			ctx.String(http.StatusNotFound, "绑定form失败", err)
			return
		}
		//txid, err2 := setup.SaveCar(NewCar)
		//if err2 != nil {
		//	ctx.JSON(http.StatusBadRequest, gin.H{
		//		"msg": "上链失败",
		//		"err": err2,
		//	})
		//}
		if err := db.Create(&NewCar).Error; err != nil {
			ctx.String(http.StatusNotFound, "插入数据库失败")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "数据库插入成功",
			"txid": "txid",
		})

	})
	//获取所有在售车辆
	r.POST("/GetAllCarOnSell", func(ctx *gin.Context) {

		var Result []service.Car
		db.Where("state = ?", "1").Find(&Result)
		fmt.Println(Result)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "查询成功",
			"data": Result,
		})

	})
	//获取所有上链车辆
	r.POST("/GetAllCar", func(ctx *gin.Context) {

		var Result []service.Car
		db.Find(&Result)
		fmt.Println(Result)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "查询成功",
			"data": Result,
		})

	})
	//获取所有等待等待定价信息
	r.POST("/GetAllCarOnCheck", func(ctx *gin.Context) {

		var Result []service.Car
		db.Where("state = ?", "0").Find(&Result)
		fmt.Println(Result)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "查询成功",
			"data": Result,
		})

	})
	//通过车id查询车辆
	r.POST("/FindByCarId", func(ctx *gin.Context) {

		var NewCar = service.Car{}
		err := ctx.ShouldBind(&NewCar)
		if err != nil {
			ctx.String(http.StatusNotFound, "绑定form失败", err)
			return
		}
		db.First(&NewCar, "car_id = ?", NewCar.CarId)

		fmt.Println(NewCar)
		if NewCar.IsEmpty() {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "差无此车",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "查询成功",
			"data": NewCar,
		})

	})
	//通过车品牌搜索在售车辆
	r.POST("/FindByCarBrand", func(ctx *gin.Context) {

		var Result []service.Car
		var GetCar = service.Car{}
		err := ctx.ShouldBind(&GetCar)
		if err != nil {
			ctx.String(http.StatusNotFound, "绑定form失败", err)
			return
		}
		fmt.Println(GetCar)
		db.Where("car_brand = ?", GetCar.CarBrand).Find(&Result)

		fmt.Println(Result)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "查询成功",
			"data": Result,
		})

	})
	//通过价格搜索在售车辆
	r.POST("/FindByPriceRange", func(ctx *gin.Context) {

		var Result []service.Car
		var PriceRange = service.PriceRange{}
		err := ctx.ShouldBind(&PriceRange)
		if err != nil {
			ctx.String(http.StatusNotFound, "绑定form失败", err)
			return
		}
		fmt.Println(PriceRange)
		db.Where("guide_price > ? AND guide_price < ?", PriceRange.StartKey, PriceRange.EndKey).Find(&Result)

		fmt.Println(Result)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "查询成功",
			"data": Result,
		})

	})
	//上传图片信息
	r.POST("/UploadImg", func(c *gin.Context) {
		file, err := c.FormFile("upload")
		carid := c.PostForm("carId")

		if err != nil {
			c.String(http.StatusBadRequest, "请求失败")
			return
		}
		//获取后缀
		//filesuffix := path.Ext(file.Filename)

		if err := c.SaveUploadedFile(file, "img/"+carid+".jpg"); err != nil {
			c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "上传成功",
			"data": "img/" + carid + ".jpg",
		})

	})
	//获取图片
	r.GET("/GetImage", func(c *gin.Context) {
		imageName := c.Query("imageName")     //截取get请求参数，也就是图片的路径，可是使用绝对路径，也可使用相对路径
		file, _ := ioutil.ReadFile(imageName) //把要显示的图片读取到变量中
		_, err := c.Writer.WriteString(string(file))
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "获取成功",
			"data": "",
		})
	})
	//同意发行车辆
	r.POST("/AgreePublish", func(ctx *gin.Context) {
		var NewCar = service.Car{}
		var GetCar = service.Car{}
		err := ctx.ShouldBind(&NewCar)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "绑定失败",
			})
			return
		}
		db.First(&GetCar, "car_id = ?", NewCar.CarId)

		fmt.Println(GetCar)
		if GetCar.IsEmpty() {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "该车不存在",
			})
			return
		}

		fmt.Println(NewCar)

		//设置在售
		if err :=
			db.Model(service.Car{}).Where("car_id = ?", NewCar.CarId).Update("state", "1").Error; err != nil {
			ctx.String(http.StatusNotFound, "修改失败")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "同意在售成功",
		})

	})
	//消费者购买车辆
	r.POST("/BuyCar", func(ctx *gin.Context) {
		var NewOrder = service.Order{}
		var GetCar = service.Car{}
		var GetUser = service.User{}
		err := ctx.ShouldBind(&NewOrder)
		if err != nil {
			return
		}
		if NewOrder.IsEmpty() {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "绑定form失败",
			})
			return
		}

		CarId := NewOrder.CarId
		BuyerName := NewOrder.BuyerName

		db.First(&GetCar, "car_id = ?", CarId)
		fmt.Println(GetCar)
		if GetCar.CarOwner == "0" {
			NewOrder.SellerName = GetCar.Manufacture
		} else {
			NewOrder.SellerName = GetCar.CarOwner
		}
		if GetCar.IsEmpty() || GetCar.CarState == "2" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "车辆不存在或车辆已经售出",
			})
			return
		}

		db.First(&GetUser, "account = ?", BuyerName)
		fmt.Println(GetUser)
		if GetUser.IsEmpty() {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "用户不存在",
			})
			return
		}

		//修改车辆的状态和拥有者
		if err :=
			db.Model(service.Car{}).Where("car_id = ?", GetCar.CarId).Update("state", "2").Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "修改失败",
			})
			return
		}
		if err :=
			db.Model(service.Car{}).Where("car_id = ?", GetCar.CarId).Update("car_owner", GetUser.Identity).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "修改失败",
			})
			return
		}
		//插入order表
		//还差交易哈西设置

		timeUnix := time.Now().Unix() //已知的时间戳

		formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")

		//txid, err1 := setup.AddOrder(NewOrder)
		//if err1 != nil {
		//	ctx.JSON(http.StatusOK, gin.H{
		//		"code": http.StatusExpectationFailed,
		//		"msg":  "交易失败",
		//		"err":  err1,
		//	})
		//}
		NewOrder.Time = formatTimeStr
		NewOrder.Price = GetCar.GuidePrice
		NewOrder.TxHash = "txid"
		if err := db.Create(&NewOrder).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "修改失败",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"msg":     "订单生成成功",
			"data":    "tx_hash" + "txid",
			"CarName": GetCar.CarName,
		})

	})
	//获取个人资产
	r.POST("/GetMyCar", func(ctx *gin.Context) {
		var MyCarList []service.Car
		var User = service.User{}
		err := ctx.ShouldBind(&User)
		if err != nil {
			ctx.String(http.StatusNotFound, "绑定form失败", err)
			return
		}
		fmt.Println(User)
		db.Where("car_owner = ?", User.Identity).Find(&MyCarList)

		fmt.Println(MyCarList)
		ctx.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"msg":    "查询成功",
			"data":   MyCarList,
			"Amount": len(MyCarList),
		})
	})
	//上传某车维修记录
	r.POST("/UpdateRepair", func(ctx *gin.Context) {

		var Repair = service.Repair{}
		err := ctx.ShouldBind(&Repair)
		if err != nil {
			ctx.String(http.StatusNotFound, "绑定form失败", err)
			return
		}
		fmt.Println(Repair)
		if err := db.Create(&Repair).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "修改失败",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"msg":     "维修信息上链成功",
			"tx_hash": "hash",
		})
	})
	//获取一个车的维修记录
	r.POST("/GetCarRepairInfo", func(ctx *gin.Context) {

		//输入Vid
		var Repair = service.Repair{}
		var GetRepair []service.Repair
		err := ctx.ShouldBind(&Repair)
		if err != nil {
			ctx.String(http.StatusNotFound, "绑定form失败", err)
			return
		}
		fmt.Println(Repair)
		db.Where("car_id = ?", Repair.CarId).Find(&GetRepair)

		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "维修信息上链成功",
			"data": GetRepair,
		})
	})
	//出售车辆
	r.POST("/SellCar", func(ctx *gin.Context) {

		var NewCar = service.Car{}
		var GetCar = service.Car{}
		err := ctx.ShouldBind(&NewCar)
		if err != nil {
			return
		}
		fmt.Println(NewCar)
		db.First(&GetCar, "car_id = ?", NewCar.CarId)

		fmt.Println(GetCar)
		if GetCar.IsEmpty() {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusExpectationFailed,
				"msg":  "汽车不存在",
			})
			return
		}

		fmt.Println(NewCar)
		//需要改的字段内容
		GetCar.SecondHandLevel = "2"
		GetCar.CarOwner = "0"
		GetCar.CarState = "0"

		fmt.Println(NewCar)
		GetCar.CarAges = NewCar.CarAges
		GetCar.LossAmount = NewCar.LossAmount
		GetCar.GuidePrice = NewCar.GuidePrice
		GetCar.ExDate = NewCar.ExDate

		//删除后再上传
		db.Where("car_id = ?", NewCar.CarId).Delete(&service.Car{})

		if err := db.Create(&GetCar).Error; err != nil {
			ctx.String(http.StatusNotFound, "插入数据库失败")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "数据库插入成功",
		})

	})
	//获取所有交易信息
	r.POST("/GetAllTX", func(ctx *gin.Context) {
		var Result []service.Order
		db.Find(&Result)
		fmt.Println(Result)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "查询成功",
			"data": Result,
		})

	})
	return r
}
