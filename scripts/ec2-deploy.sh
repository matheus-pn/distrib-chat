#! /bin/bash

source "./scripts/.env"


echo "UPLOADING app ..."
scp -i $EC2_CERTIFICATE ./distrib-chat "$EC2_USER_HOST:~/distrib-chat"

