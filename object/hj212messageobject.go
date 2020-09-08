package object

import (
	"errors"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	log "github.com/sirupsen/logrus"
	"github.com/wanghaichaoAAA/iot-component/utils"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

var hjLog = log.WithFields(log.Fields{"method": "Hj 212 message"})

var MultiPkgMessagesMap cmap.ConcurrentMap

func init() {
	MultiPkgMessagesMap = cmap.New()
}

const (
	Suffix     = ";"
	PrefixQN   = "QN="
	PrefixST   = "ST="
	PrefixCN   = "CN="
	PrefixPW   = "PW="
	PrefixMN   = "MN="
	PrefixFlag = "Flag="
	PrefixCP   = "CP=&&"
	SuffixCP   = "&&"
	StartMark  = "##"
	PrefixPNUM = "PNUM="
	PrefixPNO  = "PNO="
)

type Hj212Message struct {
	QN          string      `json:"qn"`
	ST          string      `json:"st"`
	CN          string      `json:"cn"`
	PW          string      `json:"pw"`
	MN          string      `json:"mn"`
	Flag        string      `json:"flag"`
	Version     int         `json:"-"` //消息版本号 `json:"-"`
	Response    int         `json:"-"` //是否应答 `json:"-"`
	Unpack      int         `json:"-"` //是否有数据包序号
	PNUM        int         `json:"-"` //总包数
	PNO         int         `json:"-"` //包号
	CP          interface{} `json:"cp"`
	CRC         string      `json:"-"`
	OriginalMsg string      `json:"original_msg"`
}

type MultiPackageMessageArr []Hj212Message

func (m MultiPackageMessageArr) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m MultiPackageMessageArr) Len() int      { return len(m) }

type SortByPNO struct{ MultiPackageMessageArr }

func (b SortByPNO) Less(i, j int) bool {
	return b.MultiPackageMessageArr[i].PNO < b.MultiPackageMessageArr[j].PNO
}

func NewMessage(message string) (*Hj212Message, error) {
	startIndex := strings.Index(message, "##")
	if startIndex < 0 {
		return nil, errors.New("No start with ##")
	}
	message = message[startIndex:]
	qnStr := utils.SubstringBetween(message, PrefixQN, Suffix)
	if qnStr == "" {
		return nil, errors.New("missing qn field")
	}
	stStr := utils.SubstringBetween(message, PrefixST, Suffix)
	if stStr == "" {
		return nil, errors.New("missing st field")
	}
	cnStr := utils.SubstringBetween(message, PrefixCN, Suffix)
	if cnStr == "" {
		return nil, errors.New("missing cn field")
	}
	mnStr := utils.SubstringBetween(message, PrefixMN, Suffix)
	if mnStr == "" {
		return nil, errors.New("missing mn field")
	}
	flagStr := utils.SubstringBetween(message, PrefixFlag, Suffix)
	if flagStr == "" {
		return nil, errors.New("missing flag field")
	}
	cpStr := utils.SubstringBetween(message, PrefixCP, "&&")
	if cpStr == "" {
		return nil, errors.New("missing cp field")
	}
	pwStr := utils.SubstringBetween(message, PrefixPW, Suffix)
	crcStr := utils.GetCRCString(message)

	if !strings.HasPrefix(cnStr, "9003") {
		crcCheckSuccess := utils.CRC16Check(message)
		if !crcCheckSuccess {
			return nil, errors.New("CRC verification failed")
		}
	}

	binaryStrArr := utils.ConvertToBinaryArr(flagStr)
	if len(binaryStrArr) != 8 {
		return nil, errors.New("Flag format error")
	}
	// is need respone
	var respInt int
	// is need packageing multi package messages
	var unpackInt int

	if string(binaryStrArr[7]) == "1" {
		respInt = 1
	}
	if string(binaryStrArr[6]) == "1" {
		unpackInt = 1
	}
	msgObj := Hj212Message{
		QN:          qnStr,
		ST:          stStr,
		CN:          cnStr,
		PW:          pwStr,
		MN:          mnStr,
		Flag:        flagStr,
		Response:    respInt,
		Unpack:      unpackInt,
		CP:          cpStr,
		CRC:         crcStr,
		OriginalMsg: message,
	}

	if unpackInt == 0 {
		return &msgObj, nil
	}
	pnumStr := utils.SubstringBetween(message, PrefixPNUM, Suffix)
	pnoStr := utils.SubstringBetween(message, PrefixPNO, Suffix)
	pnumInt, _ := strconv.Atoi(pnumStr)
	msgObj.PNUM = pnumInt
	pnoInt, _ := strconv.Atoi(pnoStr)
	msgObj.PNO = pnoInt
	msgs, err := packageingMultiPackageMessages(&msgObj)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func packageingMultiPackageMessages(msg *Hj212Message) (*Hj212Message, error) {
	mapName := msg.MN + "-" + msg.QN
	go time.AfterFunc(time.Minute*5, func() {
		MultiPkgMessagesMap.Remove(mapName)
	})
	getValue, isExist := MultiPkgMessagesMap.Get(mapName)
	if !isExist {
		msgMap := cmap.New()
		msgMap.Set(strconv.Itoa(msg.PNO), *msg)
		MultiPkgMessagesMap.Set(mapName, msgMap)
		return nil, errors.New("receive remaining multi package messages")
	}
	multiPkgMsgMap := getValue.(cmap.ConcurrentMap)
	multiPkgMsgMap.Set(strconv.Itoa(msg.PNO), *msg)
	count := multiPkgMsgMap.Count()
	if count == msg.PNUM {
		var msgArr MultiPackageMessageArr
		for item := range multiPkgMsgMap.IterBuffered() {
			obj := item.Val.(Hj212Message)
			msgArr = append(msgArr, obj)
		}
		sort.Sort(SortByPNO{msgArr})
		var cps string
		for _, msg := range msgArr {
			cps += msg.CP.(string)
		}
		msg.CP = cps
		MultiPkgMessagesMap.Remove(mapName)
		return msg, nil
	}
	return nil, errors.New("multi package messages")
}

// max length for cp
var CP_MAX_LENGTH = 950

func MakeMessage(qnStr, stStr, cnStr, pwStr, mnStr, cpStr, protocolVersion, isResp string) []string {
	if len([]rune(cpStr)) > CP_MAX_LENGTH {
		return subpackage(stStr, cnStr, pwStr, mnStr, cpStr, protocolVersion, isResp)
	}
	commandStr := PrefixQN + qnStr + Suffix +
		PrefixST + stStr + Suffix +
		PrefixCN + cnStr + Suffix +
		PrefixPW + pwStr + Suffix +
		PrefixMN + mnStr + Suffix +
		PrefixFlag + getFlag(protocolVersion, isResp, "0") + Suffix +
		PrefixCP + cpStr + SuffixCP
	commandLength := len([]rune(commandStr))
	commandLengthStr := fmt.Sprintf("%04d", commandLength)
	crcCode := utils.GenerateCRCCode(commandStr)
	commandStr = StartMark + commandLengthStr + commandStr + crcCode
	return []string{commandStr}
}

func subpackage(stStr, cnStr, pwStr, mnStr, cpStr, protocolVersion, isResp string) []string {
	cpLength := len([]rune(cpStr))
	totalPkgF := float64(cpLength) / float64(CP_MAX_LENGTH)
	totalPkg := int(math.Ceil(totalPkgF))
	var messages []string
	for i := 0; i < totalPkg; i++ {
		qn := utils.GenerateQNStr()
		start := i * CP_MAX_LENGTH
		end := start + CP_MAX_LENGTH
		if end > cpLength {
			end = cpLength
		}
		commandStr := PrefixQN + qn + Suffix +
			PrefixST + stStr + Suffix +
			PrefixCN + cnStr + Suffix +
			PrefixPW + pwStr + Suffix +
			PrefixMN + mnStr + Suffix +
			PrefixFlag + getFlag(protocolVersion, isResp, "1") + Suffix +
			PrefixPNUM + strconv.Itoa(totalPkg) + Suffix +
			PrefixPNO + strconv.Itoa(i+1) + Suffix +
			PrefixCP + cpStr[start:end] + SuffixCP
		commandLength := len([]rune(commandStr))
		commandLengthStr := fmt.Sprintf("%04d", commandLength)
		crcCode := utils.GenerateCRCCode(commandStr)
		commandStr = StartMark + commandLengthStr + commandStr + crcCode
		messages = append(messages, commandStr)
	}
	return messages
}

func getFlag(protocolVersion, isResp, isPackage string) string {
	versionArr := utils.ConvertToBinaryArr(protocolVersion)
	flagStr := string(versionArr) + isPackage + isResp
	i, _ := strconv.ParseInt(flagStr, 2, 32)
	return strconv.Itoa(int(i))
}

func CleanMultiPkgMessagesMap() {
	count := MultiPkgMessagesMap.Count()
	if count == 0 {
		return
	}
	nowTime := time.Now()
	for item := range MultiPkgMessagesMap.IterBuffered() {
		mapName := item.Key
		strArr := strings.Split(mapName, "-")
		if len(strArr) != 2 {
			MultiPkgMessagesMap.Remove(mapName)
			continue
		}
		qn := strArr[1]
		qnTime, err := utils.GetTimeByQN(qn)
		if err != nil {
			MultiPkgMessagesMap.Remove(mapName)
			continue
		}
		if nowTime.Sub(qnTime).Hours() > 1 {
			MultiPkgMessagesMap.Remove(mapName)
			continue
		}
	}
}

func MakeHeartbeatMsg(mnStr string) string {
	qnStr := utils.GenerateQNStr()
	commandStr := PrefixQN + qnStr + Suffix +
		PrefixCN + CN9020.Code + Suffix +
		PrefixFlag + Flag100.Code + Suffix +
		PrefixMN + mnStr + Suffix +
		PrefixCP + SuffixCP
	commandLength := len([]rune(commandStr))
	commandLengthStr := fmt.Sprintf("%04d", commandLength)
	crcCode := utils.GenerateCRCCode(commandStr)
	commandStr = StartMark + commandLengthStr + commandStr + crcCode
	return commandStr
}
