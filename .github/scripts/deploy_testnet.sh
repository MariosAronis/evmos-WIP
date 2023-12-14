#!/bin/bash

VALUE="evmos-validator"

get_instance_id () {
  InstanceIDs=`aws ec2 describe-instances \
      --filters "Name=tag:Name,Values=$VALUE" \
                "Name=instance-state-name,Values=running" \
      --output json \
      --query 'Reservations[*].Instances[*].{InstanceId:InstanceId}' \
      | jq '.[]'`   
  echo $InstanceIDs 
}

get_instance_id2 () {
  InstanceIDs=`aws ec2 describe-instances \
      --filters "Name=tag:Name,Values=$VALUE" \
                "Name=instance-state-name,Values=running" \
      --output json \
      --query 'Reservations[*].Instances[*].{InstanceId:InstanceId}'`   
  echo $InstanceIDs 
}

# Discover Testnet Nodes
get_instance_id

Instances=`get_instance_id2`

Length=`echo $Instances | jq '. | length'`
echo $Length

if [[ $Length -eq 0 ]] ; then
  echo "No validator nodes available. Please scale the validators' cluster"
  exit 1
else
  echo "Proceeding with deployment"
  exit 0

# Copy evmosd binary to testnet nodes