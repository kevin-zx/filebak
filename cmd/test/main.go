package main

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"

	//"github.com/tidwall/gjson"
	"os"
)

func main() {
	data := `{"a":"1", "b":"2"}`
	data2 := `{"c":"1"}`
	d1 := gjson.Parse(data)
	d2 := gjson.Parse(data2)
	md1 := d1.Value().(map[string]interface{})
	md2 := d2.Value().(map[string]interface{})
	md1["a"] = md2
	dd, err := json.Marshal(md1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dd))
	//	data := []byte(`{
	//  "person": {
	//    "name": {
	//      "first": "Leonid",
	//      "last": "Bugaev",
	//      "fullName": "Leonid Bugaev"
	//    },
	//    "github": {
	//      "handle": "buger",
	//      "followers": 109
	//    },
	//    "avatars": [
	//      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
	//    ]
	//  },
	//  "company": {
	//    "name": "Acme"
	//  }
	//}`)
	//
	//	v,err := jsonparser.Set(data,[]byte("111"),"person","name","first")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(string(v))
	//r := gjson.Parse(`{"name":{"first":"Janet","last":"Prichard"},"age":47}`)
	//gjs
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
	//testgzip()
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
