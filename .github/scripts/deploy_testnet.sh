#!/bin/bash

BINARY=$1
VALUE="evmos-validator"

fetch_binary () {
      aws ssm send-command \
    --instance-ids $instance \
    --document-name "AWS-RunShellScript" \
    --parameters commands="aws s3 cp://evmosd-binaries/$BINARY /home/ubuntu/go/bin/evmosd"
}

get_instance_id () {
  InstanceIDs=`aws ec2 describe-instances \
      --filters "Name=tag:Name,Values=$VALUE" \
                "Name=instance-state-name,Values=running" \
      --output json \
      --query 'Reservations[*].Instances[*].{InstanceId:InstanceId}'`   
  echo $InstanceIDs 
}

# Discover Testnet Nodes - Abort if node count is zero
Instances=`get_instance_id`

Length=`echo $Instances | jq '. | length'`

if [[ $Length -eq 0 ]] ; then
  echo "No validator nodes available. Please scale the validators' cluster"
  exit 1
else
  echo "Proceeding with deployment"
  for instance in `echo $Instances | jq -r ' .[] | .[]."InstanceId"'`;
  do 
    # Copy evmosd binary to testnet nodes
    fetch_binary
  done
fi

