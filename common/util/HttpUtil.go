package util

import (
	"cracker.com/base/common/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(config.GetCommonConfig().HttpTimeout()) * time.Second,
	}
}

type HttpException struct {
	error string
}

func (h *HttpException) Error() string {
	return h.error
}

func Get(url string) (string, error) {
	client := GetHttpClient()
	response, err := client.Get(url)
	if err != nil {
		log.Println("request err:", err)
		return "", &HttpException{
			error: err.Error(),
		}
	}
	defer func() {
		if response != nil && response.Body != nil {
			response.Body.Close()
		}
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("response read err:", err)
		return "", &HttpException{
			error: err.Error(),
		}
	}
	result := string(body)
	log.Println(result)
	return result, nil
}
