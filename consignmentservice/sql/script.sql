-- Drop table

-- DROP TABLE public.consignments

CREATE TABLE public.consignments (
	id serial NOT NULL,
	description varchar(500) NULL,
	weight int4 NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamp NULL,
	CONSTRAINT consignments_pk PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE public.consignments OWNER TO postgres;
GRANT ALL ON TABLE public.consignments TO postgres;

-- Drop table

-- DROP TABLE public.containers

CREATE TABLE public.containers (
	id serial NOT NULL,
	customer_id varchar(50) NULL,
	user_id varchar(50) NULL,
	origin varchar(100) NULL,
	consignment_id int4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at timestamp NULL,
	CONSTRAINT containers_pk PRIMARY KEY (id),
	CONSTRAINT containers_consignments_fk FOREIGN KEY (consignment_id) REFERENCES consignments(id)
);

-- Permissions

ALTER TABLE public.containers OWNER TO postgres;
GRANT ALL ON TABLE public.containers TO postgres;
