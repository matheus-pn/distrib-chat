#! /bin/bash

source "./scripts/.env"


echo "UPLOADING app ..."
# scp -i $EC2_CERTIFICATE /tmp/distrib-chat.zip "$EC2_USER_HOST:~/app/app.zip"
scp -i $EC2_CERTIFICATE  "$EC2_USER_HOST:~/ruby.zip"


# # Extract the contents of the zip file on the server into the ~/app/ directory
# ssh -i "$EC2_CERTIFICATE" "$EC2_USER_HOST" '
#   unzip -qo ~/app/app.zip -d ~/app
#   rm ~/app/app.zip"
# '
