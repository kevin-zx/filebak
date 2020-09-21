package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	backup_rpc "filebak/pkg/service/grpc"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

type XHSData struct {
	Data string    `json:"data"`
	Seed string    `json:"seed"`
	Time time.Time `json:"time"`
	Type string    `json:"-"`
}

func main() {
	port := 11858
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer conn.Close()
	c := backup_rpc.NewBackupServiceClient(conn)

	dataFileName := "spider_result_20200907_202009181302.csv"
	dataFile, err := os.Open(dataFileName)
	if err != nil {
		panic(err)
	}
	defer dataFile.Close()
	csvF := csv.NewReader(dataFile)
	taskName := "xhs_spider"
	gs, err := c.SaveData(context.Background())
	if err != nil {
		panic(err)
	}
	//rd := XHSData{}
	typeNames := map[string]byte{}
	tType := ""
	for {
		rs, err := csvF.Read()
		if err == io.EOF {
			r, err := gs.CloseAndRecv()
			if err != nil {
				panic(err)
			}
			fmt.Printf("err:%v\n", r.Err)
			break
		}
		if err != nil {
			panic(err)
		}
		if rs[0] == "data" {
			continue
		}
		rd := XHSData{}
		rd.Data = rs[0]
		rd.Seed = rs[1]

		rd.Time, err = time.Parse("2006-01-02 15:04:05.0", rs[2])
		if err != nil {
			panic(err)
		}
		//rd.Time = t.Unix()
		spiderType := rs[3]
		if spiderType == "" {
			fmt.Printf("%+v\n", spiderType)
			continue
		}
		if spiderType != tType {
			typeNames[spiderType] = 1
			tType = spiderType
		}
		rawDb, err := json.Marshal(&rd)
		if err != nil {
			panic(err)
		}
		err = gs.Send(&backup_rpc.SaveRequest{
			Data: &backup_rpc.Data{
				DataType: spiderType,
				Time:     rd.Time.Unix(),
				Data:     rawDb,
			},
			TaskName: taskName,
		})

		if err != nil {
			gs.CloseSend()
		}

	}
	for tn := range typeNames {
		r, err := c.CloseType(context.Background(), &backup_rpc.CloseRequest{
			TaskName: taskName,
			TypeName: tn,
		})
		if err != nil {
			panic(err)
		}
		if r.Err != nil && r.Err.Info != "" {
			panic(r.Err.Info)
		}
	}

}
