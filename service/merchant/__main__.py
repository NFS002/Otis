from concurrent import futures
import time
import math
import logging
import utils

import grpc

import proto.merchant.merchant_pb2 as merchant_service_pb2
import proto.merchant.merchant_pb2_grpc as merchant_service_pb2_grpc


class MerchantService(merchant_service_pb2_grpc.MerchantServiceServicer):
    """Provides methods that implement functionality of the grpc merchant service."""


    def CreateGeneralMerchant(self, request, context):
        # Parse request to dtype format
        print("Looking good boyos")
        # Insert request to db as new general merchant
        return request


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    merchant_service_pb2_grpc.add_MerchantServiceServicer_to_server(
        MerchantService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()