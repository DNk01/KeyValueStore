# Key Value Storage

Key-value storage (redis-like) that supports Set, Get, and Delete methods and store data in RAM

## Download and install

1. Clone repository:

        git clone https://github.com/DNk01/KeyValueStore.git


2. Go to directory:

        cd KeyValueStore


3. Launch project:

        go run main.go


## Swagger usage to test API

You can use included Swagger UI to test API.

1. Launch the project.
2. Open the browser and go to the address `http://localhost:8080/swagger/index.html` (Port can be changed in config.yaml).
3. Use interface Swagger UI for requests to API.

## How it works

1. HashMap for Storage: 
The core data structure for storing key-value pairs is a map. 
This structure provides O(1) average time complexity for all operations (Set/Get/Delete).
2. TTL Handling with a PriorityQueue:
To manage the Time To Live (TTL) for each key-value pair, a priority queue is utilized. 
This priority queue is sorted based on the expiration times of the items, with the nearest expiration time at the top of the queue.
