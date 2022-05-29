package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type NameApi struct {
	Err  int        `json:"err"`
	Name [][]string `json:name`
}

// var jsonstr = `
// callback({
// 	"err": 0,
// 	"name": [
// 	[
// 	"鈴木 真紀",
// 	"すずき まき",
// 	"Suzuki Maki"
// 	],
// 	[
// 	"木村 久美子",
// 	"きむら くみこ",
// 	"Kimura Kumiko"
// 	],
// 	[
// 	"田島 好治",
// 	"たしま よしはる",
// 	"Tasima Yosiharu"
// 	]
// 	]
// 	});
// `

// type JSONPWrapper struct {
// 	Prefix     string
// 	Underlying io.Reader
// 	gotPrefix  bool
// }

// func (jpw *JSONPWrapper) Read(b []byte) (int, error) {
// 	if jpw.gotPrefix {
// 		return jpw.Underlying.Read(b)
// 	}

// 	prefix := make([]byte, len(jpw.Prefix))
// 	n, err := io.ReadFull(jpw.Underlying, prefix)
// 	if err != nil {
// 		return n, err
// 	}

// 	if string(prefix) != jpw.Prefix {
// 		return n, fmt.Errorf("JSONP prefix mismatch: expected %q, got %q", jpw.Prefix, prefix)
// 	}

// 	char := make([]byte, 1)
// 	for char[0] != '(' {
// 		n, err = jpw.Underlying.Read(char)
// 		if n == 0 || err != nil {
// 			return n, err
// 		}
// 	}

// 	jpw.gotPrefix = true
// 	return jpw.Underlying.Read(b)
// }

// type jsonExample struct {
// 	Hello int `json:hello`
// 	Num   int `json:num`
// }

func main() {
	// jsonpStr := `callback({
	// 	"hello": "world!!",
	// 	"num": 123456
	// });`
	res, _ := http.Get("https://green.adam.ne.jp/roomazi/cgi-bin/randomname.cgi?n=3")
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	jsonStr := string(body)[strings.Index(string(body), "(")+1 : strings.Index(string(body), ")")]
	fmt.Println(jsonStr)
	// var v interface{}
	var name NameApi
	json.Unmarshal([]byte(jsonStr), &name)
	fmt.Printf("%+v", name)
	// res, _ := http.Get("https://green.adam.ne.jp/roomazi/cgi-bin/randomname.cgi?n=3")
	// defer res.Body.Close()

	// body, _ := io.ReadAll(res.Body)

	// prefix := "callback"
	// jsonp := bytes.NewBuffer([]byte(jsonstr))
	// var decoded jsonExample

	// decoder := json.NewDecoder(&JSONPWrapper{Prefix: prefix, Underlying: jsonp})
	// decoder.Decode(&decoded)
	// fmt.Println(decoded)

	// fmt.Println(string(body))
	// Jsonpがかえってくるため、トリミングがラップしないとParseできなさそう
	// var name NameApi
	// json.Unmarshal([]byte(string(body)), &name)
	for i := 0; i < len(name.Name[0]); i++ {
		fmt.Printf("名前: %s\n", name.Name[i][0])
	}

}
