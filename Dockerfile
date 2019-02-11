FROM golang:latest
COPY hostings.json ./
COPY main.go ./ 
RUN go get github.com/gorilla/mux
EXPOSE 8000
CMD go run main.go FOREGROUND

