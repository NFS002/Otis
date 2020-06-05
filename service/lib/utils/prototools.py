"""
Utility functions to convert a manipulate protobuf messages
-*- coding:utf-8 -*-
"""
import google.protobuf.json_format as google_pb_tools


def message_to_dict(message):
    return google_pb_tools.MessageToDict(message)
