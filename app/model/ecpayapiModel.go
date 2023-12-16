package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"myapp/app/structs"
	"net/http"
)

func ECPayApiModel(ecpayId string) (*[]structs.GreenWorld, error) {
	greenWorld := []structs.GreenWorld{}
	url := "https://payment.ecpay.com.tw/Broadcaster/CheckDonate/" + ecpayId
	postBody, _ := json.Marshal([]string{})
	responseBody := bytes.NewBuffer(postBody)
	response, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	json.Unmarshal(body, &greenWorld)
	return &greenWorld, nil
}
