import psycopg2

# Данные для подключения к твоей локальной базе
connection_params = {
    "dbname": "postgres",        # Название твоей БД (часто по умолчанию "postgres")
    "user": "postgres",          # Твой пользователь (часто по умолчанию "postgres")
    "password": "postgres", # Твой пароль, который ты ставил при установке
    "host": "localhost",         # Адрес сервера
    "port": "5432"               # Порт (стандартный для PostgreSQL)
}

try:
    # Пытаемся соединиться
    connection = psycopg2.connect(**connection_params)
    print("✅ Успех! Соединение с PostgreSQL установлено.")
    
    # Закрываем соединение
    connection.close()
    
except Exception as error:
    print(f"❌ Ошибка при подключении к PostgreSQL: {error}")