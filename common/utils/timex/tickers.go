package timex

import (
	"errors"
	"github.com/oxcaffee/caffee-go-middleware/common/global"
	"time"
)

type Ticker interface {
	Chan() <-chan time.Time
	Stop()
}

// SimulateTicker 测试专用接口
type SimulateTicker interface {
	Ticker
	Done()
	DoTick()
	Wait(d time.Duration) error
}

type simulateTicker struct {
	timeChan chan time.Time
	doneChan chan global.UniversalType
}

func (s *simulateTicker) Chan() <-chan time.Time {
	return s.timeChan
}

func (s *simulateTicker) Stop() {
	close(s.timeChan)
}

func (s *simulateTicker) Done() {
	s.doneChan <- global.UniversalType{}
}

func (s *simulateTicker) DoTick() {
	// @todo 暂时计数添加的是当前时间，并且计时间隔为1s
	s.timeChan <- TickerTime(time.Second)
}

func (s *simulateTicker) Wait(d time.Duration) error {
	select {
	case <-time.After(d):
		return errors.New("[common/utils/time]: timeout")
	case <-s.doneChan:
		// @todo 添加统一的日志输出
		return nil
	}
}

type realTicker struct {
	*time.Ticker
}

func (r *realTicker) Chan() <-chan time.Time {
	return r.C
}

func NewTicker(d time.Duration) Ticker {
	return &realTicker{
		Ticker: time.NewTicker(d),
	}
}

func NewSimulateTicker() SimulateTicker {
	return &simulateTicker{
		timeChan: make(chan time.Time, 1),
		doneChan: make(chan global.UniversalType, 1),
	}
}

func TickerTime(d time.Duration) time.Time {
	<-time.After(d)
	return time.Now()
}
