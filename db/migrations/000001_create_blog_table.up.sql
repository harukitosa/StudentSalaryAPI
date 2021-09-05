CREATE TABLE IF NOT EXISTS blogs(
    id serial PRIMARY KEY,
    `url` VARCHAR (255) UNIQUE NOT NULL,
    title VARCHAR (255),
    `description` VARCHAR (255),
    company_name VARCHAR (255) NOT NULL,
    `year` int,
    season VARCHAR(20),
    image_url VARCHAR(255)
);