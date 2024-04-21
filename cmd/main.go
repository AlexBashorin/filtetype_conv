package main

import (
	"net/http"
	"s3_getter/pkg/xlsx_json/excel_json"
	"s3_getter/pkg/xlsx_json/json_excel"
)

func main() {
	http.HandleFunc("/excel_json", excel_json.WriteJson)
	http.HandleFunc("/json_excel", json_excel.WriteExcel)
}
