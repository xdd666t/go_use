package base


//基本返回类型,该类型固定,切勿随意改动
type Response struct {
	//返回数据类,可谓任意类型
	Data interface{}  `json:"data"`
	//接口返回是否成功
	Success bool	`json:"success"`
	//如果返回状态位失败,msg中必须返回相关失败信息
	Msg string		`json:"msg"`
}