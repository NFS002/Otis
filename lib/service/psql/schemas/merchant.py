""" Merchant ORM class definition """

from sqlalchemy import Column, String
from sqlalchemy.ext.declarative import declarative_base

from lib.service.prototools.common import try_get_property_or_default_null as get_or_default
from lib.service.psql.mixins import IdMixin, CreatedAtMixin, UpdatedAtMixin, DeletedAtMixin, AddressMixin


MerchantBase = declarative_base()

AddressMixin.set_defaults(nullable=True, default_country=None)
IdMixin.set_defaults(id_col_name='id')


class Merchant(MerchantBase, IdMixin, CreatedAtMixin, UpdatedAtMixin, DeletedAtMixin, AddressMixin):
    """ Merchant ORM class definition used to map to psql statements """
    __tablename__ = 'merchants'

    name = Column(String, nullable=False)
    sector = Column(String, nullable=False)

    def __init__(self, **kwargs):
        super().__init__(**kwargs)

    def __repr__(self):
        return "<Merchant(id='%s', name='%s'...)>" % (self.id, self.name)

    def to_dict(self):
        default = getattr(self, 'default_null', 'Null')

        return {
            'id': get_or_default(obj=self, prop='id', default_null=default),
            'name': get_or_default(obj=self, prop='name', default_null=default),
            'sector': get_or_default(obj=self, prop='sector', default_null=default),
            'created_at': get_or_default(obj=self, prop='created_at', default_null=default, date_format='%c'),
            'updated_at': get_or_default(obj=self, prop='deleted_at', default_null=default, date_format='%c'),
            'deleted_at': get_or_default(obj=self, prop='updated_at', default_null=default, date_format='%c'),
            'street_address': get_or_default(obj=self, prop='street_address', default_null=default),
            'postcode': get_or_default(obj=self, prop='postcode', default_null=default),
            'country': get_or_default(obj=self, prop='country', default_null=default),
            'city': get_or_default(obj=self, prop='city', default_null=default)
        }
