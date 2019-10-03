package bandiera

import (
	"reflect"
	"testing"
)

type FakeHttpClient struct {
	Body string
	Url  string
}

func (f *FakeHttpClient) GetUrlContent(url string, _ Params) ([]byte, error) {
	f.Url = url
	return []byte(f.Body), nil
}

func TestClient_GetAll(t *testing.T) {
	type fields struct {
		httpClient FakeHttpClient
	}
	type args struct {
		params Params
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes AllResponse
	}{
		{
			"happy path",
			fields{
				FakeHttpClient{
					Body: `{
  "response": {
    "pubserv": {
      "show-article-metrics": false, 
      "show-new-search": false, 
      "show-subject-pages": true
    }, 
    "shunter": {
      "themes": false
    }
  }
}
`,
				},
			},
			args{
				map[string]string{},
			},
			AllResponse{
				Warning: "",
				GroupFlags: GroupFlags{
					"pubserv": {
						"show-article-metrics": false,
						"show-new-search":      false,
						"show-subject-pages":   true,
					},
					"shunter": {
						"themes": false,
					},
				},
			},
		},
		{
			"with warning",
			fields{
				FakeHttpClient{
					Body: `{
  "response": {
    "pubserv": {
      "show-subject-pages": true
    }, 
    "shunter": {
      "themes": false
    }
  },
  "warning": "warn"
}
`,
				},
			},
			args{
				map[string]string{},
			},
			AllResponse{
				Warning: "warn",
				GroupFlags: GroupFlags{
					"pubserv": {
						"show-subject-pages": true,
					},
					"shunter": {
						"themes": false,
					},
				},
			},
		},
		{
			"invalid response",
			fields{
				FakeHttpClient{
					Body: `this is not json`,
				},
			},
			args{
				map[string]string{},
			},
			AllResponse{},
		},
		{
			"empty response",
			fields{
				FakeHttpClient{
					Body: `{"response": {}}`,
				},
			},
			args{
				map[string]string{},
			},
			AllResponse{
				Warning:    "",
				GroupFlags: GroupFlags{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				httpClient: &tt.fields.httpClient,
			}
			if gotRes := c.GetAll(tt.args.params); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetAll() = %v, want %v", gotRes, tt.wantRes)
			}
			if tt.fields.httpClient.Url != "/api/v2/all" {
				t.Errorf("URL = %s, want %s", tt.fields.httpClient.Url, "/api/v2/all")
			}
		})
	}
}

func TestClient_GetFeaturesForGroup(t *testing.T) {
	type fields struct {
		httpClient FakeHttpClient
	}
	type args struct {
		group  string
		params Params
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes GroupResponse
	}{
		{
			"happy path",
			fields{
				httpClient: FakeHttpClient{
					Body: `{"response": {"foo": true, "bar": false}}`,
				},
			},
			args{
				group:  "foo",
				params: nil,
			},
			GroupResponse{
				Warning: "",
				Flags: Flags{
					"foo": true,
					"bar": false,
				},
			},
		},
		{
			name: "warning",
			fields: fields{
				httpClient: FakeHttpClient{
					Body: `{"warning": "warning", "response": {}}`,
				},
			},
			args: args{
				group:  "foo",
				params: nil,
			},
			wantRes: GroupResponse{
				Warning: "warning",
				Flags:   Flags{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				httpClient: &tt.fields.httpClient,
			}
			if gotRes := c.GetFeaturesForGroup(tt.args.group, tt.args.params); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetFeaturesForGroup() = %v, want %v", gotRes, tt.wantRes)
			}
			if tt.fields.httpClient.Url != "/api/v2/groups/foo/features" {
				t.Errorf("URL = %s, want %s", tt.fields.httpClient.Url, "/api/v2/all")
			}
		})
	}
}

func TestClient_IsEnabled(t *testing.T) {
	type fields struct {
		httpClient FakeHttpClient
	}
	type args struct {
		group   string
		feature string
		params  Params
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes bool
		wantUrl string
	}{
		{
			name: "feature enabled",
			fields: fields{
				httpClient: FakeHttpClient{
					Body: `{"response": true}`,
				},
			},
			args: args{
				group:   "foo",
				feature: "bar",
				params:  nil,
			},
			wantRes: true,
			wantUrl: "/api/v2/groups/foo/features/bar",
		},
		{
			name: "feature disabled",
			fields: fields{
				httpClient: FakeHttpClient{
					Body: `{"response": false}`,
				},
			},
			args: args{
				group:   "bar",
				feature: "foo",
				params:  nil,
			},
			wantRes: false,
			wantUrl: "/api/v2/groups/bar/features/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				httpClient: &tt.fields.httpClient,
			}
			if gotRes := c.IsEnabled(tt.args.group, tt.args.feature, tt.args.params); gotRes != tt.wantRes {
				t.Errorf("IsEnabled() = %v, want %v", gotRes, tt.wantRes)
			}
			if tt.fields.httpClient.Url != tt.wantUrl {
				t.Errorf("URL = %s, want %s", tt.fields.httpClient.Url, tt.wantUrl)
			}
		})
	}
}
