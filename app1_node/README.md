# Node server

Simple node server that does the following
`curl localhost:3000` --> returns Hello World!
`curl localhost:3000/eth` --> returns a new ethereum address!

Spin it up by running the following

```bash
docker build -t app1:node -f Dockerfile .
docker run app1:node
```
