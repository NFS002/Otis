"""Class to format log record to JSON"""
import logging
import threading
import traceback
from datetime import datetime
import psutil
from json_log_formatter import JSONFormatter


class LogFormatter(JSONFormatter):
    """Custom log record formatter to emit JSON records"""

    def __init__(self, meta=False):
        super().__init__()
        self.meta = meta

    def move_to_back(self, extra, key):
        r = extra[key]
        del extra[key]
        extra[key] = r

    def get_thread_info(self):
        thread_info = dict()
        thread_info['id'] = threading.get_ident()
        thread_info['active_count'] = threading.active_count()
        return thread_info

    def get_process_info(self):
        ps_wl = ['pid', 'name', 'username', 'status', 'create_time']
        ps_meta = psutil.Process().as_dict(attrs=ps_wl)
        return ps_meta

    def get_os_info(self):
        memory = psutil.virtual_memory()
        os_wl = ['active', 'free', 'available', 'inactive', 'percent', 'total', 'used']
        os_meta = dict(
            [attr, getattr(memory, attr)] for attr in dir(memory) if (not attr.startswith('_')) and (attr in os_wl))
        return os_meta

    def json_record(self, message: str, extra: dict, record: logging.LogRecord) -> dict:
        """ Format a log record """
        # Include builtins
        extra['message'] = message
        extra['level'] = record.levelname
        extra['name'] = record.name
        extra['time'] = datetime.now().isoformat()

        # Move request filed to after builtins
        if 'request' in extra:
            self.move_to_back(extra, 'request')

        # Move rollbar_uuid field
        if 'rollbar_uuid' in extra:
            self.move_to_back(extra, 'rollbar_uuid')

        # Exception/error info goes last
        if record.exc_info:
            extra['traceback'] = traceback.format_exc()

        if self.meta is True:
            # Add process, os, and thread info
            extra['meta'] = {'thread': self.get_thread_info(), 'ps': self.get_process_info(),
                             'os': self.get_os_info()}

        return extra
