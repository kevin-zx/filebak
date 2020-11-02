package main

import (
	"context"
	"database/sql"
	"encoding/json"
	backup_rpc "filebak/pkg/service/grpc"
	"fmt"
	"github.com/kevin-zx/sqlbase"
	"google.golang.org/grpc"
	"log"
	"time"
)

type XHSData struct {
	Data string    `json:"data"`
	Seed string    `json:"seed"`
	Time time.Time `json:"time"`
	Type string    `json:"-"`
}

func main() {
	s, err := sqlbase.NewStorage("mysql", 3306, "spider_xhs", "spider_xhs", "47.103.97.36", "Z3ujLXYZ2UDwjHuY", sqlbase.DefaultConfig)
	if err != nil {
		panic(err)
	}

	port := 11858
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()
	c := backup_rpc.NewBackupServiceClient(conn)
	gs, err := c.SaveData(context.Background())
	if err != nil {
		panic(err)
	}
	offset := 0
	limit := 1000

	typeNames := map[string]byte{}
	tType := ""
	taskName := "xhs_spider"
	max := -1

	defer func() {
		for tn := range typeNames {
			r, err := c.CloseType(context.Background(), &backup_rpc.CloseRequest{
				TaskName: taskName,
				TypeName: tn,
			})
			if err != nil {
				panic(err)
			}
			if r.Err != nil && r.Err.Info != "" {
				log.Printf("close type %s err:%s", tn, r.Err.Info)
			}
		}
	}()

	for {
		if max != -1 && offset+limit > max {
			r, err := gs.CloseAndRecv()
			if err != nil {
				panic(err)
			}
			if r != nil {
				fmt.Printf("err:%v\n", r.Err)
			}
			break
		}
		fmt.Printf("loop:%d offset:%d \n", offset/limit+1, offset)
		data, err := GetData(s, limit, offset)
		if err != nil {
			panic(err)
		}
		for _, rd := range data {
			if rd.Type == "" {
				fmt.Printf("%+v\n", rd)
				continue
			} else {
				if rd.Type != tType {
					typeNames[rd.Type] = 1
					tType = rd.Type
				}
			}

			rawDb, err := json.Marshal(&rd)
			if err != nil {
				panic(err)
			}
			err = gs.Send(&backup_rpc.SaveRequest{
				Data: &backup_rpc.Data{
					DataType: rd.Type,
					Time:     rd.Time.Unix(),
					Data:     rawDb,
				},
				TaskName: taskName,
			})

			if err != nil {
				gs.CloseSend()
			}
		}
		if len(data) != limit {
			r, err := gs.CloseAndRecv()
			if err != nil {
				panic(err)
			}
			fmt.Printf("err:%v\n", r.Err)
			break
		}
		offset += limit
	}

	//for  {
	//
	//}
	//
}

func GetData(storage *sqlbase.Storage, limit int, offset int) ([]XHSData, error) {
	xhsdata := []XHSData{}
	err := storage.RawScan("SELECT seed,data,spider_type,created_at FROM spider_xhs.spider_result_20200914",
		" LIMIT ?,?", func(rows *sql.Rows) error {
			xhs := XHSData{
				//Data: "",
				//Seed: "",
				//Time: time.Time{},
				//Type: "",
			}
			err := rows.Scan(&xhs.Seed, &xhs.Data, &xhs.Type, &xhs.Time)
			if err != nil {
				return err
			}
			xhsdata = append(xhsdata, xhs)
			return nil
		}, offset, limit)
	if err != nil {
		return nil, err
	}
	return xhsdata, nil
}
