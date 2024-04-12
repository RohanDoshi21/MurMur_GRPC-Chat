## MurMur: A GRPC + MQTT Chat Server

This project implements a group chat application with features like invite-based group creation, real-time messaging, and notifications.

### Technologies Used

* **Authentication:** Casdoor OAuth 2.0
* **Server-Side:** Go
* **Group Management:** Invite System
* **Real-time Messaging:** gRPC
* **Notifications:** MQTT

### Features

* User authentication using Casdoor
* Group creation with an invite system
* Real-time messaging for group and individual chats
* Notification system using MQTT

### Getting Started
Starting casdoor & postgres docker container
```bash
docker compose up -d
```

Run Server
```bash
go run main.go
```

### Using the Application:
* Users will need to authenticate using a valid Casdoor account.
* The application will allow the creation of groups and inviting users through the invite system.
* Users can participate in real-time messaging within groups and for individual chats.
* The notification system will send alerts through MQTT.

### Screenshots:
![Screenshot from 2024-04-12 12-21-58](https://github.com/RohanDoshi21/messaging-platform/assets/63660267/afee9531-3f08-4e40-afdc-51a465f3c189)
![Screenshot from 2024-04-12 12-22-13](https://github.com/RohanDoshi21/messaging-platform/assets/63660267/5d5997f6-0d8c-415c-9d9a-e7bdcdf1aba8)
![Screenshot from 2024-04-12 12-23-21](https://github.com/RohanDoshi21/messaging-platform/assets/63660267/f9850591-ef6e-47f6-ac0e-87add6a1aae7)
![Screenshot from 2024-04-12 12-23-32](https://github.com/RohanDoshi21/messaging-platform/assets/63660267/d6a40b24-0c5e-4aab-af80-6693747e7f1f)
![Screenshot from 2024-04-12 12-23-41](https://github.com/RohanDoshi21/messaging-platform/assets/63660267/4837fe27-888d-4b1e-a000-6eddfcf2700d)
![Screenshot from 2024-04-12 12-23-57](https://github.com/RohanDoshi21/messaging-platform/assets/63660267/7fc2a5e4-5a7a-4336-b7f8-abcf6f2954ce)

