#!/bin/bash

# Run grpcurl to get the UUID
UUID=$(./grpcurl -d '{"code_archive": "Ut ut adipisicing sint"}' -plaintext -proto checkerSystem.proto 127.0.0.1:4000 CheckerSystem.CheckerSystem.RegisterChecker | ./jq -r '.id')

# Export the UUID as an environment variable
export UUID

# Run the application
./application