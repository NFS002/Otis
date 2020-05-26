""" Creating loggers, processing log records, etc... """
import sys
import os
import logging
from logging.handlers import RotatingFileHandler
import rollbar
from rollbar.logger import RollbarHandler
from .formatter import LogFormatter


def request_to_dict(request):
    return {"request": str(request)}


def add_rollbar_handler(logger, history_size=3):
    """ Log handler that sends log records to rollbar.com """
    access_token = os.environ["OTIS_ROLLBAR_ACCESS_TOKEN"]
    environ = os.getenv("OTIS_ENV", default='development')
    rollbar.init(access_token, environ)
    rollbar_handler = RollbarHandler(history_size=history_size)
    rollbar_handler.setLevel(logging.ERROR)

    # gather history for DEBUG+ log messages
    rollbar_handler.setHistoryLevel(logging.DEBUG)

    # attach the history handler to the root logger
    logger.addHandler(rollbar_handler)


def add_console_handler(logger, formatter):
    sh = logging.StreamHandler(sys.stdout)
    sh.setFormatter(formatter)
    sh.setLevel(logging.DEBUG)
    logger.addHandler(sh)


def add_file_handler(logger, path, formatter, max_bytes=None, backups=None, level=None):
    if max_bytes is not None and backups is not None:
        fh = RotatingFileHandler(path, maxBytes=max_bytes, backupCount=backups)
    else:
        fh = logging.FileHandler(path)
    if level is not None:
        fh.setLevel(level)
    fh.setFormatter(formatter)
    logger.addHandler(fh)


def get_logger(logging_conf, log_name):
    """ Method return a logger instance with configuration options
        from ./service-config.json """
    log = logging.getLogger(log_name)
    log.propagate = False
    use_meta = 'meta' in logging_conf and logging_conf['meta'] is True
    json_formatter = LogFormatter(meta=use_meta)
    simple_formatter = logging.Formatter('[%(name)s:%(levelname)s] %(asctime)s - %(message)s')
    log.setLevel(logging.DEBUG)

    if logging_conf['console'] is True:
        add_console_handler(log, simple_formatter)

    if logging_conf['rollbar'] is True:
        add_rollbar_handler(log)

    for f in logging_conf['files']:
        filename = f['filename']
        level = f.get('level')
        max_bytes = f.get('maxsize')
        backups = f.get('backups')
        add_file_handler(log, filename, json_formatter, max_bytes=max_bytes, backups=backups, level=level)

    return log
