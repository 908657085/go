cd src/LocationServer
CGO_ENABLED=0 go get
cd ../../
docker build -t locationserver:V1 .
docker run --rm -it -d -p 8080:8080 locationserver:V1
