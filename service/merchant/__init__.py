"""The grpc merchant service"""
from lib.service.utils import get_logger, get_value
from service.merchant.service_config import SERVICE_CONFIG

log = get_logger(get_value(SERVICE_CONFIG, 'logs'), 'merchant-service')
