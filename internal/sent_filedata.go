package internal

import (
	"bytes"
	"fmt"
	"net/http"
)

func SentB64(byteArr string) {

	url := "https://4gso5mt55fz3w.elma365.ru/api/extensions/fd7decc3-08bd-450a-9fe6-f7ea5772bd14/script/getbig"

	fmt.Printf("size: %v\n", len(byteArr))

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(byteArr))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Чтение тела ответа
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	fmt.Printf("Ответ: %v\n", buf.String())
}
