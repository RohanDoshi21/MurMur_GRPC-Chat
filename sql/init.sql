create database casdoor;
create database messagedb;

\c messagedb;

create table messages (
  id SERIAL PRIMARY KEY,
  sender text NOT NULL,
  receiver text NOT NULL,
  content text NOT NULL,
  timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE messages ADD COLUMN is_group BOOLEAN DEFAULT FALSE;

create table groups (
  id text PRIMARY KEY,
  name text NOT NULL,
  users text[] NOT NULL,
  owner text NOT NULL,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP  
);

create table invites (
  id text PRIMARY KEY,
  sender text NOT NULL,
  receiver text[] NOT NULL,
  times_used int NOT NULL DEFAULT 0,
  group_id text NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
