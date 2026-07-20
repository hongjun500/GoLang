package main

import "log"

func main() {
	log.Println("反射第一定律：接口值 → 反射对象（TypeOf/ValueOf、Type vs Kind）----------")
	reflectBasic()
	log.Println("反射第二定律：反射对象 → 接口值（Interface 装箱回原类型）----------")
	reflectBackToInterface()
	log.Println("反射第三定律：可设置性 settability =（可寻址 + 已导出）----------")
	reflectSettable()
	log.Println("遍历结构体字段 + 读取 struct tag（json/gorm/validator 的底层）----------")
	reflectStruct()
	log.Println("动态调用方法（RPC / DI 框架的底层）----------")
	reflectCall()
	log.Println("反射访问未导出字段（能读，不能 Interface，不能 Set）----------")
	reflectUnexported()
}
