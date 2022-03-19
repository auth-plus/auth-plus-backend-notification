package test

import (
	"testing"

	u "auth-plus-notification/core/usecases"
	m "auth-plus-notification/test/mock/usecases/driven"
)

func TestPushNotificationSuccess(t *testing.T) {
	mocked := new(m.SendingPushNotificationMock)
	deviceId_test := "algum.email@test.com"
	title_test := "titulo"
	content_test := "conteúdo teste"
	mocked.On("SendPN", deviceId_test, title_test, content_test).Return()

	usecase := u.NewPushNotificationUsecase(mocked)
	usecase.Send(deviceId_test, title_test, content_test)
	mocked.AssertCalled(t, "SendPN", deviceId_test, title_test, content_test)
}
