# UpdateUser Service

This is the UpdateUser service

Generated with

```
micro new ihomegit/ihome/service/updateUser --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.updateUser
- Type: srv
- Alias: updateUser

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./updateUser-srv
```

Build a docker image
```
make docker
```