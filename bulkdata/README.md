# Insert products data into mongo database

## Requirements

1. Mongo database running (Recommend to use the `docker-compose` file at `../api/` folder)
2. `.env` file created at `bulkdata/` folder

The `.env` file should contain the following fields: 

```
MONGO_USER=yourUser
MONGO_PASSWORD=yourPassword
MONGO_HOST=yourHost
MONGO_PORT=yourPort
```

If you are using the provided `docker-compose` file as development environment, your `MONGO_USER` and `MONGO_PASSWORD` are both `admin`, `MONGO_HOST` is `localhost` and `MONGO_PORT` is `27017`.

3. NodeJS installed on your system.

4. Install npm packages: 

```
npm install --production
```

## Inser data

From `bulkdata/` folder, just run: 

```
npm run bulk
```
