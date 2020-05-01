#!/usr/bin/env bash
set -e

AWS_ACCOUNT_ID=""
DOMAIN=""
CONCEALED_EMAIL_PREFIX=""
RECEIVING_EMAIL=""

while getopts "a:d:c:r:" opt; do
    case "$opt" in
        a)
            AWS_ACCOUNT_ID="$OPTARG"
            ;;
        d)
            DOMAIN="$OPTARG"
            ;;
        c)
            CONCEALED_EMAIL_PREFIX="$OPTARG"
            ;;
        r)
            RECEIVING_EMAIL="$OPTARG"
            ;;
        *)
            echo "Unknown argument"
            exit 1
            ;;
    esac
done

if [[ -z "${AWS_ACCOUNT_ID}" || -z "${DOMAIN}" || -z "${CONCEALED_EMAIL_PREFIX}" || -z "${RECEIVING_EMAIL}" ]]; then
    echo "All the arguments are required"
    exit 2
fi

pushd ./iac/common/

echo "Terraform the common infrastructure"
terraform init
terraform apply -auto-approve

popd

GIT_REF=$(git rev-parse --short HEAD)
DOCKER_IMAGE_REPOSITORY="${AWS_ACCOUNT_ID}.dkr.ecr.us-east-1.amazonaws.com/email-conceal-forwarder"
DOCKER_IMAGE="${DOCKER_IMAGE_REPOSITORY}:${GIT_REF}"

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin "${DOCKER_IMAGE_REPOSITORY}"

docker build -f ForwarderDockerfile -t email-conceal-forwarder .

docker tag email-conceal-forwarder:latest "${DOCKER_IMAGE}"

docker push "${DOCKER_IMAGE}"

pushd ./iac/environments/dev/

echo "Terraform the dev infrastructure"
terraform init
terraform apply -auto-approve -var docker_image="${DOCKER_IMAGE}" -var domain="${DOMAIN}" -var concealed_email_prefix="${CONCEALED_EMAIL_PREFIX}" -var receiving_email="${RECEIVING_EMAIL}"

popd
