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