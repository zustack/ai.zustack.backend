## Endpoints

- Generate image 
```bash
curl -X POST "http://localhost:8080/generate/image" \
     -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQxMDgxMTEsImlhdCI6MTcyMTUxNjExMSwibmJmIjoxNzIxNTE2MTExLCJzdWIiOjN9.fm_VNKAlFyy49EyZDWlJOGi3x_Ti-KBlhOClppoFm0w" \
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

- Get the images
```bash
curl -X GET "http://localhost:8080/get/user/images?cursor=0" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQxMDgxMTEsImlhdCI6MTcyMTUxNjExMSwibmJmIjoxNzIxNTE2MTExLCJzdWIiOjN9.fm_VNKAlFyy49EyZDWlJOGi3x_Ti-KBlhOClppoFm0w" \
| jq
```
