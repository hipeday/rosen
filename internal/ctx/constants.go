package ctx

type (
	Key string
)

const (
	LocalizeKey    Key = Key("localize")
	RequestIdKey   Key = Key("X-Request-Id")
	RedisClientKey Key = Key("RedisClient")
	OneIdKey       Key = Key("OneId")
	UsernameKey    Key = Key("Username")
)

func (k Key) String() string {
	return string(k)
}
