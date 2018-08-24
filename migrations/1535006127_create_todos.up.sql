CREATE TABLE todos(
	id serial NOT NULL,
    name text NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    updated_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,

    PRIMARY KEY(id)
);