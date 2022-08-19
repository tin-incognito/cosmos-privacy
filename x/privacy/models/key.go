package models

import (
	"privacy/x/privacy/common"
	"privacy/x/privacy/repos/key"
)

func GeneratePrivateKey() key.PrivateKey {
	b := common.RandBytes(common.HashSize)
	return key.GeneratePrivateKey(b)

}
