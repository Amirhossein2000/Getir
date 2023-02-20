## Run the application
```bash
docker compose up -d
```

## Endpoints

| URL        | Method(s)    |
|------------|--------------|
| /in-memory | GET and POST |
| /fetch     | POST         |

## Error Table
| error code | description            |
|------------|------------------------|
| 1          | Bad Request (400)      |
| 2          | Internal Server (500)  |
| 3          | Not Found (404)        |

## API Usage
For api usage import Getir.postman_collection.json which is in docs directory. It includes all of the requests with a few examples. 

## AWS Deployment Guideline
1. Create VPC with Nat
2. Create elastic-cache (redis) on aws
3. Create VPC endpoint for elastic-cache
4. Build docker image and push it to ECR
5. Use aws app runner to create a container from ECR image

## Deployment Domain
For address of the endpoints instead of localhost:8080 you can use `https://z4tdfk5eu2.eu-west-1.awsapprunner.com`.
Check [this](https://z4tdfk5eu2.eu-west-1.awsapprunner.com/in-memory?key=testKey) for example.