package test

import (
	"github.com/stretchr/testify/mock"
)

type SendingTelegramMock struct {
	mock.Mock
}

func (m *SendingTelegramMock) SendTele(phone string, content string) {
	m.Called(phone, content)
}
