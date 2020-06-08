"""  SQLAlchemy address columns mixin """
from sqlalchemy import Column, String
from sqlalchemy.ext.declarative import declared_attr


# pylint: disable=no-self-argument
class AddressMixin():
    """
    This mixin (for use with SQLAlchemy ORM) includes sql columns
    that can be used to represent the physical address of an entity
    """
    nullable = True
    default_country = None

    @declared_attr
    def street_address(cls):
        return Column(String, nullable=cls.nullable)

    @declared_attr
    def postcode(cls):
        return Column(String(10), nullable=cls.nullable)

    @declared_attr
    def country(cls):
        return Column(String, nullable=cls.nullable, default=cls.default_country)

    @declared_attr
    def city(cls):
        return Column(String(20), nullable=cls.nullable)

    @classmethod
    def set_defaults(cls, nullable=True, default_country=None):
        cls.nullable = nullable
        cls.default_country = default_country
