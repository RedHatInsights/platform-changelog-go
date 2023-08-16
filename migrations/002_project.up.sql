-- add project abstraction
CREATE TABLE projects (
    id bigint NOT NULL PRIMARY KEY,
    service_id bigint NOT NULL,
    name text NOT NULL,
    repo text,
    deploy_file text,
    namespaces text[],
    branches text[]
);

CREATE SEQUENCE projects_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE projects_id_seq OWNED BY projects.id;
ALTER TABLE ONLY projects ALTER COLUMN id SET DEFAULT nextval('projects_id_seq'::regclass);

ALTER TABLE services 
    DROP gh_repo, 
    DROP gl_repo, 
    DROP deploy_file, 
    DROP namespace, 
    DROP branch;
