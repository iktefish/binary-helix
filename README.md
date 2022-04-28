# Docker Container

## Compute Node

- Pull **Ubuntu** docker image:
```sh
$ docker pull golang
```

- Create container from image:
```sh
$ docker run -dit -p 4040:4040 --name binary-helix_c1 --hostname binary-helix_c1 golang
```

- Enter container:
```sh
$ docker exec -it binary-helix_c1 bash 
```

- Contents of `/etc/hosts`:
```
127.0.0.1	localhost
::1	localhost ip6-localhost ip6-loopback
fe00::0	ip6-localnet
ff00::0	ip6-mcastprefix
ff02::1	ip6-allnodes
ff02::2	ip6-allrouters
172.17.0.2	binary-helix_c1
```

## MongoDB

- Pull MongoDB image:
```sh
$ docker pull mongo
```

- Create container from image named `binary-helix_db`:
```sh
$ docker run -d -p 27017:27017 -v $PROJECT_ROOT/db:/data/db --name binary-helix_db --hostname binary-helix_db mongo:latest
```

- Enter container:
```sh
$ docker exec -it binary-helix_c1 bash 
```

- Contents of `/etc/hosts`:
```
127.0.0.1	localhost
::1	localhost ip6-localhost ip6-loopback
fe00::0	ip6-localnet
ff00::0	ip6-mcastprefix
ff02::1	ip6-allnodes
ff02::2	ip6-allrouters
172.17.0.3	binary-helix_db
```
