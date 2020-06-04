""" This mixin automatically records the start time an SQL entity was updated """
from dataclasses import dataclass

from sqlalchemy.sql import func
from sqlalchemy import Column, DateTime


@dataclass()
class UpdatedAtMixin():
    updated_at = Column(DateTime(timezone=True), nullable=False, onupdate=func.now())
