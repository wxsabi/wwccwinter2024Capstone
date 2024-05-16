curl -X POST -H "Content-Type: application/json" -d '{"Name": "John","LastName": "Salchichón","Email": "john@laconchadelalora.com","Password": "rambo",}' http://localhost:8888/signup


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

To login:
curl -X POST -d "email=john@laconchadelalora.com&password=rambo&remembertoken=false" http://localhost:8888/signin

to update user:

curl -X PUT -H "Content-Type: application/json" -d '{"Name": "John","LastName": "Salchichón","Email": "john@laconchadelalora.com","ID": 1,}' http://localhost:8888/signup 

MySQL Commands:

Show users and their data: 
SELECT * FROM Users LEFT JOIN Sessions ON Users.ID = Sessions.UserID LEFT JOIN Items ON Users.ID = Items.UserID LEFT JOIN ItemPhotos ON Items.ItemID = ItemPhotos.ItemID;

same but also the 4 separate tables:
SELECT * FROM Users LEFT JOIN Sessions ON Users.ID = Sessions.UserID LEFT JOIN Items ON Users.ID = Items.UserID LEFT JOIN ItemPhotos ON Items.ItemID = ItemPhotos.ItemID; SELECT * FROM Users; SELECT * FROM Sessions; SELECT * FROM Items; SELECT * FROM ItemPhotos;