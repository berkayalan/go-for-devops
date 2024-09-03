## Query data from Redis

##### Go inside the container

docker exec -it golang-redis-cache redis-cli

##### Get the Key

GET keyName

##### List all keys

KEYS *