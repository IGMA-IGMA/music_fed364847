from faker import Faker

fake = Faker("en_US")
users = [{"username": fake.name(), "email": fake.email(), "pwd": fake.password()}
         for _ in range(100)]
print(users)
