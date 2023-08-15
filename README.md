# Instalasi Menggunakan Docker
## Intoduction

Hanya untuk bejar golang sampai docker dan message brocker menggunakan kafka

Install docker pada ubuntu 22.04
```bash
    sudo apt-get update
    sudo apt-get install ca-certificates curl gnupg
    sudo install -m 0755 -d /etc/apt/keyrings
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
    sudo chmod a+r /etc/apt/keyrings/docker.gpg
    echo \
    "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
    "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
    sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

## Install Container
#### Step 1 : Pastikan memiliki akun docker hub
```bash
    https://hub.docker.com/
```

#### Step 2 : Membuat repository pada docker hub 
- Pilih menu repositories 
- Create repository

#### Step 3 : Build Image
Membuat Images
```bash
    docker build -t username_dockerhub/repository_dockerhub:tag_yang_diinginkan .
```

Untuk melihat image
```bash
    docker images
```

Untuk melihat container yang sedang running
```bash
    docker container ls
```

Untuk melihat semua container
```bash
    docker container ls -a
```

#### Step 4 : Push ke docker hub
Push ke docker hub
```bash
    docker push username_dockerhub/repository_dockerhub:tag_yang_diinginkan
```

#### Step 5 : Hapus image
Tujuan menghapus image untuk pull dari docker hub
```bash
    docker image prune -a --force
```

#### Step 6 : Menggunakan docker compose
Contoh docker compose
```bash
version: '3.8'
services:
  backend:
    image: NAMAIMAGE
    container_name: NAMACONTAINER
    restart: always
    ports:
      - 8080:8080
    environment:
      - ENVIRONMENT=DEV
      - DATABASE_PORT=PORT
      - DATABASE_USER=USER
      - DATABASE_PASSWRD=PASSWORD
      - DATABASE_NAME=DATABASE
      - DATABASE_HOSTNAME=1HOST
      - DATABASE_HOST=postgres
      - DATABASE_SSLMODE=disable
      - DATABASE_TIMEZONE=Asia/Jakarta
      - SWAGGER=1
      - PORT_SWAGGER=4000
      - PATH_LOG=/home/arief/log_microservice/other
      - AUTO_DELETE_LOG=7
      - IP_SWAGGER=172.16.2.220
      - MIGRATION=INACTIVE
```

Jika menggunakan docker compose maka akan dipull image yang ada pada dicker hub 
```bash
    docker compose up --build -d
```

Jika ingin rebuild
```bash
    docker compose up --build -d
```

Jika ingin down
```bash
    docker compose down
```


### Perintah Tmabahan
#### Hapus container
```bash
    docker system prune --volumes
```