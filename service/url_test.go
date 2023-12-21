package service

import (
	"reflect"
	"sync"
	"testing"
	"zlink/domain"
	"zlink/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestService_AddUrl(t *testing.T) {
	log := zap.NewExample()

	memStore := domain.NewMemoryStore(make(map[string]string), new(sync.Mutex))

	mapRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectCount := domain.NewUrlCollectionCount(model.Collections{
		URLs:   make([]string, 0),
		Counts: make([]int, 0)},
		new(sync.Mutex))

	store := domain.NewStore(log, memStore, cntStore, mapRevStore, collectCount)

	type fields struct {
		log   *zap.Logger
		store domain.IStore
	}
	type args struct {
		ctx    *gin.Context
		addUrl model.AddUrl
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
				log:   log,
				store: store,
			},
			args: args{
				ctx:    &gin.Context{},
				addUrl: model.AddUrl{URL: "wwww.google.com"},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Failed1",
			fields: fields{
				log:   log,
				store: store,
			},
			args: args{
				ctx:    &gin.Context{},
				addUrl: model.AddUrl{URL: "http://localhost:8080"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:   tt.fields.log,
				store: tt.fields.store,
			}
			_, err := s.AddUrl(tt.args.ctx, tt.args.addUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_GetUrl(t *testing.T) {
	log := zap.NewExample()

	memStore := domain.NewMemoryStore(map[string]string{
		"xyz": "http://www.google.com",
	}, new(sync.Mutex))

	mapRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectCount := domain.NewUrlCollectionCount(model.Collections{
		URLs:   make([]string, 0),
		Counts: make([]int, 0)},
		new(sync.Mutex))

	store := domain.NewStore(log, memStore, cntStore, mapRevStore, collectCount)

	type fields struct {
		log   *zap.Logger
		store domain.IStore
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
			name: "Pass",
			fields: fields{
				log:   log,
				store: store,
			},
			args: args{
				path: "xyz",
			},
			want:    "http://www.google.com",
			wantErr: false,
		},
		{
			name: "Fail",
			fields: fields{
				log:   log,
				store: store,
			},
			args: args{
				path: "yz",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:   tt.fields.log,
				store: tt.fields.store,
			}
			got, err := s.GetUrl(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.GetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_MostVisitUrl(t *testing.T) {
	log := zap.NewExample()

	memStore := domain.NewMemoryStore(map[string]string{"xyz": "xyz", "abc": "abc", "sfd": "sfd"}, new(sync.Mutex))

	mapRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectCount := domain.NewUrlCollectionCount(model.Collections{URLs: []string{"xyz", "abc", "sfd"}, Counts: []int{12, 34, 56}}, new(sync.Mutex))

	store := domain.NewStore(log, memStore, cntStore, mapRevStore, collectCount)

	type fields struct {
		log   *zap.Logger
		store domain.IStore
	}
	type args struct {
		count int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]int
	}{
		{
			name: "Pass",
			fields: fields{
				log:   log,
				store: store,
			},
			args: args{count: 2},
			want: map[string]int{
				"sfd": 56,
				"abc": 34,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:   tt.fields.log,
				store: tt.fields.store,
			}
			if got := s.MostVisitUrl(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.MostVisitUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
