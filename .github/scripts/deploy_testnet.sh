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
get_instance_id2

# Copy evmosd binary to testnet nodes