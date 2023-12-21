package service

import (
	"reflect"
	"sync"
	"testing"
	"zlink/domain"
	"zlink/model"

	"go.uber.org/zap"
)

func TestNewService(t *testing.T) {
	log := zap.NewExample()

	memStore := domain.NewMemoryStore(make(map[string]string), new(sync.Mutex))

	mapRevStore := domain.NewMappingRev(make(map[string]string), new(sync.Mutex))

	cntStore := domain.NewCountStore(0, new(sync.Mutex))

	collectCount := domain.NewUrlCollectionCount(model.Collections{
		URLs:   make([]string, 0),
		Counts: make([]int, 0)},
		new(sync.Mutex))

	store := domain.NewStore(log, memStore, cntStore, mapRevStore, collectCount)

	type args struct {
		logger *zap.Logger
		store  domain.IStore
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "Pass 1",
			args: args{
				logger: log,
				store:  store,
			},
			want: &Service{
				log:   log,
				store: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.logger, tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}
