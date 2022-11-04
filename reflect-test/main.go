package main

import (
	"log"
	"reflect"
)

/*
	反射类 相关学习
 */

func main() {
	var u1 *User

	u1 = GenStruct(u1).(*User)
	log.Printf("GenStruct方法后返回结果:%+v", u1)

	var name = "张三"
	var id = 123
	var isBoy = true
	var u2 User
	Call(&u2, name, id, isBoy)
	log.Printf("Call方法后user:%+v", u2)

	sliceResult2 := GenSlice(u2)
	log.Printf("GenSlice方法传入User结构体后返回结果:%+v", sliceResult2)

	var u3 *User
	sliceResult3 := GenSlice(u3)
	log.Printf("GenSlice方法传入User结构体空空指针后返回结果:%+v", sliceResult3)

	stringResult := GenSlice("string")
	log.Printf("GenSlice方法传入string类型后返回结果:%+v", stringResult)

	genArrayResult := GenArray(u2, 1)
	log.Printf("GenArray方法后返回结果:%+v", genArrayResult.(*[1]User))
}

type User struct {
	Id    int  `json:"id"`
	Name  string `json:"name"`
	IsBoy bool   `json:"is_boy"`
}

func (u *User) GetName() string {
	log.Println("执行GetName方法")
	return u.Name
}

func (u *User) SetName(name string) {
	log.Println("执行SetName方法")
	u.Name = name
}

func (u *User) GetId() int {
	log.Println("执行GetId方法")
	return u.Id
}

func (u *User) SetId(id int) {
	log.Println("执行SetId方法")
	u.Id = id
}

func (u *User) GetIsBoy() bool {
	log.Println("执行GetIsBoy方法")
	return u.IsBoy
}

func (u *User) SetIsBoy(isBoy bool) {
	log.Println("执行SetIsBoy方法")
	u.IsBoy = isBoy
}

// 反射调用结构体对象方法
func Call(typ interface{}, name string, id int, isBoy bool) {
	val := reflect.ValueOf(typ)

	// 调用SetName方法
	_ = val.MethodByName("SetName").Call([]reflect.Value{reflect.ValueOf(name)})

	// 调用GetName方法
	nameVals := val.MethodByName("GetName").Call([]reflect.Value{})
	for _, v := range nameVals {
		log.Println("执行SetName方法后 GetName方法返回结果值:", v.Interface().(string))
	}

	// 调用 SetId 方法
	_ = val.MethodByName("SetId").Call([]reflect.Value{reflect.ValueOf(id)})

	// 调用 GetId 方法
	idVals := val.MethodByName("GetId").Call([]reflect.Value{})
	for _, v := range idVals {
		log.Println("执行SetId方法后 GetId方法返回结果值:", v.Interface().(int))
	}

	// 调用SetName方法
	_ = val.MethodByName("SetIsBoy").Call([]reflect.Value{reflect.ValueOf(isBoy)})

	// 调用GetName方法
	isBoyVals := val.MethodByName("GetIsBoy").Call([]reflect.Value{})
	for _, v := range isBoyVals {
		log.Println("执行SetIsBoy方法后 GetIsBoy方法返回结果值:", v.Interface().(bool))
	}
}

// 反射生成结构体切片 根据传入类型生成对应切片, 并把req的值放到切片里
func GenSlice(req interface{}) interface{} {

	typ := reflect.TypeOf(req)

	// 通过类型 产生切片类型
	sliceType := reflect.SliceOf(typ)

	// 创建切片
	sliceVal := reflect.MakeSlice(sliceType, 0, 0)

	// 赋值
	vals := reflect.Append(sliceVal, reflect.ValueOf(req))

	return vals.Interface()
}

func GenArray(req interface{}, len int) interface{} {

	typ := reflect.TypeOf(req)

	typArr := reflect.ArrayOf(len, typ)

	return reflect.New(typArr).Interface()
}

// 反射生成结构体
func GenStruct(req interface{}) interface{} {
	typ := reflect.TypeOf(req)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return reflect.New(typ).Interface()
}