""" This mixin includes columns that can be used to represent the physical address of an entity """
from dataclasses import dataclass

from sqlalchemy import Column, String


@dataclass()
class AddressMixin():

    @classmethod
    def set_defaults(cls, nullable=False, default_country=None):
        cls.street_adress = Column(String, nullable=nullable)
        cls.city = Column(String(20), nullable=nullable)
        cls.country = Column(String, nullable=nullable, default=default_country)
        cls.postcode = Column(String(10), nullable=nullable)
