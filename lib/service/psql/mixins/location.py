""" This mixin includes columns that can be used to represent the physical location of an entity """
from sqlalchemy import Column, String


class LocationMixin():

    def __init__(self, nullable=True):
        self.longitude = Column(String, nullable=nullable, default='0')
        self.latitude = Column(String, nullable=nullable, default='0')
