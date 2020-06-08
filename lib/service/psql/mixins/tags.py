"""
This mixin includes columns used to represent a one-to-many relationship between the parent table  and the tags table
"""
from sqlalchemy.orm import relationship


class TagsMixin():
    def __init__(self):
        self.tags = relationship("Tags", back_populates=self.__tablename__)
