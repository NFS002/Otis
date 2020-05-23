""" Initialise and run the grpc merchant service """
from concurrent import futures
import grpc

from service.lib.utils import get_value, gen_creds
from service.merchant import SERVICE_CONFIG, log
from impl import MerchantService
import proto.merchant.merchant_pb2_grpc as merchant_service_pb2_grpc


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
    merchant_service_pb2_grpc.add_MerchantServiceServicer_to_server(MerchantService(), server)
    server.start()
    log.info('Listening on %s', addr)
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
