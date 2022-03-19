package test

import (
	"testing"

	u "auth-plus-notification/core/usecases"
	m "auth-plus-notification/test/mock/usecases/driven"
)

func TestWhatsappSuccess(t *testing.T) {
	mocked := new(m.SendingWhatsappMock)
	phone_test := "123456789"
	content_test := "conteúdo teste"
	mocked.On("SendWhats", phone_test, content_test).Return()

	usecase := u.NewWhatsappUsecase(mocked)
	usecase.Send(phone_test, content_test)
	mocked.AssertCalled(t, "SendWhats", phone_test, content_test)
}
