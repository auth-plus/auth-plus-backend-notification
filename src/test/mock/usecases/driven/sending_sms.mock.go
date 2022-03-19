package test

import (
	"github.com/stretchr/testify/mock"
)

type SendingSmsMock struct {
	mock.Mock
}

func (m *SendingSmsMock) SendSms(phone string, content string) {
	m.Called(phone, content)
}
