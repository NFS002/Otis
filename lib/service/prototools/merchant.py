"""
Functions to manipulate gRPC Merchant messages
"""
from lib.types.partnermerchant.proto import partnermerchant_pb2 as partnermerchant
from lib.types.generalmerchant.proto import generalmerchant_pb2 as generalmerchant

from lib.proto.merchant.merchant_pb2 import MerchantsResponse

from lib.service.psql.schemas.merchant import Merchant


def get_merchants_response(general_merchants=None, partner_merchants=None, executed=True):
    return MerchantsResponse(generalMerchants=general_merchants, partnerMerchants=partner_merchants, executed=executed)


def get_merchants_response_from_orm_merchants(orm_merchants, executed=True):
    partner_merchants = []
    for orm_merchant in orm_merchants:
        orm_merchant_dict = orm_merchant.to_dict()
        print(orm_merchant_dict)
        partner_merchant = partnermerchant.PartnerMerchant(**orm_merchant_dict)
        partner_merchants.append(partner_merchant)
    return get_merchants_response(partner_merchants=partner_merchants, executed=executed)


def dict_to_general_merchant(general_merchant_dict):
    return generalmerchant.GeneralMerchant(**general_merchant_dict)


def dict_to_partner_merchant(partner_merchant_dict):
    return partnermerchant.PartnerMerchant(**partner_merchant_dict)


def dict_to_orm_merchant(merchant_dict):
    return Merchant(**merchant_dict)
