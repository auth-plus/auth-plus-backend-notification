package test

import (
	"testing"

	u "auth-plus-notification/core/usecases"
	m "auth-plus-notification/test/mock/usecases/driven"
)

func TestTelegramSuccess(t *testing.T) {
	mocked := new(m.SendingTelegramMock)
	phone_test := "123456789"
	content_test := "conteúdo teste"
	mocked.On("SendTele", phone_test, content_test).Return()

	usecase := u.NewTelegramUsecase(mocked)
	usecase.Send(phone_test, content_test)
	mocked.AssertCalled(t, "SendTele", phone_test, content_test)
}
