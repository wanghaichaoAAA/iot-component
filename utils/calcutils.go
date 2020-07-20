package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//十六进制字符串转换成浮点数
// 41360000H --> 11.375
func HexToFloatString(hexStr string) string {
	n, _ := strconv.ParseUint(hexStr, 16, 32)
	f := math.Float32frombits(uint32(n))
	return fmt.Sprintf("%f", f)
}

//将flag值转换成8位二进制字符串
func ConvertToBinaryArr(numStr string) []rune {
	if numStr == "" {
		numStr = "4"
	}
	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		//utilsLogger.Error("不是数字")
		return []rune("")
	}
	if numInt < 0 || numInt > 255 {
		return []rune("")
	}
	var s string
	for ; numInt > 0; numInt /= 2 {
		lsb := numInt % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return []rune(fmt.Sprintf("%08s", s))
}

// 格式化小数位数
func FormatFloat(floatNum float64, capacity int) string {
	formatStr := "%." + strconv.Itoa(capacity) + "f"
	return fmt.Sprintf(formatStr, floatNum)
}

// 字节数组转换成整数
func GetIntValue(data []byte) string {
	length := len(data)
	switch length {
	case 2:
		return strconv.Itoa(int(binary.BigEndian.Uint16(data)))
	case 4:
		return strconv.Itoa(int(binary.BigEndian.Uint32(data)))
	default:
		utilLogger.Error("将字节数组转换成整数失败")
		return "0"
	}
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		utilLogger.Error("StringToInt error ,", err)
		return 0
	}
	return i
}

func StringToFloat64(str string) float64 {
	floatA, err := strconv.ParseFloat(str, 64)
	if err != nil {
		utilLogger.Error("StringToFloat64 error ,", err)
		return 0
	}
	return floatA
}

func Get2BitByteArray(str string) []byte {
	i := StringToInt(str)
	s1 := make([]byte, 1)
	if i < 0 || i > 255 {
		return s1
	}
	buf := bytes.NewBuffer(s1)
	binary.Write(buf, binary.BigEndian, uint8(i))
	return buf.Bytes()
}

func IsNumber(str string) bool {
	if str == "" {
		return false
	}
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}

//10转16
func GetHexString(arr []byte) string {
	var tmp string
	for _, one := range arr {
		s := hex.EncodeToString([]byte{one})
		sprintf := fmt.Sprintf("%4s", s)
		tmp += strings.ReplaceAll(sprintf, " ", "0")

	}

	return tmp
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}
