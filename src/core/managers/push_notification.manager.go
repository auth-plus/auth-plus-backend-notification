package managers

import (
	d "auth-plus-notification/core/usecases/driven"
	"math/rand"
)

//Class for PushNotificationManager
type PushNotificationManager struct {
	firebase d.SendingPushNotification
	braze    d.SendingPushNotification
}

func NewPushNotificationManager(firebase d.SendingPushNotification, braze d.SendingPushNotification) *PushNotificationManager {
	return &PushNotificationManager{firebase: firebase, braze: braze}
}

func (e *PushNotificationManager) SendPN(deviceId string, title string, content string) {
	choosedProvider := choosePushNotificationProvider(deviceId, title, content)
	switch choosedProvider {
	case "Firebase":
		e.firebase.SendPN(deviceId, title, content)
	case "Braze":
		e.braze.SendPN(deviceId, title, content)
	}
}

//Function for choosing a provider, it can be by IP warming, Limit, timeout
func choosePushNotificationProvider(deviceId string, title string, content string) EnumPushNotificationProvider {
	if rand.Float64() > 0.5 {
		return "Firebase"
	}
	return "Braze"
}
