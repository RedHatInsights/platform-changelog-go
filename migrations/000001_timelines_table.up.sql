CREATE TABLE timelines (
    id bigint NOT NULL,
    service_id bigint NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
    type text NOT NULL,
    repo text NOT NULL,
    ref text,
    author text,
    merged_by text,
    message text,
    deploy_namespace text,
    cluster text,
    image text,
    triggered_by text,
    status text
);

CREATE SEQUENCE public.timelines_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TYPE timeline_type AS ENUM (
    'unknown',
    'commit',
    'deploy'
);
