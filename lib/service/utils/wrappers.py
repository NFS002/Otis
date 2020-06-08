""" Function wrappers for logging and handling exceptions"""
from functools import wraps
from lib.service.prototools.common import message_to_dict


def default_uncurried_logging_wrapper(func, log=None, to_dict=True, default_return_value=None):
    """ Function decorator to wrap gRPC endpoints, log arguments, and handle/log exceptions"""
    @wraps(func)
    def wrap(*args, **kwargs):
        raw_request = str(args[1])
        request_dict = message_to_dict(args[1]) if to_dict else None
        extra_logging_args = {"request": {"raw": raw_request, "optional_request_dict": request_dict}}
        func_name = func.__name__ or func.func_name
        if log is not None:
            log.info(f"Called {func_name}", extra=extra_logging_args)
        return_value = default_return_value
        try:
            return_value = func(*args, **kwargs, optional_request_dict=request_dict)
        except Exception as e:
            if log is not None:
                log.error(f"Error in call to {func_name}: %s", repr(e), extra=extra_logging_args, exc_info=e)
        return return_value
    return wrap
