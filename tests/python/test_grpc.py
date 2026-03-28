import grpc
import sys
import os

sys.path.insert(0, os.path.dirname(__file__))

import message_pb2
import message_pb2_grpc

def test_grpc_server():
    # Подключение к gRPC серверу
    channel = grpc.insecure_channel('127.0.0.1:50051')
    stub = message_pb2_grpc.ApiUsersStub(channel)
    
    print("="*60)
    print("ТЕСТИРОВАНИЕ gRPC СЕРВЕРА")
    print("="*60)
    
    # Тест 1: Создание пользователя
    print("\n[1] Регистрация пользователя...")
    try:
        create_request = message_pb2.CreateUserRequest(
            user_name="testuser",
            user_email="test@example.com",
            user_password="password123"
        )
        create_response = stub.CreateUser(create_request)
        print(f"    Статус: {create_response.code}")
        if create_response.code == 201:
            print("    ✅ Пользователь успешно создан")
        elif create_response.code == 400:
            print("    ⚠️ Ошибка валидации")
        else:
            print(f"    Код ответа: {create_response.code}")
    except grpc.RpcError as e:
        print(f"    ❌ RPC ошибка: {e.code()} - {e.details()}")
    except Exception as e:
        print(f"    ❌ Ошибка: {e}")
    
    # Тест 2: Вход в систему
    print("\n[2] Вход в систему...")
    try:
        login_request = message_pb2.UserInput(
            user_name="testuser",
            user_email="test@example.com",
            user_password="password123"
        )
        login_response = stub.Login(login_request)
        print(f"    Статус: {login_response.status.code}")
        print(f"    Сообщение: {login_response.message}")
        if login_response.token:
            print(f"    ✅ Токен получен: {login_response.token[:50]}...")
    except grpc.RpcError as e:
        print(f"    ❌ RPC ошибка: {e.code()} - {e.details()}")
    except Exception as e:
        print(f"    ❌ Ошибка: {e}")
    
    # Тест 3: Вход с неверным паролем
    print("\n[3] Вход с неверным паролем...")
    try:
        login_request = message_pb2.UserInput(
            user_name="testuser",
            user_email="test@example.com",
            user_password="wrongpassword"
        )
        login_response = stub.Login(login_request)
        print(f"    Статус: {login_response.status.code}")
        print(f"    Сообщение: {login_response.message}")
        if login_response.status.code == 401:
            print("    ✅ Неверный пароль правильно отклонен")
        else:
            print(f"    Ожидался 401, получен {login_response.status.code}")
    except grpc.RpcError as e:
        print(f"    ❌ RPC ошибка: {e.code()} - {e.details()}")
    except Exception as e:
        print(f"    ❌ Ошибка: {e}")
    
    # Тест 4: Вход с пустыми данными
    print("\n[4] Вход с пустыми данными...")
    try:
        login_request = message_pb2.UserInput(
            user_name="",
            user_email="",
            user_password=""
        )
        login_response = stub.Login(login_request)
        print(f"    Статус: {login_response.status.code}")
        print(f"    Сообщение: {login_response.message}")
        if login_response.status.code == 400:
            print("    ✅ Валидация работает правильно")
        else:
            print(f"    Ожидался 400, получен {login_response.status.code}")
    except grpc.RpcError as e:
        print(f"    ❌ RPC ошибка: {e.code()} - {e.details()}")
    except Exception as e:
        print(f"    ❌ Ошибка: {e}")
    
    print("\n" + "="*60)
    print("ТЕСТИРОВАНИЕ ЗАВЕРШЕНО")
    print("="*60)

if __name__ == "__main__":
    # Проверка подключения к серверу
    print("Проверка подключения к серверу...")
    try:
        # Простая проверка - создаем канал и ждем
        channel = grpc.insecure_channel('127.0.0.1:50051')
        grpc.channel_ready_future(channel).result(timeout=3)
        print("✅ Сервер доступен\n")
        test_grpc_server()
    except grpc.FutureTimeoutError:
        print("❌ Сервер не доступен!")
        print("Убедитесь, что gRPC сервер запущен на 127.0.0.1:50051")
        print("Для запуска сервера выполните в другом терминале:")
        print("cd go_service && go run .")
    except Exception as e:
        print(f"❌ Ошибка подключения: {e}")