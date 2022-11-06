# Buenavida back-end

📜 You can find more documentation on `/docs` folder.

## Setup

To start the project at the first time:

### Create the `.env` files:

From the `api/` folder, create a `.env` file and define the following values (**for development environment using the provided `docker-compose` file):

```
MONGO_USER=admin
MONGO_PASSWORD=admin
MONGO_HOST=localhost
MONGO_PORT=27017
PG_USER=admin
PG_PASSWORD=admin
PG_HOST=localhost
PG_PORT=5432
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=admin
REDIS_DATABASE=0
JWT_KEY={4Z!502lmtA4Guo6shSnaC+(8$bpR4Q+
```

From the `bulkdata/` folder, create a `.env`file and define the following values (**for development environment using the provided `docker-compose` file):

```
MONGO_USER=admin
MONGO_PASSWORD=admin
MONGO_HOST=localhost
MONGO_PORT=27018
```

### Setup postgres database

Start docker containers: 

```bash
docker-compose up
```

The **first time** you start the docker-compose file, run the following commands to create the database and it's tables:

1. Create an interactive shell into the postgres container:

```bash
docker exec -it $(docker ps --filter NAME=pg_database --format "{{.ID}}") /bin/bash
```

2. Execute the init.sql script:

```bash
psql -U admin -a -f files/init.sql
```

3. Stop and restart the container (`Ctrol + C`, `docker-compose up`).

### Start development environment

1. Run the docker containers:

```bash
docker-compose up
```

2. Start the golang api: 

You have to install [air](https://github.com/cosmtrek/air) to complete the next step:

```bash
go install github.com/cosmtrek/air@latest
```

**From `/api` folder:**

```bash
./listen.sh
```
