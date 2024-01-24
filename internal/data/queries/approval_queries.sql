-- name: UpdateEnrollmentApplication :exec
UPDATE enrollment_application SET status = 'A' WHERE id = ?;

-- name: CreateEnrollmentGenerated :execresult
INSERT INTO enrollment_generated (enrollment_application, approved_by, observation) VALUES (?, ?, ?);

-- name: GetEnrollmentGenerated :one
SELECT eg.id                     AS id,
	   eg.enrollment_application AS application_id,
	   eg.generated_at           AS generated_at,
	   p.id                      AS project_id,
	   p.description             AS project_description,
	   ps.schedule               AS project_schedule,
	   p.starts                  AS project_starts,
	   p.ends                    AS project_ends,
	   c.id                      AS company_id,
	   c.name                    AS company_name,
	   c.ruc                     AS company_ruc,
	   su.id                     AS approver_id,
	   su.id_card_number         AS approver_card_number,
	   su.name                   AS approver_name,
	   su.surname                AS approver_surname
FROM enrollment_generated eg
		 INNER JOIN enrollment_application ea ON eg.enrollment_application = ea.id
		 INNER JOIN project p ON ea.project = p.id
		 INNER JOIN project_schedule ps ON ea.schedule = ps.id
		 INNER JOIN company c ON p.company = c.id
		 INNER JOIN user su ON eg.approved_by = su.id
WHERE eg.id = ?
LIMIT 1;

-- name: GetEnrollmentApplications :many
SELECT ea.id 					 AS application_id,
	   ea.status				 AS application_status,
	   sus.id			 	 	 AS student_id,
	   sus.name			 		 AS student_name,
	   sus.surname				 AS student_surname,
	   p.id                      AS project_id,
	   p.name 				  	 AS project_name,	
	   p.description             AS project_description,
	   ps.schedule               AS project_schedule,
	   p.starts                  AS project_starts,
	   p.ends                    AS project_ends,
	   c.id                      AS company_id,
	   c.name                    AS company_name,
	   c.ruc                     AS company_ruc
FROM enrollment_application ea
		 INNER JOIN project p ON ea.project = p.id
		 INNER JOIN project_schedule ps ON ea.schedule = ps.id
		 INNER JOIN company c ON p.company = c.id
		 INNER JOIN user sus ON ea.student = sus.id;
		 