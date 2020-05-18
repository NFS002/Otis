""" Initialise and run the grpc merchant service """
from concurrent import futures
import grpc

from service.lib.utils import get_value
from service.merchant import GLOBAL_CONF, log
from impl import MerchantService
import proto.merchant.merchant_pb2_grpc as merchant_service_pb2_grpc


def serve():
    """ Initialise and run a grpc server for the merchant service """
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    merchant_service_pb2_grpc.add_MerchantServiceServicer_to_server(
        MerchantService(), server)
    ip = get_value(GLOBAL_CONF, 'network::address')
    port = get_value(GLOBAL_CONF, 'network::port')
    addr = '%s:%d' % (ip, port)
    server.add_insecure_port(addr)
    server.start()
    log.info('Listening on %s', addr)
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
