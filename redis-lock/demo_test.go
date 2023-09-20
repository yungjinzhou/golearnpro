package redis_lock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golearnpro/mocks"
	"testing"
	"time"
)

func TestClient_TryLock(t *testing.T) {
	var testClass []struct {
		name       string
		mock func() redis.Cmdable
		key        string
		expiration time.Duration
		wantErr    error
		wantLock   *Lock
	}{
		{
			name: "locked",
			key: "locked_key",
			expiration: time.Minute,
			mock: func() redis.Cmdable {
				rdb := mocks.NewMockCmdable(ctrl)
				rdb.EXPECT().SetNX(gomock.Any(), "locked_key", gomock.Any(), time.Minute).Return()
				return rdb
		}
		},
	},
	for _, tc := range testClass {
		c := NewClient(tc.mock())
		t.Run(tc.name, func(t *testing.T) {
			var c Client
			l, err := c.TryLock(context.Background(), tc.key, tc.expiration)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, l.client, c.client)
			assert.Equal(t, tc.key, l.key)
			assert.NotEmpty(t, l.value)
		})
	}
}
