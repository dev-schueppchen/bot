FROM golang:1.13.1-buster

LABEL maintainer="dev-schueppchen <schuppen@zekro.de>" \
      description="A community Discord bot by the zekro's Dev-Schuppen guild."

WORKDIR /var/bot

ADD . .

RUN go get -v ./...

RUN go mod verify

# When we are adding ldflags, here
# we need to insert the full package
# paths to the fields with the 
# desired string value
RUN go build -v -o ./bin/server \
        -ldflags "" \
        ./cmd/server/main.go

# Add docker flag if we somehow need
# to run the bot differently inside
# a docker container
CMD ./bin/server -docker
