""" Initialise and run the grpc merchant service """
from concurrent import futures
import grpc

from lib.service.utils import get_value, gen_creds
from lib.service.psql import get_engine

from service.merchant import SERVICE_CONFIG, log
from service.merchant.impl import MerchantService

import lib.proto.merchant.merchant_pb2_grpc as unimplemented_merchant_service


def serve():
    """ Initialise and run a grpc server for the merchant service """
    ip = get_value(SERVICE_CONFIG, 'network::address')
    port = get_value(SERVICE_CONFIG, 'network::port')
    addr = '%s:%d' % (ip, port)
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    creds = gen_creds(SERVICE_CONFIG)
    if creds is not None:
        server.add_secure_port(addr, creds)
    else:
        server.add_insecure_port(addr)
    rdb_engine_conf = get_value(SERVICE_CONFIG, 'rdb_engine', default_value={})
    db_engine = get_engine(**rdb_engine_conf)
    unimplemented_merchant_service.add_MerchantServiceServicer_to_server(MerchantService(db_engine=db_engine), server)
    server.start()
    log.info('Listening on %s', addr)
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
