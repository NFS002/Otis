""" Merchant ORM class definition """
from dataclasses import dataclass

from sqlalchemy import Column, String
from sqlalchemy.ext.declarative import declarative_base

from service.lib.psql.mixins import IdMixin, CreatedAtMixin, UpdatedAtMixin, DeletedAtMixin, AddressMixin


MerchantBase = declarative_base()
AddressMixin.set_defaults(nullable=True)
IdMixin.set_defaults(name='merchantID')


@dataclass()
class Merchant(MerchantBase, IdMixin, CreatedAtMixin, UpdatedAtMixin, DeletedAtMixin, AddressMixin):
    """ Merchant class definition used to map to psql statements """
    __tablename__ = 'merchants'

    name = Column(String, nullable=False)
    sector = Column(String, nullable=False)

    def __init__(self, **kwargs):
        super().__init__(**kwargs)

    def __repr__(self):
        return "<Merchant(merchantID='%s', name='%s'...)>" % (self.merchantID, self.name)
