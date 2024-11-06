package services

import (
	"TranslatorAPI/models"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func YoudaoSendRequests(word, sourceLang, targetLang string) (models.YoudaoResponse, error) {
	// 获取配置文件路径
	configPath := "./config/config.json"
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return models.YoudaoResponse{}, err
	}

	// 打开配置文件
	keyFile, err := os.Open(absPath)
	if err != nil {
		return models.YoudaoResponse{}, err
	}
	defer keyFile.Close()

	// 解析配置文件
	decoder := json.NewDecoder(keyFile)
	var config models.Config
	err = decoder.Decode(&config)
	if err != nil {
		return models.YoudaoResponse{}, err
	}

	appKey := config.Youdao.AppKey
	appSecret := config.Youdao.AppSecret

	// 检查配置是否为空
	if appKey == "" || appSecret == "" {
		return models.YoudaoResponse{}, errors.New("appKey or appSecret is empty in config file")
	}

	salt := strconv.FormatInt(time.Now().UnixNano(), 10)
	curtime := strconv.FormatInt(time.Now().Unix(), 10)

	// 签名计算
	signStr := appKey + truncate(word) + salt + curtime + appSecret
	hash := sha256.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))

	// 设置请求参数
	data := url.Values{}
	data.Set("q", word)
	data.Set("appKey", appKey)
	data.Set("salt", salt)
	data.Set("from", sourceLang)
	data.Set("to", targetLang)
	data.Set("signType", "v3")
	data.Set("curtime", curtime)
	data.Set("sign", sign)

	apiURL := "https://openapi.youdao.com/api"
	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return models.YoudaoResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.YoudaoResponse{}, errors.New("unexpected status code: " + strconv.Itoa(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.YoudaoResponse{}, err
	}
	var result models.YoudaoResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return models.YoudaoResponse{}, err
	}

	return result, nil
}

func truncate(q string) string {
	r := []rune(q)
	length := len(r)
	if length <= 20 {
		return q
	}
	return string(r[:10]) + strconv.Itoa(length) + string(r[length-10:])
}
