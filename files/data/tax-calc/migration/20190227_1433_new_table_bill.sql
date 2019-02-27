CREATE TABLE bills (
    id       SERIAL,       -- Row primary key.
    name     VARCHAR(100), -- Tax object's product name.
    tax_code SMALLINT,     -- Tax object's category code (sync value with what's available in codebase).
    price    DECIMAL(19,2) -- Tax object's price amount.
);
