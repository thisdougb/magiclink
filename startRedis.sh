# refresh the redis container

docker-compose pull redis && docker-compose up -d --force-recreate --no-deps --build redis
