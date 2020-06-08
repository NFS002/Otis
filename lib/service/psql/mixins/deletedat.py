""" This mixin implements soft deletes to automatically records the time an SQL entity was deleted """
from sqlalchemy import Column, DateTime


class DeletedAtMixin():
    deleted_at = Column(DateTime(timezone=True), nullable=True)
