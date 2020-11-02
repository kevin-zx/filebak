package service

import (
	"filebak/pkg/backup"
	"filebak/pkg/config"
	"filebak/pkg/data"
	"fmt"
	"time"
)

type Service interface {
	Save(taskName string, data data.Data) error
	Close(taskName string, typeName string, time int64) error
	Destroy(taskName string, typeName string, t time.Time) error
}

func NewService() (Service, error) {
	backups := map[string]backup.Backup{}
	for _, task := range config.Config.Tasks {
		b, err := backup.NewBackup(config.Config.BaseDir, task)
		if err != nil {
			return nil, err
		}
		backups[task.Name] = b
	}
	return &service{backups: backups}, nil
}

type service struct {
	backups map[string]backup.Backup
}

func (s *service) Close(taskName string, typeName string, time int64) error {
	if b, ok := s.backups[taskName]; ok {
		return b.Close(typeName, time)
	} else {
		return fmt.Errorf("unknown task name: %s", taskName)
	}
}

func (s *service) Save(taskName string, data data.Data) error {
	if b, ok := s.backups[taskName]; ok {
		return b.Save(data)
	} else {
		return fmt.Errorf("unknown task name: %s", taskName)
	}
}

func (s *service) Destroy(taskName string, typeName string, t time.Time) error {
	if b, ok := s.backups[taskName]; ok {
		return b.Destroy(typeName, t)
	} else {
		return fmt.Errorf("unknown task name: %s", taskName)
	}
}
