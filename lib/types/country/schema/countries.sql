-- countries table schema
CREATE TABLE IF NOT EXISTS countries (
    code                    VARCHAR(4)            not null  unique, -- ISO 3166-1 alpha-2
    name                    VARCHAR               not null  unique,

    CONSTRAINT              pk_countries_code     PRIMARY KEY (code)
);