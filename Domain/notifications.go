package domain

import "time"

type Notification struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	User_id   int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationsRepository interface {
	All() ([]*Notification, error)
	FetchById(int64) (*Notification, error)
	Insert(notification *Notification) error
	Update(notification *Notification) error
}

type NotificationInteractor interface {
	GetAllNotifications() (GetAllNotificationsResponse, error)
	CreateNotification(title string) (CreateNotificationResponse, error)
}

// Responses (only for UseCase layer)
type GetAllNotificationsResponse struct {
	StatusCode    string
	Notifications []*Notification
}

type CreateNotificationResponse struct {
	StatusCode string
	Id         int64
}
