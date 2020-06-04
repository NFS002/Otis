""" This mixin automatically records the start time an SQL entity was created """
from dataclasses import dataclass

from sqlalchemy.sql import func
from sqlalchemy import Column, DateTime


@dataclass()
class CreatedAtMixin():
    created_at = Column(DateTime(timezone=True), nullable=False, server_default=func.now())
