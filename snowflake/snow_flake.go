package snowflake

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/xlizy/common-go/utils"
	"github.com/xlizy/common-go/zlog"
	"sync"
	"time"
)

type SnowFlake struct {
	epoch     int64 // 起始时间戳
	timestamp int64 // 当前时间戳，毫秒
	centerId  int64 // 数据中心机房ID
	workerId  int64 // 机器ID
	sequence  int64 // 毫秒内序列号

	timestampBits  int64 // 时间戳占用位数
	centerIdBits   int64 // 数据中心id所占位数
	workerIdBits   int64 // 机器id所占位数
	sequenceBits   int64 // 序列所占的位数
	lastTimestamp  int64 // 上一次生成ID的时间戳
	sequenceMask   int64 // 生成序列的掩码最大值
	workerIdShift  int64 // 机器id左移偏移量
	centerIdShift  int64 // 数据中心机房id左移偏移量
	timestampShift int64 // 时间戳左移偏移量
	maxTimeStamp   int64 // 最大支持的时间

	lock sync.Mutex // 锁
}

var s = &SnowFlake{}

// Init 初始化
func Init(centerId, workerId int64) error {
	s.epoch = int64(1636373513000) //设置起始时间戳：2021-11-08 20:11:53
	s.centerId = centerId
	s.workerId = workerId
	s.centerIdBits = 4   // 支持的最大机房ID占位数，最大是15
	s.workerIdBits = 6   // 支持的最大机器ID占位数，最大是63
	s.timestampBits = 41 // 时间戳占用位数
	s.maxTimeStamp = -1 ^ (-1 << s.timestampBits)

	maxWorkerId := -1 ^ (-1 << s.workerIdBits)
	maxCenterId := -1 ^ (-1 << s.centerIdBits)

	//此处写死start
	for _, b := range []byte(utils.GetLocalIp()) {
		centerId += int64(b)
		workerId += int64(b)
	}
	centerId = centerId%s.centerIdBits + 1
	workerId = workerId%s.workerIdBits + 1
	//此处写死end

	// 参数校验
	if int(centerId) > maxCenterId || centerId < 0 {
		return errors.New(fmt.Sprintf("Center ID can't be greater than %d or less than 0", maxCenterId))
	}
	if int(workerId) > maxWorkerId || workerId < 0 {
		return errors.New(fmt.Sprintf("Worker ID can't be greater than %d or less than 0", maxWorkerId))
	}

	s.sequenceBits = 12 // 序列在ID中占的位数,最大为4095
	s.sequence = -1

	s.lastTimestamp = -1                                                // 上次生成 ID 的时间戳
	s.sequenceMask = -1 ^ (-1 << s.sequenceBits)                        // 计算毫秒内，最大的序列号
	s.workerIdShift = s.sequenceBits                                    // 机器ID向左移12位
	s.centerIdShift = s.sequenceBits + s.workerIdBits                   // 机房ID向左移18位
	s.timestampShift = s.sequenceBits + s.workerIdBits + s.centerIdBits // 时间截向左移22位

	return nil
}

func NextId() int64 {
	s.lock.Lock() //设置锁，保证线程安全
	defer s.lock.Unlock()

	now := time.Now().UnixNano() / 1000000 // 获取当前时间戳，转毫秒
	if now < s.lastTimestamp {             // 如果当前时间小于上一次 ID 生成的时间戳，说明发生时钟回拨
		zlog.Error(fmt.Sprintf("Clock moved backwards. Refusing to generate id for %d milliseconds", s.lastTimestamp-now))
		return 0
	}

	t := now - s.epoch
	if t > s.maxTimeStamp {
		zlog.Error(fmt.Sprintf("epoch must be between 0 and %d", s.maxTimeStamp-1))
		return 0
	}

	// 同一时间生成的，则序号+1
	if s.lastTimestamp == now {
		s.sequence = (s.sequence + 1) & s.sequenceMask
		// 毫秒内序列溢出：超过最大值; 阻塞到下一个毫秒，获得新的时间戳
		if s.sequence == 0 {
			for now <= s.lastTimestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0 // 时间戳改变，序列重置
	}
	// 保存本次的时间戳
	s.lastTimestamp = now

	// 根据偏移量，向左位移达到
	return (t << s.timestampShift) | (s.centerId << s.centerIdShift) | (s.workerId << s.workerIdShift) | s.sequence
}
