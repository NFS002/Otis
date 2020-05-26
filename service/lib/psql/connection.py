""" Helper functions to connect and query the psql database """
from urllib import parse
import os
import sqlalchemy

CONN_STRING = None

def get_new_connection_string():
    db_name = os.environ["OTIS_DB_NAME"]
    username =  os.environ["OTIS_DB_USER"]
    password = os.environ["OTIS_DB_PASSWORD"]
    endpoint = os.environ["OTIS_DB_ENDPOINT"]
    port = os.environ["OTIS_DB_PORT"]
    return parse.quote_plus(f"postgresql://{username}:{password}@{endpoint}:{port}/{db_name}")

def get_connection_string():
    global CONN_STRING
    if CONN_STRING is None:
        CONN_STRING = get_new_connection_string()
    return CONN_STRING

def get_connection():
    """Open a new connection to the database"""
    db_conn_string = get_connection_string()
    engine = sqlalchemy.create_engine(db_conn_string, echo=True)