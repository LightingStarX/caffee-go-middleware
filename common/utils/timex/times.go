package timex

import (
	"fmt"
	"time"
)

// RelativeTimeSince 从t时间算起的相对时间间隔
func RelativeTimeSince(t time.Time) time.Duration {
	return time.Since(t)
}

// TransferDuration2Ms 将给定的时间间隔转化为ms表示的字符串格式，精确到小数点后1位
func TransferDuration2Ms(d time.Duration) string {
	return fmt.Sprintf("%.1fms", float32(d)/float32(time.Millisecond))
}

// ElapsedTimer 用来跟踪begin开始算起过去了多少时间
type ElapsedTimer struct {
	begin time.Time
}

func NewElapsedTimer() *ElapsedTimer {
	return &ElapsedTimer{
		begin: time.Now(),
	}
}

// ElapsedDuration 判断计时器过去了多少时间，返回 time.Duration 格式
func (et *ElapsedTimer) ElapsedDuration() time.Duration {
	return RelativeTimeSince(et.begin)
}

// ElapsedString 判断计时器过去了多少时间，返回 string 格式
func (et *ElapsedTimer) ElapsedString() string {
	return RelativeTimeSince(et.begin).String()
}

// ElapsedMsString 判断计时器过去了多少时间，返回ms格式的字符串
func (et *ElapsedTimer) ElapsedMsString() string {
	return TransferDuration2Ms(RelativeTimeSince(et.begin))
}

// CurrentMicroTime 当前时间的以微秒为单位的int64格式
func CurrentMicroTime() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// CurrentMilliTime 当前时间以毫秒为单位的int64格式
func CurrentMilliTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
