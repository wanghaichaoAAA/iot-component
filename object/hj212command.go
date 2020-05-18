package object

type Hj212Code struct {
	Code string
	Name string
}

//----------------------------------执行结果---------------------------------------
var ExeRtn1 = Hj212Code{"1", "执行成功"}
var ExeRtn2 = Hj212Code{"2", "执行失败，但不知道原因"}
var ExeRtn3 = Hj212Code{"3", "命令请求条件错误"}
var ExeRtn99 = Hj212Code{"99", "执行失败数据库更新失败"}

//----------------------------------flag编码---------------------------------------
var Flag101 = Hj212Code{"5", "hj2122017版协议，不需要分包，需要应答"}
var Flag100 = Hj212Code{"4", "hj2122017版协议，不需要分包，不需要应答"}

//----------------------------------系统编码---------------------------------------
var ST21 = Hj212Code{"21", "地表水质量监测"}
var ST22 = Hj212Code{"22", "空气质量监测"}
var ST23 = Hj212Code{"23", "声环境质量监测"}
var ST31 = Hj212Code{"31", "大气环境污染源"}
var ST32 = Hj212Code{"32", "地表水质污染源"}
var ST99 = Hj212Code{"99", "餐饮油烟污染源"}
var ST91 = Hj212Code{"91", "系统交互"}

//----------------------------------命令编码---------------------------------------
//-------------初始化命令-------------
var CN1000 = Hj212Code{"1000", "设置超时时间及重发次数"}
var CN1011 = Hj212Code{"1011", "提取现场机时间"}
var CN1012 = Hj212Code{"1012", "设置现场机时间"}
var CN1061 = Hj212Code{"1061", "提取实时数据间隔"}
var CN1062 = Hj212Code{"1062", "设置实时数据间隔"}
var CN1063 = Hj212Code{"1063", "提取分钟数据间隔"}
var CN1064 = Hj212Code{"1064", "设置分钟数据间隔"}
var CN1072 = Hj212Code{"1072", "设置现场机密码"}

//-------------数据命令-------------
//实时数据
var CN2011 = Hj212Code{"2011", "污染物实时数据"}
var CN2012 = Hj212Code{"2012", "停止察看污染物实时数据"}

//设备状态
var CN2021 = Hj212Code{"2021", "设备运行状态数据"}
var CN2022 = Hj212Code{"2022", "停止察看设备运行状态"}

//日数据
var CN2031 = Hj212Code{"2031", "染物日历史数据"}
var CN2041 = Hj212Code{"2041", "设备运行时间日历史数据"}

//分钟数据
var CN2051 = Hj212Code{"2051", "污染物分钟数据"}

//小时数据
var CN2061 = Hj212Code{"2061", "污染物小时数据"}

//其它数据
var CN2081 = Hj212Code{"2081", "数采仪开机时间"}

//动态管控功能
var CN2082 = Hj212Code{"2082", "打开动态管控功能"}
var CN2083 = Hj212Code{"2083", "关闭动态管控功能"}

//-------------控制命令-------------
var CN3011 = Hj212Code{"3011", "零点校准量程校准"}
var CN3012 = Hj212Code{"3012", "即时采样"}
var CN3013 = Hj212Code{"3013", "启动清洗/反吹"}
var CN3014 = Hj212Code{"3014", "比对采样"}
var CN3015 = Hj212Code{"3015", "超标留样"}
var CN3016 = Hj212Code{"3016", "零设置采样时间周期"}
var CN3017 = Hj212Code{"3017", "提取采样时间周期"}
var CN3018 = Hj212Code{"3018", "提取出样时间"}
var CN3019 = Hj212Code{"3019", "提取设备唯一标识"}
var CN3020 = Hj212Code{"3020", "提取现场机信息"}
var CN3021 = Hj212Code{"3021", "设置现场机参数"}

var CN3030 = Hj212Code{"3030", "标定"}
var CN3031 = Hj212Code{"3031", "标定"}
var CN3041 = Hj212Code{"3041", "重启数采仪系统"}

//-------------交互命令-------------
var CN9011 = Hj212Code{"9011", "请求应答，用于现场机回应接收的上位机请求命令是否有效"}
var CN9012 = Hj212Code{"9012", "执行结果，用于现场机回应接收的上位机请求命令执行结果"}
var CN9013 = Hj212Code{"9013", "通知应答"}
var CN9014 = Hj212Code{"9014", "数据应答,数据应答命令"}

//-------------自定义命令-------------
var CN2072 = Hj212Code{"2072", "自定义命令-报警数据"}
var CN9020 = Hj212Code{"9020", "心跳消息"}

var CN9030 = Hj212Code{"90030", "反控指令"}
var CN9031 = Hj212Code{"90031", "开机信息，接受私有反控指令"}
var CN9032 = Hj212Code{"90032", "反控应答信息，是否接收到私有反控指令"}
var CN9033 = Hj212Code{"90033", "反控执行结果信息，私有反控指令执行结果"}
