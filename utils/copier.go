package utils

import "github.com/jinzhu/copier"

func Copy(toValue interface{}, fromValue interface{}) (err error) {
	return copier.Copy(toValue, fromValue)
}

func CopyWithIgnoreEmpty(toValue interface{}, fromValue interface{}) (err error) {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{IgnoreEmpty: true})
}

func DeepCopy(toValue interface{}, fromValue interface{}) (err error) {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{DeepCopy: true})
}

func DeepCopyWithIgnoreEmpty(toValue interface{}, fromValue interface{}) (err error) {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{IgnoreEmpty: true, DeepCopy: true})
}
