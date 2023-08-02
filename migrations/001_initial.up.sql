-- services
CREATE TABLE services (
    id bigint NOT NULL PRIMARY KEY,
    name text NOT NULL,
    display_name text NOT NULL,
    tenant text NOT NULL,
    gh_repo text,
    gl_repo text,
    deploy_file text,
    namespace text,
    branch text DEFAULT 'master'::text
);

CREATE SEQUENCE services_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE services_id_seq OWNED BY services.id;
ALTER TABLE ONLY services ALTER COLUMN id SET DEFAULT nextval('services_id_seq'::regclass);

-- timelines
CREATE TABLE timelines (
    id bigint NOT NULL PRIMARY KEY,
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

CREATE TYPE timeline_type AS ENUM (
    'unknown',
    'commit',
    'deploy'
);

CREATE SEQUENCE timelines_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE timelines_id_seq OWNED BY timelines.id;
ALTER TABLE ONLY timelines ALTER COLUMN id SET DEFAULT nextval('timelines_id_seq'::regclass);
