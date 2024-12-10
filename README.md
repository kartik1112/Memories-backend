Memories Backend
=============

**Memories** is a collaborative event photo-sharing platform where users can upload photos, tag themselves using facial recognition, and relive shared memories. This repository contains the backend API built with **Go** using the **Gin framework**, **GORM** for ORM, and **MySQL** as the database.

* * * * *

Memories Frontend
=============

Link to frontend repo (Flutter) - https://github.com/kartik1112/Memories-flutter 

Features
--------

-   **User Authentication:** Secure login/signup with JWT-based authentication.
-   **Event Management:** Create and join events using a 6-digit unique event code.
-   **Photo Management:** Upload photos linked to specific events.
-   **Facial Recognition:** Automatically tags users in photos.
-   **Scalable Architecture:** Built with best practices for modular and maintainable code.

* * * * *

Tech Stack
----------

-   **Language:** Go (Golang)
-   **Framework:** Gin
-   **Database:** MySQL (with GORM ORM)
-   **Authentication:** JWT
-   **Cloud Storage:** For photos (e.g., AWS S3, Firebase Storage - TBD)
-   **Facial Recognition:** TBD (e.g., AWS Rekognition, Google Vision API, or a pre-trained model like Dlib/FaceNet)

* * * * *

API Endpoints
-------------

### **Public Endpoints**

| Method | Endpoint | Description |
| --- | --- | --- |
| POST | `/signup` | Create a new user account. |
| POST | `/login` | Authenticate and get a JWT. |

### **Protected Endpoints**

| Method | Endpoint | Description |
| --- | --- | --- |
| POST | `/event/create` | Create a new event. |
| POST | `/events/join` | Join an event using a 6-digit code. |
| GET | `/events/fetch` | Fetch all events for the logged-in user. |
| POST | `/photos/upload` | Upload a photo to an event (WIP). |

* * * * *

Contributing
------------

Contributions are welcome! Please fork the repository and create a pull request.

* * * * *

License
-------

This project is licensed under the MIT License.
