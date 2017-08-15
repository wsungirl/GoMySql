# Go Interface for MySql
## Rest Api Info
* **GET**
    * /dbs -- _Получение списка баз данных пользователя_
    * /dbs/{db_id} -- _Получение данных о выбраной БД_
    * /dbs/{db_id}/tables -- _Получение списка таблиц выбраной базы данных_
    * /dbs/{db_id}/tables/{table_id} -- _Получение данных о таблице_
    * /dbs/{db_id}/tables/{table_id}/rows[?offset={offset}&limit={limit}] -- _Получение строк из таблицы_
    * /dbs/{db_id}/tables/{table_id}/rows/{row_id} -- _Получение конкретной строки из таблицы_
    * /users/{user_id}/auth -- _Получение авторизационного токена_
* **POST**
    * /dbs -- _Создание новой БД_
    * /dbs/{db_id}/tables -- _Создание новой таблици в БД_
    * /dbs/{db_id}/tables/{table_id} -- _Создание новой строки в таблице в БД_
    * /users -- _Создание нового пользователя_
* **PUT**
    * /dbs/{db_id}/tables/{table_id}/rows/{row_id} -- _Обновление данных в строке_
* **DELETE**
    * /dbs/{db_id} -- _Удалить базу_
    * /dbs/{db_id}/tables/{table_id} -- _Удалить таблицу_
    * /dbs/{db_id}/tables/{table_id}/rows/{row_id} -- _Удалить строку_
