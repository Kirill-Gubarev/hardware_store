CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS units(
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS products(
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description VARCHAR(255),
	unit_id UUID REFERENCES units(id),
	image_id UUID,
	manufacturer_id UUID
);
