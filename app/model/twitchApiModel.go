package model

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"myapp/app/structs"
	"net/http"
	"os"
)

var twitchToken string = ""

func TwitchEventSubModel(twitchBody []byte) structs.EventSubRespose {
	var eventSubJson structs.EventSubRespose
	json.Unmarshal(twitchBody, &eventSubJson)
	// responseStr := fmt.Sprintf("%v", testJson["challenge"])
	return eventSubJson
}

func SetSubscriptions() structs.SubscriptionList {
	if twitchToken == "" {
		getTwitchToken()
	}

	twitchId := os.Getenv("TWITCHID")
	secret := os.Getenv("SECRET")
	callbackUrl := os.Getenv("CallBackUrl")
	twitchSubscriptionsUrl := os.Getenv("TwitchSubscriptionsUrl")
	var newSubscriptions structs.SubscriptionRequset
	newSubscriptions.Type = "channel.chat.notification"
	newSubscriptions.Version = "1"
	newSubscriptions.Condition.BroadcasterUserId = twitchId
	newSubscriptions.Condition.UserId = twitchId
	newSubscriptions.Transport.Method = "webhook"
	newSubscriptions.Transport.Callback = callbackUrl + "eventsub/callback"
	newSubscriptions.Transport.Secret = secret
	jsonBody, _ := json.Marshal(newSubscriptions)

	twitchApiUrl := twitchSubscriptionsUrl

	clientId := os.Getenv("ClientId")
	req, err := http.NewRequest("POST", twitchApiUrl, bytes.NewReader([]byte(jsonBody)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-Id", clientId)
	req.Header.Set("Authorization", "Bearer "+twitchToken)
	if err != nil {
		panic("create twitch eventsub error")
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic("create twitch eventsub error")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic("create twitch eventsub error")
	}
	var createResponse structs.SubscriptionList
	json.Unmarshal(body, &createResponse)
	return createResponse
}

func SubscriptionsList() structs.SubscriptionList {
	if twitchToken == "" {
		getTwitchToken()
	}
	return twitchSubscriptionsGet()
}

func DeleteUnUseSubscriptions() bool {
	subscriptionList := twitchSubscriptionsGet()
	clientId := os.Getenv("ClientId")
	twitchSubscriptionsUrl := os.Getenv("TwitchSubscriptionsUrl")
	for _, element := range subscriptionList.Data {
		if element.Status != "enabled" {
			twitchApiUrl := twitchSubscriptionsUrl
			twitchApiUrl = twitchApiUrl + "?id=" + element.Id
			req, err := http.NewRequest("DELETE", twitchApiUrl, nil)
			if err != nil {
				return false
			}
			req.Header.Set("Client-Id", clientId)
			req.Header.Set("Authorization", "Bearer "+twitchToken)
			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				return false
			}
			defer response.Body.Close()
			if response.StatusCode == 401 {
				getTwitchToken()
				return false
			}

			body, err := io.ReadAll(response.Body)
			if err != nil {
				panic("get twitch SubscriptionsList error")
			}
			fmt.Print(string(body))
		} else {
			continue
		}
	}
	return true
}

func twitchSubscriptionsGet() structs.SubscriptionList {
	var subscriptionList structs.SubscriptionList
	clientId := os.Getenv("ClientId")
	twitchSubscriptionsUrl := os.Getenv("TwitchSubscriptionsUrl")
	twitchApiUrl := twitchSubscriptionsUrl
	req, err := http.NewRequest("GET", twitchApiUrl, nil)
	if err != nil {
		panic("get twitch SubscriptionsList error")
	}
	req.Header.Set("Authorization", "Bearer "+twitchToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Client-Id", clientId)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic("get twitch SubscriptionsList error")
	}
	defer response.Body.Close()

	if response.StatusCode == 401 {
		getTwitchToken()
		panic("need get token")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic("get twitch SubscriptionsList error")
	}
	json.Unmarshal(body, &subscriptionList)
	return subscriptionList
}

func getSecret() []byte {
	secret := os.Getenv("SECRET")
	return []byte(secret)
}

func getHmac(callbackBody string) string {
	mac := hmac.New(sha256.New, getSecret())
	mac.Write([]byte(callbackBody))
	expectedMAC := mac.Sum(nil)
	return hex.EncodeToString(expectedMAC)
}

func getTwitchToken() {
	twitchTokenUrl := os.Getenv("TwitchTokenUrl")
	twitchApiUrl := twitchTokenUrl
	clientId := os.Getenv("ClientId")
	clientSecret := os.Getenv("ClientSecret")

	postData := "client_id=" + clientId + "&client_secret=" + clientSecret + "&grant_type=client_credentials"

	response, err := http.Post(twitchApiUrl, "application/x-www-form-urlencoded", bytes.NewBufferString(postData))
	if err != nil {
		panic("get twitch token error")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic("get twitch token error")
	}
	var tokenResponse structs.TokenRes
	json.Unmarshal(body, &tokenResponse)
	twitchToken = tokenResponse.AccessToken
}
