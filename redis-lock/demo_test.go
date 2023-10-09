package redis_lock

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golearnpro/mocks"
	"testing"
	"time"
)

func TestClient_TryLock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testClass := []struct {
		name       string
		mock       func() redis.Cmdable
		key        string
		expiration time.Duration
		wantErr    error
		wantLock   *Lock
	}{
		{
			name:       "locked",
			key:        "locked-key",
			expiration: time.Minute,
			mock: func() redis.Cmdable {
				rdb := mocks.NewMockCmdable(ctrl)
				res := redis.NewBoolResult(true, nil)
				rdb.EXPECT().SetNX(gomock.Any(), "locked_key", gomock.Any(), time.Minute).Return(res)
				return rdb
			},
			wantLock: &Lock{
				key: "locked-key",
				//expiration: time.Minute,
			},
		},
		{
			name:       "network error",
			key:        "network-key",
			expiration: time.Minute,
			mock: func() redis.Cmdable {
				rdb := mocks.NewMockCmdable(ctrl)
				res := redis.NewBoolResult(false, errors.New("network error"))
				rdb.EXPECT().SetNX(gomock.Any(), "network-key", gomock.Any(), time.Minute).Return(res)
				return rdb
			},
			wantErr: errors.New("network error"),
		},
		{
			name:       "failed to lock",
			key:        "failed-key",
			expiration: time.Minute,
			mock: func() redis.Cmdable {
				rdb := mocks.NewMockCmdable(ctrl)
				res := redis.NewBoolResult(false, nil)
				rdb.EXPECT().SetNX(gomock.Any(), "failed-key", gomock.Any(), time.Minute).Return(res)
				return rdb
			},
			wantErr: ErrFailedToPreemptLock,
		},
	}
	for _, tc := range testClass {
		t.Run(tc.name, func(t *testing.T) {
			//var c Client
			c := NewClient(tc.mock())
			l, err := c.TryLock(context.Background(), tc.key, tc.expiration)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, l.client, c.client)
			assert.Equal(t, tc.wantLock.key, l.key)
			assert.NotEmpty(t, l.value)
		})
	}
}
