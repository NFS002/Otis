""" This mixin includes columns that can be used to represent the physical address of an entity """
from dataclasses import dataclass

from sqlalchemy import Column, String


@dataclass()
class AddressMixin():

    def __init__(self, nullable=True, default_country='GB'):
        self.street_adress = Column(String, nullable=nullable)
        self.city = Column(String(20), nullable=nullable)
        self.country = Column(String, nullable=nullable, default=default_country)
        self.postcode = Column(String(10), nullable=nullable)
