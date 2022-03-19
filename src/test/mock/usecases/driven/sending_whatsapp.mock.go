package test

import (
	"github.com/stretchr/testify/mock"
)

type SendingWhatsappMock struct {
	mock.Mock
}

func (m *SendingWhatsappMock) SendWhats(phone string, content string) {
	m.Called(phone, content)
}
