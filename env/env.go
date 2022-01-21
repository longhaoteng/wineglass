package env

import (
	"os"
	"strings"

	"github.com/longhaoteng/wineglass/conv"
)

func Lookup(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

func GetString(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

func GetStrings(key string) []string {
	val := os.Getenv(key)
	return strings.Split(val, ",")
}

func GetInt64(key string, def int64) (int64, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseInt64(val)
}

func GetInt64s(key string) ([]int64, error) {
	val := os.Getenv(key)
	return conv.ParseInt64s(strings.Split(val, ","))
}

func GetInt32(key string, def int32) (int32, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseInt32(val)
}

func GetInt32s(key string) ([]int32, error) {
	val := os.Getenv(key)
	return conv.ParseInt32s(strings.Split(val, ","))
}

func GetInt16(key string, def int16) (int16, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseInt16(val)
}

func GetInt16s(key string) ([]int16, error) {
	val := os.Getenv(key)
	return conv.ParseInt16s(strings.Split(val, ","))
}

func GetInt8(key string, def int8) (int8, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseInt8(val)
}

func GetInt8s(key string) ([]int8, error) {
	val := os.Getenv(key)
	return conv.ParseInt8s(strings.Split(val, ","))
}

func GetInt(key string, def int) (int, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseInt(val)
}

func GetInts(key string) ([]int, error) {
	val := os.Getenv(key)
	return conv.ParseInts(strings.Split(val, ","))
}

func GetBool(key string, def bool) (bool, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseBool(val)
}

func GetBools(key string) ([]bool, error) {
	val := os.Getenv(key)
	return conv.ParseBools(strings.Split(val, ","))
}

func GetFloat32(key string, def float32) (float32, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseFloat32(val)
}

func GetFloat32s(key string) ([]float32, error) {
	val := os.Getenv(key)
	return conv.ParseFloat32s(strings.Split(val, ","))
}

func GetFloat64(key string, def float64) (float64, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseFloat64(val)
}

func GetFloat64s(key string) ([]float64, error) {
	val := os.Getenv(key)
	return conv.ParseFloat64s(strings.Split(val, ","))
}

func GetUint(key string, def uint) (uint, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseUint(val)
}

func GetUints(key string) ([]uint, error) {
	val := os.Getenv(key)
	return conv.ParseUints(strings.Split(val, ","))
}

func GetUint8(key string, def uint8) (uint8, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseUint8(val)
}

func GetUint8s(key string) ([]uint8, error) {
	val := os.Getenv(key)
	return conv.ParseUint8s(strings.Split(val, ","))
}

func GetUint16(key string, def uint16) (uint16, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseUint16(val)
}

func GetUint16s(key string) ([]uint16, error) {
	val := os.Getenv(key)
	return conv.ParseUint16s(strings.Split(val, ","))
}

func GetUint32(key string, def uint32) (uint32, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseUint32(val)
}

func GetUint32s(key string) ([]uint32, error) {
	val := os.Getenv(key)
	return conv.ParseUint32s(strings.Split(val, ","))
}

func GetUint64(key string, def uint64) (uint64, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def, nil
	}
	return conv.ParseUint64(val)
}

func GetUint64s(key string) ([]uint64, error) {
	val := os.Getenv(key)
	return conv.ParseUint64s(strings.Split(val, ","))
}
