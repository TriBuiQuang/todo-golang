-- public.todos definition

-- Drop table

-- DROP TABLE public.todos;

CREATE TABLE public.todos (
	id text NOT NULL,
	title text NULL,
	body text NULL,
	completed text NULL,
	created_at timestamptz NULL DEFAULT now(),
	updated_at timestamptz NULL DEFAULT now(),
	CONSTRAINT todos_pkey PRIMARY KEY (id)
);


-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id text NOT NULL,
	username text NOT NULL,
	limit_per_day int8 NULL DEFAULT 10,
	created_at timestamptz NULL DEFAULT now(),
	updated_at timestamptz NULL DEFAULT now(),
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_username_key UNIQUE (username)
);

-- public.tasks definition

-- Drop table

-- DROP TABLE public.tasks;

CREATE TABLE public.tasks (
	id text NOT NULL,
	user_id text NOT NULL,
	"name" text NULL,
	created_at timestamptz NULL DEFAULT now(),
	updated_at timestamptz NULL DEFAULT now(),
	CONSTRAINT tasks_pkey PRIMARY KEY (id)
);


-- Permissions

ALTER TABLE public.users OWNER TO username;
GRANT ALL ON TABLE public.users TO username;

ALTER TABLE public.tasks OWNER TO username;
GRANT ALL ON TABLE public.tasks TO username;

ALTER TABLE public.todos OWNER TO username;
GRANT ALL ON TABLE public.todos TO username;