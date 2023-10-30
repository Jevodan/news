-- Table: public.news

-- DROP TABLE IF EXISTS public.news;

CREATE TABLE IF NOT EXISTS public.news
(
    id integer NOT NULL DEFAULT nextval('news_id_seq'::regclass),
    title character varying(255) COLLATE pg_catalog."default",
    description character varying(4000) COLLATE pg_catalog."default",
    link character varying(255) COLLATE pg_catalog."default",
    publish bigint,
    CONSTRAINT news_pkey PRIMARY KEY (id),
    CONSTRAINT link UNIQUE (link)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.news
    OWNER to leka;