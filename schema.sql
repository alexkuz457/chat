CREATE DATABASE chat
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251@icu'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
	
	CREATE TABLE public.chats
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name character varying(64) COLLATE pg_catalog."default",
    created_at timestamp with time zone DEFAULT now(),
    users integer[],
    CONSTRAINT chats_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.chats
    OWNER to postgres;
	
	CREATE TABLE public.messages
(
    autor character varying(64) COLLATE pg_catalog."default",
    created_at timestamp with time zone DEFAULT now(),
    id bigint NOT NULL DEFAULT nextval('messages_id_seq'::regclass),
    text text COLLATE pg_catalog."default",
    chat integer,
    CONSTRAINT messages_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.messages
    OWNER to postgres;
	
	CREATE TABLE public.users
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    username character varying(64) COLLATE pg_catalog."default",
    created_at timestamp with time zone DEFAULT now(),
    CONSTRAINT id PRIMARY KEY (id)
        INCLUDE(id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;