# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from proto.merchant import merchant_pb2 as proto_dot_merchant_dot_merchant__pb2


class MerchantServiceStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateGeneralMerchant = channel.unary_unary(
                '/merchant.MerchantService/CreateGeneralMerchant',
                request_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantRequest.SerializeToString,
                response_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
                )
        self.CreatePartnerMerchant = channel.unary_unary(
                '/merchant.MerchantService/CreatePartnerMerchant',
                request_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantRequest.SerializeToString,
                response_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
                )
        self.GetGeneralMerchant = channel.unary_unary(
                '/merchant.MerchantService/GetGeneralMerchant',
                request_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantQuery.SerializeToString,
                response_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
                )
        self.GetPartnerMerchant = channel.unary_unary(
                '/merchant.MerchantService/GetPartnerMerchant',
                request_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantQuery.SerializeToString,
                response_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
                )
        self.DeleteGeneralMerchant = channel.unary_unary(
                '/merchant.MerchantService/DeleteGeneralMerchant',
                request_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantQuery.SerializeToString,
                response_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
                )
        self.DeletePartnerMerchant = channel.unary_unary(
                '/merchant.MerchantService/DeletePartnerMerchant',
                request_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantQuery.SerializeToString,
                response_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
                )


class MerchantServiceServicer(object):
    """Missing associated documentation comment in .proto file"""

    def CreateGeneralMerchant(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreatePartnerMerchant(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetGeneralMerchant(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPartnerMerchant(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteGeneralMerchant(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeletePartnerMerchant(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MerchantServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateGeneralMerchant': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateGeneralMerchant,
                    request_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantRequest.FromString,
                    response_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.SerializeToString,
            ),
            'CreatePartnerMerchant': grpc.unary_unary_rpc_method_handler(
                    servicer.CreatePartnerMerchant,
                    request_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantRequest.FromString,
                    response_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.SerializeToString,
            ),
            'GetGeneralMerchant': grpc.unary_unary_rpc_method_handler(
                    servicer.GetGeneralMerchant,
                    request_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantQuery.FromString,
                    response_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.SerializeToString,
            ),
            'GetPartnerMerchant': grpc.unary_unary_rpc_method_handler(
                    servicer.GetPartnerMerchant,
                    request_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantQuery.FromString,
                    response_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.SerializeToString,
            ),
            'DeleteGeneralMerchant': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteGeneralMerchant,
                    request_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantQuery.FromString,
                    response_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.SerializeToString,
            ),
            'DeletePartnerMerchant': grpc.unary_unary_rpc_method_handler(
                    servicer.DeletePartnerMerchant,
                    request_deserializer=proto_dot_merchant_dot_merchant__pb2.MerchantQuery.FromString,
                    response_serializer=proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'merchant.MerchantService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class MerchantService(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def CreateGeneralMerchant(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/merchant.MerchantService/CreateGeneralMerchant',
            proto_dot_merchant_dot_merchant__pb2.MerchantRequest.SerializeToString,
            proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreatePartnerMerchant(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/merchant.MerchantService/CreatePartnerMerchant',
            proto_dot_merchant_dot_merchant__pb2.MerchantRequest.SerializeToString,
            proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetGeneralMerchant(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/merchant.MerchantService/GetGeneralMerchant',
            proto_dot_merchant_dot_merchant__pb2.MerchantQuery.SerializeToString,
            proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPartnerMerchant(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/merchant.MerchantService/GetPartnerMerchant',
            proto_dot_merchant_dot_merchant__pb2.MerchantQuery.SerializeToString,
            proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteGeneralMerchant(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/merchant.MerchantService/DeleteGeneralMerchant',
            proto_dot_merchant_dot_merchant__pb2.MerchantQuery.SerializeToString,
            proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeletePartnerMerchant(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/merchant.MerchantService/DeletePartnerMerchant',
            proto_dot_merchant_dot_merchant__pb2.MerchantQuery.SerializeToString,
            proto_dot_merchant_dot_merchant__pb2.MerchantsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
