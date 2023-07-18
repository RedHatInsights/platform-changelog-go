CREATE TABLE services (
    id bigint NOT NULL,
    name text NOT NULL,
    display_name text NOT NULL,
    tenant text NOT NULL,
    gh_repo text,
    gl_repo text,
    deploy_file text,
    namespace text,
    branch text DEFAULT 'master'::text
);

CREATE SEQUENCE public.services_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
