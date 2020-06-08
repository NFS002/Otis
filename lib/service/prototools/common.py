""" Utility functions common to all gRPC services"""
from datetime import datetime
import google.protobuf.json_format as google_pb_tools


def message_to_dict(message):
    return google_pb_tools.MessageToDict(message)


def get_property_from_message(message, prop, default_value=None):
    message_dict = message_to_dict(message)
    return get_property_from_message_dict(message_dict, prop, default_value=default_value)


def get_property_from_message_dict(message_dict, prop, default_value=None):
    return prop if prop in message_dict else default_value


def try_get_property_or_default_null(obj, prop, default_null='Null', date_format='%c'):
    value = getattr(obj, prop, None)
    if value is None:
        return default_null
    if isinstance(value, datetime):
        return value.strftime(date_format)
    return value
