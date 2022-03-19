package test

import (
	"github.com/stretchr/testify/mock"
)

type SendingPushNotificationMock struct {
	mock.Mock
}

func (m *SendingPushNotificationMock) SendPN(deviceId string, title string, content string) {
	m.Called(deviceId, title, content)
}
