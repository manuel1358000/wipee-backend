#!/bin/bash

# Name of the LocalStack container
CONTAINER_NAME=wipee-localstack

# Stop the container if it's already running
echo "Stopping container if it's already running..."
docker kill $CONTAINER_NAME

# Run LocalStack in Docker
echo "Starting LocalStack..."
docker run --rm -it -e DYNAMODB_SHARE_DB=1 -e AWS_CBOR_DISABLE=true -p 4566:4566 -d --name $CONTAINER_NAME localstack/localstack:0.14.4
echo 'LocalStack is running and listening on port 4566'
echo "To stop LocalStack, run: docker kill $CONTAINER_NAME"

# Wait for LocalStack to fully initialize
echo 'Waiting 20 seconds for LocalStack to start...'
sleep 20

# Configure DynamoDB resources
echo 'Setting up resources...'
echo '***************************************'
echo 'Creating DynamoDB tables...'
echo '***************************************'

# Create DynamoDB table for main CRUD
echo 'Creating DynamoDB table for Wipee...'
aws --endpoint-url=http://localhost:4566 dynamodb create-table --table-name UserProfileTable \
  --attribute-definitions AttributeName=PK,AttributeType=S AttributeName=SK,AttributeType=S \
  --key-schema AttributeName=PK,KeyType=HASH AttributeName=SK,KeyType=RANGE \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
echo 'UserProfileTable created.'

echo '***************************************'
echo 'Table creation completed.'


# Script completion message
echo 'LocalStack is ready. To stop it, run: docker kill '$CONTAINER_NAME
