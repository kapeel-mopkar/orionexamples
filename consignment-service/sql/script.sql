-- Drop table

-- DROP TABLE shippingapp.consignments

CREATE TABLE shippingapp.consignments (
	id int4 NOT NULL,
	description varchar(500) NULL,
	weight int4 NULL,
	CONSTRAINT consignments_pk PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE shippingapp.consignments OWNER TO postgres;
GRANT ALL ON TABLE shippingapp.consignments TO postgres;

-- Drop table

-- DROP TABLE shippingapp.containers

CREATE TABLE shippingapp.containers (
	id int4 NOT NULL,
	customer_id varchar(50) NULL,
	user_id varchar(50) NULL,
	origin varchar(100) NULL,
	consignment_id int4 NOT NULL,
	CONSTRAINT containers_pk PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE shippingapp.containers OWNER TO postgres;
GRANT ALL ON TABLE shippingapp.containers TO postgres;
