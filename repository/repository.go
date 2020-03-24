package repository

// общие методы с информацией о клиентах
type UserRepo interface {
	Get(url string, body string) ([]byte, error)
	GetUserOrders(clientID string) ([]byte, error)
	SendOSMICard(clientID string) error
}

// методы для мобильного приложения
type MobileRepo interface {
	Creation(deviceUUID string, id string, vendor string, model string) ([]byte, error)
	Authorization(deviceUUID string) ([]byte, error)
	Registration(deviceUUID string, user map[string]string) ([]byte, error)
	Code(phone string) ([]byte, error)
	CheckCode(numberPhone string, code string) ([]byte, error)
	EditUser(deviceUUID string, userID string, numberPhone string) ([]byte, error)
}
