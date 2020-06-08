-- plans table schema
CREATE TABLE IF NOT EXISTS plans (
      id                     SERIAL,
      name                   VARCHAR        not null,
      description            VARCHAR        not null,
      start_date             DATE           not null,
      street_address         VARCHAR        not null,
      postcode               VARCHAR(10)    not null,
      city                   VARCHAR        not null,
      longitude              VARCHAR        not null         default '0',
      latitude               VARCHAR        not null         default '0',
      group_id               INTEGER        not null,
      pot_id                 INTEGER        not null,
      created_at             DATE           not null         default current_date,
      deleted_at             DATE

      CHECK                  ( name <> '' ),
      CHECK                  ( street_address <> '' ),
      CHECK                  ( city <> '' ),

      CONSTRAINT             pk_plans_id          PRIMARY KEY (id),
      CONSTRAINT             fk_plans_grous_id    FOREIGN KEY (id) REFERENCES groups (id) ON UPDATE CASCADE ON DELETE RESTRICT,
      CONSTRAINT             fk_plans_pot_id      FOREIGN KEY (id) REFERENCES pots   (id) ON UPDATE CASCADE ON DELETE RESTRICT
);