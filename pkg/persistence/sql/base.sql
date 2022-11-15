USE swapiapp;

CREATE TABLE IF NOT EXISTS 'film' (
    'id' INT PRIMARY KEY,
    'title' VARCHAR(255) NOT NULL,
    'episodeid' INT NOT NULL,
    'director' VARCHAR(50) NOT NULL,
    'created' VARCHAR(40) NOT NULL,
    'url' VARCHAR(255) NOT NULL,
    INDEX ('title')
);

CREATE TABLE IF NOT EXISTS 'planet' (
    'id' INT PRIMARY KEY,
    'name' VARCHAR(255) NOT NULL,
    'climate' VARCHAR(50) NOT NULL,
    'terrain' VARCHAR(40),
    'url' VARCHAR(255),
    INDEX ('name')
);

CREATE TABLE IF NOT EXISTS 'planet_film' (
    'filmid' INT NOT NULL,
    'planetid' INT NOT NULL,
    INDEX ('filmid', 'planetid')
);

