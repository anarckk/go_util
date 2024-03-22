package go_net

import (
	"context"
	"fmt"
	"io"
	"time"

	"gitea.bee.anarckk.me/anarckk/go_util/go_cli"
	"gitea.bee.anarckk.me/anarckk/go_util/go_time"
	"gitea.bee.anarckk.me/anarckk/go_util/go_unit"
)

var _ io.Writer = &transportSizeOutputStream{}

type transportSizeOutputStream struct {
	writer io.Writer
	ch     chan<- int
}

func (t *transportSizeOutputStream) Write(p []byte) (n int, err error) {
	t.ch <- len(p)
	return t.writer.Write(p)
}

type manager struct {
	ctx   context.Context
	inCh  chan int
	outCh chan int
}

func (m *manager) run() {
	var inputTmp = 0
	var inputTotal uint64 = 0
	var outputTmp = 0
	var outputTotal uint64 = 0
	ticker := time.NewTicker(time.Second)
	printer := go_cli.GetCliPrinter()
FF:
	for {
		select {
		case _inputNum := <-m.inCh:
			inputTmp += _inputNum
			inputTotal += uint64(_inputNum)
		case _outputNum := <-m.outCh:
			outputTmp += _outputNum
			outputTotal += uint64(_outputNum)
		case <-ticker.C:
			printer(fmt.Sprintf("%s 入站速度:%12s/s(%12d), 出站速度:%12s/s(%9d), 入站总量:%12s(%9d), 出站总量:%12s(%12d)",
				go_time.CurrentDatetimeStr(), go_unit.HumanReadableByteCountBin(int64(inputTmp)), inputTmp,
				go_unit.HumanReadableByteCountBin(int64(outputTmp)), outputTmp,
				go_unit.HumanReadableByteCountBin(int64(inputTotal)), inputTotal,
				go_unit.HumanReadableByteCountBin(int64(outputTotal)), outputTotal,
			))
			inputTmp = 0
			outputTmp = 0
		case <-m.ctx.Done():
			close(m.inCh)
			close(m.outCh)
			ticker.Stop()
			break FF
		}
	}
}

// 主要作用是给反向代理服务器提供实时速度显示
// 给两边的输出流套一层代理对象
// inOutStream 入站输出流
// outOutStream 出站输出流
func NewReportReverseProxy(ctx context.Context, in io.Writer, out io.Writer) (io.Writer, io.Writer) {
	_inCh := make(chan int, 256)
	_outCh := make(chan int, 256)
	inWrapper := &transportSizeOutputStream{in, _inCh}
	outWrapper := &transportSizeOutputStream{out, _outCh}
	m := manager{ctx, _inCh, _outCh}
	go m.run()
	return inWrapper, outWrapper
}
