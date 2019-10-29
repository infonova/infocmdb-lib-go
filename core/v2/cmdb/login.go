package cmdb

import (
	"github.com/infonova/infocmdb-lib-go/core/v1/cmdb"
	"github.com/infonova/infocmdb-lib-go/core/v2/cmdb/client"
	log "github.com/sirupsen/logrus"
)

type LoginTokenReturn struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func (i *Cmdb) Login() (err error) {

	if i.Config.ApiKey != "" {
		log.Debug("already logged in")
		return nil
	}
	if i.Config.Username == "" || i.Config.Password == "" {
		return cmdb.ErrNoCredentials
	}

	var loginResult LoginTokenReturn
	params := map[string]string{
		"username": i.Config.Username,
		"password": i.Config.Password,
		"lifetime": "600",
	}

	i.Client = client.NewClient(i.Config.Url)

	var errResp client.ResponseError
	resp, err := i.Client.NewRequest().
		SetError(&errResp).
		SetResult(&loginResult).
		SetFormData(params).
		Post("/apiV2/auth/token")

	if err != nil {
		return err
	}

	if resp != nil && resp.IsError() {
		return errResp
	}

	if loginResult.Data.Token == "" {
		return cmdb.ErrLoginFailed
	}

	i.Config.ApiKey = loginResult.Data.Token
	i.Client.SetAuthToken(i.Config.ApiKey)
	return
}
