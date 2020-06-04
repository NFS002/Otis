""" This mixin implements soft deletes to automatically records the time an SQL entity was deleted """
import dataclasses

from sqlalchemy import Column, DateTime


@dataclasses.dataclass()
class DeletedAtMixin():
    deleted_at = Column(DateTime(timezone=True), nullable=True)
