## Endpoints

- Generate image with stable-diffusion-3-medium
```bash
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{
          "prompt": "a photograph of an astronaut riding a horse",
          "cfg_scale": 5,
          "aspect_ratio": "16:9",
          "seed": 0,
          "steps": 50,
          "negative_prompt": ""
     }' \
http://localhost:8080/image/stable-diffusion-3-medium
```
