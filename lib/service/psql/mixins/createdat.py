""" This mixin automatically records the start time an SQL entity was created """
from sqlalchemy.sql import func
from sqlalchemy import Column, DateTime


class CreatedAtMixin():
    created_at = Column(DateTime(timezone=True), nullable=False, server_default=func.now())
