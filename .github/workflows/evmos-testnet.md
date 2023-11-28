# Deployment Workflow for an emvmos testnet with multiple validators

This workflow is currently configured to run only on workflow dispatch (mannual trigger)
Executes the following actions:

1. Creates a new release. User selects the target branch and release tag/name
2. Builds the evmosd binaries
3. Uploads build artifacts to action workspace
4. Assumes AWS short lived credentials against certain resources (AWS CodeArtifact, AWS SSM)
    **TO DO**
5. Uploads evmosd binary to AWS CodeArtifact
6. Checks if an evmos private testnet is deployed according to https://github.com/MariosAronis/evmos-testnet.
    - If there is not one, workflow will trigger a TF cloud workspace run to deploy needed resources.
    - Otherwise proceed
7. Remote script execution against validator nodes via AWS SSM (SLCs needed also):
    - download binary from CodeArtifact
    - prepare the testnet config files for each node:
        - either run evmosd testnet init-files on each host and then start evmosd service as a system daemon
        - or use one of the nodes to bootstrap the evmos chain with script localnet-start; create and distributing node configs across the cluster and run each validator as an individual container in each ec2 host 

https://github.com/MariosAronis/evmos-WIP/issues/4
