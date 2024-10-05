FROM golang:alpine

# setting up the working directory
WORKDIR /app

# copying go modules in working directory and downloading them
COPY go.mod go.sum ./
RUN go mod download

# copying application code from root folder in local to docker root folder
COPY . .

# using go build, build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# port to be used in the container
EXPOSE 8080

# command to run the go application
CMD ["./main"]