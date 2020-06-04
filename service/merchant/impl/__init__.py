""" The gRPC merchant service implementation. CRUD requests on merchant objects. """
from functools import partial

from service.merchant import log as _log
from service.lib.utils import default_uncurried_logging_wrapper as _default_uncurried_logging_wrapper
from service.lib.psql import mock_psql_engine
from service.lib.psql.schemas.merchant import Merchant

import proto.merchant.merchant_pb2 as merchant_service_pb2
import proto.merchant.merchant_pb2_grpc as merchant_service_pb2_grpc

default_logging_wrapper = partial(_default_uncurried_logging_wrapper, log=_log,
                                  default_return_value=merchant_service_pb2.MerchantsResponse())


class MerchantService(merchant_service_pb2_grpc.MerchantServiceServicer):
    """ Class that provides methods that implement functionality of the gRPC merchant service. """

    def __init__(self, db_engine=mock_psql_engine()):
        self.db_engine = db_engine

    @default_logging_wrapper
    def CreateGeneralMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called create general", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()

    @default_logging_wrapper
    def CreatePartnerMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called create partner", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()

    @default_logging_wrapper
    def GetGeneralMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called get general", "->", request, "<-")
        print(Merchant.metadata.create_all(self.db_engine, checkfirst=False))
        return merchant_service_pb2.MerchantsResponse()

    @default_logging_wrapper
    def GetPartnerMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called get partner", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()

    @default_logging_wrapper
    def DeleteGeneralMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called delete general", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()

    @default_logging_wrapper
    def DeletePartnerMerchant(self, request, unused_context):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        print("called delete partner", "->", request, "<-")
        return merchant_service_pb2.MerchantsResponse()
