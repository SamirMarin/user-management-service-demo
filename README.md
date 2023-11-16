# user-management-service
User management service, manages creation and querying of user accounts.

## Running the service locally
1. Start local running dynamodb
```bash
# must be running docker to run this command
docker-compose up
```
2. Build the service
```bash
go build -o user-management-service
```

3. Run the service
```bash
./user-management-service
```

## Making requests to the service locally
1. Create a user
```bash
curl -X POST http://localhost:1324/create \
-H "Content-Type: application/json" \
-d '{
      "username": "samir@gmail.com",
      "email": "samir@gmail.com",
      "firstname": "Samir",
      "lastname": "Marin",
      "dob": "1989-01-01T00:00:00Z",
      "membership": {
          "kind": "unlimited",
          "owner": true,
          "joined": "2023-01-01T00:00:00Z",
          "renewal": "2024-01-01T00:00:00Z"
      }
    }'
```
2. Get a workout
```bash
curl -X POST http://localhost:1324/get \
-H "Content-Type: application/json" \
-d '{
  "username": "samir@gmail.com"
}'
```
