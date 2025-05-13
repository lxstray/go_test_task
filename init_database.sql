-- DROP TABLE IF EXISTS "banners";

CREATE TABLE banners (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    image TEXT NOT NULL,
    cpm NUMERIC(10,2) NOT NULL CHECK (cpm >= 0 AND cpm <= 1000),
    geo VARCHAR(2) NOT NULL, --BY, US, RU, UK etc
    feature INTEGER NOT NULL CHECK (feature >= 0 AND feature <= 100)
);

INSERT INTO banners (name, image, cpm, geo, feature)
SELECT 
    'Banner_' || g AS name,
    'https://example.com/banners/image' || g || '.jpg' AS image,
    ROUND((RANDOM() * 1000)::NUMERIC, 2) AS cpm,
    (ARRAY['BY', 'US', 'RU', 'UK', 'DE', 'FR', 'CN', 'JP', 'IN', 'BR'])[FLOOR(RANDOM() * 10 + 1)] AS geo,
    FLOOR(RANDOM() * 101)::INTEGER AS feature
FROM generate_series(1, 10000) g;

EXPLAIN ANALYZE SELECT * FROM banners WHERE geo = 'US' AND feature = 50 ORDER BY cpm DESC LIMIT 1;

CREATE INDEX idx_geo_feature ON banners (geo, feature);