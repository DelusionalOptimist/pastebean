# pastebean
Implementing a pastebin clone as a microservice architecture.
Written using [go-gin](https://github.com/gin-gonic/gin) and [gorm](https://github.com/go-gorm/gorm) alongwith many other awesome open source Go libaries.
Served using [envoy proxy](https://envoyproxy.io/).
Containerized and orchestrated using [docker and docker-compose](https://www.docker.com/).

### Try it out
Run using:
```bash
docker-compose -f docker-compose.yaml up
```

Create a paste:
```bash
curl -XPOST localhost:8000/create --data-raw '{"name": "paste1", "body": "This is a big paste body."}'
```

Retrieve the above created paste using the returned uuid:
```bash
curl localhost:8000/<16-digit-uuid>
```

### Services
* Create: A service that creates paste.
* Read: A service that reads created paste.

#### Architecture
[pastebean architecture](docs/images/pastebean-1.png)

### Development
* Requirements:
	- make
	- go 1.17
	- docker
	- docker-compose

* Build and run on your host and be ready for being haunted by cloud native evangelists:
	```bash
	make pastebean-dirty
	```

* Build and run it in containers using docker compose:
	```bash
	make pastebean-dc-dev
	```

### TODO
- [] Blog/tutorial series
- [] More services using different protocols
- [] Dyanmic Envoy configuration and service discovery
- [] More container orchestrators
- [] Service mesh
