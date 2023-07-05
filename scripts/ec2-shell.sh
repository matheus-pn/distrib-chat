#! /bin/bash

source "./scripts/.env"

ssh -i $EC2_CERTIFICATE $EC2_USER_HOST