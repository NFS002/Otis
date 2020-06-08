""" SQLAlchemy id column mixin """
from sqlalchemy import Column, Integer
from sqlalchemy.ext.declarative import declared_attr


# pylint: disable=no-self-argument
class IdMixin():
    """ An ORM mixin class (for use with SQLAlchemy) to give a unique 'id' column to an sql table """
    id_col_name = 'id'

    @declared_attr
    def id(cls):
        return Column(cls.id_col_name, Integer, primary_key=True, autoincrement=True)

    @classmethod
    def set_defaults(cls, id_col_name='id'):
        cls.id_col_name = id_col_name
