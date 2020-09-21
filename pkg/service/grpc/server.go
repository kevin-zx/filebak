package backup_rpc

import (
	"context"
	"filebak/pkg/data"
	"filebak/pkg/service"
	"io"
	"log"
	"time"
)

func NewBackupRPCService(s service.Service) BackupServiceServer {
	return &backup{s: s}
}

type backup struct {
	s service.Service
}

func (b *backup) CloseType(ctx context.Context, request *CloseRequest) (*CloseResponse, error) {
	err := b.s.Close(request.TaskName, request.TypeName, request.Time)
	var errInfo *Error
	if err != nil {
		errInfo = &Error{
			Info: err.Error(),
		}
	}
	return &CloseResponse{
		Err: errInfo,
	}, nil
}

func (b *backup) SaveData(server BackupService_SaveDataServer) error {
	var tmpData *SaveRequest
	for {
		log.Printf("I'm in\n")
		dataRequest, err := server.Recv()
		if err == io.EOF {
			if tmpData != nil && tmpData.TaskName != "" {
				b.s.Close(tmpData.TaskName, tmpData.Data.DataType, tmpData.Data.Time)
			}
			err = server.SendAndClose(&SaveResponse{
				Err: nil,
			})
			if err != nil {
				return nil
			}
			break
		}
		tmpData = dataRequest
		if err != nil {
			err = server.SendAndClose(&SaveResponse{
				Err: &Error{
					Info: err.Error(),
				},
			})
			if err != nil {
				return nil
			}
			break
		}

		d := data.Data{
			DataType: dataRequest.Data.DataType,
			Time:     time.Unix(dataRequest.Data.Time, 0),
			Data:     dataRequest.Data.Data,
		}

		err = b.s.Save(dataRequest.TaskName, d)
		if err != nil {
			return err
		}
	}
	return nil
}
