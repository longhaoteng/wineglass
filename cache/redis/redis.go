package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/longhaoteng/wineglass/config"
	"github.com/longhaoteng/wineglass/consts"
)

const (
	Nil = redis.Nil
)

var (
	r *Redis
)

type Redis struct {
	client    redis.Cmdable
	keyPrefix string
	addrs     []string
	pipeline  redis.Pipeliner
}

func Init() error {
	var client redis.Cmdable
	addrs := config.Redis.Addrs
	poolSize := 512
	poolTimeout := 10 * time.Second
	idleTimeout := 10 * time.Second
	dialTimeout := 10 * time.Second
	readTimeout := 3 * time.Second
	writeTimeout := 3 * time.Second
	password := config.Redis.Password

	if len(addrs) == 1 {
		client = redis.NewClient(&redis.Options{
			DB:           config.Redis.DB,
			Addr:         addrs[0],
			PoolSize:     poolSize,
			PoolTimeout:  poolTimeout,
			IdleTimeout:  idleTimeout,
			DialTimeout:  dialTimeout,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			Password:     password,
		})
	} else {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        addrs,
			PoolSize:     poolSize,
			PoolTimeout:  poolTimeout,
			IdleTimeout:  idleTimeout,
			DialTimeout:  dialTimeout,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			Password:     password,
		})
	}

	r = &Redis{
		client:    client,
		keyPrefix: config.Redis.KeyPrefix,
		addrs:     config.Redis.Addrs,
	}

	return nil
}

func IsNil(err error) bool {
	return errors.Is(err, Nil)
}

func IsNotNil(err error) bool {
	return !errors.Is(err, Nil)
}

func Client() redis.Cmdable {
	return r.client
}

func Close() error {
	if r.pipeline != nil {
		return r.pipeline.Close()
	}
	return r.client.(redis.UniversalClient).Close()
}

func KeyPrefix() string {
	return r.keyPrefix
}

func WithPrefix(k string) string {
	return r.keyPrefix + consts.Delimiter + k
}

func WithPrefixes(keys []string) []string {
	withPrefixKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		withPrefixKeys = append(withPrefixKeys, WithPrefix(k))
	}
	return withPrefixKeys
}

type MetricsPipeline struct {
	*Redis
}

func (c *MetricsPipeline) metric(begin time.Time, cmd string) {

}

func (c *MetricsPipeline) Exec(ctx context.Context) ([]redis.Cmder, error) {
	defer func(begin time.Time) { c.metric(begin, "Exec") }(time.Now())
	return c.pipeline.Exec(ctx)
}

func GetClusterPipeline() *MetricsPipeline {
	raw := r.client.Pipeline()
	return &MetricsPipeline{
		Redis: &Redis{
			client:    raw,
			keyPrefix: r.keyPrefix,
			addrs:     r.addrs,
			pipeline:  raw,
		},
	}
}

func Command(ctx context.Context) (map[string]*redis.CommandInfo, error) {
	return r.client.Command(ctx).Result()
}

func ClientGetName(ctx context.Context) (string, error) {
	return r.client.ClientGetName(ctx).Result()
}

func Echo(ctx context.Context, message interface{}) (string, error) {
	return r.client.Echo(ctx, message).Result()
}

func Ping(ctx context.Context) (string, error) {
	return r.client.Ping(ctx).Result()
}

func Del(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Del(ctx, WithPrefixes(keys)...).Result()
}

func Unlink(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Unlink(ctx, WithPrefixes(keys)...).Result()
}

func Dump(ctx context.Context, key string) (string, error) {
	return r.client.Dump(ctx, WithPrefix(key)).Result()
}

func Exists(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Exists(ctx, WithPrefixes(keys)...).Result()
}

func Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return r.client.Expire(ctx, WithPrefix(key), expiration).Result()
}

func ExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	return r.client.ExpireAt(ctx, WithPrefix(key), tm).Result()
}

func Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.client.Keys(ctx, pattern).Result()
}

func Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) (string, error) {
	return r.client.Migrate(ctx, host, port, WithPrefix(key), db, timeout).Result()
}

func Move(ctx context.Context, key string, db int) (bool, error) {
	return r.client.Move(ctx, WithPrefix(key), db).Result()
}

func ObjectRefCount(ctx context.Context, key string) (int64, error) {
	return r.client.ObjectRefCount(ctx, WithPrefix(key)).Result()
}

func ObjectEncoding(ctx context.Context, key string) (string, error) {
	return r.client.ObjectEncoding(ctx, WithPrefix(key)).Result()
}

func ObjectIdleTime(ctx context.Context, key string) (time.Duration, error) {
	return r.client.ObjectIdleTime(ctx, WithPrefix(key)).Result()
}

func Persist(ctx context.Context, key string) (bool, error) {
	return r.client.Persist(ctx, WithPrefix(key)).Result()
}

func PExpire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return r.client.PExpire(ctx, WithPrefix(key), expiration).Result()
}

func PExpireAt(ctx context.Context, key string, tm time.Time) (bool, error) {
	return r.client.PExpireAt(ctx, WithPrefix(key), tm).Result()
}

func PTTL(ctx context.Context, key string) (time.Duration, error) {
	return r.client.PTTL(ctx, WithPrefix(key)).Result()
}

func RandomKey(ctx context.Context) (string, error) {
	return r.client.RandomKey(ctx).Result()
}

func Rename(ctx context.Context, key, newkey string) (string, error) {
	return r.client.Rename(ctx, WithPrefix(key), newkey).Result()
}

func RenameNX(ctx context.Context, key, newkey string) (bool, error) {
	return r.client.RenameNX(ctx, WithPrefix(key), newkey).Result()
}

func Restore(ctx context.Context, key string, ttl time.Duration, value string) (string, error) {
	return r.client.Restore(ctx, WithPrefix(key), ttl, value).Result()
}

func RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) (string, error) {
	return r.client.RestoreReplace(ctx, WithPrefix(key), ttl, value).Result()
}

func Sort(ctx context.Context, key string, sort *redis.Sort) ([]string, error) {
	return r.client.Sort(ctx, WithPrefix(key), sort).Result()
}

func SortStore(ctx context.Context, key, store string, sort *redis.Sort) (int64, error) {
	return r.client.SortStore(ctx, WithPrefix(key), store, sort).Result()
}

func SortInterfaces(ctx context.Context, key string, sort *redis.Sort) ([]interface{}, error) {
	return r.client.SortInterfaces(ctx, WithPrefix(key), sort).Result()
}

func Touch(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Touch(ctx, WithPrefixes(keys)...).Result()
}

func TTL(ctx context.Context, key string) (time.Duration, error) {
	return r.client.TTL(ctx, WithPrefix(key)).Result()
}

func Type(ctx context.Context, key string) (string, error) {
	return r.client.Type(ctx, WithPrefix(key)).Result()
}

func Append(ctx context.Context, key, value string) (int64, error) {
	return r.client.Append(ctx, WithPrefix(key), value).Result()
}

func Decr(ctx context.Context, key string) (int64, error) {
	return r.client.Decr(ctx, WithPrefix(key)).Result()
}

func DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	return r.client.DecrBy(ctx, WithPrefix(key), decrement).Result()
}

func Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, WithPrefix(key)).Result()
}

func GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	return r.client.GetRange(ctx, WithPrefix(key), start, end).Result()
}

func GetSet(ctx context.Context, key string, value interface{}) (string, error) {
	return r.client.GetSet(ctx, WithPrefix(key), value).Result()
}

func Incr(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, WithPrefix(key)).Result()
}

func IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return r.client.IncrBy(ctx, WithPrefix(key), value).Result()
}

func IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	return r.client.IncrByFloat(ctx, WithPrefix(key), value).Result()
}

func MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	return r.client.MGet(ctx, WithPrefixes(keys)...).Result()
}

func MSet(ctx context.Context, values ...interface{}) (string, error) {
	return r.client.MSet(ctx, values...).Result()
}

func MSetNX(ctx context.Context, values ...interface{}) (bool, error) {
	return r.client.MSetNX(ctx, values...).Result()
}

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	return r.client.Set(ctx, WithPrefix(key), value, expiration).Result()
}

func SetArgs(ctx context.Context, key string, value interface{}, a redis.SetArgs) (string, error) {
	return r.client.SetArgs(ctx, WithPrefix(key), value, a).Result()
}

func SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	return r.client.SetEX(ctx, WithPrefix(key), value, expiration).Result()
}

func SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.client.SetNX(ctx, WithPrefix(key), value, expiration).Result()
}

func SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.client.SetXX(ctx, WithPrefix(key), value, expiration).Result()
}

func SetRange(ctx context.Context, key string, offset int64, value string) (int64, error) {
	return r.client.SetRange(ctx, WithPrefix(key), offset, value).Result()
}

func StrLen(ctx context.Context, key string) (int64, error) {
	return r.client.StrLen(ctx, WithPrefix(key)).Result()
}

func GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	return r.client.GetBit(ctx, WithPrefix(key), offset).Result()
}

func SetBit(ctx context.Context, key string, offset int64, value int) (int64, error) {
	return r.client.SetBit(ctx, WithPrefix(key), offset, value).Result()
}

func BitCount(ctx context.Context, key string, bitCount *redis.BitCount) (int64, error) {
	return r.client.BitCount(ctx, WithPrefix(key), bitCount).Result()
}

func BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return r.client.BitOpAnd(ctx, destKey, WithPrefixes(keys)...).Result()
}

func BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return r.client.BitOpOr(ctx, destKey, WithPrefixes(keys)...).Result()
}

func BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error) {
	return r.client.BitOpXor(ctx, destKey, WithPrefixes(keys)...).Result()
}

func BitOpNot(ctx context.Context, destKey string, key string) (int64, error) {
	return r.client.BitOpNot(ctx, destKey, WithPrefix(key)).Result()
}

func BitPos(ctx context.Context, key string, bit int64, pos ...int64) (int64, error) {
	return r.client.BitPos(ctx, WithPrefix(key), bit, pos...).Result()
}

func BitField(ctx context.Context, key string, args ...interface{}) ([]int64, error) {
	return r.client.BitField(ctx, WithPrefix(key), args...).Result()
}

func Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return r.client.Scan(ctx, cursor, match, count).Result()
}

func ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) ([]string, uint64, error) {
	return r.client.ScanType(ctx, cursor, match, count, keyType).Result()
}

func SScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return r.client.SScan(ctx, WithPrefix(key), cursor, match, count).Result()
}

func HScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return r.client.HScan(ctx, WithPrefix(key), cursor, match, count).Result()
}

func ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return r.client.ZScan(ctx, WithPrefix(key), cursor, match, count).Result()
}

func HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return r.client.HDel(ctx, WithPrefix(key), fields...).Result()
}

func HExists(ctx context.Context, key, field string) (bool, error) {
	return r.client.HExists(ctx, WithPrefix(key), field).Result()
}

func HGet(ctx context.Context, key, field string) (string, error) {
	return r.client.HGet(ctx, WithPrefix(key), field).Result()
}

func HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, WithPrefix(key)).Result()
}

func HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	return r.client.HIncrBy(ctx, WithPrefix(key), field, incr).Result()
}

func HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error) {
	return r.client.HIncrByFloat(ctx, WithPrefix(key), field, incr).Result()
}

func HKeys(ctx context.Context, key string) ([]string, error) {
	return r.client.HKeys(ctx, WithPrefix(key)).Result()
}

func HLen(ctx context.Context, key string) (int64, error) {
	return r.client.HLen(ctx, WithPrefix(key)).Result()
}

func HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return r.client.HMGet(ctx, WithPrefix(key), fields...).Result()
}

func HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return r.client.HSet(ctx, WithPrefix(key), values...).Result()
}

func HMSet(ctx context.Context, key string, values ...interface{}) (bool, error) {
	return r.client.HMSet(ctx, WithPrefix(key), values...).Result()
}

func HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	return r.client.HSetNX(ctx, WithPrefix(key), field, value).Result()
}

func HVals(ctx context.Context, key string) ([]string, error) {
	return r.client.HVals(ctx, WithPrefix(key)).Result()
}

func BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return r.client.BLPop(ctx, timeout, WithPrefixes(keys)...).Result()
}

func BRPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	return r.client.BRPop(ctx, timeout, WithPrefixes(keys)...).Result()
}

func BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) (string, error) {
	return r.client.BRPopLPush(ctx, source, destination, timeout).Result()
}

func LIndex(ctx context.Context, key string, index int64) (string, error) {
	return r.client.LIndex(ctx, WithPrefix(key), index).Result()
}

func LInsert(ctx context.Context, key, op string, pivot, value interface{}) (int64, error) {
	return r.client.LInsert(ctx, WithPrefix(key), op, pivot, value).Result()
}

func LInsertBefore(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	return r.client.LInsertBefore(ctx, WithPrefix(key), pivot, value).Result()
}

func LInsertAfter(ctx context.Context, key string, pivot, value interface{}) (int64, error) {
	return r.client.LInsertAfter(ctx, WithPrefix(key), pivot, value).Result()
}

func LLen(ctx context.Context, key string) (int64, error) {
	return r.client.LLen(ctx, WithPrefix(key)).Result()
}

func LPop(ctx context.Context, key string) (string, error) {
	return r.client.LPop(ctx, WithPrefix(key)).Result()
}

func LPos(ctx context.Context, key string, value string, args redis.LPosArgs) (int64, error) {
	return r.client.LPos(ctx, WithPrefix(key), value, args).Result()
}

func LPosCount(ctx context.Context, key string, value string, count int64, args redis.LPosArgs) ([]int64, error) {
	return r.client.LPosCount(ctx, WithPrefix(key), value, count, args).Result()
}

func LPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return r.client.LPush(ctx, WithPrefix(key), values...).Result()
}

func LPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return r.client.LPushX(ctx, WithPrefix(key), values...).Result()
}

func LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return r.client.LRange(ctx, WithPrefix(key), start, stop).Result()
}

func LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error) {
	return r.client.LRem(ctx, WithPrefix(key), count, value).Result()
}

func LSet(ctx context.Context, key string, index int64, value interface{}) (string, error) {
	return r.client.LSet(ctx, WithPrefix(key), index, value).Result()
}

func LTrim(ctx context.Context, key string, start, stop int64) (string, error) {
	return r.client.LTrim(ctx, WithPrefix(key), start, stop).Result()
}

func RPop(ctx context.Context, key string) (string, error) {
	return r.client.RPop(ctx, WithPrefix(key)).Result()
}

func RPopLPush(ctx context.Context, source, destination string) (string, error) {
	return r.client.RPopLPush(ctx, source, destination).Result()
}

func RPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return r.client.RPush(ctx, WithPrefix(key), values...).Result()
}

func RPushX(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return r.client.RPushX(ctx, WithPrefix(key), values...).Result()
}

func SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return r.client.SAdd(ctx, WithPrefix(key), members...).Result()
}

func SCard(ctx context.Context, key string) (int64, error) {
	return r.client.SCard(ctx, WithPrefix(key)).Result()
}

func SDiff(ctx context.Context, keys ...string) ([]string, error) {
	return r.client.SDiff(ctx, WithPrefixes(keys)...).Result()
}

func SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return r.client.SDiffStore(ctx, destination, WithPrefixes(keys)...).Result()
}

func SInter(ctx context.Context, keys ...string) ([]string, error) {
	return r.client.SInter(ctx, WithPrefixes(keys)...).Result()
}

func SInterStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return r.client.SInterStore(ctx, destination, WithPrefixes(keys)...).Result()
}

func SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return r.client.SIsMember(ctx, WithPrefix(key), member).Result()
}

func SMembers(ctx context.Context, key string) ([]string, error) {
	return r.client.SMembers(ctx, WithPrefix(key)).Result()
}

func SMembersMap(ctx context.Context, key string) (map[string]struct{}, error) {
	return r.client.SMembersMap(ctx, WithPrefix(key)).Result()
}

func SMove(ctx context.Context, source, destination string, member interface{}) (bool, error) {
	return r.client.SMove(ctx, source, destination, member).Result()
}

func SPop(ctx context.Context, key string) (string, error) {
	return r.client.SPop(ctx, WithPrefix(key)).Result()
}

func SPopN(ctx context.Context, key string, count int64) ([]string, error) {
	return r.client.SPopN(ctx, WithPrefix(key), count).Result()
}

func SRandMember(ctx context.Context, key string) (string, error) {
	return r.client.SRandMember(ctx, WithPrefix(key)).Result()
}

func SRandMemberN(ctx context.Context, key string, count int64) ([]string, error) {
	return r.client.SRandMemberN(ctx, WithPrefix(key), count).Result()
}

func SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return r.client.SRem(ctx, WithPrefix(key), members...).Result()
}

func SUnion(ctx context.Context, keys ...string) ([]string, error) {
	return r.client.SUnion(ctx, WithPrefixes(keys)...).Result()
}

func SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error) {
	return r.client.SUnionStore(ctx, destination, WithPrefixes(keys)...).Result()
}

func XAdd(ctx context.Context, a *redis.XAddArgs) (string, error) {
	return r.client.XAdd(ctx, a).Result()
}

func XDel(ctx context.Context, stream string, ids ...string) (int64, error) {
	return r.client.XDel(ctx, stream, ids...).Result()
}

func XLen(ctx context.Context, stream string) (int64, error) {
	return r.client.XLen(ctx, stream).Result()
}

func XRange(ctx context.Context, stream, start, stop string) ([]redis.XMessage, error) {
	return r.client.XRange(ctx, stream, start, stop).Result()
}

func XRangeN(ctx context.Context, stream, start, stop string, count int64) ([]redis.XMessage, error) {
	return r.client.XRangeN(ctx, stream, start, stop, count).Result()
}

func XRevRange(ctx context.Context, stream string, start, stop string) ([]redis.XMessage, error) {
	return r.client.XRevRange(ctx, stream, start, stop).Result()
}

func XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) ([]redis.XMessage, error) {
	return r.client.XRevRangeN(ctx, stream, start, stop, count).Result()
}

func XRead(ctx context.Context, a *redis.XReadArgs) ([]redis.XStream, error) {
	return r.client.XRead(ctx, a).Result()
}

func XReadStreams(ctx context.Context, streams ...string) ([]redis.XStream, error) {
	return r.client.XReadStreams(ctx, streams...).Result()
}

func XGroupCreate(ctx context.Context, stream, group, start string) (string, error) {
	return r.client.XGroupCreate(ctx, stream, group, start).Result()
}

func XGroupCreateMkStream(ctx context.Context, stream, group, start string) (string, error) {
	return r.client.XGroupCreateMkStream(ctx, stream, group, start).Result()
}

func XGroupSetID(ctx context.Context, stream, group, start string) (string, error) {
	return r.client.XGroupSetID(ctx, stream, group, start).Result()
}

func XGroupDestroy(ctx context.Context, stream, group string) (int64, error) {
	return r.client.XGroupDestroy(ctx, stream, group).Result()
}

func XGroupDelConsumer(ctx context.Context, stream, group, consumer string) (int64, error) {
	return r.client.XGroupDelConsumer(ctx, stream, group, consumer).Result()
}

func XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) ([]redis.XStream, error) {
	return r.client.XReadGroup(ctx, a).Result()
}

func XAck(ctx context.Context, stream, group string, ids ...string) (int64, error) {
	return r.client.XAck(ctx, stream, group, ids...).Result()
}

func XPending(ctx context.Context, stream, group string) (*redis.XPending, error) {
	return r.client.XPending(ctx, stream, group).Result()
}

func XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) ([]redis.XPendingExt, error) {
	return r.client.XPendingExt(ctx, a).Result()
}

func XClaim(ctx context.Context, a *redis.XClaimArgs) ([]redis.XMessage, error) {
	return r.client.XClaim(ctx, a).Result()
}

func XClaimJustID(ctx context.Context, a *redis.XClaimArgs) ([]string, error) {
	return r.client.XClaimJustID(ctx, a).Result()
}

func XTrim(ctx context.Context, key string, maxLen int64) (int64, error) {
	return r.client.XTrim(ctx, WithPrefix(key), maxLen).Result()
}

func XTrimApprox(ctx context.Context, key string, maxLen int64) (int64, error) {
	return r.client.XTrimApprox(ctx, WithPrefix(key), maxLen).Result()
}

func XInfoGroups(ctx context.Context, key string) ([]redis.XInfoGroup, error) {
	return r.client.XInfoGroups(ctx, WithPrefix(key)).Result()
}

func XInfoStream(ctx context.Context, key string) (*redis.XInfoStream, error) {
	return r.client.XInfoStream(ctx, WithPrefix(key)).Result()
}

func BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) (*redis.ZWithKey, error) {
	return r.client.BZPopMax(ctx, timeout, WithPrefixes(keys)...).Result()
}

func BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) (*redis.ZWithKey, error) {
	return r.client.BZPopMin(ctx, timeout, WithPrefixes(keys)...).Result()
}

func ZAdd(ctx context.Context, key string, members ...*redis.Z) (int64, error) {
	return r.client.ZAdd(ctx, WithPrefix(key), members...).Result()
}

func ZAddNX(ctx context.Context, key string, members ...*redis.Z) (int64, error) {
	return r.client.ZAddNX(ctx, WithPrefix(key), members...).Result()
}

func ZAddXX(ctx context.Context, key string, members ...*redis.Z) (int64, error) {
	return r.client.ZAddXX(ctx, WithPrefix(key), members...).Result()
}

func ZAddCh(ctx context.Context, key string, members ...*redis.Z) (int64, error) {
	return r.client.ZAddCh(ctx, WithPrefix(key), members...).Result()
}

func ZAddNXCh(ctx context.Context, key string, members ...*redis.Z) (int64, error) {
	return r.client.ZAddNXCh(ctx, WithPrefix(key), members...).Result()
}

func ZAddXXCh(ctx context.Context, key string, members ...*redis.Z) (int64, error) {
	return r.client.ZAddXXCh(ctx, WithPrefix(key), members...).Result()
}

func ZIncr(ctx context.Context, key string, member *redis.Z) (float64, error) {
	return r.client.ZIncr(ctx, WithPrefix(key), member).Result()
}

func ZIncrNX(ctx context.Context, key string, member *redis.Z) (float64, error) {
	return r.client.ZIncrNX(ctx, WithPrefix(key), member).Result()
}

func ZIncrXX(ctx context.Context, key string, member *redis.Z) (float64, error) {
	return r.client.ZIncrXX(ctx, WithPrefix(key), member).Result()
}

func ZCard(ctx context.Context, key string) (int64, error) {
	return r.client.ZCard(ctx, WithPrefix(key)).Result()
}

func ZCount(ctx context.Context, key, min, max string) (int64, error) {
	return r.client.ZCount(ctx, WithPrefix(key), min, max).Result()
}

func ZLexCount(ctx context.Context, key, min, max string) (int64, error) {
	return r.client.ZLexCount(ctx, WithPrefix(key), min, max).Result()
}

func ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	return r.client.ZIncrBy(ctx, WithPrefix(key), increment, member).Result()
}

func ZInterStore(ctx context.Context, destination string, store *redis.ZStore) (int64, error) {
	return r.client.ZInterStore(ctx, destination, store).Result()
}

func ZMScore(ctx context.Context, key string, members ...string) ([]float64, error) {
	return r.client.ZMScore(ctx, WithPrefix(key), members...).Result()
}

func ZPopMax(ctx context.Context, key string, count ...int64) ([]redis.Z, error) {
	return r.client.ZPopMax(ctx, WithPrefix(key), count...).Result()
}

func ZPopMin(ctx context.Context, key string, count ...int64) ([]redis.Z, error) {
	return r.client.ZPopMin(ctx, WithPrefix(key), count...).Result()
}

func ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return r.client.ZRange(ctx, WithPrefix(key), start, stop).Result()
}

func ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	return r.client.ZRangeWithScores(ctx, WithPrefix(key), start, stop).Result()
}

func ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	return r.client.ZRangeByScore(ctx, WithPrefix(key), opt).Result()
}

func ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	return r.client.ZRangeByLex(ctx, WithPrefix(key), opt).Result()
}

func ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return r.client.ZRangeByScoreWithScores(ctx, WithPrefix(key), opt).Result()
}

func ZRank(ctx context.Context, key, member string) (int64, error) {
	return r.client.ZRank(ctx, WithPrefix(key), member).Result()
}

func ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	return r.client.ZRem(ctx, WithPrefix(key), members...).Result()
}

func ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	return r.client.ZRemRangeByRank(ctx, WithPrefix(key), start, stop).Result()
}

func ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error) {
	return r.client.ZRemRangeByScore(ctx, WithPrefix(key), min, max).Result()
}

func ZRemRangeByLex(ctx context.Context, key, min, max string) (int64, error) {
	return r.client.ZRemRangeByLex(ctx, WithPrefix(key), min, max).Result()
}

func ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return r.client.ZRevRange(ctx, WithPrefix(key), start, stop).Result()
}

func ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	return r.client.ZRevRangeWithScores(ctx, WithPrefix(key), start, stop).Result()
}

func ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	return r.client.ZRevRangeByScore(ctx, WithPrefix(key), opt).Result()
}

func ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) ([]string, error) {
	return r.client.ZRevRangeByLex(ctx, WithPrefix(key), opt).Result()
}

func ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return r.client.ZRevRangeByScoreWithScores(ctx, WithPrefix(key), opt).Result()
}

func ZRevRank(ctx context.Context, key, member string) (int64, error) {
	return r.client.ZRevRank(ctx, WithPrefix(key), member).Result()
}

func ZScore(ctx context.Context, key, member string) (float64, error) {
	return r.client.ZScore(ctx, WithPrefix(key), member).Result()
}

func ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) (int64, error) {
	return r.client.ZUnionStore(ctx, dest, store).Result()
}

func PFAdd(ctx context.Context, key string, els ...interface{}) (int64, error) {
	return r.client.PFAdd(ctx, WithPrefix(key), els...).Result()
}

func PFCount(ctx context.Context, keys ...string) (int64, error) {
	return r.client.PFCount(ctx, WithPrefixes(keys)...).Result()
}

func PFMerge(ctx context.Context, dest string, keys ...string) (string, error) {
	return r.client.PFMerge(ctx, dest, WithPrefixes(keys)...).Result()
}

func BgRewriteAOF(ctx context.Context) (string, error) {
	return r.client.BgRewriteAOF(ctx).Result()
}

func BgSave(ctx context.Context) (string, error) {
	return r.client.BgSave(ctx).Result()
}

func ClientKill(ctx context.Context, ipPort string) (string, error) {
	return r.client.ClientKill(ctx, ipPort).Result()
}

func ClientKillByFilter(ctx context.Context, keys ...string) (int64, error) {
	return r.client.ClientKillByFilter(ctx, WithPrefixes(keys)...).Result()
}

func ClientList(ctx context.Context) (string, error) {
	return r.client.ClientList(ctx).Result()
}

func ClientPause(ctx context.Context, dur time.Duration) (bool, error) {
	return r.client.ClientPause(ctx, dur).Result()
}

func ClientID(ctx context.Context) (int64, error) {
	return r.client.ClientID(ctx).Result()
}

func ConfigGet(ctx context.Context, parameter string) ([]interface{}, error) {
	return r.client.ConfigGet(ctx, parameter).Result()
}

func ConfigResetStat(ctx context.Context) (string, error) {
	return r.client.ConfigResetStat(ctx).Result()
}

func ConfigSet(ctx context.Context, parameter, value string) (string, error) {
	return r.client.ConfigSet(ctx, parameter, value).Result()
}

func ConfigRewrite(ctx context.Context) (string, error) {
	return r.client.ConfigRewrite(ctx).Result()
}

func DBSize(ctx context.Context) (int64, error) {
	return r.client.DBSize(ctx).Result()
}

func FlushAll(ctx context.Context) (string, error) {
	return r.client.FlushAll(ctx).Result()
}

func FlushAllAsync(ctx context.Context) (string, error) {
	return r.client.FlushAllAsync(ctx).Result()
}

func FlushDB(ctx context.Context) (string, error) {
	return r.client.FlushDB(ctx).Result()
}

func FlushDBAsync(ctx context.Context) (string, error) {
	return r.client.FlushDBAsync(ctx).Result()
}

func Info(ctx context.Context, section ...string) (string, error) {
	return r.client.Info(ctx).Result()
}

func LastSave(ctx context.Context) (int64, error) {
	return r.client.LastSave(ctx).Result()
}

func Save(ctx context.Context) (string, error) {
	return r.client.Save(ctx).Result()
}

func Shutdown(ctx context.Context) (string, error) {
	return r.client.Shutdown(ctx).Result()
}

func ShutdownSave(ctx context.Context) (string, error) {
	return r.client.ShutdownSave(ctx).Result()
}

func ShutdownNoSave(ctx context.Context) (string, error) {
	return r.client.ShutdownNoSave(ctx).Result()
}

func SlaveOf(ctx context.Context, host, port string) (string, error) {
	return r.client.SlaveOf(ctx, host, port).Result()
}

func Time(ctx context.Context) (time.Time, error) {
	return r.client.Time(ctx).Result()
}

func DebugObject(ctx context.Context, key string) (string, error) {
	return r.client.DebugObject(ctx, WithPrefix(key)).Result()
}

func ReadOnly(ctx context.Context) (string, error) {
	return r.client.ReadOnly(ctx).Result()
}

func ReadWrite(ctx context.Context) (string, error) {
	return r.client.ReadWrite(ctx).Result()
}

func MemoryUsage(ctx context.Context, key string, samples ...int) (int64, error) {
	return r.client.MemoryUsage(ctx, WithPrefix(key), samples...).Result()
}

func Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	return r.client.Eval(ctx, script, WithPrefixes(keys), args...).Result()
}

func EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return r.client.EvalSha(ctx, sha1, WithPrefixes(keys), args...).Result()
}

func ScriptExists(ctx context.Context, hashes ...string) ([]bool, error) {
	return r.client.ScriptExists(ctx, hashes...).Result()
}

func ScriptFlush(ctx context.Context) (string, error) {
	return r.client.ScriptFlush(ctx).Result()
}

func ScriptKill(ctx context.Context) (string, error) {
	return r.client.ScriptKill(ctx).Result()
}

func ScriptLoad(ctx context.Context, script string) (string, error) {
	return r.client.ScriptLoad(ctx, script).Result()
}

func Publish(ctx context.Context, channel string, message interface{}) (int64, error) {
	return r.client.Publish(ctx, channel, message).Result()
}

func PubSubChannels(ctx context.Context, pattern string) ([]string, error) {
	return r.client.PubSubChannels(ctx, pattern).Result()
}

func PubSubNumSub(ctx context.Context, channels ...string) (map[string]int64, error) {
	return r.client.PubSubNumSub(ctx, channels...).Result()
}

func PubSubNumPat(ctx context.Context) (int64, error) {
	return r.client.PubSubNumPat(ctx).Result()
}

func ClusterSlots(ctx context.Context) ([]redis.ClusterSlot, error) {
	return r.client.ClusterSlots(ctx).Result()
}

func ClusterNodes(ctx context.Context) (string, error) {
	return r.client.ClusterNodes(ctx).Result()
}

func ClusterMeet(ctx context.Context, host, port string) (string, error) {
	return r.client.ClusterMeet(ctx, host, port).Result()
}

func ClusterForget(ctx context.Context, nodeID string) (string, error) {
	return r.client.ClusterForget(ctx, nodeID).Result()
}

func ClusterReplicate(ctx context.Context, nodeID string) (string, error) {
	return r.client.ClusterReplicate(ctx, nodeID).Result()
}

func ClusterResetSoft(ctx context.Context) (string, error) {
	return r.client.ClusterResetSoft(ctx).Result()
}

func ClusterResetHard(ctx context.Context) (string, error) {
	return r.client.ClusterResetHard(ctx).Result()
}

func ClusterInfo(ctx context.Context) (string, error) {
	return r.client.ClusterInfo(ctx).Result()
}

func ClusterKeySlot(ctx context.Context, key string) (int64, error) {
	return r.client.ClusterKeySlot(ctx, WithPrefix(key)).Result()
}

func ClusterGetKeysInSlot(ctx context.Context, slot int, count int) ([]string, error) {
	return r.client.ClusterGetKeysInSlot(ctx, slot, count).Result()
}

func ClusterCountFailureReports(ctx context.Context, nodeID string) (int64, error) {
	return r.client.ClusterCountFailureReports(ctx, nodeID).Result()
}

func ClusterCountKeysInSlot(ctx context.Context, slot int) (int64, error) {
	return r.client.ClusterCountKeysInSlot(ctx, slot).Result()
}

func ClusterDelSlots(ctx context.Context, slots ...int) (string, error) {
	return r.client.ClusterDelSlots(ctx, slots...).Result()
}

func ClusterDelSlotsRange(ctx context.Context, min, max int) (string, error) {
	return r.client.ClusterDelSlotsRange(ctx, min, max).Result()
}

func ClusterSaveConfig(ctx context.Context) (string, error) {
	return r.client.ClusterSaveConfig(ctx).Result()
}

func ClusterSlaves(ctx context.Context, nodeID string) ([]string, error) {
	return r.client.ClusterSlaves(ctx, nodeID).Result()
}

func ClusterFailover(ctx context.Context) (string, error) {
	return r.client.ClusterFailover(ctx).Result()
}

func ClusterAddSlots(ctx context.Context, slots ...int) (string, error) {
	return r.client.ClusterAddSlots(ctx, slots...).Result()
}

func ClusterAddSlotsRange(ctx context.Context, min, max int) (string, error) {
	return r.client.ClusterAddSlotsRange(ctx, min, max).Result()
}

func GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) (int64, error) {
	return r.client.GeoAdd(ctx, WithPrefix(key), geoLocation...).Result()
}

func GeoPos(ctx context.Context, key string, members ...string) ([]*redis.GeoPos, error) {
	return r.client.GeoPos(ctx, WithPrefix(key), members...).Result()
}

func GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) ([]redis.GeoLocation, error) {
	return r.client.GeoRadius(ctx, WithPrefix(key), longitude, latitude, query).Result()
}

func GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) (int64, error) {
	return r.client.GeoRadiusStore(ctx, WithPrefix(key), longitude, latitude, query).Result()
}

func GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) ([]redis.GeoLocation, error) {
	return r.client.GeoRadiusByMember(ctx, WithPrefix(key), member, query).Result()
}

func GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) (int64, error) {
	return r.client.GeoRadiusByMemberStore(ctx, WithPrefix(key), member, query).Result()
}

func GeoDist(ctx context.Context, key string, member1, member2, unit string) (float64, error) {
	return r.client.GeoDist(ctx, WithPrefix(key), member1, member2, unit).Result()
}

func GeoHash(ctx context.Context, key string, members ...string) ([]string, error) {
	return r.client.GeoHash(ctx, WithPrefix(key), members...).Result()
}
