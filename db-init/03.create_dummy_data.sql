use ppugenrollment;

/* Default companies */
insert into company (name, ruc, image_url)
values ('Viamatica', '0928192031001', 'https://viamatica.com/wp-content/uploads/2021/05/Logo-Viamatica.png');


/* Default projects */
insert into project (company, name, description, starts, ends)
values (1, 'Adopt Pet', 'Programar un sistema de adopción de mascotas', str_to_date('5,2,2024', '%d,%m,%Y'),
        str_to_date('7,4,2024', '%d,%m,%Y'));

insert into project_schedule (project, schedule)
values (1, 'M');
insert into project_schedule (project, schedule)
values (1, 'E');