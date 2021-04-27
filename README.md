# Code Challenge

## How to run the app
First of all, you need to set the environment variables, you can use the default values in the `.env.example` file, just rename it to `.env`.

Now you can start the app by running:

`docker compose up`

---

## API docs

### Upload a file

**Request:**

`POST` `/files`

```bash
curl --request POST \
  --url http://127.0.0.1:8000/files \
  --header 'Content-Type: multipart/form-data' \
  --form file=@<PATH_TO_YOUR_FILE>
```

### List all files

**Request:**

`GET` `/files`

```bash
curl --request GET \
  --url http://127.0.0.1:8000/files
```

### Download a file

**Request:**

`GET` `/files/<id>`

```bash
curl --request GET \
  --url http://127.0.0.1:8000/files/<FILE_ID>
```
