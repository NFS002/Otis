-- tags table schema
CREATE TABLE IF NOT EXISTS tags (
         name                    VARCHAR               not null  unique, -- ISO 3166-1 alpha-2
         info                    VARCHAR,
         created_at              DATE                  not null  default CURRENT_DATE,

         CONSTRAINT              pk_tags_name          PRIMARY KEY (name)
);