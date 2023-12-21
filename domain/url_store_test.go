package domain

import (
	"sync"
	"testing"
	"zlink/model"

	"go.uber.org/zap"
)

func TestStore_UrlStore(t *testing.T) {
	memStore := NewMemoryStore(make(map[string]string), new(sync.Mutex))
	mapRevStore := NewMappingRev(make(map[string]string), new(sync.Mutex))
	cntStore := NewCountStore(0, new(sync.Mutex))

	type fields struct {
		log *zap.Logger
		ms  *model.MemoryStore
		sc  *model.CountStore
		mr  *model.MappingRev
	}
	type args struct {
		data map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Pass1",
			fields: fields{
				log: &zap.Logger{},
				ms:  memStore,
				sc:  cntStore,
				mr:  mapRevStore,
			},
			args: args{
				data: map[string]string{"skm": "http://www.google.com"},
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
			s.UrlStore(tt.args.data)
		})
	}
}

func TestStore_GetUrl(t *testing.T) {
	memStore := NewMemoryStore(map[string]string{"skm": "google.com"}, new(sync.Mutex))
	mapRevStore := NewMappingRev(make(map[string]string), new(sync.Mutex))
	cntStore := NewCountStore(0, new(sync.Mutex))

	type fields struct {
		log *zap.Logger
		ms  *model.MemoryStore
		sc  *model.CountStore
		mr  *model.MappingRev
	}
	type args struct {
		path string
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
				log: zap.NewExample(),
				ms:  memStore,
				sc:  cntStore,
				mr:  mapRevStore,
			},
			args: args{
				path: "skm",
			},
			want:    "google.com",
			wantErr: false,
		},
		{
			name: "Failed",
			fields: fields{
				log: zap.NewExample(),
				ms:  memStore,
				sc:  cntStore,
				mr:  mapRevStore,
			},
			args: args{
				path: "sk",
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
			got, err := s.GetUrl(tt.args.path)
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
			},
			want: "xyz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				mr: tt.fields.mr,
			}
			if got := s.GetUrlMapping(tt.args.url); got != tt.want {
				t.Errorf("Store.GetUrlMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
