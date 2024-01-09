USE ppugenrollment;

CREATE TABLE user_role (
	code        CHAR        NOT NULL,
	description VARCHAR(15) NOT NULL,
	PRIMARY KEY (code));

CREATE TABLE user (
	id             INT          NOT NULL AUTO_INCREMENT,
	id_card_number VARCHAR(10)  NOT NULL UNIQUE,
	name           VARCHAR(50)  NOT NULL,
	surname        VARCHAR(50)  NOT NULL,
	email          VARCHAR(50)  NOT NULL,
	password       VARCHAR(300) NOT NULL,
	role           CHAR         NOT NULL,
	date_of_birth  DATE,
	is_a_graduate  BOOLEAN DEFAULT FALSE,
	level          INT     DEFAULT 0,
	PRIMARY KEY (id),
	FOREIGN KEY (role) REFERENCES user_role (code));

CREATE TABLE schedule (
	code       CHAR        NOT NULL,
	desciption VARCHAR(15) NOT NULL,
	PRIMARY KEY (code));

CREATE TABLE company (
	id        INT         NOT NULL AUTO_INCREMENT,
	name      VARCHAR(50) NOT NULL,
	ruc       VARCHAR(13) NOT NULL,
	image_url TEXT,
	PRIMARY KEY (id));

CREATE TABLE project (
	id          INT         NOT NULL AUTO_INCREMENT,
	company     INT         NOT NULL,
	name        VARCHAR(50) NOT NULL,
	description TEXT        NOT NULL,
	starts      DATETIME    NOT NULL,
	ends        DATETIME    NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (company) REFERENCES company (id));

CREATE TABLE project_schedule (
	id       INT  NOT NULL AUTO_INCREMENT,
	project  INT  NOT NULL,
	schedule CHAR NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (project) REFERENCES project (id),
	FOREIGN KEY (schedule) REFERENCES schedule (code));

CREATE TABLE enrollment_application (
	id         INT  NOT NULL AUTO_INCREMENT,
	student    INT  NOT NULL,
	project    INT  NOT NULL,
	schedule   INT  NOT NULL,
	applied_on DATETIME      DEFAULT NOW(),
	status     CHAR NOT NULL DEFAULT 'P', # P: Pendiente, A: Approved
	PRIMARY KEY (id),
	FOREIGN KEY (student) REFERENCES user (id),
	FOREIGN KEY (project) REFERENCES project (id),
	FOREIGN KEY (schedule) REFERENCES project_schedule (id));

CREATE TABLE enrollment_generated (
	id                     INT      NOT NULL AUTO_INCREMENT,
	enrollment_application INT      NOT NULL,
	approved_by            INT      NOT NULL,
	generated_at           DATETIME NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id),
	FOREIGN KEY (approved_by) REFERENCES user (id));
