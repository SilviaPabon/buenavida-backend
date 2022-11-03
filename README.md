# Buenavida back-end

## Create .env file

Create a `.env` file inside `api/` folder. You should specify the following values:

```
MONGO_USER=youruser
MONGO_PASSWORD=yourpassword
MONGO_HOST=yourhost
MONGO_PORT=yourport
PG_USER=youruser
PG_PASSWORD=yourpassword
PG_HOST=yourhost
PG_PORT=yourport
REDIS_HOST=yourhost
REDIS_PORT=yourport
REDIS_DATABASE=yourdatabase
JWT_KEY=yoursecret
```

If you want to use the provided docker-compose file as development environment, your `MONGO_USER`, `MONGO_PASSWORD`, `PG_USER` and `PG_PASSWORD` are all `admin` and all the hosts are `localhost`.

**Recommended**: Use some kind of hash / MD5 as JWT_KEY.

## Build docker images

Build docker-compose file

```
docker-compose build
```

Run docker-compose file

```
docker-compose up
```

You should see the following messages on `buenavida-api` docker console: 

```
🟩 Connected to mongo 🟩
```

```
🐘 Connected to postgresSQL
```

The **first time** you start the docker-compose file, you probably will see a postgres connection error on `buenavida-api` docker console. Follow this steps to fix it:

1. Create an interactive shell into postgres container: 

```
docker exec -it $(docker ps --filter NAME=pg_database --format "{{.ID}}") /bin/bash
```

2. Move to files folder:

```
cd files
```

3. Execute init.sql script: 

```
psql -U admin -a -f ./init.sql
```

4. Finally, restart the docker-compose (`Ctrol + C`, `docker-compose up`).

## Start development server

1. Install air package: 

Read [here](https://github.com/cosmtrek/air) for more information or troubleshooting.

```
go install github.com/cosmtrek/air@latest
```

2. Run listen script: 


From /api folder:

```
./listen.sh

```
