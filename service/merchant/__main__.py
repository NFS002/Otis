""" The grpc merchant service which services CRUD requests on merchant objects from the api-gateway """
import os
from concurrent import futures
import logging

import grpc

import proto.merchant.merchant_pb2 as merchant_service_pb2
import proto.merchant.merchant_pb2_grpc as merchant_service_pb2_grpc


class MerchantService(merchant_service_pb2_grpc.MerchantServiceServicer):
    """Provides methods that implement functionality of the grpc merchant service. """

    def CreateGeneralMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant

        print("called create general", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()

    def CreatePartnerMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called create partner", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()

    def GetGeneralMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called get general", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()

    def GetPartnerMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called get partner", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()

    def DeleteGeneralMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called delete general", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()

    def DeletePartnerMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called delete partner", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()


def serve():
    address = os.environ["OTIS_SERVICE_MERCHANT_ADDRESS"]
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    merchant_service_pb2_grpc.add_MerchantServiceServicer_to_server(
        MerchantService(), server)
    server.add_insecure_port(address)
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
