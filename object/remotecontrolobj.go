package object

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"io"
)

type RemoteControlObj struct {
	ResponseType string      `json:"response_type"`
	ResponseData interface{} `json:"response_data"`
}

func (obj *RemoteControlObj) GetJSONString() string {
	bytes, err := json.Marshal(*obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (obj *RemoteControlObj) GetBase64String() (string, error) {
	bytes, err := json.Marshal(*obj)
	if err != nil {
		return "", err
	}
	tmpByte := CompressBase64Str(bytes)
	base64Str := base64.URLEncoding.EncodeToString(tmpByte)

	return base64Str, nil
}

func DeCompress(in []byte) string {
	buffer := bytes.NewBuffer(in)
	var out bytes.Buffer
	r, _ := zlib.NewReader(buffer)
	io.Copy(&out, r)
	return out.String()
}

func CompressBase64Str(in []byte) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(in)
	w.Close()
	return b.Bytes()
}

type ResendDataObj struct {
	RtuMN     string `json:"rtuMN"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	DataType  int    `json:"dataType"`
	UUID      string `json:"uuid"`
}

type RemoteControlResendDataObj struct {
	ResponseType string        `json:"response_type"`
	ResponseData ResendDataObj `json:"response_data"`
}

type DeviceObj struct {
	Id                    string            `json:"id"`
	Name                  string            `json:"name"`
	AccessType            string            `json:"access_type"`
	St                    string            `json:"st"`
	IntervalCollect       int               `json:"interval_collect"`
	FailTime              int               `json:"fail_time"`
	CommunicationProtocol string            `json:"communication_protocol"`
	DeviceAddress         int               `json:"device_address"`
	Com                   string            `json:"com"`
	Baudrate              string            `json:"baudrate"`
	DataBit               string            `json:"data_bit"`
	CheckBit              string            `json:"check_bit"`
	StopBit               string            `json:"stop_bit"`
	IpAddr                string            `json:"ip_addr"`
	Params                string            `json:"params"` //设备读指令参数
	Factors               []DeviceFactorObj `json:"factors"`
}

type DeviceFactorObj struct {
	Id             int     `json:"id" gorm:"id primary_key"`
	DeviceId       string  `json:"device_id"`                          //设备id
	St             string  `json:"st"`                                 //系统编码(哪个设备)
	FactorCode     string  `json:"factor_code"`                        //因子编码
	FactorAlias    string  `json:"factor_alias"`                       //因子别名（自定义因子名称）
	FactorId       int     `json:"factor_id"`                          //因子id：当因子编码重复时使用因子id来区分
	Ratio          string  `json:"ratio"`                              //系数
	DataBitNo      string  `json:"data_bit_no"`                        //数据位序号
	IsModbus       int     `json:"is_modbus"`                          //是否modbus协议
	RegisterAddr   int     `json:"register_addr"`                      //寄存器地址
	RegisterLength int     `json:"register_length"`                    //寄存器个数
	IsInteger      int     `json:"is_integer"`                         //是否整数
	Decimals       int     `json:"decimals"`                           //小数位
	IsAnalog       int     `json:"is_analog"`                          //模拟量 1开启模拟量，2不开器模拟量
	AnalogUpper    float32 `json:"analog_upper" gorm:"default:'0.0'" ` //量程上限（选择模拟量后必填）
	AnalogLower    float32 `json:"analog_lower" gorm:"default:'0.0'" ` //量程下线（选择模拟量后必填）
	AlarmUpper     float32 `json:"alarm_upper"`                        //警告上限
	AlarmLower     float32 `json:"alarm_lower"`                        //警告下线
	IsSendMsg      int     `json:"is_send_msg"`                        //发送报警信息
	IsControlRelay int     `json:"is_control_relay"`                   //是否控制继电器，1控制继电器，2不控制继电器
	RelayId        string  `json:"relay_id"`                           //继电器id（选择继电器后必填）
	RelayAction    string  `json:"relay_action"`                       //继电器动作（选择继电器后必填）
	ActionTime     string  `json:"action_time"`                        //继电器动作时间（选择继电器后必填）
	FactorName     string  `json:"-" gorm:"-"`                         //因子名称
	FactorUnit     string  `json:"-" gorm:"-"`                         //计量单位（浓度）
}

type RemControlSerialPortDebugObj struct {
	RtuMN      string `json:"rtuMN"`
	Serial     string `json:"serial"`      //串口
	Baudrate   string `json:"baudrate"`    //波特率
	DataBit    string `json:"data_bit"`    //数据位
	CheckBit   string `json:"check_bit"`   //校验位
	StopBit    string `json:"stop_bit"`    //停止位
	VerifyType string `json:"verify_type"` //验证方式
	Message    string `json:"message"`
}

type TFactorShowConfigReq struct {
	Id          int64  `json:"id"`
	FactorCode  string `json:"factor_code"`
	FactorName  string `json:"factor_name"`
	FactorId    int    `json:"factor_id"`
	St          string `json:"st"`
	FactorValue string `json:"factor_value"`
	Contents    string `json:"contents"`
}
