#!/bin/bash
# Compiles all proto files to python

# Requirements:
# pip install grpcio
# pip install grpcio-tools

cd $OTIS_HOME

# Types
python -m grpc_tools.protoc -I $OTIS_HOME lib/types/generalmerchant/proto/generalmerchant.proto --python_out=.
python -m grpc_tools.protoc -I $OTIS_HOME lib/types/partnermerchant/proto/partnermerchant.proto --python_out=.

# Merchant service
python -m grpc_tools.protoc -I $OTIS_HOME lib/proto/merchant/merchant.proto --python_out=. --grpc_python_out=.