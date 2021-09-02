CREATE SEQUENCE user_id_seq;
CREATE TABLE User (
    ID BIGINT NOT NULL DEFAULT nextval('user_id_seq'),
    Username VARCHAR (127) NOT NULL UNIQUE,
    Password VARCHAR (127) NOT NULL,
    PRIMARY KEY (ID)
);

CREATE SEQUENCE links_id_seq;
CREATE TABLE IF NOT EXISTS Links(
    ID BIGINT NOT NULL DEFAULT nextval('links_id_seq'),
    Title VARCHAR (255) ,
    Address VARCHAR (255) ,
    UserID INT ,
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,
    PRIMARY KEY (ID)
)