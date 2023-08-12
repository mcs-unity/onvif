![mcs logo](https://github.com/mcs-unity/go_ocpp/blob/main/resources/logo.png)
<br/><br/>
[![APM](https://img.shields.io/apm/l/vim-mode)](https://github.com/mcs-unity/go_ocpp/blob/main/LICENSE)

## Introduction

A project that was aimed at myself where i practiced golang programing. The goal was to add an alternative option to
C++ it needed to be simple and easy to use. Any one who wishes to read contribute or review the code and provide
suggestion is welcome to do so.

<!--ts-->

## Table of Contents

- [Getting started](#getting-started)
- [Prerequisites](#prerequisites)
- [Built With](#built-with)
- [Authors](#authors)
- [License](#license)
- [Acknowledgments](#acknowledgments)

<!--te-->

### Getting started

### Prerequisites

The get everything running smoothy without any hiccups will requires some tools they might
come preinstalled with your OS and if not I have listed the tools you will need to run service
on your machine.

- [golang](https://golang.org/) - golang

And then lastly a text editor of your choosing I recommend VS code but any text editor will do.

- [Visual Studio Code](https://code.visualstudio.com/) - Code Editor

### Standards

#### RFC

For more information as to how the application works please visit the following links to
respective RFC standard with more information.

| RFC  |     Name      | url                                                       |
| ---- | :-----------: | --------------------------------------------------------- |
| 2616 |   HTTP 1.1    | [RFC 2616](https://tools.ietf.org/html/rfc2616)           |
| 2818 | HTTP over TLS | [RFC 2818](https://tools.ietf.org/html/rfc2818)           |
| 3339 |  Timestamps   | [RFC 3339](https://tools.ietf.org/html/rfc3339)           |
| 3986 |      URI      | [RFC 3986](https://tools.ietf.org/html/rfc3986)           |
| 4122 |     UUID      | [RFC 4122](https://tools.ietf.org/html/rfc4122)           |
| 6455 |   Websocket   | [RFC 6455](https://tools.ietf.org/html/rfc6455)           |
| 7159 |     JSON      | [RFC 7159](https://tools.ietf.org/html/rfc7159)           |
| 5364 |      XML      | [RFC 5364](https://datatracker.ietf.org/doc/html/rfc5364) |
| 7523 |     SOAP      | [RFC 7523](https://datatracker.ietf.org/doc/html/rfc4227) |
| 7616 |  Digest Auth  | [RFC 7616](https://datatracker.ietf.org/doc/html/rfc7616) |

### Development

Continuing development on this service will require a couple of settings to be made
for it to function properly.

#### Enable TLS

To enable TLS you will need to import a certificate(.crt) file and a private key
both file should be added to the root folder. use the command bellow to generate
self signed certificates in linux environment.

```
openssl req -newkey rsa:4096 \
            -x509 \
            -sha256 \
            -days 3650 \
            -nodes \
            -out server.crt \
            -keyout server.key
```

Keep in mind that openssl library must be installed in your system.

#### Editor Config

Make sure that you install Editor Config for your IDE before modifying the project and
pushing github. This is to make it easier for all to read and modify your code.

### Deployment

We make use of makefile in this project so deploying the service has been simplified
allowing individuals with close to no knowledge the ability to deploy the service.

#### Docker

- Prerequisites
  - [make](https://www.gnu.org/software/make/) - makefile

##### Configure variables

before you run the service make sure to go to the dockerfile and modify the
"ENV" variables to your specification then go to the next step Build and run.

```
Variables

```

##### Build and run

to deploy the service using docker is easy cake all you need to do is make sure
that you have make installed. Once installed simply run "make build_docker" and
then "make run_docker" and this will run the image.

### Built With

The practice golang code was built with the following tools on GNU/Linux with Fedora 32
as distribution.

- [Visual Studio Code](https://code.visualstudio.com/) - Code Editor
- [Golang](https://golang.org/) - golang
- [GIT](https://git-scm.com/) - Version Control
- [Github](https://github.com/) - Version Control provider
- [Docker](https://www.docker.com/) - Docker container
- [Editor config](https://editorconfig.org/) - Editor Config
- [Mac os Ventura](https://www.apple.com/th/macos/ventura/) - Mac os ventura

### Authors

- **Nasar Eddaoui** - _Initial work_ - [Nasar Eddaoui](https://github.com/Nasar165)

See also the list of [contributors](https://github.com/mcs-unity/onvif/graphs/contributors) who participated in this project.

### License

This project is licensed under the GNU GENERAL PUBLIC V3 - see the [LICENSE](LICENSE) file for details

[![Go](https://github.com/mcs-unity/onvif/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/mcs-unity/onvif/actions/workflows/go.yml)
![GitHub](https://img.shields.io/github/license/mcs-unity/onvif)

