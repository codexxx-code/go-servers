package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// IsFraudReq описывает данные для проверки фраудскором.
type IsFraudReq struct {
	Event     string // Тип события
	IP        string // IP-адрес
	UserAgent string // User-Agent пользователя
	AT        int64  // Временная метка
}

func (s *IsFraudReq) GetHash() string {
	hash := md5.Sum([]byte(fmt.Sprintf("%v%v%v", s.Event, s.IP, s.UserAgent)))
	return hex.EncodeToString(hash[:])
}
