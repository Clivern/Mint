### Deployment with Docker and Docker Compose

1. Update and upgrade your server packages.

```zsh
$ apt-get update
$ apt-get upgrade
```

2. You might need to reboot if the kernel got upgraded

```zsh
$ reboot
```

3. Install `docker` and `docker-compose`.

```zsh
$ apt-get install git
$ apt-get install docker.io -y
$ systemctl enable docker
$ apt-get install docker-compose -y
```

4. Clone the repository.

```zsh
$ git clone git@gitlab.com:clivern/mint.git
$ cd mint
```

5. Start the docker containers and wait until they are healthy.

```zsh
$ docker-compose up -d
$ docker ps
```
