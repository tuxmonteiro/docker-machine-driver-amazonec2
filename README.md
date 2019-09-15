# Docker Machine AWS EC2 Driver

Docker Machine AWS EC2 Driver is a driver for [Docker Machine](https://docs.docker.com/machine/).
It allows to create Docker hosts on [Amazon EC2](https://aws.amazon.com/pt/ec2/).

## Requirements

* [Docker Machine](https://docs.docker.com/machine/) 0.16.0 or later

## Installation

Download the binary from follwing link and put it within your PATH (ex. `/usr/local/bin`)

https://github.com/tuxmonteiro/docker-machine-driver-awsec2/releases/latest

## Usage

```bash
docker-machine create -d awsec2 \

  docker-machine
```

## Acknowledgement

The driver is originally written by [Docker Machine Team](https://github.com/docker/machine)
