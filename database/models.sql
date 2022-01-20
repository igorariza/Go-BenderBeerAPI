CREATE TABLE IF NOT EXISTS beers (
    id serial NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now(),
    name VARCHAR(150) NOT NULL,
    brewery VARCHAR(150) NOT NULL,
    country VARCHAR(150) NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    currency NUMERIC(10,2) NOT NULL
);