BEGIN;

CREATE TABLE IF NOT EXISTS countries (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS counties (
	id SERIAL PRIMARY KEY,
	country_id INTEGER NOT NULL,
	name VARCHAR(100) NOT NULL,

	CONSTRAINT fk_country
		FOREIGN KEY(country_id)
			REFERENCES countries(id)
			ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS contact_data (
	id SERIAL PRIMARY KEY,
	country_id INTEGER NOT NULL,
	county_id INTEGER NOT NULL,
	phone_number VARCHAR(50) NOT NULL,
	address VARCHAR(300) NOT NULL,

	CONSTRAINT fk_country
		FOREIGN KEY(country_id)
			REFERENCES countries(id)
			ON DELETE CASCADE,

	CONSTRAINT fk_county
		FOREIGN KEY(county_id)
			REFERENCES counties(id)
			ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	contact_data_id INTEGER NOT NULL,
	email VARCHAR(300) UNIQUE NOT NULL,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	password VARCHAR(255) NOT NULL,
	admin BOOLEAN DEFAULT false NOT NULL,
	create_time TIMESTAMPTZ DEFAULT NOW() NOT NULL,

	CONSTRAINT fk_contact_data
		FOREIGN KEY(contact_data_id)
			REFERENCES contact_data(id)
			ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS services (
	id SERIAL PRIMARY KEY,
	contact_data_id INTEGER NOT NULL,
	email VARCHAR(300) UNIQUE NOT NULL,
	name VARCHAR(50) NOT NULL,
	password VARCHAR(255) NOT NULL,
	create_time TIMESTAMPTZ DEFAULT NOW() NOT NULL,

	CONSTRAINT fk_contact_data
		FOREIGN KEY(contact_data_id)
			REFERENCES contact_data(id)
			ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS ratings (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	service_id INTEGER NOT NULL,
	value SMALLINT CHECK(value >= 1 AND value <= 5 AND value IS NOT NULL),
	message TEXT NULL,
	create_time TIMESTAMPTZ DEFAULT NOW() NOT NULL,

	CONSTRAINT fk_service
		FOREIGN KEY(service_id)
			REFERENCES services(id)
			ON DELETE CASCADE,
	
	CONSTRAINT fk_user
		FOREIGN KEY(user_id)
			REFERENCES users(id)
			ON DELETE CASCADE
);

CREATE TYPE appointment_status AS ENUM (
	'pending',
	'done',
	'refused',
	'accepted'
);

CREATE TABLE IF NOT EXISTS appointments (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	service_id INTEGER NOT NULL,
	data TIMESTAMPTZ NOT NULL,
	status appointment_status NOT NULL,
	details TEXT NULL,
	create_time TIMESTAMPTZ DEFAULT NOW() NOT NULL,

	CONSTRAINT fk_service
		FOREIGN KEY(service_id)
			REFERENCES services(id)
			ON DELETE CASCADE,
	
	CONSTRAINT fk_user
		FOREIGN KEY(user_id)
			REFERENCES users(id)
			ON DELETE CASCADE

);

CREATE TABLE IF NOT EXISTS reports (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NULL,
	details TEXT NULL,
	title VARCHAR(100) NOT NULL,
	create_time TIMESTAMPTZ DEFAULT NOW() NOT NULL,

	CONSTRAINT fk_user
		FOREIGN KEY(user_id)
			REFERENCES users(id)
			ON DELETE SET NULL
);

COMMIT;
