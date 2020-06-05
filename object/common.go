package object

type DataType string

//心跳循环定时器ID前缀
const Prefix_timer_heartbeat = "heartbeat_"

//采集循环定时器ID前缀
const Prefix_timer_collect = "collect_"

//存储循环定时器ID前缀
const Prefix_timer_store = "store_crontab"

//上传数据循环定时器ID前缀
const Prefix_timer_upload = "upload_"

//上传分钟数据循环定时器ID前缀
const Prefix_timer_minute = "minute_crontab"

//上传小时数据循环定时器ID前缀
const Prefix_timer_hour = "hour_crontab"

//上传天数据循环定时器ID前缀
const Prefix_timer_day = "day_crontab"

const Prefix_timer_active_upload = "active_load"

const Data_rtd_type = "1"

const Data_calc_minute_type = "2"

const Data_calc_hour_type = "2"

const Data_calc_day_type = "2"

const Data_type_rtd DataType = "Rtd"

const Data_type_min DataType = "Min"

const Data_type_max DataType = "Max"

const Data_type_avg DataType = "Avg"

const Data_statistics_rtd DataType = "rtd"

const Data_statistics_minute DataType = "minute"

const Data_statistics_hour DataType = "hour"

const Data_statistics_day DataType = "day"

const Username = "admin"

const SuperUsername = "superadmin"

const SuperUserPw = "lczn321"

func (d DataType) String() string {
	switch d {
	case Data_type_rtd:
		return "Rtd"
	case Data_type_min:
		return "Min"
	case Data_type_max:
		return "Max"
	case Data_type_avg:
		return "Avg"
	case Data_statistics_minute:
		return "minute"
	case Data_statistics_hour:
		return "hour"
	case Data_statistics_day:
		return "day"
	}
	return ""
}

const OverTime string = "OverTime="
const Recount string = "Recount="
const PolId string = "PolId="
const SystemTime string = "SystemTime="
const RtdInterval string = "RtdInterval="
const MinInterval string = "MinInterval="
const NewPW string = "NewPW="
const DataTime string = "DataTime="
const RestartTime string = "RestartTime="
const BeginTime string = "BeginTime="
const EndTime string = "EndTime="

const SmokeGasPollutantSTs string = "22,27,31,"

const ModbusRtu string = "modbus-rtu"
const ModbusAscii string = "modbus-ascii"
const ModbusTCP string = "modbus-tcp"
const Hj212ActiveUpload string = "hj212-active"
const ModbusDynamicControl string = "modbus-dynamicControl"

const MnHj21217 string = "hj212-17"
const MnHj21205 string = "hj212-05"

const Yes string = "1"
const No string = "0"

const RtuWholeLog = "rtuWholeLog"
const RtuRunLog = "rtuRunLog"
const RtuRealTimeData = "rtuRealTimeData"
const RtuTcpLog = "rtuTcpLog"
const RtuStatisticsLog = "rtuStatisticsLog"
const RtuAlarmLog = "rtuAlarmLog"
const RtuMnList = "rtuMnList"
const RtuNewMnList = "rtuNewMnList"
const RtuResendData = "rtuResendData"
const RtuDeviceList = "rtuDeviceList"
const RtuNewDeviceList = "rtuNewDeviceList"
const RtuBasic = "rtuBasicInfo"
const RtuNewBasicSetting = "rtuNewBasicSetting"
const RtuParameter = "rtuParameter"
const RtuConfiguredFactors = "rtuConfiguredFactors"
const RtuSerialDebug = "rtuSerialDebug"

const RFID = "RFD"
