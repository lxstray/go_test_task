CREATE TABLE banners (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    image TEXT NOT NULL,
    cpm NUMERIC(10,2) NOT NULL CHECK (cpm >= 0 AND cpm <= 1000),
    geo VARCHAR(2) NOT NULL, --BY, US, RU, UK etc
    feature INTEGER NOT NULL CHECK (feature >= 0 AND feature <= 100)
);