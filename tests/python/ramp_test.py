import grpc
import time
import threading
import statistics
import sys
import message_pb2
import message_pb2_grpc

class MaxCapacityTest:
    def __init__(self):
        self.results = []
        self.errors = []
        self.lock = threading.Lock()
        self.stop = threading.Event()
        
    def worker(self, stub, user_id):
        while not self.stop.is_set():
            try:
                start = time.time()
                response = stub.Login(message_pb2.UserInput(
                    user_name=f"user{user_id}",
                    user_email=f"user{user_id}@example.com",
                    user_password="password123"
                ))
                latency = (time.time() - start) * 1000
                
                with self.lock:
                    if response.status.code == 200:
                        self.results.append(latency)
                    else:
                        self.errors.append(1)
            except Exception as e:
                with self.lock:
                    self.errors.append(1)
            
            time.sleep(0)  # Без задержки - максимальная нагрузка
    
    def test_capacity(self, duration=10):
        """Тест максимальной емкости"""
        channel = grpc.insecure_channel('127.0.0.1:50051')
        stub = message_pb2_grpc.ApiUsersStub(channel)
        
        print(f"\n{'='*60}")
        print(f"ТЕСТ МАКСИМАЛЬНОЙ ПРОПУСКНОЙ СПОСОБНОСТИ")
        print(f"Длительность: {duration} секунд")
        print(f"{'='*60}")
        
        # Создаем тестовых пользователей
        print("Создание тестовых пользователей...")
        for i in range(100):
            try:
                stub.CreateUser(message_pb2.CreateUserRequest(
                    user_name=f"user{i}",
                    user_email=f"user{i}@example.com",
                    user_password="password123"
                ))
            except:
                pass
        
        # Постепенно увеличиваем количество воркеров
        max_workers = 500
        best_rps = 0
        best_workers = 0
        
        for workers in [500, 1000]:
            print(f"\n▶ Тест с {workers} параллельными соединениями...")
            
            self.results = []
            self.errors = []
            self.stop.clear()
            
            threads = []
            for i in range(workers):
                t = threading.Thread(target=self.worker, args=(stub, i % 100))
                t.daemon = True
                threads.append(t)
                t.start()
            
            time.sleep(duration)
            self.stop.set()
            
            for t in threads:
                t.join(timeout=1)
            
            total = len(self.results) + len(self.errors)
            rps = total / duration
            success_rate = (len(self.results) / total * 100) if total > 0 else 0
            
            print(f"   ✅ Успешно: {len(self.results)}")
            print(f"   ❌ Ошибки: {len(self.errors)}")
            print(f"   📊 RPS: {rps:.2f}")
            print(f"   📈 Успешность: {success_rate:.1f}%")
            
            if self.results:
                avg_latency = statistics.mean(self.results)
                p95 = sorted(self.results)[int(len(self.results) * 0.95)]
                print(f"   ⏱️  Средняя задержка: {avg_latency:.2f}мс")
                print(f"   ⏱️  95-й перцентиль: {p95:.2f}мс")
            
            if success_rate > 95 and rps > best_rps:
                best_rps = rps
                best_workers = workers
            elif success_rate < 80:
                print(f"\n⚠️ Качество упало ниже 80% при {workers} соединениях")
                break
        
        print(f"\n{'='*60}")
        print(f"РЕЗУЛЬТАТЫ ТЕСТИРОВАНИЯ")
        print(f"{'='*60}")
        print(f"🏆 Максимальный RPS: {best_rps:.2f}")
        print(f"👥 Оптимальное кол-во соединений: {best_workers}")
        print(f"{'='*60}")

if __name__ == "__main__":
    test = MaxCapacityTest()
    test.test_capacity(duration=10)