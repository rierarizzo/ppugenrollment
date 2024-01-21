USE ppugenrollment;

INSERT INTO
	company (name, ruc, image_url)
VALUES
	('Cobis', '0987293812001', 'no-image'),
	('Banco Pichincha', '0982239281001', 'no-image'),
	('Pronaca', '0922819283001', 'no-image');

INSERT INTO
	project (company, name, description, starts, ends)
VALUES
	(
		1,
		'Sistema financiero',
		'Sistema financiero desc',
		now(),
		now()
	),
	(
		2,
		'Sistema financiero 2',
		'Sistema financiero 2 desc',
		now(),
		now()
	);