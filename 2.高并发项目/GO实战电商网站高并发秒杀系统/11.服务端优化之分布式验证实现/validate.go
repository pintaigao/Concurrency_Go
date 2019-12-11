package main

import (
	"fmt"
	"imooc-product/common"
	"imooc-product/encrypt"
	"net/http"
	"sync"
	"strconv"
	"io/ioutil"
)

//设置集群地址，最好内外IP
var hostArray= []string{"127.0.0.1","127.0.0.1"}

var localHost = "127.0.0.1"

var port = "8081"

var hashConsistent *common.Consistent

//用来存放控制信息，
type AccessControl struct {
	//用来存放用户想要存放的信息
	sourcesArray map[int]interface{}
	sync.RWMutex
}

//创建全局变量
var accessControl = &AccessControl{sourcesArray:make(map[int]interface{})}

//获取制定的数据
func (m *AccessControl) GetNewRecord(uid int) interface{} {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	data:=m.sourcesArray[uid]
	return data
}

//设置记录
func (m *AccessControl) SetNewRecord(uid int) {
	m.RWMutex.Lock()
	m.sourcesArray[uid]="hello imooc"
	m.RWMutex.Unlock()
}

func (m *AccessControl) GetDistributedRight(req *http.Request) bool {
	//获取用户UID
	uid ,err := req.Cookie("uid")
	if err !=nil {
		return false
	}

	//采用一致性hash算法，根据用户ID，判断获取具体机器
	hostRequest,err:=hashConsistent.Get(uid.Value)
	if err !=nil {
		return false
	}

	//判断是否为本机
	if hostRequest == localHost {
		//执行本机数据读取和校验
		return m.GetDataFromMap(uid.Value)
	} else {
		//不是本机充当代理访问数据返回结果
		return GetDataFromOtherMap(hostRequest,req)
	}

}

//获取本机map，并且处理业务逻辑，返回的结果类型为bool类型
func (m *AccessControl) GetDataFromMap(uid string) (isOk bool) {
	uidInt,err := strconv.Atoi(uid)
	if err !=nil {
		return false
	}
	data:=m.GetNewRecord(uidInt)

	//执行逻辑判断
	if data !=nil {
		return true
	}
	return
}

//获取其它节点处理结果
func GetDataFromOtherMap(host string,request *http.Request) bool  {
	//获取Uid
	uidPre,err := request.Cookie("uid")
	if err !=nil {
		return false
	}
	//获取sign
	uidSign,err:=request.Cookie("sign")
	if err !=nil {
		return  false
	}

	//模拟接口访问，
	client :=&http.Client{}
	req,err:= http.NewRequest("GET","http://"+host+":"+port+"/check",nil)
	if err !=nil {
		return false
	}

	//手动指定，排查多余cookies
	cookieUid :=&http.Cookie{Name:"uid",Value:uidPre.Value,Path:"/"}
	cookieSign :=&http.Cookie{Name:"sign",Value:uidSign.Value,Path:"/"}
	//添加cookie到模拟的请求中
	req.AddCookie(cookieUid)
	req.AddCookie(cookieSign)

	//获取返回结果
	response,err :=client.Do(req)
	if err !=nil {
		return false
	}
	body,err:=ioutil.ReadAll(response.Body)
	if err !=nil {
		return false
	}

	//判断状态
	if response.StatusCode == 200 {
		if string(body) == "true" {
			return true
		} else {
			return false
		}
	}
	return false
}

//执行正常业务逻辑
func Check(w http.ResponseWriter, r *http.Request) {
	//执行正常业务逻辑
	fmt.Println("执行check！")
}

//统一验证拦截器，每个接口都需要提前验证
func Auth(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("执行验证！")
	//添加基于cookie的权限验证
	err := CheckUserInfo(r)
	if err != nil {
		return err
	}
	return nil
}

//身份校验函数
func CheckUserInfo(r *http.Request) error {
	//获取Uid，cookie
	uidCookie, err := r.Cookie("uid")
	if err != nil {
		//return errors.New("用户UID Cookie 获取失败！")
	}
	//获取用户加密串
	signCookie, err := r.Cookie("sign")
	if err != nil {
		//return errors.New("用户加密串 Cookie 获取失败！")
	}

	//对信息进行解密
	signByte, err := encrypt.DePwdCode(signCookie.Value)
	if err != nil {
		//return errors.New("加密串已被篡改！")
	}

	//fmt.Println("结果比对")
	//fmt.Println("用户ID：" + uidCookie.Value)
	//fmt.Println("解密后用户ID：" + string(signByte))
	if checkInfo(uidCookie.Value, string(signByte)) {
		//return nil
	}
	//return errors.New("身份校验失败！")
	return nil
}

//自定义逻辑判断
func checkInfo(checkStr string, signStr string) bool {
	if checkStr == signStr {
		return true
	}
	return false
}

func main() {
	//负载均衡器设置
	//采用一致性哈希算法
	hashConsistent =common.NewConsistent()
	//采用一致性hash算法，添加节点
	for _,v :=range hostArray {
		hashConsistent.Add(v)
	}


	//1、过滤器
	filter := common.NewFilter()
	//注册拦截器
	filter.RegisterFilterUri("/check", Auth)
	//2、启动服务
	http.HandleFunc("/check", filter.Handle(Check))
	//启动服务
	http.ListenAndServe(":8083", nil)
}
