""" An ORM mixin class to give a unique 'id' column to an sql table """
from dataclasses import dataclass

from sqlalchemy import Column, Integer


@dataclass()
class IdMixin():
    id = Column(Integer, primary_key=True, autoincrement=True)
