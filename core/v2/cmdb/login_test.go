package cmdb

import (
	"git.appteam.infonova.cloud/infocmdb/library/core/v2/cmdb/client"
	"testing"

	"github.com/patrickmn/go-cache"
)

func TestInfoCMDB_LoginWithUserPass(t *testing.T) {
	type fields struct {
		Config Config
		Cache  *cache.Cache
		Client *client.Client
	}
	type args struct {
		url      string
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"valid login admin//admin",
			fields{Config: Config{Url: "http://localhost", Username: "admin", Password: "admin"}},
			args{},
			false,
		},
		{
			"invalid login admin//noadmin",
			fields{Config: Config{Url: "http://localhost", Username: "admin", Password: "noadmin"}},
			args{},
			true,
		},
		{
			"invalid no username",
			fields{Config: Config{Url: "http://localhost", Username: "", Password: "noadmin"}},
			args{},
			true,
		},
		{
			"invalid no password",
			fields{Config: Config{Url: "http://localhost", Username: "admin", Password: ""}},
			args{},
			true,
		},
		{
			"invalid no data",
			fields{Config: Config{Url: "", Username: "", Password: ""}},
			args{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := new(InfoCMDB)
			i.Config = tt.fields.Config
			i.Cache = tt.fields.Cache
			i.Client = client.NewClient(i.Config.Url)

			if err := i.Login(); (err != nil) != tt.wantErr {
				t.Errorf("InfoCMDB.LoginWithUserPass() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
