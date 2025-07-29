package util

import (
	"context"
	"flowing/internal/repository"
	"github.com/mojocn/base64Captcha"
	"time"
)

const captchaKeyPrefix = "captcha:"

type captchaStore struct{}

func (c *captchaStore) Set(id string, value string) error {
	return repository.Redis().
		SetEx(context.Background(), captchaKeyPrefix+id, value, time.Minute*1).
		Err()
}

func (c *captchaStore) Get(id string, clear bool) string {
	var value string
	var err error
	if clear {
		value, err = repository.Redis().GetDel(context.Background(), captchaKeyPrefix+id).Result()
	} else {
		value, err = repository.Redis().Get(context.Background(), captchaKeyPrefix+id).Result()
	}
	if err != nil {
		return ""
	}
	return value
}

func (c *captchaStore) Verify(id, answer string, clear bool) bool {
	value := c.Get(id, clear)
	return answer == value
}

var captchaImpl = base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, &captchaStore{})

func GenCaptcha() (string, string, error) {
	id, img, _, err := captchaImpl.Generate()
	if err != nil {
		return "", "", err
	}
	return id, img, nil
}

func VerifyCaptcha(id, answer string) bool {
	return captchaImpl.Verify(id, answer, true)
}
