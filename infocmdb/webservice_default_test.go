package infocmdb

import (
	v2 "github.com/infonova/infocmdb-sdk-go/infocmdb/v2/infocmdb"
	utilTesting "github.com/infonova/infocmdb-sdk-go/util/testing"
	"reflect"
	"testing"
)

func TestInfoCMDB_CreateCi(t *testing.T) {
	infocmdbUrl := utilTesting.New().GetUrl()
	type fields struct {
		v2Config v2.Config
	}
	type args struct {
		ciTypeID  int
		icon      string
		historyId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantR   CreateCi
		wantErr bool
	}{
		{
			"v2 create CI",
			fields{
				v2.Config{
					Url:      infocmdbUrl,
					Username: "admin",
					Password: "admin",
					BasePath: "/app/",
				},
			},
			args{
				ciTypeID:  476,
				icon:      "",
				historyId: 0,
			},
			CreateCi{
				ID:        617827,
				CiTypeID:  476,
				Icon:      "",
				HistoryID: 59529024,
				ValidFrom: "2020-01-13 15:14:05",
				CreatedAt: "2020-01-13 15:14:05",
				UpdatedAt: "2020-01-13 15:14:05",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdbV2 := v2.New()
			if err := cmdbV2.LoadConfig(tt.fields.v2Config); err != nil {
				t.Fatalf("LoadConfig failed: %v\n", err)
			}
			cmdb := &Client{
				v2: cmdbV2,
			}
			gotR, err := cmdb.CreateCi(tt.args.ciTypeID, tt.args.icon, tt.args.historyId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("CreateCi() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestInfoCMDB_AddCiProjectMapping(t *testing.T) {
	infocmdbUrl := utilTesting.New().GetUrl()
	type fields struct {
		v2Config v2.Config
	}
	type args struct {
		ciID      int
		projectID int
		historyID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"v2 create CI",
			fields{
				v2.Config{
					Url:      infocmdbUrl,
					Username: "admin",
					Password: "admin",
					BasePath: "/app/",
				},
			},
			args{
				ciID:      617830,
				projectID: 33,
				historyID: 59529029,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdbV2 := v2.New()
			if err := cmdbV2.LoadConfig(tt.fields.v2Config); err != nil {
				t.Fatalf("LoadConfig failed: %v\n", err)
			}
			cmdb := &Client{
				v2: cmdbV2,
			}
			err := cmdb.AddCiProjectMapping(tt.args.ciID, tt.args.projectID, tt.args.historyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddCiProjectMapping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestInfoCMDB_GetListOfCiIdsOfCiType(t *testing.T) {
	infocmdbUrl := utilTesting.New().GetUrl()
	type fields struct {
		v2Config v2.Config
	}
	type args struct {
		ciTypeID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantR   CiIds
		wantErr bool
	}{
		{
			"v2 List Ci's pf Type '1' with wrong Credentials (fail)",
			fields{
				v2.Config{
					Url:      infocmdbUrl,
					Username: "fail",
					Password: "fail",
					BasePath: "/app/",
				},
			},
			args{ciTypeID: 1},
			nil,
			true,
		},
		{
			"v2 List Ci's of Type '1' (demo)",
			fields{
				v2.Config{
					Url:      infocmdbUrl,
					Username: "admin",
					Password: "admin",
					BasePath: "/app/",
				},
			},
			args{ciTypeID: 1},
			CiIds{1, 2},
			false,
		},
		{
			"v2 List Ci's of Type '-1' (error)",
			fields{
				v2.Config{
					Url:      infocmdbUrl,
					Username: "admin",
					Password: "admin",
					BasePath: "/app/",
				},
			},
			args{ciTypeID: -1},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdbV2 := v2.New()
			if err := cmdbV2.LoadConfig(tt.fields.v2Config); err != nil {
				t.Fatalf("LoadConfig failed: %v\n", err)
			}
			cmdb := &Client{
				v2: cmdbV2,
			}
			gotR, err := cmdb.GetListOfCiIdsOfCiType(tt.args.ciTypeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getListOfCiIdsOfCiType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("getListOfCiIdsOfCiType() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestInfoCMDB_GetListOfCiIdsOfCiTypeV2(t *testing.T) {
	infocmdbUrl := utilTesting.New().GetUrl()
	type fields struct {
		v2Config v2.Config
	}
	type args struct {
		ciTypeID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantR   CiIds
		wantErr bool
	}{
		{
			"v2 List Ci's pf Type '1' with wrong Credentials (fail)",
			fields{
				v2.Config{
					Url:      infocmdbUrl,
					Username: "false",
					Password: "false",
					BasePath: "/app/",
				},
			},
			args{ciTypeID: 1},
			nil,
			true,
		},
		{
			"v2 List CIs of Type 1 (demo)",
			fields{
				v2.Config{
					Url:      infocmdbUrl,
					Username: "admin",
					Password: "admin",
					BasePath: "/app/",
				},
			},
			args{ciTypeID: 1},
			CiIds{1, 2},
			false,
		},
		{
			"v2 List Ci's of Type '-1' (error)",
			fields{
				v2.Config{
					Url:      infocmdbUrl,
					Username: "admin",
					Password: "admin",
					BasePath: "/app/",
				},
			},
			args{ciTypeID: -1},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdbV2 := v2.New()
			if err := cmdbV2.LoadConfig(tt.fields.v2Config); err != nil {
				t.Fatalf("LoadConfig failed: %v\n", err)
			}
			cmdb := &Client{
				v2: cmdbV2,
			}
			gotR, err := cmdb.GetListOfCiIdsOfCiTypeV2(tt.args.ciTypeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getListOfCiIdsOfCiType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("getListOfCiIdsOfCiType() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestInfoCMDB_QueryWebservice(t *testing.T) {
	infocmdbUrl := utilTesting.New().GetUrl()
	type fields struct {
		v2Config v2.Config
	}
	type args struct {
		ws     string
		params map[string]string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantR   string
		wantErr bool
	}{
		{
			"v2 List CIs of Type 1 (demo)",
			fields{
				v2.Config{
					Url:      infocmdbUrl,
					Username: "admin",
					Password: "admin",
					BasePath: "/app/",
				},
			},
			args{ws: "int_getCi", params: map[string]string{"argv1": "1"}},
			`{"success":true,"message":"Query executed successfully","data":[{"ci_id":"1","ci_type_id":"1","ci_type":"demo","project":"springfield","project_id":"4"}]}`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdbV2 := v2.New()
			if err := cmdbV2.LoadConfig(tt.fields.v2Config); err != nil {
				t.Fatalf("LoadConfig failed: %v\n", err)
			}
			cmdb := &Client{
				v2: cmdbV2,
			}
			gotR, err := cmdb.QueryWebservice(tt.args.ws, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryWebservice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR != tt.wantR {
				t.Errorf("QueryWebservice() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestInfoCMDB_UpdateCiAttribute(t *testing.T) {
	infocmdbUrl := utilTesting.New().GetUrl()
	type fields struct {
		v2Config v2.Config
	}
	type args struct {
		ci int
		ua []v2.UpdateCiAttribute
	}

	cmdbConfigValid := v2.Config{
		Url:      infocmdbUrl,
		Username: "admin",
		Password: "admin",
		BasePath: "/app/",
	}
	baseCiID := 14 // ci to base this tests on
	baseAttributeName := "emp_lastname"

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"v2 Delete CI Attribute - fail - requires ciattributeid",
			fields{cmdbConfigValid},
			args{ci: baseCiID, ua: []v2.UpdateCiAttribute{
				{Mode: v2.UPDATE_MODE_DELETE, Name: baseAttributeName},
			}},
			true,
		},
		{
			"v2 Update CI Attribute",
			fields{cmdbConfigValid},
			args{ci: baseCiID, ua: []v2.UpdateCiAttribute{
				{Mode: v2.UPDATE_MODE_SET, Name: "emp_firstname", Value: "22322"},
			}},
			false,
		},
		{
			"v2 Update CI Attribute - wrong attribute name",
			fields{cmdbConfigValid},
			args{ci: baseCiID, ua: []v2.UpdateCiAttribute{
				{Mode: v2.UPDATE_MODE_SET, Name: baseAttributeName + "_NOT_EXISTING", Value: "1"},
			}},
			true,
		},
		{
			"v2 Insert CI Attribute",
			fields{cmdbConfigValid},
			args{ci: baseCiID, ua: []v2.UpdateCiAttribute{
				{Mode: v2.UPDATE_MODE_INSERT, Name: baseAttributeName, Value: "New1"},
			}},
			false,
		},
		{
			"v2 Update CI Attribute - fail - multiple attributes with the same name",
			fields{cmdbConfigValid},
			args{ci: baseCiID, ua: []v2.UpdateCiAttribute{
				{Mode: v2.UPDATE_MODE_SET, Name: baseAttributeName, Value: "22322"},
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdbV2 := v2.New()
			if err := cmdbV2.LoadConfig(tt.fields.v2Config); err != nil {
				t.Fatalf("LoadConfig failed: %v\n", err)
			}
			cmdb := &Client{
				v2: cmdbV2,
			}
			err := cmdb.UpdateCiAttribute(tt.args.ci, tt.args.ua)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCiAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestInfoCMDB_GetAttrDefaultOptionIdByAttrId(t *testing.T) {
	infocmdbUrl := utilTesting.New().GetUrl()
	type fields struct {
		v2Config v2.Config
	}
	type args struct {
		attrId      int
		optionValue string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantR   int
		wantErr bool
	}{
		{
			"v2 GetAttributeDefaultOptionId",
			fields{
				v2.Config{
					Url:      infocmdbUrl,
					Username: "admin",
					Password: "admin",
					BasePath: "/app/",
				},
			},
			args{
				attrId:      438,
				optionValue: "IN PROGRESS",
			},
			1329,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdbV2 := v2.New()
			if err := cmdbV2.LoadConfig(tt.fields.v2Config); err != nil {
				t.Fatalf("LoadConfig failed: %v\n", err)
			}
			cmdb := &Client{
				v2: cmdbV2,
			}
			gotR, err := cmdb.GetAttrDefaultOptionIdByAttrId(tt.args.attrId, tt.args.optionValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAttrDefaultOptionIdByAttrId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR != tt.wantR {
				t.Errorf("GetAttrDefaultOptionIdByAttrId() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
