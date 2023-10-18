package logic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httputil"
	"strconv"

	"chative-server-go/utils/crypto"

	"github.com/zeromicro/go-zero/core/logx"
)

// add new webauthn user (调用创建接口&保存密码)
// hanko host + email
func CreateWebauthnUser(host, email string) (userID, password string, err error) {
	var body = struct {
		Email string `json:"email"`
	}{email}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return
	}
	res, err := http.Post("http://"+host+"/users",
		"application/json", bytes.NewReader(bodyJson))
	if err != nil {
		return
	}
	token := res.Header.Get("X-Auth-Token")
	if token == "" {
		err = errors.New("can not get token")
		return
	}
	var resJson = struct {
		UserID  string `json:"user_id"`
		EmailID string `json:"email_id"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&resJson)
	if err != nil {
		return
	}
	// 保存密码
	// 1. 生成随机密码
	password = crypto.GenRandomString(16)
	var setPass = struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}{UserID: resJson.UserID, Password: password}
	jsonSetPass, err := json.Marshal(setPass)
	if err != nil {
		return
	}
	req, err := http.NewRequest(http.MethodPut, "http://"+host+"/password",
		bytes.NewReader(jsonSetPass))
	if err != nil {
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	res, err = http.DefaultClient.Do(req)
	if res.StatusCode >= 300 {
		dump, _ := httputil.DumpRequest(req, true)
		err = errors.New("set password failed,StatusCode:" + strconv.Itoa(res.StatusCode) + ",Dump:" + string(dump))
		return
	}
	userID = resJson.UserID
	return
}

// fetch token
func getWebauthnToken(host, userID, password string) (token string, err error) {
	var body = struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}{userID, password}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return
	}
	res, err := http.Post("http://"+host+"/password/login",
		"application/json", bytes.NewReader(bodyJson))
	if err != nil {
		return
	}
	defer res.Body.Close()
	token = res.Header.Get("X-Auth-Token")
	if token == "" {
		err = errors.New("can not get token")
		return
	}
	return
}

func WebauthnRegInit(host, userID, password string) (resBody []byte, err error) {
	token, err := getWebauthnToken(host, userID, password)
	if err != nil {
		return
	}
	req, err := http.NewRequest(http.MethodPost, "http://"+host+"/webauthn/registration/initialize",
		nil)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}

func addWebauthnCredential(host string, token string, regFinReq []byte) (resBody []byte, code int, err error) {
	req, err := http.NewRequest(http.MethodPost, "http://"+host+"/webauthn/registration/finalize",
		bytes.NewReader(regFinReq))
	if err != nil {
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	code = res.StatusCode
	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}

//
func WebauthnRegFin(logger logx.Logger, host string, userID, password string, regFinReq []byte) (code int, err error) {
	// 添加这次的
	token, err := getWebauthnToken(host, userID, password)
	if err != nil {
		logger.Errorw("WebauthnRegFin get token failed", logx.Field("err", err))
		return
	}
	resBody, httpCode, err := addWebauthnCredential(host, token, regFinReq)
	code = httpCode
	if err != nil || httpCode >= 300 {
		logger.Errorw("WebauthnRegFin add credential failed", logx.Field("err", err))
		return
	}
	var createRes = struct {
		CredentialID string `json:"credential_id"`
		UserID       string `json:"user_id"`
	}{}
	err = json.Unmarshal(resBody, &createRes)
	if err != nil {
		logger.Errorw("WebauthnRegFin unmarshal failed", logx.Field("err", err),
			logx.Field("resBody", string(resBody)), logx.Field("httpCode", httpCode))
		return
	}
	var allCredential = []struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		PublicKey  string   `json:"public_key"`
		Aaguid     string   `json:"aaguid"`
		Transports []string `json:"transports"`
		CreatedAt  string   `json:"created_at"`
	}{} // 从webauthn获取所有的
	req, err := http.NewRequest(http.MethodGet, "http://"+host+"/webauthn/credentials", nil)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&allCredential)
	if err != nil {
		return
	}
	// 删除之前的
	for _, credential := range allCredential {
		if createRes.CredentialID == credential.ID {
			continue
		}
		delReq, err2 := http.NewRequest(http.MethodDelete, "http://"+host+"/webauthn/credentials/"+credential.ID, nil)
		if err2 != nil {
			println(err2.Error())
			// return nil, err
		}
		delReq.Header.Add("Authorization", "Bearer "+token)
		res, err = http.DefaultClient.Do(delReq)
		if err != nil {
			println(err.Error())
			continue
			// return nil, err
		}
		defer res.Body.Close()
		if res.StatusCode >= 300 {
			println(res.Status)
			// return nil, errors.New("delete credential failed")
		}
	}
	return
}

func WebauthnDeleteCredentials(host string, userID, password string) (err error) {
	token, err := getWebauthnToken(host, userID, password)
	if err != nil {
		return
	}
	// 获取creadential id
	var allCredential = []struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		PublicKey  string   `json:"public_key"`
		Aaguid     string   `json:"aaguid"`
		Transports []string `json:"transports"`
		CreatedAt  string   `json:"created_at"`
	}{} // 从webauthn获取所有的
	req, err := http.NewRequest(http.MethodGet, "http://"+host+"/webauthn/credentials", nil)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&allCredential)
	if err != nil {
		return
	}
	// 删除
	for _, credential := range allCredential {
		delReq, err := http.NewRequest(http.MethodDelete, "http://"+host+"/webauthn/credentials/"+credential.ID, nil)
		if err != nil {
			return err
		}
		delReq.Header.Add("Authorization", "Bearer "+token)
		res, err := http.DefaultClient.Do(delReq)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode >= 300 {
			err = errors.New("delete credential failed")
			return err
		}
	}
	return
}

//
