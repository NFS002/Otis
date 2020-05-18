"""The grpc merchant service"""
import os
from service.lib.utils import load_config, get_logger, get_value

d = os.path.dirname(os.path.realpath(__file__))
p = os.path.join(d, './service-config.json')

GLOBAL_CONF = load_config(p)
log = get_logger(get_value(GLOBAL_CONF, 'logs'), d, 'merchant-service')
