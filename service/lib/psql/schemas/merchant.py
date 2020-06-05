""" Merchant ORM class definition """
from dataclasses import dataclass

from sqlalchemy import Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base

from service.lib.psql.mixins import IdMixin, CreatedAtMixin, UpdatedAtMixin, DeletedAtMixin, AddressMixin


MerchantBase = declarative_base()


@dataclass()
class Merchant(MerchantBase, IdMixin, CreatedAtMixin, UpdatedAtMixin, DeletedAtMixin, AddressMixin):
    """ Merchant class definition used to map to psql statements """
    __tablename__ = 'merchants'

    id = Column(Integer, primary_key=True)
    name = Column(String, nullable=False)
    sector = Column(String, nullable=False)

    def __init__(self, **kwargs):
        AddressMixin.__init__(self, nullable=False)
        super().__init__(**kwargs)

    def __repr__(self):
        return "<Merchant(id='%s', name='%s')>" % (self.id, self.name)
