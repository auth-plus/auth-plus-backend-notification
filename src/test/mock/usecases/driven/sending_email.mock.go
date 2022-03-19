package test

import (
	"github.com/stretchr/testify/mock"
)

type SendingEmailMock struct {
	mock.Mock
}

func (m *SendingEmailMock) SendEmail(email string, content string) {
	m.Called(email, content)
}
