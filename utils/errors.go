package utils

// Содержание ошибок

var (
	//Server
	ErrServerError = map[string]interface{}{"success":false,"msg":"Server error"}
	//Auth
	ErrIncorrectLogin = map[string]interface{}{"success":false,"msg":"Неправильные логин или пароль"}
	ErrPassMatch = map[string]interface{}{"success":false,"msg":"Пароли не совпадают"}
	ErrRegExist = map[string]interface{}{"success":false,"msg":"Такой пользователь уже есть"}
	ErrSmsConfirm = map[string]interface{}{"success":false,"msg":"Вы не подтвердили смс"}
	ErrUserNotFound = map[string]interface{}{"success":false,"msg":"Пользователь не найден"}
	//Sms
	ErrSmsTimelimit = map[string]interface{}{"success":false,"msg":"Повторное смс можно будет отправить через 60 секунд"}
	ErrSmsApprove = map[string]interface{}{"success":false,"msg":"Пользователь уже прошел проверку"}
	ErrSmsCode = map[string]interface{}{"success":false,"msg":"Код смс введен неверно"}
	ErrSmsExpired = map[string]interface{}{"success":false,"msg":"Проверочный код истек, повторите отправку"}
	//Transactions
	ErrTransactionId = map[string]interface{}{"success":false,"msg":"Параметр id передан в неправильном формате"}
	//User
	ErrFullname = map[string]interface{}{"success":false,"msg":"ФИО Введено некорректно"}
	//Files
	ErrFileError = map[string]interface{}{"success":false,"msg":"Не удалось создать файл"}
	ErrCSVError = map[string]interface{}{"success":false,"msg":"Ошабка записи в CSV файл"}
	//Token
	ErrCreateToken = map[string]interface{}{"success":false,"msg":"Ошибка создания токена"}
	ErrTokenMissing = map[string]interface{}{"success":false,"msg":"Токен не передан"}
	ErrSendToken = map[string]interface{}{"success":false,"msg":"Токен передан неправильно"}
	ErrAuthType = map[string]interface{}{"success":false,"msg":"Неправильный тип авторизации"}
	ErrMalformedToken = map[string]interface{}{"success":false,"msg":"Токен сформирован неправильно"}
	ErrInvalidToken = map[string]interface{}{"success":false,"msg":"Неправильный токен"}
	ErrExpiredToken = map[string]interface{}{"success":false,"msg":"Срок токена истек"}
)