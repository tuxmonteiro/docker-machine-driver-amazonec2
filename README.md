# Docker Machine AWS EC2 Driver

Docker Machine AWS EC2 Driver is a driver for [Docker Machine](https://docs.docker.com/machine/).
It allows to create Docker hosts on [Amazon EC2](https://aws.amazon.com/pt/ec2/).

## Requirements

* [Docker Machine](https://docs.docker.com/machine/) 0.16.0 or later

## Installation

Download the binary from follwing link and put it within your PATH (ex. `/usr/local/bin`)

https://github.com/tuxmonteiro/docker-machine-driver-amazonec2/releases/latest

## Usage

```bash
docker-machine create -d amazonec2 \
  --amazonec2-*** \
  my-docker-machine
```

## Params

```bash
--amazonec2-access-key                                                                                  AWS Access Key [$AWS_ACCESS_KEY_ID]
--amazonec2-ami                                                                                         AWS machine image [$AWS_AMI]
--amazonec2-block-duration-minutes "0"                                                                  AWS spot instance duration in minutes (60, 120, 180, 240, 300, or 360)
--amazonec2-device-name "/dev/sda1"                                                                     AWS root device name [$AWS_DEVICE_NAME]
--amazonec2-endpoint                                                                                    Optional endpoint URL (hostname only or fully qualified URI) [$AWS_ENDPOINT]
--amazonec2-iam-instance-profile                                                                        AWS IAM Instance Profile [$AWS_INSTANCE_PROFILE]
--amazonec2-insecure-transport                                                                          Disable SSL when sending requests [$AWS_INSECURE_TRANSPORT]
--amazonec2-instance-type "t2.micro"                                                                    AWS instance type [$AWS_INSTANCE_TYPE]
--amazonec2-keypair-name                                                                                AWS keypair to use; requires --amazonec2-ssh-keypath [$AWS_KEYPAIR_NAME]
--amazonec2-monitoring                                                                                  Set this flag to enable CloudWatch monitoring
--amazonec2-open-port [--amazonec2-open-port option --amazonec2-open-port option]                             Make the specified port number accessible from the Internet
--amazonec2-private-address-only                                                                        Only use a private IP address
--amazonec2-region "us-east-1"                                                                          AWS region [$AWS_DEFAULT_REGION]
--amazonec2-request-spot-instance                                                                       Set this flag to request spot instance
--amazonec2-retries "5"                                                                                 Set retry count for recoverable failures (use -1 to disable)
--amazonec2-root-size "16"                                                                              AWS root disk size (in GB) [$AWS_ROOT_SIZE]
--amazonec2-secret-key                                                                                  AWS Secret Key [$AWS_SECRET_ACCESS_KEY]
--amazonec2-security-group [--amazonec2-security-group option --amazonec2-security-group option]              AWS VPC security group [$AWS_SECURITY_GROUP]
--amazonec2-security-group-readonly                                                                     Skip adding default rules to security groups [$AWS_SECURITY_GROUP_READONLY]
--amazonec2-session-token                                                                               AWS Session Token [$AWS_SESSION_TOKEN]
--amazonec2-spot-price "0.50"                                                                           AWS spot instance bid price (in dollar)
--amazonec2-ssh-keypath                                                                                 SSH Key for Instance [$AWS_SSH_KEYPATH]
--amazonec2-ssh-port "22"                                                                               SSH port [$AWS_SSH_PORT]
--amazonec2-ssh-user "ubuntu"                                                                           SSH username [$AWS_SSH_USER]
--amazonec2-subnet-id                                                                                   AWS VPC subnet id [$AWS_SUBNET_ID]
--amazonec2-tags                                                                                        AWS Tags (e.g. key1,value1,key2,value2) [$AWS_TAGS]
--amazonec2-use-ebs-optimized-instance                                                                  Create an EBS optimized instance
--amazonec2-use-private-address                                                                         Force the usage of private IP address
--amazonec2-userdata                                                                                    path to file with cloud-init user data [$AWS_USERDATA]
--amazonec2-userdata-base64                                                                             Cloud-init User data Base64 (ignored if amazonec2-userdata is defined) [$AWS_USERDATA_BASE64]
--amazonec2-volume-type "gp2"                                                                           Amazon EBS volume type [$AWS_VOLUME_TYPE]
--amazonec2-vpc-id                                                                                      AWS VPC id [$AWS_VPC_ID]
--amazonec2-zone "a"                                                                                    AWS zone for instance (i.e. a,b,c,d,e) [$AWS_ZONE]
```

## Acknowledgement

The driver is originally written by [Docker Machine Team](https://github.com/docker/machine)
