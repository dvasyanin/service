package repository

// общие методы с информацией о клиентах
type UserRepo interface {
	// получить основную информацию по клиенту
	Get(url string, body string) ([]byte, error)
	// история покупок клиента
	GetUserOrders(clientID string) ([]byte, error)
	// выпустить эл.карту клиенту
	SendOSMICard(clientID string) error
}

// методы для мобильного приложения
type MobileRepo interface {
	// Передача информации об устройстве клиента при запуске приложения
	Creation(deviceUUID string, id string, vendor string, model string) ([]byte, error)
	// Авторизация пользователя по ранее заполненному deviceUUID.
	Authorization(deviceUUID string) ([]byte, error)
	// Регистрация пользователя и связка с deviceUUID, для дальнейшей авторизации только по deviceUUID
	Registration(deviceUUID string, user map[string]string) ([]byte, error)
	// Генерация ключа авторизации и отправка его на номер телефона клиента
	Code(phone string) ([]byte, error)
	// Проверка ключа авторизации
	CheckCode(numberPhone string, code string) ([]byte, error)
	// Редактирование учетных данных клиента. (Смена номера клиента)
	EditUser(deviceUUID string, userID string, numberPhone string) ([]byte, error)
}
