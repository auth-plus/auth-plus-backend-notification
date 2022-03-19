package test

import (
	"testing"

	u "auth-plus-notification/core/usecases"
	m "auth-plus-notification/test/mock/usecases/driven"
)

func TestSmsSuccess(t *testing.T) {
	mocked := new(m.SendingSmsMock)
	phone_test := "123456789"
	content_test := "conteúdo teste"
	mocked.On("SendSms", phone_test, content_test).Return()

	usecase := u.NewSmsUsecase(mocked)
	usecase.Send(phone_test, content_test)
	mocked.AssertCalled(t, "SendSms", phone_test, content_test)
}
