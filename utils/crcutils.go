package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/sigurn/crc16"
	log "github.com/sirupsen/logrus"
	"strings"
)

var utilLogger = log.WithFields(log.Fields{"method": "工具类"})

func CRC16Check(msg string) bool {
	index := strings.LastIndex(msg, "&&")
	if index < 0 {
		utilLogger.Error("没有找到标志位&&，", msg)
		return false
	}
	checkStr := msg[index+2:]
	if checkStr == "" {
		utilLogger.Error("没有校验码，", msg)
		return false
	}
	vaildSegment := msg[6 : index+2]
	checkCode := calculateCRC(vaildSegment)
	checkBytesArr, e := hex.DecodeString(checkStr)
	if e != nil {
		return false
	}

	if bytes.Equal(checkCode, checkBytesArr) {
		return true
	}
	return false
}

func GetCRCString(msg string) string {
	index := strings.LastIndex(msg, "&&")
	if index < 0 {
		utilLogger.Error("没有找到标志位&&，", msg)
		return ""
	}
	checkStr := msg[index+2 : len(msg)]
	return checkStr
}
func calculateCRC(msg string) []byte {
	data := []byte(msg)
	var high uint16
	var flag uint16
	// 16位寄存器，所有数位均为1
	wcrc := 0xFFFF
	for i := 0; i < len(data); i++ {
		// 16 位寄存器的高位字节
		high = uint16(wcrc >> 8)
		// 取被校验串的一个字节与 16 位寄存器的高位字节进行“异或”运算
		wcrc = int(high ^ uint16(data[i]))

		for j := 0; j < 8; j++ {
			flag = uint16(wcrc & 0x0001)
			// 把这个 16 寄存器向右移一位
			wcrc >>= 1
			// 若向右(标记位)移出的数位是 1,则生成多项式 1010 0000 0000 0001 和这个寄存器进行“异或”运算
			if flag == 0x0001 {
				wcrc ^= 0xA001
			}
		}
	}
	bytes := Int16ToBytes(wcrc)
	//s := hex.EncodeToString(bytes)
	return bytes
}

func GenerateCRCCode(msg string) string {
	crcByteArr := calculateCRC(msg)
	str := hex.EncodeToString(crcByteArr)
	return strings.ToUpper(str)
}

func Int16ToBytes(i int) []byte {
	var buf = make([]byte, 2)
	//var buf []byte
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

func GenerateCRCModbusCode(bytes []byte) []byte {
	table := crc16.MakeTable(crc16.CRC16_MODBUS)
	checksum := crc16.Checksum(bytes, table)
	var buf = make([]byte, 2)
	//var buf []byte
	binary.LittleEndian.PutUint16(buf, uint16(checksum))
	return buf
}
