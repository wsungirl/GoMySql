# Go Interface for MySql
## Rest Api Info
* **GET**
    * /dbs -- _Получение списка баз данных пользователя_
    * /dbs/{db_id} -- _Получение данных о выбраной БД_
    * /dbs/{db_id}/tables -- _Получение списка таблиц выбраной базы данных_
    * /dbs/{db_id}/tables/{table_id} -- _Получение данных о таблице_
    * /dbs/{db_id}/tables/{table_id}/rows[?offset={offset}&limit={limit}&fields[]={FieldName}...] -- _Получение строк из таблицы_
    * /dbs/{db_id}/tables/{table_id}/rows/{row_id} -- _Получение конкретной строки из таблицы_
    * /users/{user_id}/auth -- _Получение авторизационного токена_
    * /users/{user_id} -- _Получение имени пользователя_
    * /users/{user_id}/revoke -- _Отзыв авторизационного ключа_
* **POST**
    * /dbs -- _Создание новой БД_
    * /dbs/{db_id}/tables -- _Создание новой таблици в БД_
    * /dbs/{db_id}/tables/{table_id} -- _Создание новой строки в таблице в БД_
    * /users -- _Создание нового пользователя_
* **PUT**
    * /dbs/{db_id}/tables/{table_id}/rows/{row_id} -- _Обновление данных в строке_
    * /dbs/{db_id}/tables/{table_id}/rows -- _Обновление данных в нескольких строках_
* **DELETE**
    * /dbs/{db_id} -- _Удалить базу_
    * /dbs/{db_id}/tables/{table_id} -- _Удалить таблицу_
    * /dbs/{db_id}/tables/{table_id}/rows/{row_id} -- _Удалить строку_

## Packages
* go get "github.com/go-sql-driver/mysql"
* go get "github.com/gorilla/mux"
* go get "golang.org/x/crypto/bcrypt"


## Permissions

 * Actions
    * read
    * create
    * update
 * Entities
    * table
    * row
