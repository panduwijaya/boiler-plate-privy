package main

import (
	"net/http"
	"testing"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"bytes"
	"log"
	"github.com/xuri/excelize/v2"
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

	const is_excel_sheet = true
	file, _ := ioutil.ReadFile("test.json")
	
	xlsx, err := excelize.OpenFile("./test_case.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	sheet1Name := "Sheet1"
	
	if is_excel_sheet == true {
		for i := 3; i < 15; i++ {
			var client = &http.Client{}
			number, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i))
			method, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i))
			module, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i))
			url, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i))
			payloads, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i))
			expecteds, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("F%d", i))

			var jsonStr = []byte(strings.ReplaceAll( payloads, `'`, `"`))
			fmt.Println("NO : ", number)
			fmt.Println("Method : ", method)
			fmt.Println("Module : ", module)
			fmt.Println("Url : ", url)

			request, err := http.NewRequest(string(method), string(url), bytes.NewBuffer(jsonStr))
			request.Header.Set("Content-Type", "application/json")
			response, err := client.Do(request)

			defer response.Body.Close()
			if err != nil {
				fmt.Println("No response from request")
			}
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body) // response body is []byte
			//fmt.Println(string(body))  

			expected := strings.ReplaceAll( expecteds, `'`, `"`)
			if string(body) != expected {
				t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
			}else{
				fmt.Println("Testing Passed : ", module )
			}
		}
	}else{
		var client = &http.Client{}
		data := ListTest{}
		_ = json.Unmarshal([]byte(file), &data)
		
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

}
