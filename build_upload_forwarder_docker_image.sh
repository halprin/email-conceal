#!/usr/bin/env bash
set -e

AWS_ACCOUNT_ID=$1

GIT_REF=$(git rev-parse --short HEAD)
DOCKER_IMAGE_REPOSITORY="${AWS_ACCOUNT_ID}.dkr.ecr.us-east-1.amazonaws.com/email-conceal-forwarder"
DOCKER_IMAGE="${DOCKER_IMAGE_REPOSITORY}:${GIT_REF}"

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin "${DOCKER_IMAGE_REPOSITORY}"

docker build -f ForwarderDockerfile -t email-conceal-forwarder .

docker tag email-conceal-forwarder:latest "${DOCKER_IMAGE}"

docker push "${DOCKER_IMAGE}"

pushd ./iac/environments/dev/

terraform init
terraform apply -var docker_image="${DOCKER_IMAGE}"

popd
