# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/gitlab.com/lukma/sample-go-clean

# Add library dependencies
RUN go get github.com/joho/godotenv
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get gopkg.in/mgo.v2
RUN go get github.com/dgrijalva/jwt-go
RUN go get -u golang.org/x/crypto/bcrypt
RUN go get firebase.google.com/go

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/lukma/sample-go-clean

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/sample-go-clean

# Document that the service listens on port 8080.
EXPOSE 8080