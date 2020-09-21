package main

import (
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {

	//base64Test()
	//t := path.Join("t?","=","a")

	//fmt.Println(t)
	//b := path.Base("tmo/test.txt")
	//fmt.Println(b)
	//d, b1 := path.Split("tmp/1/2/test.txt")
	//fmt.Println(d, b1)
	//d1 := path.Dir("tmp/1/2/test.txt")
	//fmt.Println(d1)
	//fmt.Println(config.Config.BaseDir)
	//for _, task := range config.Config.Tasks {
	//	fmt.Printf("%+v\n",task)
	//}
	testgzip()
}

func testgzip() {
	f, err := os.Create("data/testgzip/test.gz")
	if err != nil {
		panic(err)
	}

	w, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	w.Name = "test.txt"
	w.Write([]byte("test asateat adtaet adtaes"))
	w.Flush()
	w.Close()
	f.Close()

}

func base64Test() {
	var str = "ZHONGGUOnihao123"
	strbytes := []byte(str)
	encoded := base64.StdEncoding.EncodeToString(strbytes)
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	decodestr := string(decoded)
	fmt.Println(decodestr, err)
}
