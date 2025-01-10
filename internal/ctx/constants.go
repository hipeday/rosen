package ctx

type (
	Key string
)

const (
	LocalizeKey    Key = Key("localize")
	RequestIdKey   Key = Key("X-Request-Id")
	RedisClientKey Key = Key("RedisClient")
)

func (k Key) String() string {
	return string(k)
}
