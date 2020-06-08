""" The gRPC merchant service implementation, which can service CRUD requests on merchant objects. """
from functools import partial
from sqlalchemy.orm import sessionmaker

from service.merchant import log as _log
from lib.service.utils import default_uncurried_logging_wrapper as _default_uncurried_logging_wrapper
from lib.service.psql import mock_psql_engine
from lib.service.psql.schemas.merchant import Merchant

import lib.proto.merchant.merchant_pb2 as merchant_service_pb2
import lib.proto.merchant.merchant_pb2_grpc as merchant_service_pb2_grpc

from lib.service.prototools.merchant import get_merchants_response_from_orm_merchants, dict_to_orm_merchant

default_logging_wrapper = partial(_default_uncurried_logging_wrapper, log=_log,
                                  default_return_value=merchant_service_pb2.MerchantsResponse())


class MerchantService(merchant_service_pb2_grpc.MerchantServiceServicer):
    """ Class that provides methods that implement functionality of the gRPC merchant service. """

    def __init__(self, db_engine=mock_psql_engine()):
        self.db_engine = db_engine
        self.get_session = sessionmaker(bind=self.db_engine)

    @default_logging_wrapper
    def CreateGeneralMerchant(self, request, unused_context, optional_request_dict):
        """ Parse request to dtype format
        Insert request to db as new general merchant """
        session = self.get_session()
        print("called create general", "->", request, "<-")
        print("called create general as dict", "->", optional_request_dict, "<-")
        merchant = dict_to_orm_merchant(optional_request_dict['partnerMerchant'])
        session.add(merchant)
        session.commit()
        session.close()
        return merchant_service_pb2.MerchantsResponse(executed=True)

    @default_logging_wrapper
    def CreatePartnerMerchant(self, request, unused_context, optional_request_dict):
        """ Parse request to dtype format
        Insert request to db as new general merchant """
        # Merchant.metadata.create_all(self.db_engine, checkfirst=False)
        session = self.get_session()
        print("called create partner", "->", request, "<-")
        print("called create partner as dict", "->", optional_request_dict, "<-")
        merchant = dict_to_orm_merchant(optional_request_dict['partnerMerchant'])
        session.add(merchant)
        session.commit()
        session.close()
        return merchant_service_pb2.MerchantsResponse(executed=True)

    @default_logging_wrapper
    def GetGeneralMerchant(self, request, unused_context, optional_request_dict):
        """ Parse request to dtype format
        Insert request to db as new general merchant """
        print("called get general", "->", request, "<-")
        print("called get general as dict", "->", optional_request_dict, "<-")
        session = self.get_session()
        merchants = session.query(Merchant).filter_by(**optional_request_dict).all()
        session.close()
        return get_merchants_response_from_orm_merchants(orm_merchants=merchants)

    @default_logging_wrapper
    def GetPartnerMerchant(self, request, unused_context, optional_request_dict):
        """ Parse request to dtype format
        Insert request to db as new general merchant """
        print("called get partner", "->", request, "<-")
        print("called get partner as dict", "->", optional_request_dict, "<-")
        session = self.get_session()
        merchants = session.query(Merchant).filter_by(**optional_request_dict).all()
        session.close()
        return get_merchants_response_from_orm_merchants(orm_merchants=merchants)

    @default_logging_wrapper
    def DeleteGeneralMerchant(self, request, unused_context, optional_request_dict):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        session = self.get_session()
        print("called delete general", "->", request, "<-")
        print("called delete general as dict", "->", optional_request_dict, "<-")
        session.query(Merchant).filter_by(**optional_request_dict).delete()
        session.commit()
        session.close()
        return merchant_service_pb2.MerchantsResponse(executed=True)

    @default_logging_wrapper
    def DeletePartnerMerchant(self, request, unused_context, optional_request_dict):
        # Parse request to dtype format
        # Insert request to db as new general merchant
        session = self.get_session()
        print("called delete partner", "->", request, "<-")
        print("called delete partner as dict", "->", optional_request_dict, "<-")
        session.query(Merchant).filter_by(**optional_request_dict).delete()
        session.commit()
        session.close()
        return merchant_service_pb2.MerchantsResponse(executed=True)
