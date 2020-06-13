package modules

import (
	"crypto/md5"
	"encoding/hex"
	E "github.com/yuw-mvc/yuw/exceptions"
)

type (
	TokenPoT struct {
		PubKey []byte
		PrvKey []byte
	}

	Token struct {
		Keys *TokenPoT
	}
)

func NewToken() *Token {
	return &Token{}
}

func (token *Token) Pwd(keys ...string) (res string, err error) {
	if len(keys) < 1 || keys[0] == "" {
		err = E.Err("yuw^m_token_a", E.ErrPosition())
		return
	}

	var strKey string = ""
	for _, key := range keys {
		strKey = strKey + key
	}

	h := md5.New()
	h.Write([]byte(strKey))
	res = hex.EncodeToString(h.Sum([]byte("")))

	return
}


