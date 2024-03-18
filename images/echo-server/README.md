docker build -f images/echo-server/Dockerfile -t docker.io/laurentsimon/echo-server images/echo-server/
docker run --network=host laurentsimon/echo-server:latest
curl -s 127.0.0.1:8000/ -d "some-data"