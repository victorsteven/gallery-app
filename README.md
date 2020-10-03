
###Photo Gallery Application

####Installation
Clone the application using:


###The Server
Change to the ``server`` directory:
```
cd server
```
- Update the ``.env`` file with correct details. Both for the real database and the testing database.

- Run unit and integration test from the root of the server directory using:
```
go test -v ./...
```
- In the root of the server directory, start the server using:
```
go run main.go
```
The server will start running on port 7070(the default port from the .env file)

###The Client(VueJS)
In a different terminal window, change  to the ``client`` directory
````.env
cd client
````
- Installation:
```
npm install
```
- Start the client using:
```.env
npm run serve
```
- The client will start running on port 8080

###Content Display

Visit the browser with the address:
```.env
localhost:8080
```
![image here](https://res.cloudinary.com/chikodi/image/upload/v1601689035/Screenshot_2020-10-03_at_02.36.09.png)

###Customization
- If you wish to change the port the server listens to, you will need to change that in the ``.env`` file in the server directory and also in the ``src/.env.js`` file in the client directory.