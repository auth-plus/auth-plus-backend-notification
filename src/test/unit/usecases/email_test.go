package test

import (
	"testing"

	u "auth-plus-notification/core/usecases"
	m "auth-plus-notification/test/mock/usecases/driven"
)

func TestEmailSuccess(t *testing.T) {
	mocked := new(m.SendingEmailMock)
	email_test := "algum.email@test.com"
	content_test := "conteúdo teste"
	mocked.On("SendEmail", email_test, content_test).Return()

	usecase := u.NewEmailUsecase(mocked)
	usecase.Send(email_test, content_test)
	mocked.AssertCalled(t, "SendEmail", email_test, content_test)
}
