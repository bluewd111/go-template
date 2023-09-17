FROM golang:1.21.1-bullseye

WORKDIR /workspace

RUN apt-get update && apt-get install -y \
    git \
    make \
    curl \
    unzip \
    && rm -rf /var/lib/apt/lists/*

# Visual Studio Code Go extension dependencies
# Ctrl/Cmd + Shift + P -> Go: Install/Update Tools
RUN go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/gopls@latest

