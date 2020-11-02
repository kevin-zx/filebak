package backup

import (
	"filebak/pkg/data"
	"filebak/pkg/filechannel"
	"filebak/pkg/infrastructure/filetool"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Backup interface {
	Save(data data.Data) error
	Close(dataType string, t int64) error
	Destroy(dataType string, t time.Time) error
}

func NewBackup(baseDir string, task data.Task) (Backup, error) {
	err := filetool.CheckDirName(baseDir)
	if err != nil {
		return nil, fmt.Errorf("base dir is illegal, %s", err.Error())
	}
	err = filetool.CheckFileName(task.Name)
	if err != nil {
		return nil, fmt.Errorf("task name is illegal, %s", err.Error())
	}
	if !strings.HasSuffix(baseDir, "/") {
		baseDir += "/"
	}

	taskDir := baseDir + task.Name + "/"
	exist, err := filetool.DirIsExist(taskDir)
	if err != nil {
		return nil, err
	}
	if !exist {
		err = os.MkdirAll(taskDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	return &backup{task: task, taskDir: taskDir, writers: make(map[string]filechannel.WriteChannel)}, nil
}

type backup struct {
	taskDir string
	task    data.Task
	writers map[string]filechannel.WriteChannel
}

func (b *backup) Close(dataType string, t int64) error {
	if wc, ok := b.writers[dataType]; ok {
		err := wc.Close()
		if err != nil {
			return err
		}
		delete(b.writers, dataType)
		return nil
	} else {
		return fmt.Errorf("dataType %s not open", dataType)
	}

}

var blank = []byte("\n")

func (b *backup) Save(data data.Data) error {
	c, err := b.getChannel(data)
	if err != nil {
		return err
	}
	_, err = c.Write(append(data.Data, blank...))
	//log.Printf("write file success\n")
	return err
}

func (b *backup) getChannel(d data.Data) (filechannel.WriteChannel, error) {
	fn, err := b.fileName(d)
	if err != nil {
		return nil, err
	}
	if wc, ok := b.writers[d.DataType]; ok {
		if wc.FileName() != fn {
			log.Printf("task:%s type: %s change name from %s to %s\n", b.task.Name, d.DataType, wc.FileName(), fn)
			_ = wc.Close()
			delete(b.writers, d.DataType)
		} else {

			return wc, nil
		}
	}

	var wc filechannel.WriteChannel
	if b.task.Compress == true {
		wc, err = filechannel.NewGzipWriteChannel(fn)
	} else {
		wc, err = filechannel.NewWriteChannel(fn)
	}
	if err != nil {
		return nil, err
	}
	b.writers[d.DataType] = wc
	log.Printf("task:%s type: %s new name %s\n", b.task.Name, d.DataType, fn)
	//log.Printf("get file channel success,filename: %s\n", wc.FileName())
	return wc, nil
}

var fileTimeFmt = "%s%s.dat"

func (b *backup) fileName(d data.Data) (string, error) {
	if err := filetool.CheckDirName(d.DataType); err != nil {
		return "", fmt.Errorf("data type err, %s", err.Error())
	}
	typeDir := b.taskDir + d.DataType + "/"
	exist, err := filetool.DirIsExist(typeDir)
	if err != nil {
		return "", err
	}
	if !exist {
		err = os.Mkdir(typeDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	if d.Time.IsZero() {
		return "", fmt.Errorf("data time not set")
	}
	fileName := ""
	switch b.task.RollingType {
	case data.RollingTimeDay:
		fileName = fmt.Sprintf(fileTimeFmt, typeDir, d.Time.Format("2006_01_02"))
	case data.RollingTimeHour:
		fileName = fmt.Sprintf(fileTimeFmt, typeDir, d.Time.Format("2006_01_02_15"))
	case data.RollingTimeMin:
		fileName = fmt.Sprintf(fileTimeFmt, typeDir, d.Time.Format("2006_01_02_15_04"))
	default:
		if b.task.RollingType == 0 {
			return "", fmt.Errorf("rollingType not set")
		}
		return "", fmt.Errorf("unknow rollingType")
	}
	return fileName, nil
}

func (b *backup) Destroy(dataType string, t time.Time) error {
	d := data.Data{
		DataType: dataType,
		Time:     t,
		Data:     nil,
	}
	fn, err := b.fileName(d)
	if err != nil {
		return err
	}
	exist, err := filetool.FileIsExist(fn)
	if err != nil {
		return err
	}
	if exist {
		return os.Remove(fn)
	}
	return nil
}
