package Deliverylayer

import (
	domain "Egor_notifications/Domain"
	"Egor_notifications/app"
)

type NotificationsDeliveryService struct {
	notificationsUCase domain.NotificationInteractor
	log                app.Logger
}

func NewDeliveryService(log app.Logger, notificationsUCase domain.NotificationInteractor) *NotificationsDeliveryService {
	return &NotificationsDeliveryService{
		log:                log,
		notificationsUCase: notificationsUCase,
	}
}
