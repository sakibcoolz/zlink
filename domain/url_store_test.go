package domain

import (
	"net/http"
	"sync"
	"testing"
	"zlink/log"
	"zlink/model"

	"github.com/gin-gonic/gin"
)

func TestStore_UrlStore(t *testing.T) {
	memStore := NewMemoryStore(make(map[string]string), new(sync.Mutex))
	mapRevStore := NewMappingRev(make(map[string]string), new(sync.Mutex))
	cntStore := NewCountStore(0, new(sync.Mutex))

	type fields struct {
		log *log.Log
		ms  *model.MemoryStore
		sc  *model.CountStore
		mr  *model.MappingRev
	}
	type args struct {
		data map[string]string
		ctx  *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Pass1",
			fields: fields{
				log: log.New(),
				ms:  memStore,
				sc:  cntStore,
				mr:  mapRevStore,
			},
			args: args{
				data: map[string]string{"skm": "http://www.google.com"},
				ctx: &gin.Context{
					Request: &http.Request{
						Header: http.Header{
							"skm": []string{""},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				log: tt.fields.log,
				ms:  tt.fields.ms,
				sc:  tt.fields.sc,
				mr:  tt.fields.mr,
			}
			s.UrlStore(tt.args.ctx, tt.args.data)
		})
	}
}

func TestStore_GetUrl(t *testing.T) {
	memStore := NewMemoryStore(map[string]string{"skm": "google.com"}, new(sync.Mutex))
	mapRevStore := NewMappingRev(make(map[string]string), new(sync.Mutex))
	cntStore := NewCountStore(0, new(sync.Mutex))

	type fields struct {
		log *log.Log
		ms  *model.MemoryStore
		sc  *model.CountStore
		mr  *model.MappingRev
	}
	type args struct {
		path string
		ctx  *gin.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Pass1",
			fields: fields{
				log: log.New(),
				ms:  memStore,
				sc:  cntStore,
				mr:  mapRevStore,
			},
			args: args{
				path: "skm",
				ctx: &gin.Context{
					Request: &http.Request{
						Header: http.Header{
							"skm": []string{""},
						},
					},
				},
			},
			want:    "google.com",
			wantErr: false,
		},
		{
			name: "Failed",
			fields: fields{
				log: log.New(),
				ms:  memStore,
				sc:  cntStore,
				mr:  mapRevStore,
			},
			args: args{
				path: "sk",
				ctx: &gin.Context{
					Request: &http.Request{
						Header: http.Header{
							"skm": []string{""},
						},
					},
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				log: tt.fields.log,
				ms:  tt.fields.ms,
				sc:  tt.fields.sc,
				mr:  tt.fields.mr,
			}
			got, err := s.GetUrl(tt.args.ctx, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.GetUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Store.GetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_GetUrlMapping(t *testing.T) {
	mapRevStore := NewMappingRev(map[string]string{"google.com": "xyz"}, new(sync.Mutex))
	type fields struct {
		mr *model.MappingRev
	}
	type args struct {
		url string
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Pass1",
			fields: fields{
				mr: mapRevStore,
			},
			args: args{
				url: "google.com",
				ctx: &gin.Context{
					Request: &http.Request{
						Header: http.Header{
							"skm": []string{""},
						},
					},
				},
			},
			want: "xyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				mr: tt.fields.mr,
			}
			if got := s.GetUrlMapping(tt.args.ctx, tt.args.url); got != tt.want {
				t.Errorf("Store.GetUrlMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
