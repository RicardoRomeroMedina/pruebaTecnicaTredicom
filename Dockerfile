FROM golang:latest AS build
WORKDIR /github.com/RicardoRomeroMedina/pruebaTecnicaTredicom
COPY . .
# COPY go.mod go.sum ./
# RUN go mod download
# COPY *.go ./
# RUN go get -d -v
RUN go build -o /github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/main.go
# RUN go build -v
# RUN go install
# EXPOSE 8080
ENTRYPOINT ["/github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/main.go"]

# RUN go get -d -v
# RUN go build -v