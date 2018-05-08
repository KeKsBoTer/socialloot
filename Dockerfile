FROM library/golang

# Godep for vendoring
RUN go get github.com/tools/godep

# Recompile the standard library without CGO
RUN CGO_ENABLED=0 go install -a std

ENV APP_DIR $GOPATH/src/github.com/KeKsBoTer/socialloot
RUN mkdir -p $APP_DIR

# Set the entrypoint
ENTRYPOINT (cd $APP_DIR && ./socialloot)
ADD . $APP_DIR

RUN cd $APP_DIR && godep restore

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=1 godep go build -ldflags ' -w -s'

EXPOSE 8080
EXPOSE 8088