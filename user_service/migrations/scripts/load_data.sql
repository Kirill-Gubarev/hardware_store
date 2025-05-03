CREATE TEMP TABLE temp_roles (
	name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TEMP TABLE temp_users (
	login VARCHAR(255) NOT NULL UNIQUE,
	password TEXT NOT NULL,
	role_name VARCHAR(255) NOT NULL
);

\copy temp_roles(name) FROM 'data/roles.csv' DELIMITER ',' CSV HEADER;
\copy temp_users(login, password, role_name) FROM 'data/users.csv' DELIMITER ',' CSV HEADER;

INSERT INTO roles(name)
SELECT name
FROM temp_roles
ON CONFLICT (name) DO NOTHING;

INSERT INTO users(login, password, role_id)
SELECT u.login, crypt(u.password, gen_salt('bf')), r.id
FROM temp_users u
LEFT JOIN roles r ON u.role_name = r.name;

DROP TABLE temp_roles;
DROP TABLE temp_users;
