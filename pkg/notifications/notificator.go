package notifications

type Notification struct{
	Message string
	Code string
}

type NotificatorInterface interface{
	AddNotification(notification Notification)
	HasNotification() bool
	GetNotifications() []Notification
}

type Notificator struct {
	notifications []Notification
}

func NewNotificator() *Notificator{
	return &Notificator{
		notifications: []Notification{},
	}
}

func (notificator *Notificator) AddNotification(notification Notification){
	notificator.notifications = append(notificator.notifications, notification)
}

func (notificator *Notificator) HasNotification() bool{
	return len(notificator.notifications) > 0
}

func (notificator *Notificator) GetNotifications() []Notification{
	return notificator.notifications
}


