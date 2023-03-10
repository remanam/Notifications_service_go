package Interactors

import (
	"Egor_notifications/Domain"
	"Egor_notifications/app"
	"github.com/pkg/errors"
)

type NotificationsInteractor struct {
	log               app.Logger
	notificationsRepo domain.NotificationsRepository
}

func NewNotificationsUCase(log app.Logger, notificationsRepo domain.NotificationsRepository) domain.NotificationInteractor {
	return &NotificationsInteractor{
		log:               log,
		notificationsRepo: notificationsRepo,
	}
}

func (i *NotificationsInteractor) GetAllNotifications() (domain.GetAllNotificationsResponse, error) {
	notifications, err := i.notificationsRepo.All()
	if err != nil {
		return domain.GetAllNotificationsResponse{}, errors.Wrap(err, "Cannot fetch all notifications")
	}

	return domain.GetAllNotificationsResponse{
		StatusCode:    domain.Success,
		Notifications: notifications,
	}, nil
}

func (i *NotificationsInteractor) CreateNotification(name string) (domain.CreateNotificationResponse, error) {
	notification := &domain.Notification{
		Title: "Title",
	}
	err := i.notificationsRepo.Insert(notification)
	if err != nil {
		return domain.CreateNotificationResponse{}, errors.Wrap(err, "Cannot insert task")
	}

	return domain.CreateNotificationResponse{
		StatusCode: domain.Success,
		Id:         notification.Id,
	}, nil
}
