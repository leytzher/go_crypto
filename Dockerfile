# Specify the base image for the go app
FROM golang:1.15
# Specify that we now need to execute any commands in this directory
WORKDIR /go/src/github.com/go_crypto
# Copy everything from this project into the filesystem of the container.
COPY . .
# Obtain the package needed to run code. Alternative we can use GO modules.
RUN go get -u github.com/lib/pq
RUN go get -u github.com/gorilla/mux
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/postgres
RUN go get -u github.com/joho/godotenv

# Compile the binary exe for out app.
RUN go build -o main .
# start the application
CMD ["./main"]
