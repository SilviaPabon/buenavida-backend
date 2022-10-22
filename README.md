# Buenavida back-end

## Create .env file

Create a `.env` file inside `api/` folder. You should specify the following values:

```
MONGO_USER=youruser
MONGO_PASSWORD=yourpassword
MONGO_HOST=yourhost
MONGO_PORT=yourport
```

If you want to use the provided docker-compose file as development environment, your `MONGO_USER` and `MONGO_PASSWORD` are both `admin`.

## Build docker images

Create api / golang image

```
docker build -t buenavida/api .
```

Build docker-compose file

```
docker-compose build
```

Run docker-compose file

```
docker-compose up
```

You should see some MongoDB logging messages on your `docker-compose` console and some message similar to: 

```
🟩 Connected to mongo 🟩
```

Otherwise, you should see an error message. 
