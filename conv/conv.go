package conv

import (
	"strconv"
)

func ParseBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func ParseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

func ParseInt8(s string) (int8, error) {
	i, err := strconv.ParseInt(s, 10, 8)
	return int8(i), err
}

func ParseInt16(s string) (int16, error) {
	i, err := strconv.ParseInt(s, 10, 16)
	return int16(i), err
}

func ParseInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	return int32(i), err
}

func ParseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func ParseFloat32(s string) (float32, error) {
	i, err := strconv.ParseFloat(s, 32)
	return float32(i), err
}

func ParseFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func MustParseBool(s string) bool {
	i, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return i
}

func MustParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func MustParseInt8(s string) int8 {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		panic(err)
	}
	return int8(i)
}

func MustParseInt16(s string) int16 {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		panic(err)
	}
	return int16(i)
}

func MustParseInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(i)
}

func MustParseInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func MustParseFloat32(s string) float32 {
	i, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}
	return float32(i)
}

func MustParseFloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseBools(s []string) ([]bool, error) {
	nums := make([]bool, 0, len(s))
	for _, a := range s {
		n, err := ParseBool(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseInts(s []string) ([]int, error) {
	nums := make([]int, 0, len(s))
	for _, a := range s {
		n, err := ParseInt(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseInt8s(s []string) ([]int8, error) {
	nums := make([]int8, 0, len(s))
	for _, a := range s {
		n, err := ParseInt8(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseInt16s(s []string) ([]int16, error) {
	nums := make([]int16, 0, len(s))
	for _, a := range s {
		n, err := ParseInt16(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseInt32s(s []string) ([]int32, error) {
	nums := make([]int32, 0, len(s))
	for _, a := range s {
		n, err := ParseInt32(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseInt64s(s []string) ([]int64, error) {
	nums := make([]int64, 0, len(s))
	for _, a := range s {
		n, err := ParseInt64(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseFloat32s(s []string) ([]float32, error) {
	nums := make([]float32, 0, len(s))
	for _, a := range s {
		n, err := ParseFloat32(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseFloat64s(s []string) ([]float64, error) {
	nums := make([]float64, 0, len(s))
	for _, a := range s {
		n, err := ParseFloat64(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 0)
	return uint(i), err
}

func MustParseUint(s string) uint {
	i, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		panic(err)
	}
	return uint(i)
}

func ParseUints(s []string) ([]uint, error) {
	nums := make([]uint, 0, len(s))
	for _, a := range s {
		n, err := ParseUint(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseUint8(s string) (uint8, error) {
	i, err := strconv.ParseUint(s, 10, 8)
	return uint8(i), err
}

func MustParseUint8(s string) uint8 {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		panic(err)
	}
	return uint8(i)
}

func ParseUint8s(s []string) ([]uint8, error) {
	nums := make([]uint8, 0, len(s))
	for _, a := range s {
		n, err := ParseUint8(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseUint16(s string) (uint16, error) {
	i, err := strconv.ParseUint(s, 10, 16)
	return uint16(i), err
}

func MustParseUint16(s string) uint16 {
	i, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(i)
}

func ParseUint16s(s []string) ([]uint16, error) {
	nums := make([]uint16, 0, len(s))
	for _, a := range s {
		n, err := ParseUint16(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseUint32(s string) (uint32, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	return uint32(i), err
}

func MustParseUint32(s string) uint32 {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(i)
}

func ParseUint32s(s []string) ([]uint32, error) {
	nums := make([]uint32, 0, len(s))
	for _, a := range s {
		n, err := ParseUint32(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ParseUint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	return i, err
}

func MustParseUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseUint64s(s []string) ([]uint64, error) {
	nums := make([]uint64, 0, len(s))
	for _, a := range s {
		n, err := ParseUint64(a)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

// formatter

func FormatBool(i bool) string {
	return strconv.FormatBool(i)
}

func FormatBools(i []bool) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatBool(a))
	}
	return strs
}

func FormatInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatInts(i []int) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatInt(int64(a), 10))
	}
	return strs
}

func FormatInt8(i int8) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatInt8s(i []int8) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatInt(int64(a), 10))
	}
	return strs
}

func FormatInt16(i int16) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatInt16s(i []int16) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatInt(int64(a), 10))
	}
	return strs
}

func FormatInt32(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatInt32s(i []int32) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatInt(int64(a), 10))
	}
	return strs
}

func FormatInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FormatInt64s(i []int64) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatInt(a, 10))
	}
	return strs
}

func FormatFloat32(i float32) string {
	return strconv.FormatFloat(float64(i), 'f', -1, 32)
}

func FormatFloat32s(i []float32) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatFloat(float64(a), 'f', -1, 32))
	}
	return strs
}

func FormatFloat64(i float64) string {
	return strconv.FormatFloat(float64(i), 'f', -1, 64)
}

func FormatFloat64s(i []float64) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatFloat(float64(a), 'f', -1, 64))
	}
	return strs
}

func FormatUint(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

func FormatUints(i []uint) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatUint(uint64(a), 10))
	}
	return strs
}

func FormatUint8(i uint8) string {
	return strconv.FormatUint(uint64(i), 10)
}

func FormatUint8s(i []uint8) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatUint(uint64(a), 10))
	}
	return strs
}

func FormatUint16(i uint16) string {
	return strconv.FormatUint(uint64(i), 10)
}

func FormatUint16s(i []uint16) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatUint(uint64(a), 10))
	}
	return strs
}

func FormatUint32(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func FormatUint32s(i []uint32) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatUint(uint64(a), 10))
	}
	return strs
}

func FormatUint64(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func FormatUint64s(i []uint64) []string {
	strs := make([]string, 0, len(i))
	for _, a := range i {
		strs = append(strs, strconv.FormatUint(a, 10))
	}
	return strs
}
