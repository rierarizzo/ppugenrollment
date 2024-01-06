use ppugenrollment;

create table user_role
(
    code        char        not null,
    description varchar(15) not null,
    primary key (code)
);

create table user
(
    id             int          not null auto_increment,
    id_card_number varchar(10)  not null unique,
    name           varchar(50)  not null,
    surname        varchar(50)  not null,
    email          varchar(50)  not null,
    password       varchar(300) not null,
    role           char         not null,
    date_of_birth date,
    is_a_graduate boolean default false,
    level         int default 0,
    primary key (id),
    foreign key (role) references user_role (code)
);

create table schedule
(
    code       char        not null,
    desciption varchar(15) not null,
    primary key (code)
);

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
    id          int         not null auto_increment,
    company     int         not null,
    name        varchar(50) not null,
    description text        not null,
    starts      datetime    not null,
    ends        datetime    not null,
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
    foreign key (student) references user (id),
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
    foreign key (approved_by) references user (id)
);
