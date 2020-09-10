package uuid

import "github.com/google/uuid"

/**
获取 uuid 直接使用似乎更好？？
*/
func GetUUID() string {
	return uuid.New().String()
}
