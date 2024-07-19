## Endpoints

- Generate image 
```bash
curl -X POST "http://localhost:8080/generate/image" \
     -H "Content-Type: application/json" \
     -d '{"prompt": "a photograph of an astronaut riding a horse in space"}'
```

- Get all the images
```bash
curl -X GET "http://localhost:8080/get/images" | jq
```

