This project aims to build a similar application in different languages. The functionality is to set a key in the memcached and the ability to retrieve that key using http requests.

To have the app running for each language you will need docker cli. To run any project go inside of it and run `docker compose up --build`. You can only run one container at a time since they use the same port.

To set a key do:

`POST localhost:8080/set/testKey?value=testValue`

To retrieve a key do:

`GET localhost:8080/get/testKey`

Some of the finding using Jmeter are the following:

Number of threads 1000
Ramp up 60

Spring
Non-using: CPU 0.12% Memory 185.81 MB
Max Hike Using: CPU 40% Memory 282 MB
Load Time: 3.5
Latency: 3.5

Rust
Non-using: CPU 0.09% Memory 11.97 MB
Max Hike Using: CPU 2.11% Memory 12.25 MB
Load Time: 2
Latency: 2

Go
Non-using: CPU 0.01% Memory 14.8 MB
Max Hike Using CPU: 4.69% Memory 22.13 MB
Load Time: 2.5
Latency: 2.5

Nest
Non-using: CPU 0.18% Memory 266.42 MB
Max Hike Using CPU: 9.43% Memory 312 MB
Load Time: 5
Latency: 5
