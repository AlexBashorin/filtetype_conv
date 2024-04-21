package xlsxjson

import (
	"log"
	"net/http"

	"s3_getter/pkg/xlsx_json/excel_json"
	"s3_getter/pkg/xlsx_json/json_excel"

	"github.com/gorilla/mux"
)

func XtoJ() {
	r := mux.NewRouter()

	// JSON to EXCEL
	r.HandleFunc("/json_excel", json_excel.WriteExcel)

	// EXCEL to JSON
	r.HandleFunc("/excel_json", excel_json.WriteJson)

	log.Fatal(http.ListenAndServe(":4444", r))
}
