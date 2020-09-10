# Documentation for Starting Envoy

## Build the envoy docker image
`do this from the envoy manifest's directory`

    docker build -t envoy:v1 .

## Start the envoy container

 This is telling docker to use the network of the host machine, this will be changed later

    docker run -p 8080:8080 -p 9901:9901 envoy:v1
