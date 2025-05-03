CREATE TEMP TABLE temp_units (
    name VARCHAR(255)
);

CREATE TEMP TABLE temp_products (
    name VARCHAR(255),
    description VARCHAR(255),
    unit_name VARCHAR(255)
);

\copy temp_units(name) FROM 'data/units.csv' DELIMITER ',' CSV HEADER;
\copy temp_products(name, description, unit_name) FROM 'data/products.csv' DELIMITER ',' CSV HEADER;

INSERT INTO units(name)
SELECT name
FROM temp_units
ON CONFLICT (name) DO NOTHING;

INSERT INTO products(name, description, unit_id)
SELECT p.name, p.description, u.id
FROM temp_products p
LEFT JOIN units u ON p.unit_name = u.name;

DROP TABLE temp_units;
DROP TABLE temp_products;
