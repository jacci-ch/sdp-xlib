package cfgv

type ValueGetter interface {
	GetValue(name, key string) (*Value, bool)

	ToInt64(name, key string, dst *int64, defaultValue int64) error
	ToInt32(name, key string, dst *int32, defaultValue int32) error
	ToInt(name, key string, dst *int, defaultValue int) error
	ToBool(name, key string, dst *bool, defaultValue bool) error
	ToStr(name, key string, dst *string, defaultValue string) error
	ToStrArray(name, key string, dst *[]string, defaultValue []string) error
}
