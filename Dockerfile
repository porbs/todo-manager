FROM golang

WORKDIR /go/src

# Copy project sources
COPY . /go/src/todo-manager

# Install revel framework
RUN go get -v github.com/revel/cmd/revel

# Install golang mongo driver
RUN go get -v gopkg.in/mgo.v2

# Launch app
CMD revel run todo-manager