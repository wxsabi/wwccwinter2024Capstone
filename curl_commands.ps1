curl -X POST -H "Content-Type: application/json" -d '{
	"Name": "John",
	"LastName": "Salchichón",
	"Email": "john@laconchadelalora.com",
	"Password": "rambo",
}' http://localhost:8888/signup


Invoke-WebRequest -Uri "http://localhost:8888/signup" -Method POST -Headers @{
    "Content-Type" = "application/json"
} -Body '{
    "Name": "John",
    "LastName": "Salchichón",
    "Email": "john@laconchadelalora.com",
    "Password": "rambo"
}'

Invoke-WebRequest -Uri "http://localhost:8888/signin" -Method POST -Headers @{
    "Content-Type" = "application/json"
} -Body '{
    "Email": "john@laconchadelalora.com",
    "Password": "rambo"
}'

$ curl -X POST -d "email=john@laconchadelalora.com&password=rambo&remembertoken=false" http://localhost:8888/signin