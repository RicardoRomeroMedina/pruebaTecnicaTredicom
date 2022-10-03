#Este es el codigo necesario para poder levantar el contenedor de docker
#Para levantar el contenedor porfavor use el comando docker compose up

FROM golang:latest AS build
WORKDIR /github.com/RicardoRomeroMedina/pruebaTecnicaTredicom
COPY . .
RUN go build -o /github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/main.go
EXPOSE 8080
ENTRYPOINT ["/github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/main.go"]