# my_tg_bot
## Телеграм бот "Список задач"  
Бот разворачивается в контейнере, для хранения задач используется PosgreSQL, на данный момент поддерживает команды:  
Добавление - "Добавь <ваша задача>"  
Отображение списка задач с разделением на выполненные и актуальные - "Список"  
Пометить задачу как выполненную - "Сделал <номер задачи>"  
Команды не чувствительны к регистру  
В корневую директорию нужно добавить .env файл с параметрами запуска:  
TOKEN=<Токен вашего тг бота>  
POSTGRES_HOST=<хост или имя docker сервиса>  
POSTGRES_PORT=<порт>  
POSTGRES_USER=<пользователь>  
POSTGRES_PASSWORD=<пароль>  
POSTGRES_DB=<имя базы данных>  
