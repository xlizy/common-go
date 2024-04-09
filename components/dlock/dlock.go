package dlock

import (
	"context"
	commonConfig "github.com/xlizy/common-go/base/config"
	"github.com/xlizy/common-go/utils/zlog"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"google.golang.org/grpc/connectivity"
	"time"
)

type RootConfig struct {
	Cluster []string `yaml:"dlock"`
}

var _client *clientv3.Client

type LockObj struct {
	mutex *concurrency.Mutex
}

func (o LockObj) UnLock() {
	err := o.mutex.Unlock(context.TODO())
	if err != nil {
		zlog.Error("释放锁异常:{}", err.Error())
	}
}

func NewConfig() *RootConfig {
	return &RootConfig{
		make([]string, 0),
	}
}

func InitDLock(rc *RootConfig) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   rc.Cluster,
		DialTimeout: time.Second * 5})
	if err != nil {
		zlog.Error("连接Etcd异常:{}", err.Error())
		panic(err)
	} else {
		_client = client
	}
}

// Lock 加锁
// ttl锁租期，内部会自动续期，发生异常后在ttl秒后自动释放
// wait等待锁时间
func Lock(lockKey string, ttl, wait int) (bool, *LockObj) {
	if connectivity.Ready != _client.ActiveConnection().GetState() {
		zlog.Error("加锁失败:{}", "etcd连接异常")
		return false, nil
	}
	_client.ActiveConnection().GetState().String()
	lockKey = commonConfig.GetNacosCfg().AppName + "_" + lockKey
	session, err1 := concurrency.NewSession(_client, concurrency.WithTTL(ttl))
	if err1 != nil {
		zlog.Error("获取锁异常:{}", err1.Error())
	}
	defer func(session *concurrency.Session) {
		err2 := session.Close()
		if err2 != nil {
			zlog.Error("释放锁异常:{}", err1.Error())
		}
	}(session)

	// 获取指定前缀的锁对象
	mutex := concurrency.NewMutex(session, lockKey)

	var ctx context.Context
	var cancel context.CancelFunc
	var err3 error
	if wait > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*time.Duration(wait))
		defer cancel()
		err3 = mutex.Lock(ctx)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
		defer cancel()
		err3 = mutex.TryLock(ctx)
	}
	if err3 != nil {
		zlog.Error("加锁失败:{}", err3.Error())
		return false, nil
	}

	return true, &LockObj{mutex}
}
