""" Helper functions to connect and query the psql database """
import os

import sqlalchemy
from sqlalchemy.dialects import postgresql
from sqlalchemy.engine.url import make_url

CONN_STRING = None


def get_new_connection_string():
    db_name = os.environ["OTIS_DB_NAME"]
    username = os.environ["OTIS_DB_USER"]
    password = os.environ["OTIS_DB_PASSWORD"]
    endpoint = os.environ["OTIS_DB_ENDPOINT"]
    port = os.environ["OTIS_DB_PORT"]
    unencoded_url = f"postgresql://{username}:{password}@{endpoint}:{port}/{db_name}"
    return make_url(unencoded_url)


def get_connection_string():
    global CONN_STRING
    if CONN_STRING is None:
        CONN_STRING = get_new_connection_string()
    return CONN_STRING


def dump_psql(sql, *unused_multiparams, **unused_params):
    print(sql.compile(dialect=postgresql.dialect()))
    print(";")


def mock_psql_engine():
    return sqlalchemy.create_engine("postgresql://", strategy='mock', executor=dump_psql)


def get_engine(**kwargs):
    """Open a new connection to the database"""
    db_conn_string = get_connection_string()
    return sqlalchemy.create_engine(db_conn_string, **kwargs)
