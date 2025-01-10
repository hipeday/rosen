package ctx

type (
	Key string
)

const (
	LocalizeKey  Key = Key("localize")
	RequestIdKey Key = Key("X-Request-Id")
)

func (k Key) String() string {
	return string(k)
}
