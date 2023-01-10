package main

import (
	"net/http"
	"testing"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"bytes"
)

type ListTest struct {
	ListTest []TestList `json:"list_test"`
}
 
type TestList struct {
	Method string `json: "method"`
	Module string `json: "module"`
	Url string `json: "url"`
	Expected string `json: "expected"`
	Payload string `json: "payload"`
}

func TestGetEntries(t *testing.T) {

	file, _ := ioutil.ReadFile("test.json")
 
	data := ListTest{}

 
	_ = json.Unmarshal([]byte(file), &data)
	var client = &http.Client{}

	for i := 0; i < len(data.ListTest); i++ {
		fmt.Println("Method : ", data.ListTest[i].Method)
		fmt.Println("Module : ", data.ListTest[i].Module)
		fmt.Println("Url : ", data.ListTest[i].Url)
		//fmt.Println("Payload : ",  strings.ReplaceAll( data.ListTest[i].Payload, `'`, `"`))
		
		
		var jsonStr = []byte(strings.ReplaceAll( data.ListTest[i].Payload, `'`, `"`))

		request, err := http.NewRequest(data.ListTest[i].Method, data.ListTest[i].Url, bytes.NewBuffer(jsonStr))
		request.Header.Set("Content-Type", "application/json")
		response, err := client.Do(request)

		defer response.Body.Close()
		if err != nil {
			fmt.Println("No response from request")
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body) // response body is []byte
		//fmt.Println(string(body))  

		expected := strings.ReplaceAll( data.ListTest[i].Expected, `'`, `"`)
		if string(body) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
		}else{
			fmt.Println("Testing Passed : ", data.ListTest[i].Module )
		}

		
	}


}
