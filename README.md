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
  --awsec2-*** \
  my-docker-machine
```

## Params

```bash
--awsec2-access-key                                                                                  AWS Access Key [$AWS_ACCESS_KEY_ID]
--awsec2-ami                                                                                         AWS machine image [$AWS_AMI]
--awsec2-block-duration-minutes "0"                                                                  AWS spot instance duration in minutes (60, 120, 180, 240, 300, or 360)
--awsec2-device-name "/dev/sda1"                                                                     AWS root device name [$AWS_DEVICE_NAME]
--awsec2-endpoint                                                                                    Optional endpoint URL (hostname only or fully qualified URI) [$AWS_ENDPOINT]
--awsec2-iam-instance-profile                                                                        AWS IAM Instance Profile [$AWS_INSTANCE_PROFILE]
--awsec2-insecure-transport                                                                          Disable SSL when sending requests [$AWS_INSECURE_TRANSPORT]
--awsec2-instance-type "t2.micro"                                                                    AWS instance type [$AWS_INSTANCE_TYPE]
--awsec2-keypair-name                                                                                AWS keypair to use; requires --awsec2-ssh-keypath [$AWS_KEYPAIR_NAME]
--awsec2-monitoring                                                                                  Set this flag to enable CloudWatch monitoring
--awsec2-open-port [--awsec2-open-port option --awsec2-open-port option]                             Make the specified port number accessible from the Internet
--awsec2-private-address-only                                                                        Only use a private IP address
--awsec2-region "us-east-1"                                                                          AWS region [$AWS_DEFAULT_REGION]
--awsec2-request-spot-instance                                                                       Set this flag to request spot instance
--awsec2-retries "5"                                                                                 Set retry count for recoverable failures (use -1 to disable)
--awsec2-root-size "16"                                                                              AWS root disk size (in GB) [$AWS_ROOT_SIZE]
--awsec2-secret-key                                                                                  AWS Secret Key [$AWS_SECRET_ACCESS_KEY]
--awsec2-security-group [--awsec2-security-group option --awsec2-security-group option]              AWS VPC security group [$AWS_SECURITY_GROUP]
--awsec2-security-group-readonly                                                                     Skip adding default rules to security groups [$AWS_SECURITY_GROUP_READONLY]
--awsec2-session-token                                                                               AWS Session Token [$AWS_SESSION_TOKEN]
--awsec2-spot-price "0.50"                                                                           AWS spot instance bid price (in dollar)
--awsec2-ssh-keypath                                                                                 SSH Key for Instance [$AWS_SSH_KEYPATH]
--awsec2-ssh-port "22"                                                                               SSH port [$AWS_SSH_PORT]
--awsec2-ssh-user "ubuntu"                                                                           SSH username [$AWS_SSH_USER]
--awsec2-subnet-id                                                                                   AWS VPC subnet id [$AWS_SUBNET_ID]
--awsec2-tags                                                                                        AWS Tags (e.g. key1,value1,key2,value2) [$AWS_TAGS]
--awsec2-use-ebs-optimized-instance                                                                  Create an EBS optimized instance
--awsec2-use-private-address                                                                         Force the usage of private IP address
--awsec2-userdata                                                                                    path to file with cloud-init user data [$AWS_USERDATA]
--awsec2-userdata-base64                                                                             Cloud-init User data Base64 (ignored if awsec2-userdata is defined) [$AWS_USERDATA_BASE64]
--awsec2-volume-type "gp2"                                                                           Amazon EBS volume type [$AWS_VOLUME_TYPE]
--awsec2-vpc-id                                                                                      AWS VPC id [$AWS_VPC_ID]
--awsec2-zone "a"                                                                                    AWS zone for instance (i.e. a,b,c,d,e) [$AWS_ZONE]
```

## Acknowledgement

The driver is originally written by [Docker Machine Team](https://github.com/docker/machine)
