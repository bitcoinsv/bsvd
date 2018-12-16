# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

LABEL maintainer="Josh Ellithorpe <quest@mac.com>"

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/bitcoinsv/bsvd

# Switch to the correct working directory.
WORKDIR /go/src/github.com/bitcoinsv/bsvd

# Restore vendored packages.
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# Build the code and the cli client.
RUN go install .
RUN go install ./cmd/bsvctl

# Symlink the config to /root/.bsvd/bsvd.conf
# so bsvctl requires fewer flags.
RUN mkdir -p /root/.bsvd
RUN ln -s /data/bsvd.conf /root/.bsvd/bsvd.conf

# Create the data volume.
VOLUME ["/data"]

# Set the start command. This starts bsvd with
# flags to save the blockchain data and the
# config on a docker volume.
ENTRYPOINT ["bsvd", "--addrindex", "--txindex", "-b", "/data", "-C", "/data/bsvd.conf"]

# Document that the service listens on port 8333.
EXPOSE 8333
