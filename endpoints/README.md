## Endpoints

- Generate image 
```bash
curl -X POST "http://localhost:8080/generate/image" \
     -H "Content-Type: application/json" \
     -d '{"prompt": "a photograph of an astronaut riding a horse in space"}'
```

- Get the images
```bash
curl -X GET "http://localhost:8080/get/images?cursor=0&q=" | jq
```

- Create a new user
```bash
curl -X POST "http://localhost:8080/register" \
     -H "Content-Type: application/json" \
     -d '{"username": "test", "password": "test"}'
```

- Login
```bash
curl -X POST "http://localhost:8080/login" \
     -H "Content-Type: application/json" \
     -d '{"username": "test", "password": "test"}'
```

