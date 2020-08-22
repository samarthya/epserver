# Project Probe
In this project I will try to create a simple HTTP endpoint that will be hosted as a service that can respond to end points
- /hello
- /check

I will be using this for livenessProbe example.

## Docker build

```bash
docker build -t samarthya/epserver .
```

## /hello
Hello end point will return a response as under

```json
{
    "message": "Hello",
    "stamp": "timestamp"
}
```

## /check
Will add details later TBD

# References
I am using the [image](https://hub.docker.com/_/golang)