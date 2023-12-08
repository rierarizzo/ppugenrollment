drop database if exists ppugenrollment;

create database ppugenrollment;
use ppugenrollment;

create table user_role
(
    code        char        not null,
    description varchar(15) not null,
    primary key (code)
);

insert into user_role
values ('S', 'Estudiante'),
       ('A', 'Aprobador'),
       ('M', 'Administrador');

create table sys_user
(
    id             int          not null auto_increment,
    id_card_number varchar(10)  not null unique,
    name           varchar(50)  not null,
    surname        varchar(50)  not null,
    email          varchar(50)  not null,
    password       varchar(300) not null,
    role           char         not null,
    primary key (id),
    foreign key (role) references user_role (code)
);

create table student
(
    sys_user      int not null,
    date_of_birth date,
    is_a_graduate boolean,
    level         int,
    foreign key (sys_user) references sys_user (id)
);

create table approver
(
    sys_user int not null,
    foreign key (sys_user) references sys_user (id)
);

create table admin
(
    sys_user int not null,
    foreign key (sys_user) references sys_user (id)
);

create table schedule
(
    code       char        not null,
    desciption varchar(15) not null,
    primary key (code)
);

insert into schedule
values ('M', 'Matutino'),
       ('E', 'Vespertino'),
       ('N', 'Nocturno');

create table company
(
    id        int         not null auto_increment,
    name      varchar(50) not null,
    ruc       varchar(13) not null,
    image_url text,
    primary key (id)
);

create table project
(
    id          int      not null auto_increment,
    company     int      not null,
    description text     not null,
    starts      datetime not null,
    ends        datetime not null,
    primary key (id),
    foreign key (company) references company (id)
);

create table project_schedule
(
    id       int  not null auto_increment,
    project  int  not null,
    schedule char not null,
    primary key (id),
    foreign key (project) references project (id),
    foreign key (schedule) references schedule (code)
);

create table enrollment_application
(
    id         int  not null auto_increment,
    student    int  not null,
    project    int  not null,
    schedule   int  not null,
    applied_on datetime      default now(),
    status     char not null default 'P', # P: Pendiente, A: Approved
    primary key (id),
    foreign key (student) references student (sys_user),
    foreign key (project) references project (id),
    foreign key (schedule) references project_schedule (id)
);

create table enrollment_generated
(
    id                     int      not null auto_increment,
    enrollment_application int      not null,
    approved_by            int      not null,
    generated_at           datetime not null default now(),
    primary key (id),
    foreign key (approved_by) references approver (sys_user)
);

/* Default companies */
insert into company (name, ruc, image_url)
values ('Viamatica', '0928192031001', 'https://viamatica.com/wp-content/uploads/2021/05/Logo-Viamatica.png');


/* Default projects */
insert into project (company, description, starts, ends)
values (1, 'Programar un sistema de adopción de mascotas', str_to_date('5,2,2024', '%d,%m,%Y'),
        str_to_date('7,4,2024', '%d,%m,%Y'));

insert into project_schedule (project, schedule) values (1, 'M');
insert into project_schedule (project, schedule) values (1, 'E');
