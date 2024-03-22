package go_file

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"gitea.bee.anarckk.me/anarckk/go_util/go_bit"
	"gitea.bee.anarckk.me/anarckk/go_util/go_cli"
	"gitea.bee.anarckk.me/anarckk/go_util/go_code"
	"gitea.bee.anarckk.me/anarckk/go_util/go_math"
	"gitea.bee.anarckk.me/anarckk/go_util/go_number"
	"gitea.bee.anarckk.me/anarckk/go_util/go_random"
	"gitea.bee.anarckk.me/anarckk/go_util/go_time"
	"gitea.bee.anarckk.me/anarckk/go_util/go_unit"
)

func FromByteArrayToInputStream(_bytes []byte) io.Reader {
	return bytes.NewReader(_bytes)
}

func FromInputStreamToByteArray(is io.Reader) ([]byte, error) {
	bytes, err := io.ReadAll(is)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// FromInputStreamToFile 从输入流复制数据到文件，直到 io.EOF 结束
//
// @param inputStream
// @param absPath
// @return string 返回文件的md5
// @return error
func FromInputStreamToFile(inputStream io.Reader, absPath string) (string, error) {
	targetFile, err := os.Create(absPath)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()
	bw := bufio.NewWriter(targetFile)
	defer bw.Flush()

	md5h := md5.New()
	writer := io.MultiWriter(bw, md5h)

	Copy(writer, inputStream)
	bw.Flush()
	return hex.EncodeToString(md5h.Sum(nil)), err
}

// FromFileToOutputStream1 从文件复制到输出流，直到 io.EOF 结束
func FromFileToOutputStream(absPath string, outputStream io.Writer) error {
	return FromFileToOutputStream1(absPath, outputStream)
}

// FromFileToOutputStream1 从文件复制到输出流，直到 io.EOF 结束
func FromFileToOutputStream1(absPath string, outputStream io.Writer) error {
	inputStream, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer inputStream.Close()
	br := bufio.NewReader(inputStream)
	_, err = br.WriteTo(outputStream)
	return err
}

// FromFileToOutputStream2 从文件复制到输出流，直到 io.EOF 结束
func FromFileToOutputStream2(absPath string, outputStream io.Writer) error {
	inputStream, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer inputStream.Close()
	br := bufio.NewReader(inputStream)
	_, err = Copy(outputStream, br)
	return err
}

// Deprecated: 还是不要这个的好，直接写一个新的stream套住这个东西，
// 直接用stream层层叠叠的套就好了，也很简单
type ByteArrayCallback func(_byte []byte)

// Deprecated: 废弃
// 通过闭包返回一个函数，这个函数作为流速率的打印器，每秒在控制器台中更新当前流的速率、已传输量
func AutoPrintSpeed(ctx context.Context) ByteArrayCallback {
	ch := make(chan int64, 10)
	go func() {
		var total int64 = 0
		var size int64 = 0
		log.Println("速度计算程序启动")
		ticker := time.NewTicker(time.Second)
		printer := go_cli.GetCliPrinter()
	FF:
		for {
			select {
			case s := <-ch:
				size += s
				total += s
			case <-ticker.C:
				printer(fmt.Sprintf("%s 管道中每秒流速为:%s/s 当前传输量:%s", go_time.CurrentDatetimeStr(), go_unit.HumanReadableByteCountBin(size), go_unit.HumanReadableByteCountBin(total)))
				size = 0
			case <-ctx.Done():
				printer(fmt.Sprintf("%s 管道中每秒流速为:%s/s 当前传输量:%s", go_time.CurrentDatetimeStr(), go_unit.HumanReadableByteCountBin(size), go_unit.HumanReadableByteCountBin(total)))
				break FF
			}
		}
		fmt.Println()
		log.Println("finish.")
	}()
	return func(_byte []byte) {
		ch <- int64(len(_byte))
	}
}

// Deprecated: 废弃
// 通过闭包返回一个函数，这个函数作为流速率的打印器，每秒在控制器台中更新当前流的速率、(已传输量/总传输量）、进度百分比、剩余时间
func AutoPrintSpeedProgress(ctx context.Context, allDataLen int64) ByteArrayCallback {
	ch := make(chan int64, 10)
	go func() {
		if wg, ok := ctx.Value(wgKey{}).(*sync.WaitGroup); ok {
			wg.Add(1)
			defer wg.Done()
		}
		var total int64 = 0
		var size int64 = 0
		log.Println("AutoPrintSpeedProgress 速度计算程序启动")
		ticker := time.NewTicker(time.Second)
		printer := go_cli.GetCliPrinter()
		start := go_time.CurrentTimeSecond()
		logFn := func() {
			var p float64 = float64(total) / float64(allDataLen)
			end := go_time.CurrentTimeSecond()
			average := float64(total) / float64(end-start)
			printer(fmt.Sprintf("%s 管道中每秒流速为: %s/s, 平均流速为: %s/s, 当前传输量:%s(%d/%d), 传输进度: %s",
				go_time.CurrentDatetimeStr(),
				go_unit.HumanReadableByteCountBin(size), go_unit.HumanReadableByteCountBin(go_math.Ceil(average)),
				go_unit.HumanReadableByteCountBin(total),
				total, allDataLen, go_number.GetPercentStr(p)))
		}
	FF:
		for {
			select {
			case s := <-ch:
				size += s
				total += s
			case <-ticker.C:
				logFn()
				size = 0
			case <-ctx.Done():
				logFn()
				close(ch)
				break FF
			}
		}
		fmt.Println()
	}()
	return func(_byte []byte) {
		if ctx.Err() == nil {
			ch <- int64(len(_byte))
		}
	}
}

// Deprecated: 废弃
type IStreamCopy interface {
	Source(r io.Reader) IStreamCopy

	Target(w io.Writer) IStreamCopy

	Peek(ByteArrayCallback) IStreamCopy

	RunCopy() (int64, error)
}

var _ IStreamCopy = &streamCopy{}

// Deprecated: 废弃
type streamCopy struct {
	w    io.Writer
	r    io.Reader
	peek ByteArrayCallback
}

// Deprecated: 废弃
func NewStreamCopy() IStreamCopy {
	return &streamCopy{}
}

// Deprecated: 废弃
func (s *streamCopy) Source(r io.Reader) IStreamCopy {
	s.r = r
	return s
}

// Deprecated: 废弃
func (s *streamCopy) Target(w io.Writer) IStreamCopy {
	s.w = w
	return s
}

// Deprecated: 废弃
func (s *streamCopy) Peek(peek ByteArrayCallback) IStreamCopy {
	s.peek = peek
	return s
}

// Deprecated: 废弃
func (s *streamCopy) RunCopy() (int64, error) {
	if s.w == nil {
		return 0, errors.New("writer is nil")
	}
	if s.r == nil {
		return 0, errors.New("reader is nil")
	}
	if s.peek != nil {
		return CopyPeek(s.w, s.r, s.peek)
	} else {
		return Copy(s.w, s.r)
	}
}

func Copy(w io.Writer, r io.Reader) (int64, error) {
	return Copy2(w, r)
}

func Copy1(w io.Writer, r io.Reader) (int64, error) {
	return io.Copy(w, r)
}

func Copy2(w io.Writer, r io.Reader) (int64, error) {
	var tmp = make([]byte, 1024)
	var l int64
	for {
		n, err := r.Read(tmp)
		if err == io.EOF {
			return l, nil
		}
		if err != nil {
			return l, err
		}
		nn, err := w.Write(tmp[:n])
		if err != nil {
			return l, err
		}
		if nn != n {
			return l, errors.New("写入的数据和读取的数据不一致")
		}
		l += int64(nn)
	}
}

// Deprecated: 废弃
func CopyPeek(w io.Writer, r io.Reader, peek ByteArrayCallback) (int64, error) {
	peekWriter := NewPeekWriter(w, peek)
	return Copy(peekWriter, r)
}

func ReadInt64FromInputStream(reader io.Reader) (value int64, err error) {
	err = binary.Read(reader, binary.BigEndian, &value)
	return
}

// ReadNBytesFromInputStream 从流中读取指定的字节数
//
//	@param reader
//	@param num
//	@return []byte
//	@return error
func ReadNBytesFromInputStream(reader io.Reader, num int) ([]byte, error) {
	return ReadNBytesFromInputStream2(reader, num)
}

func ReadNBytesFromInputStream1(reader io.Reader, num int) ([]byte, error) {
	var _bytes = make([]byte, 0, num)
	total := 0
	for {
		var tmp = make([]byte, (num - total))
		n, err := reader.Read(tmp)
		if err != nil {
			return []byte{}, err
		}
		for i := 0; i < n; i++ {
			_bytes = append(_bytes, tmp[i])
		}
		total += n
		if total == num {
			break
		}
	}
	if len(_bytes) != num {
		return []byte{}, fmt.Errorf("期望读取的字节数是 %d , 实际读取的字节数是 %d", num, len(_bytes))
	}
	return _bytes, nil
}

func ReadNBytesFromInputStream2(reader io.Reader, size int) (result []byte, err error) {
	result = make([]byte, size)
	_, err = io.ReadFull(reader, result)
	return
}

// BitWriteIntToOutputStream 按照自己规则(传输指定的字节量) - 写入一个数字
func BitWriteIntToOutputStream(a int, writer io.Writer) error {
	_, err := writer.Write(go_bit.IntToBytes(a))
	return err
}

// BitWriteInt64ToOutputStream 按照自己规则(传输指定的字节量) - 写入一个数字
func BitWriteInt64ToOutputStream(a int64, writer io.Writer) error {
	_, err := writer.Write(go_bit.Int64ToBytes(a))
	return err
}

// BitReadInt32FromInputStream 按照自己规则(传输指定的字节量) - 读取一个数字
func BitReadInt32FromInputStream(reader io.Reader) (int, error) {
	bytes, err := ReadNBytesFromInputStream(reader, 4)
	if err != nil {
		return 0, err
	}
	return go_bit.BytesToInt(bytes), nil
}

// BitReadInt64FromInputStream 按照自己规则(传输指定的字节量) - 读取一个数字
func BitReadInt64FromInputStream(reader io.Reader) (int64, error) {
	bytes, err := ReadNBytesFromInputStream(reader, 8)
	if err != nil {
		return 0, err
	}
	return go_bit.BytesToInt64(bytes), nil
}

// BitWriteStringToOutputStream 按照自己规则(传输指定的字节量) - 写入一个字符串
func BitWriteStringToOutputStream(str string, writer io.Writer) error {
	bytes := []byte(str)
	err := BitWriteIntToOutputStream(len(bytes), writer)
	if err != nil {
		return err
	}
	_, err = writer.Write(bytes)
	return err
}

// BitReadStringFromStream 按照自己规则 - 读取一个字符串
func BitReadStringFromStream(reader io.Reader) (string, error) {
	_len, err := BitReadInt32FromInputStream(reader)
	if err != nil {
		return "", err
	}
	bytes, err := ReadNBytesFromInputStream(reader, _len)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// BitRawToOutputStream 把数据写入到运输流中
// 如果还需要带长度入参的方法，还是再写一个api吧
// 返回发送流的md5
func BitFromRawToOutputStream(raw io.Reader, writer io.Writer) (string, error) {
	var tmp [1024]byte
	getBinMd5 := go_code.GetBinMd5()
	for {
		n, err := raw.Read(tmp[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		_, err = writer.Write(tmp[0:n])
		if err != nil {
			return "", err
		}
		_, err = getBinMd5.Write(tmp[0:n])
		if err != nil {
			return "", err
		}
	}
	return getBinMd5.GetMd5(), nil
}

// BitFromInputStreamToOutputStream 从运输流中读取数据
func BitFromInputStreamToRaw(reader io.Reader, raw io.Writer, size int64) (string, error) {
	var tmpLen int = min(int(size), 1024)
	var tmp = make([]byte, tmpLen)
	var total int64
	getBinMd5 := go_code.GetBinMd5()
	for {
		tmpReadLen, err := reader.Read(tmp)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		_, err = raw.Write(tmp[0:tmpReadLen])
		if err != nil {
			return "", err
		}
		_, err = getBinMd5.Write(tmp[0:tmpReadLen])
		if err != nil {
			return "", err
		}
		total += int64(tmpReadLen)
		if total >= size {
			break
		}
		if int(size-total) < tmpLen {
			tmpLen = int(size - total)
			tmp = make([]byte, tmpLen)
		}
	}
	return getBinMd5.GetMd5(), nil
}

type IBitFromFileToOutputStream interface {
	Writer(io.Writer) IBitFromFileToOutputStream
	AbsPath(string) IBitFromFileToOutputStream
	ShowProgress(bool) IBitFromFileToOutputStream
	Run() error
}

var _ IBitFromFileToOutputStream = &_BitFromFileToOutputStream{}

type _BitFromFileToOutputStream struct {
	outputStream io.Writer
	absPath      string
	showProgress bool
}

func NewBitFromFileToOutputStream() IBitFromFileToOutputStream {
	return &_BitFromFileToOutputStream{}
}

func (b *_BitFromFileToOutputStream) AbsPath(absPath string) IBitFromFileToOutputStream {
	b.absPath = absPath
	return b
}

func (b *_BitFromFileToOutputStream) ShowProgress(showProgress bool) IBitFromFileToOutputStream {
	b.showProgress = showProgress
	return b
}

func (b *_BitFromFileToOutputStream) Writer(outputStream io.Writer) IBitFromFileToOutputStream {
	b.outputStream = outputStream
	return b
}

// context 中存放 sync.WaitGroup 的 key
type wgKey struct{}

func (b *_BitFromFileToOutputStream) Run() error {
	size := GetFileSize(b.absPath)
	err := BitWriteInt64ToOutputStream(size, b.outputStream)
	if err != nil {
		return err
	}

	inputStream, err := os.Open(b.absPath)
	if err != nil {
		return err
	}
	defer inputStream.Close()
	br := bufio.NewReader(inputStream)

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), wgKey{}, wg))
	defer cancel()
	var outputStream io.Writer
	if b.showProgress { // TODO 改一下这个东西，把这个改成一个流，简单点用，不要在 ctx 里套值，麻烦
		outputStream = NewPeekWriter(b.outputStream, AutoPrintSpeedProgress(ctx, size))
	} else {
		outputStream = b.outputStream
	}

	md5, err := BitFromRawToOutputStream(br, outputStream)
	if err != nil {
		return err
	}
	err = BitWriteStringToOutputStream(md5, b.outputStream)
	if err != nil {
		return err
	}
	cancel()
	wg.Wait()
	return nil
}

type IBitFromInputStreamToFile interface {
	Reader(io.Reader) IBitFromInputStreamToFile
	AbsPath(string) IBitFromInputStreamToFile
	ShowProgress(bool) IBitFromInputStreamToFile
	Run() error
}

func NewBitFromInputStreamToFile() IBitFromInputStreamToFile {
	return &_BitFromInputStreamToFile{}
}

var _ IBitFromInputStreamToFile = &_BitFromInputStreamToFile{}

type _BitFromInputStreamToFile struct {
	inputStream  io.Reader
	absPath      string
	showProgress bool
}

func (b *_BitFromInputStreamToFile) Reader(reader io.Reader) IBitFromInputStreamToFile {
	b.inputStream = reader
	return b
}

func (b *_BitFromInputStreamToFile) AbsPath(absPath string) IBitFromInputStreamToFile {
	b.absPath = absPath
	return b
}

func (b *_BitFromInputStreamToFile) ShowProgress(showProgress bool) IBitFromInputStreamToFile {
	b.showProgress = showProgress
	return b
}

func (b *_BitFromInputStreamToFile) Run() error {
	size, err := BitReadInt64FromInputStream(b.inputStream)
	if err != nil {
		return err
	}
	outputStream, err := os.Create(b.absPath)
	if err != nil {
		return err
	}
	defer outputStream.Close()
	bw := bufio.NewWriter(outputStream)
	defer bw.Flush()

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), wgKey{}, wg))
	defer cancel()
	var writer io.Writer
	if b.showProgress { // TODO 改造
		writer = NewPeekWriter(bw, AutoPrintSpeedProgress(ctx, size))
	} else {
		writer = bw
	}
	md5, err := BitFromInputStreamToRaw(b.inputStream, writer, size)
	if err != nil {
		return err
	}
	md5Next, err := BitReadStringFromStream(b.inputStream)
	if err != nil {
		return err
	}
	if md5 != md5Next {
		return errors.New("md5校验不一致!原始md5: " + md5 + ", 收到md5: " + md5Next)
	}
	cancel()
	wg.Wait()
	return nil
}

// Deprecated: 废弃
var _ io.Writer = &PeekWriter{}

// Deprecated: 废弃
type PeekWriter struct {
	writer io.Writer
	peek   ByteArrayCallback
}

// Deprecated: 废弃
func NewPeekWriter(writer io.Writer, peek ByteArrayCallback) io.Writer {
	return &PeekWriter{writer, peek}
}

// Deprecated: 废弃
func (line *PeekWriter) Write(p []byte) (n int, err error) {
	line.peek(p)
	return line.writer.Write(p)
}

var _ io.Writer = &NullWriter{}

type NullWriter struct {
}

// Write implements io.Writer.
func (nWriter *NullWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func NewNullWriter() *NullWriter {
	return &NullWriter{}
}

func NewRandomReader(size uint64) io.Reader {
	return &RandomReader{size, 0}
}

var _ io.Reader = &RandomReader{}

// 每秒流速为:725.4 MiB/s
type RandomReader struct {
	size uint64
	i    uint64 // current reading index
}

func (r *RandomReader) Read(p []byte) (n int, err error) {
	if r.i >= r.size {
		return 0, io.EOF
	}
	bufSize := len(p)
	cap := int(r.size - r.i)
	randomBytes := go_random.RandomBytes(min(bufSize, cap))
	n = copy(p, randomBytes[:])
	r.i += uint64(n)
	return
}

var _ io.ReadCloser = &AutoSpeedInputStream{}

type AutoSpeedInputStream struct {
	reader     io.Reader
	allByteLen uint64
	_ch        chan uint64
	_wg        *sync.WaitGroup
}

func NewAutoSpeedInputStream(reader io.Reader, allByteLen uint64) *AutoSpeedInputStream {
	auto := &AutoSpeedInputStream{
		reader:     reader,
		allByteLen: allByteLen,
		_ch:        make(chan uint64, 10),
		_wg:        &sync.WaitGroup{},
	}
	go auto.autoSpeed()
	return auto
}

func (is *AutoSpeedInputStream) autoSpeed() {
	wg := is._wg
	wg.Add(1)
	defer wg.Done()

	allDataLen := is.allByteLen
	var total uint64 = 0
	var size uint64 = 0
	log.Println("AutoSpeedInputStream 速度计算程序启动")
	ticker := time.NewTicker(time.Second)
	printer := go_cli.GetCliPrinter()
	start := go_time.CurrentTimeSecond()
	logFn := func() {
		var p float64 = float64(total) / float64(allDataLen)
		end := go_time.CurrentTimeSecond()
		average := float64(total) / float64(end-start)
		printer(fmt.Sprintf("%s InputStream 每秒流速为: %s/s, 平均流速为: %s/s, 当前传输量:%s(%d/%d), 传输进度: %s",
			go_time.CurrentDatetimeStr(),
			go_unit.HumanReadableByteCountBin(int64(size)), go_unit.HumanReadableByteCountBin(go_math.Ceil(average)),
			go_unit.HumanReadableByteCountBin(int64(total)),
			total, allDataLen, go_number.GetPercentStr(p)))
	}
FF:
	for {
		select {
		case s, ok := <-is._ch:
			if !ok {
				logFn()
				break FF
			}
			size += s
			total += s
		case <-ticker.C:
			logFn()
			size = 0
		}
	}
	ticker.Stop()
	fmt.Println()
}

func (is *AutoSpeedInputStream) Close() error {
	close(is._ch)
	is._wg.Wait()
	return nil
}

func (is *AutoSpeedInputStream) Read(p []byte) (n int, err error) {
	is._ch <- uint64(len(p))
	return is.reader.Read(p)
}

var _ io.WriteCloser = &AutoSpeedOutStream{}

type AutoSpeedOutStream struct {
	writer     io.Writer
	allByteLen uint64
	_ch        chan uint64
	_wg        *sync.WaitGroup
}

func NewAutoSpeedOutStream(writer io.Writer, allByteLen uint64) *AutoSpeedOutStream {
	auto := &AutoSpeedOutStream{
		writer:     writer,
		allByteLen: allByteLen,
		_ch:        make(chan uint64, 10),
		_wg:        &sync.WaitGroup{},
	}
	go auto.autoSpeed()
	return auto
}

func (is *AutoSpeedOutStream) autoSpeed() {
	wg := is._wg
	wg.Add(1)
	defer wg.Done()

	allDataLen := is.allByteLen
	var total uint64 = 0
	var size uint64 = 0
	log.Println("AutoSpeedOutStream 速度计算程序启动")
	ticker := time.NewTicker(time.Second)
	printer := go_cli.GetCliPrinter()
	start := go_time.CurrentTimeSecond()
	logFn := func() {
		var p float64 = float64(total) / float64(allDataLen)
		end := go_time.CurrentTimeSecond()
		average := float64(total) / float64(end-start)
		printer(fmt.Sprintf("%s OutputStream 每秒流速为: %s/s, 平均流速为: %s/s, 当前传输量:%s(%d/%d), 传输进度: %s",
			go_time.CurrentDatetimeStr(),
			go_unit.HumanReadableByteCountBin(int64(size)), go_unit.HumanReadableByteCountBin(go_math.Ceil(average)),
			go_unit.HumanReadableByteCountBin(int64(total)),
			total, allDataLen, go_number.GetPercentStr(p)))
	}
FF:
	for {
		select {
		case s, ok := <-is._ch:
			if !ok {
				logFn()
				break FF
			}
			size += s
			total += s
		case <-ticker.C:
			logFn()
			size = 0
		}
	}
	ticker.Stop()
	fmt.Println()
}

func (is *AutoSpeedOutStream) Close() error {
	close(is._ch)
	is._wg.Wait()
	return nil
}

func (os *AutoSpeedOutStream) Write(p []byte) (n int, err error) {
	os._ch <- uint64(len(p))
	return os.writer.Write(p)
}
