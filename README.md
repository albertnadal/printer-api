## printer-api

### Make the printer-api Docker image.

```
docker build -t printer-api:latest .
```

### Run a MQTT broker Docker container from an emqx/emqx image.

```
docker run -d --name emqx -p 1883:1883 -p 8081:8081 -p 8083:8083 -p 8883:8883 -p 8084:8084 -p 18083:18083 emqx/emqx
```

### Run a printer-api Docker container from the image printer-api:latest

```
docker run -d -p 8080:8080 --name printer-api printer-api:latest
```
