package util

import (
	"bytes"
	"fmt"
	"github.com/crack007/api-base/common/config"
	"github.com/crack007/api-base/core"
	"io/ioutil"
	"net/http"
	"net/url"
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

func Get(requestUrl string) (string, error) {
	client := GetHttpClient()
	response, err := client.Get(requestUrl)
	if err != nil {
		core.GetLogger().Error(fmt.Sprintf("request err: %s", err))
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
		core.GetLogger().Error(fmt.Sprintf("response read err: %s", err))
		return "", &HttpException{
			error: err.Error(),
		}
	}
	result := string(body)
	core.GetLogger().Info(result)
	return result, nil
}

func PostForm(requestUrl string, data url.Values) (string, error) {
	response, err := http.PostForm(requestUrl, data)
	if err != nil {
		core.GetLogger().Error(fmt.Sprintf("request err: %s", err))
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
		core.GetLogger().Error(fmt.Sprintf("response read err: %s", err))
		return "", &HttpException{
			error: err.Error(),
		}
	}
	result := string(body)
	core.GetLogger().Info(result)
	return result, nil
}

func PostJson(url string, data string) (string, error) {
	response, err := http.Post(url,
		"application/json",
		bytes.NewBuffer([]byte(data)))
	if err != nil {
		core.GetLogger().Error(fmt.Sprintf("request err: %s", err))
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
		core.GetLogger().Error(fmt.Sprintf("response read err: %s", err))
		return "", &HttpException{
			error: err.Error(),
		}
	}
	result := string(body)
	core.GetLogger().Info(result)
	return result, nil
}
