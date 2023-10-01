PRAGMA foreign_keys = OFF;
-- truncate, basically
DELETE FROM magazine_manga;
DELETE FROM manga_mangaka_job;
DELETE FROM manga_genre;
DELETE FROM magazine;
DELETE FROM mangaka;
DELETE FROM chapter;
DELETE FROM manga;
DELETE FROM genre;
DELETE FROM demo;
DELETE FROM job;
-- restart identity, basically
DELETE FROM sqlite_sequence;
PRAGMA foreign_keys = ON;

INSERT INTO demo
VALUES
('kodomo'),
('shonen'),
('shojo'),
('seinen'),
('josei'),
('seijin'),
('mina');

INSERT INTO job
VALUES
('author'),
('artist'),
('author_artist');

INSERT INTO magazine ("name", "other_names", "description", "demo")
VALUES
('Weekly Shonen Jump', '["WSJ", "Shonen Jump", "Jump"]', 'A truly venerable institution', 'shonen'),
('Weekly Morning', '["Morning"]', 'Sister magazine to Afternoon and Evening.', 'seinen'),
('Monthly Afternoon', '["Afternoon"]', 'A hardcore otaku magainze. Sister magazine to Morning and Evening.', 'seinen'),
('Evening', '[]', 'No associated manga in seed. Sister magazine to Morning and Afternoon.', 'seinen'), -- no associated manga
('Young Animal', '[]', 'They publish Berserk, sometimes.', 'seinen'),
('Monthly Animal House', '[]', 'They used to publish Berserk, regularly.', 'seinen'),
('Weekly Young Jump', '[]', 'If theres''s one thing Japanese franchises love, it''s confusing names.', 'seinen'),
('Weekly Shonen Magazine', '[]', 'Mirror universe Weekly Shonen Jump', 'shonen'),
('Big Comic', '[]', 'I''d heard of Big Comic Spirits, but never Big Comic', 'seinen');

INSERT INTO mangaka ("name", "other_names", "description")
VALUES
('Kubo Tite', '["Kubo Taito"]', 'Struck gold with the worst, but, mercifully, shortest of the big 3.'),
('Oda Eiichiro', '[]', 'A true mad genius, author of the first, last, and greatest of the big 3.'),
('Kishimoto Masashi', '[]', 'If hadn''t read the first chapter of naruto when I was like 12 we wouldn''t be here.'),
('Oba Tsugumi', '[]', 'It''s a pen name. His real name is, oddly enough, a closely held secret.'),
('Obata Takeshi', '[]', 'He''s a lot less name-shy than his boy and parter Oba Tsugumi.'),
('Miura Kentaro', '[]', 'We''ll always love him for Berserk, and always hate him for abandoning it.'),
('Samura Hiroaki', '[]', 'We''ll always love him for Blade, and always be skeeved out by everything else.'),
('Inoue Takehiko', '[]', 'Another gargantuan talent wrecked on the shoals of a brutal release schedule.'),
('Kio Shimoku', '[]', 'What would possess a man to write and publish genshiken?'),
('Kajiwara Ikki', '["Takamori Asaki", "Takamori Asao"]', 'Although a super prolific author, he''s most famous for writing Ashita no Joe.'),
('Chiba Tetsuya', '[]', 'Although he was an author_artist, he''s most famous for drawing Ashita no Joe.'), -- author_artist and artist jobs
('Murata Yusuke', '[]', 'No associated manga in seed. He''s famous for eyeshield 21 and one punch man though.'), -- no associated manga
('Shūhō Satō', '[]', 'I''m permanently indebted to him for making one of his manga, Say Hello to Black Jack, public domain, giving me an actual series to put on this site.');

INSERT INTO manga ("name", "other_names", "description", "demo", "start_date", "end_date")
VALUES
('Bleach', '[]', 'One of the big three, started strong, burned out fast.', 'shonen', 997142400, 1471824000), -- no chapters (not very realistic but useful to test)
('One Piece', '[]', 'First, last and greatest of the big three.', 'shonen', 997142400, NULL), -- no genres, "float" chapter num
('Naruto', '[]', 'The highs! So high. The lows! So low.', 'shonen', 937872000, 1415577600),
('Bakuman', '[]', 'A manga in shonen jump about writing manga for publication in shonen jump. Woah.', 'shonen', 1218412800, 1335139200), -- multiple authors
('Berserk', '[]', 'The GOAT, abandoned by its creator.', 'seinen', 623203200, NULL), -- multiple magazines
('Blade of the Immortal', '["Mugen no Junin", "The Inhabitant of Infinity"]', 'A rare classic that started strong, middled strong, and ended strong.', 'seinen', 740793600, 1356393600), -- one chapter
('SLAM DUNK', '["Slam Dunk"]', 'Sports movie; the manga!', 'shonen', 654739200, 834969600), -- same author
('Vagabond', '["Bagabondo"]', 'Described as about "life and death, the human condition, etc."', 'shonen', 905990400, NULL), -- same author
('REAL', '["Riaru", "Real"]', 'Followed up a smash hit basketball shonen sports manga with a wheelchair basketball seien drama manga.', 'seinen', 915148800, NULL), -- same author
('Genshiken', '[]', 'About Japanese college nerds who love anime. Inherently relatable.', 'seinen', 1019692800, 1148515200),
('Ashita no Joe', '["Tomorrow''s Joe"]', 'The father of all boxing manga.', 'shonen', -63158400, 106099200),
('Notari Matsutarō', '["Carefree Matsutarō"]', 'Never read it, just need a mangaka with single and joint roles.', 'seinen', 113011200, 888710400),
('Say Hello to Black Jack', '["Burakku Jakku ni Yoroshiku", "Give My Regards to Black Jack"]', 'Remarkably, the author made this public domain in 2012', 'seinen', 1009843200, 1136073600);

INSERT INTO chapter ("manga_id", "chapter_num", "name", "page_count")
VALUES
(2, 1, 'Romance Dawn', 19),
(2, 2, 'They Call Him "Straw Hat Luffy"', 20),
(2, 2.5, NULL, 21),
(3, 1, NULL, 19),
(3, 2, NULL, 20),
(4, 1, NULL, 19),
(4, 2, NULL, 20),
(5, 1, NULL, 19),
(5, 2, NULL, 20),
(6, 1, 'Blood of One Thousand', 60),
(7, 1, NULL, 19),
(7, 2, NULL, 20),
(8, 1, NULL, 19),
(8, 2, NULL, 20),
(9, 1, NULL, 19),
(9, 2, NULL, 20),
(10, 1, NULL, 19),
(10, 2, NULL, 20),
(11, 1, NULL, 19),
(11, 2, NULL, 20),
(12, 1, NULL, 19),
(12, 2, NULL, 20);

INSERT INTO genre
VALUES
('action'),
('adventure'),
('battle manga'),
('drama'),
('fantasy'),
('historical'),
('martial arts'),
('meta'),
('NC-17'),
('slice of life'),
('sports'),
('supernatural'),
('no associations');

INSERT INTO magazine_manga
VALUES
(1, 1),
(1, 2),
(1, 3),
(1, 4),
(5, 5),
(6, 5),
(3, 6),
(1, 7),
(2, 8),
(7, 9),
(3, 10),
(8, 11),
(9, 12),
(2, 13);

INSERT INTO manga_mangaka_job
VALUES
(1, 1, 'author_artist'),
(2, 2, 'author_artist'),
(3, 3, 'author_artist'),
(4, 4, 'author'),
(4, 5, 'artist'),
(5, 6, 'author_artist'),
(6, 7, 'author_artist'),
(7, 8, 'author_artist'),
(8, 8, 'author_artist'),
(9, 8, 'author_artist'),
(10, 9, 'author_artist'),
(11, 10, 'author'),
(11, 11, 'artist'),
(12, 11, 'author_artist'),
(13, 13, 'author_artist');

INSERT INTO manga_genre
VALUES
(1, 'action'),
(1, 'battle manga'),
(1, 'supernatural'),
(3, 'action'),
(3, 'adventure'),
(3, 'battle manga'),
(4, 'meta'),
(4, 'slice of life'),
(5, 'action'),
(5, 'fantasy'),
(5, 'NC-17'),
(5, 'supernatural'),
(6, 'action'),
(6, 'historical'),
(6, 'martial arts'),
(7, 'slice of life'),
(7, 'sports'),
(8, 'historical'),
(8, 'martial arts'),
(9, 'drama'),
(9, 'slice of life'),
(9, 'sports'),
(10, 'drama'),
(10, 'meta'),
(11, 'sports'),
(11, 'drama'),
(12, 'sports'),
(13, 'drama');
