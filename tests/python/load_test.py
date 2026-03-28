import grpc
import time
import threading
import statistics
import message_pb2
import message_pb2_grpc

def test_rps(rps=50, duration=10):
    channel = grpc.insecure_channel('127.0.0.1:50051')
    stub = message_pb2_grpc.ApiUsersStub(channel)
    
    results = []
    errors = []
    stop = threading.Event()
    
    def worker():
        while not stop.is_set():
            try:
                start = time.time()
                stub.Login(message_pb2.UserInput(
                    user_name="testuser",
                    user_email="test@example.com",
                    user_password="password123"
                ))
                results.append((time.time() - start) * 1000)
            except:
                errors.append(1)
            time.sleep(1.0 / rps)
    
    threads = [threading.Thread(target=worker) for _ in range(rps // 10 + 1)]
    for t in threads:
        t.start()
    
    time.sleep(duration)
    stop.set()
    
    for t in threads:
        t.join()
    
    print(f"\nRPS: {rps}, Длительность: {duration}с")
    print(f"Успешно: {len(results)}")
    print(f"Ошибок: {len(errors)}")
    if results:
        print(f"Ср. latency: {statistics.mean(results):.2f}мс")
        print(f"95-й перцентиль: {sorted(results)[int(len(results)*0.95)]:.2f}мс")

test_rps(rps=50, duration=10)
