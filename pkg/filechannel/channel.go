package filechannel

import (
	"bufio"
	"compress/gzip"
	"filebak/pkg/infrastructure/filetool"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

type WriteChannel interface {
	io.WriteCloser
	FileName() string
	Flush() error
	IsClose() bool
}

func periodFlush(channel WriteChannel) {
	for !channel.IsClose() {
		//_ = channel.Flush()
		time.Sleep(10 * time.Second)
	}
}

// ErrExist 文件已经存在了，不能重复创建
var ErrExist = fmt.Errorf("file exist")

func NewGzipWriteChannel(fileName string) (WriteChannel, error) {
	f, err := createChannelFile(fileName + ".gz")
	if err != nil {
		return nil, err
	}
	w, err := gzip.NewWriterLevel(f, gzip.BestCompression)

	if err != nil {
		return nil, err
	}
	w.Name = path.Base(fileName)
	w.ModTime = time.Now()

	gc := &gzipWriteChannel{
		closer: f,
		gzipW:  w,
		mu:     new(sync.Mutex),
		fn:     fileName,
		close:  false,
	}
	return gc, nil
}

type gzipWriteChannel struct {
	closer io.Closer
	gzipW  *gzip.Writer
	mu     *sync.Mutex
	fn     string
	close  bool
}

func (g *gzipWriteChannel) IsClose() bool {
	return g.close
}

func (g *gzipWriteChannel) Flush() error {
	return g.gzipW.Flush()
	//panic("implement me")
}

//func (g gzipWriteChannel) periodFlush() {
//	for !g.close {
//		_ = g.Flush()
//		time.Sleep(10 * time.Second)
//	}
//}

func (g *gzipWriteChannel) FileName() string {
	return g.fn
}

func (g *gzipWriteChannel) Write(p []byte) (n int, err error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.gzipW.Write(p)
}

func (g *gzipWriteChannel) Close() error {
	g.mu.Lock()
	defer g.mu.Unlock()
	log.Printf("fileName: %s close\n", g.fn)
	g.close = true
	fErr := g.gzipW.Flush()
	if fErr != nil {
		log.Printf("%+v\n", fErr)
	}
	errInfo := ""
	err2 := g.gzipW.Close()
	if err2 != nil {
		errInfo += err2.Error() + "\n"
	}
	err1 := g.closer.Close()
	if err1 != nil {
		errInfo += err1.Error() + "\n"
	}

	if errInfo != "" {
		return fmt.Errorf(errInfo)
	}
	return nil
}

func NewWriteChannel(fileName string) (WriteChannel, error) {
	f, err := createChannelFile(fileName)
	if err != nil {
		return nil, err
	}
	w := bufio.NewWriterSize(f, 4096*10)
	wc := writeChannel{write: w, closer: f, fn: fileName}
	//go periodFlush(&wc)
	return &wc, nil
}

func createChannelFile(fileName string) (*os.File, error) {
	err := existJudge(fileName)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(fileName)
	return f, nil
}

type writeChannel struct {
	closer io.Closer
	write  *bufio.Writer
	mu     *sync.Mutex
	fn     string
	close  bool
}

func (w *writeChannel) Flush() error {
	return w.write.Flush()
}

func (w *writeChannel) IsClose() bool {
	return w.close
	//panic("implement me")
}

func (w *writeChannel) FileName() string {
	return w.fn
}

func (w *writeChannel) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.write.Write(p)
}

func (w *writeChannel) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.close = true
	_ = w.write.Flush()
	return w.closer.Close()
}

func existJudge(fileName string) error {
	exist, err := filetool.FileIsExist(fileName)
	if err != nil {
		return err
	}
	if exist {
		return ErrExist
	}
	return nil
}
