create extension if not exists vector;

create table org(
  id bigserial primary key,
  name text not null
);

create table "user"(
  id bigserial primary key,
  org_id bigint references org(id),
  email text unique not null,
  password_hash text not null,
  role text not null default 'user'
);

create table matter(
  id bigserial primary key,
  org_id bigint references org(id),
  title text,
  requester text,
  contract_type text,
  due_date date,
  status text default 'open'
);

create table document(
  id bigserial primary key,
  org_id bigint references org(id),
  matter_id bigint references matter(id),
  title text,
  mime text,
  stored_path text,
  text text,
  created_at timestamptz default now()
);

create table clause_type(
  id bigserial primary key,
  name text unique,
  description text
);

create table clause(
  id bigserial primary key,
  document_id bigint references document(id),
  start_idx int,
  end_idx int,
  text text,
  clause_type_id bigint references clause_type(id)
);

create table playbook(
  id bigserial primary key,
  org_id bigint references org(id),
  name text,
  description text,
  active bool default true
);

create table rule(
  id bigserial primary key,
  playbook_id bigint references playbook(id),
  name text,
  severity text,
  pattern text,
  guidance text,
  llm_check bool default false
);

create table finding(
  id bigserial primary key,
  document_id bigint references document(id),
  clause_id bigint references clause(id),
  rule_id bigint references rule(id),
  severity text,
  rationale text,
  suggestion text,
  status text default 'open'
);

create table action_log(
  id bigserial primary key,
  finding_id bigint references finding(id),
  user_id bigint references "user"(id),
  action text,
  note text,
  ts timestamptz default now()
);

create table precedent(
  id bigserial primary key,
  org_id bigint references org(id),
  title text,
  text text,
  tags text
);

create table precedent_embedding(
  precedent_id bigint primary key references precedent(id),
  vector vector(1536)
);

create table clause_embedding(
  clause_id bigint primary key references clause(id),
  vector vector(1536)
);

insert into org(name) values ('Demo Org');
insert into "user"(org_id,email,password_hash,role) values (1,'demo@demo.dev','demo','admin');
