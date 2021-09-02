/* CREATE SCHEMA*/
    CREATE SCHEMA IF NOT EXISTS belivine AUTHORIZATION root;
    SET SEARCH_PATH = "belivine";
/* END */

/* USERS requirement relation */
CREATE SEQUENCE category_seq;
CREATE TABLE IF NOT EXISTS category(
    id int NOT NULL UNIQUE  DEFAULT nextval('category_seq'),
    name varchar,
    PRIMARY KEY (id)
);

CREATE SEQUENCE team_seq;
CREATE TABLE IF NOT EXISTS team(
    id int NOT NULL UNIQUE DEFAULT nextval('team_seq'),
    name varchar,
    scale int,
    category_id int,
    FOREIGN KEY(category_id) REFERENCES category(id),
    PRIMARY KEY (id)
);

CREATE SEQUENCE users_seq;
CREATE TABLE IF NOT EXISTS users(
    id int NOT NULL UNIQUE DEFAULT nextval('users_seq'),
    username varchar,
    password varchar(255),
    email varchar,
    team_id int,
    FOREIGN KEY(team_id) REFERENCES team(id),
    PRIMARY KEY (id)
);

/*END*/
CREATE SEQUENCE task_group_seq;
CREATE TABLE IF NOT EXISTS task_group(
    id int NOT NULL UNIQUE DEFAULT nextval('task_group_seq'),
    name varchar,
    PRIMARY KEY (id)
);

CREATE SEQUENCE tasks_seq;
CREATE TABLE IF NOT EXISTS tasks(
    id int NOT NULL UNIQUE DEFAULT nextval('tasks_seq'),
    name varchar,
    est_hours int,
    est_date date,
    task_group_id int,
    FOREIGN KEY(task_group_id) REFERENCES task_group(id),
    PRIMARY KEY (id)
);

CREATE SEQUENCE times_seq;
CREATE TABLE IF NOT EXISTS times(
    id int NOT NULL UNIQUE DEFAULT nextval('times_seq'),
    start timestamp,
    end timestamp,
    task_id int,
    user_id int,
    FOREIGN KEY(task_id) REFERENCES tasks(id), 
    FOREIGN KEY(user_id) REFERENCES users(id),
    PRIMARY KEY (id)
)
