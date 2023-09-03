#!/bin/sh
set -e

# Author : Nasar Eddaoui
# Copyright (c) MCS Unity Co.,LTD
# Data 18th Mars 2023

# Description
## The following script is used to compile the project

mkdir cert

openssl req -newkey rsa:4096 -x509 \
            -sha256 \
            -days 3650 \
            -nodes \
            -out cert/server.crt \
            -keyout cert/server.key
