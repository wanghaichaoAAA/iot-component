package utils

import (
	"fmt"
	"sync"
	"time"
)

type idWorker struct {
	startTime             int64
	workerIdBits          uint
	datacenterIdBits      uint
	maxWorkerId           int64
	maxDatacenterId       int64
	sequenceBits          uint
	workerIdLeftShift     uint
	datacenterIdLeftShift uint
	timestampLeftShift    uint
	sequenceMask          int64
	workerId              int64
	datacenterId          int64
	sequence              int64
	lastTimestamp         int64
	signMask              int64
	idLock                *sync.Mutex
}

var OneIdWorker *idWorker

func init() {
	OneIdWorker = &idWorker{
		startTime: time.Now().UnixNano(),
	}
	OneIdWorker.initIdWorker(1000, 1)
}

func (this *idWorker) initIdWorker(workerId, datacenterId int64) {

	var baseValue int64 = -1
	//this.startTime = 1463834116272
	this.workerIdBits = 5
	this.datacenterIdBits = 5
	this.maxWorkerId = baseValue ^ (baseValue << this.workerIdBits)
	this.maxDatacenterId = baseValue ^ (baseValue << this.datacenterIdBits)
	this.sequenceBits = 12
	this.workerIdLeftShift = this.sequenceBits
	this.datacenterIdLeftShift = this.workerIdBits + this.workerIdLeftShift
	this.timestampLeftShift = this.datacenterIdBits + this.datacenterIdLeftShift
	this.sequenceMask = baseValue ^ (baseValue << this.sequenceBits)
	this.sequence = 0
	this.lastTimestamp = -1
	this.signMask = ^baseValue + 1

	this.idLock = &sync.Mutex{}

	if this.workerId < 0 || this.workerId > this.maxWorkerId {
		utilLogger.Error(fmt.Sprintf("workerId[%v] is less than 0 or greater than maxWorkerId[%v].", workerId, datacenterId))
	}
	if this.datacenterId < 0 || this.datacenterId > this.maxDatacenterId {
		utilLogger.Error(fmt.Sprintf("datacenterId[%d] is less than 0 or greater than maxDatacenterId[%d].", workerId, datacenterId))
	}
	this.workerId = workerId
	this.datacenterId = datacenterId
}

func (this *idWorker) NextId() int64 {
	this.idLock.Lock()
	timestamp := time.Now().UnixNano()
	if timestamp < this.lastTimestamp {
		utilLogger.Error(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", this.lastTimestamp-timestamp))
		return time.Now().UnixNano()
	}

	if timestamp == this.lastTimestamp {
		this.sequence = (this.sequence + 1) & this.sequenceMask
		if this.sequence == 0 {
			timestamp = this.tilNextMillis()
			this.sequence = 0
		}
	} else {
		this.sequence = 0
	}

	this.lastTimestamp = timestamp

	this.idLock.Unlock()

	id := ((timestamp - this.startTime) << this.timestampLeftShift) |
		(this.datacenterId << this.datacenterIdLeftShift) |
		(this.workerId << this.workerIdLeftShift) |
		this.sequence

	if id < 0 {
		id = -id
	}

	return id
}

func (this *idWorker) tilNextMillis() int64 {
	timestamp := time.Now().UnixNano()
	if timestamp <= this.lastTimestamp {
		timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	}
	return timestamp
}
