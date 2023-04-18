<p align="right">
   <a href="./README.md">中文</a> | <strong>English</strong>
</p>

<div align="center">

# Go Relay

_✨ Golang based HTTP relay server, easy to deploy & use ✨_

</div>

<p align="center">
  <a href="https://raw.githubusercontent.com/songquanpeng/go-relay/master/LICENSE">
    <img src="https://img.shields.io/github/license/songquanpeng/go-relay?color=brightgreen" alt="license">
  </a>
  <a href="https://github.com/songquanpeng/go-relay/releases/latest">
    <img src="https://img.shields.io/github/v/release/songquanpeng/go-relay?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://github.com/songquanpeng/go-relay/releases/latest">
    <img src="https://img.shields.io/github/downloads/songquanpeng/go-relay/total?color=brightgreen&include_prereleases" alt="release">
  </a>
  <a href="https://hub.docker.com/repository/docker/justsong/go-relay">
    <img src="https://img.shields.io/docker/pulls/justsong/go-relay?color=brightgreen" alt="docker pull">
  </a>
  <a href="https://goreportcard.com/report/github.com/songquanpeng/go-relay">
  <img src="https://goreportcard.com/badge/github.com/songquanpeng/go-relay" alt="GoReportCard">
  </a>
</p>

## Features
+ [x] Easy to use
+ [x] Token authentication
+ [ ] Support IP whitelist

## Usage
### Server

```bash
# Initialize configuration file
./go-relay init
# Check and save the generated token
cat go-relay.yaml
# Start the server
./go-relay
```

Or deploy using Docker:

```bash
docker run -d --restart always --name go-relay -p 6872:6872 -v /home/ubuntu/data/go-relay:/app justsong/go-relay
```

### Client
When making an HTTP request, replace the host address and port in the request URL with your relay server address and port.

Then add the following fields to the request header:
1. `X-Relay-Token`: Token configured on Go Relay server
2. `X-Relay-Host`: Target address to request
3. `X-Relay-Protocol`: Request protocol, optional, defaults to https

## Flowchart

```mermaid
sequenceDiagram
    participant Client
    participant Relay
    participant Server
    
    Client->>Relay: HTTP Request
    Relay->>Server: Forward HTTP Request
    Server->>Relay: HTTP Response
    Relay->>Client: Forward HTTP Response
```