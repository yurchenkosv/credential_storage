**GophKeeper** - это сервис для хранения и работы с секретной информацией

Сервис поддерживает следующие типы хранимых данных
- пары логин/пароль;
- произвольные текстовые данные;
- произвольные бинарные данные;
- данные банковских карт.

Сервис распространяется в виде 2х компонент: клиента и сервера

Настройка сервера.

Для настройки могут исопльзоваться переменные окружения 

```
CRED_SERVER_DATABASE_DSN - строка подключения к БД postgres  в формате DSN
CRED_SERVER_GRPC_LISTEN - хост и порт на которые будет биндиться сервер, к примеру localhost:8080
CRED_SERVER_JWT_SECRET - секрет используемый для генерации и проверки JWT токенов
CRED_SERVER_ENCRYPTION_SECRET - секрет используемый для шифрования всей информации о хранимых данных. При утере восстановить данные не получится.
CRED_SERVER_LOCAL_STORAGE_LOCATION - локальная папка для сохранения шифрованных бинарей
CRED_SERVER_CERT_LOCATION - путь к файлу сертификата для сервера
CRED_SERVER_PRIVATE_KEY_LOCATION - путь к закрытому ключу для сервера
```

Настройка клиента:

Для настройки используются переменные окружения и флаги
Переменные окружения:
```
CRED_SERVER_ADDRESS - адрес сервера, к которому выполняется подключение (localhost:8090)
CRED_CLIENT_BINARY_STORAGE_LOCATION - путь на клиенте, куда сохранять скачанные файлы
CRED_CLIENT_CA_CERT - путь на клиенте до CA сертификата для провеки подлинности серера
```

Флаги:

```
-address  адрес сервера, к которому выполняется подключение (localhost:8090)
-u имя пользователя для аутентификации на сервере
-p пароль пользователя для аутентификации на сервере
-r необязательная переменная. Указание ее говорит о том, что вместо аутентификации нужно сделать регистрацию пользователя
-n необязательная переменная. При регистрации передает имя пользователя (не логин)
-s путь на клиенте, куда сохранять скачанные файлы
-с путь на клиенте до CA сертификата для провеки подлинности серера
```

Локальный запуск:

Сервер можно запустить командой `make run_server` 
После этого будет поднят контейнер с БД и запущен сервер с дефолтными настройками (localhost:8090)

Клиента сначала надо сбилдить командой `make build_local_client`

После этого в директории bin/client появятся бинари клиента для разных платформ
Начальный запуск подразумевает регистрацию клиента на сервере. Сделать это можно командой

`<client_bin> -u <login> -p <password> -r -n <name>`

После успешной регистрации запускаем клиент в режиме работы с секретной информацией 

`<client_bin> -u <login> -p <password> -s <storage directory>`

Для удобства имеет смысл установить переменные окружения
```
export CRED_CLIENT_CA_CERT=hack/ca-cert.pem
export CRED_SERVER_ADDRESS=localhost:8090
```

Тестирование юнит + интеграционное выполняется командой `make test`
