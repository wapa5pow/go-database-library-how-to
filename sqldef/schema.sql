CREATE TABLE public.pilots (
    id integer NOT NULL,
    name text NOT NULL
);
CREATE SEQUENCE public.pilots_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.pilots_id_seq OWNED BY public.pilots.id;
ALTER TABLE ONLY public.pilots ALTER COLUMN id SET DEFAULT nextval('public.pilots_id_seq'::regclass);
ALTER TABLE ONLY public.pilots
    ADD CONSTRAINT pilot_pkey PRIMARY KEY (id);
