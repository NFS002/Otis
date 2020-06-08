""" This mixin automatically records the start time an SQL entity was updated """

from sqlalchemy.sql import func
from sqlalchemy import Column, DateTime


class UpdatedAtMixin():
    updated_at = Column(DateTime(timezone=True), nullable=False, server_default=func.now(), onupdate=func.now())
