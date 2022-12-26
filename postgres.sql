CREATE TABLE Artists(
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE Albums(
    id SERIAL PRIMARY KEY,
    artist_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL, 
    price NUMERIC NOT NULL,
    CONSTRAINT FK_album_artist_id FOREIGN KEY (artist_id)
        REFERENCES Artists (id)
);

CREATE TABLE Songs(
    id SERIAL PRIMARY KEY,
    album_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    lyric VARCHAR(1000),
    CONSTRAINT FK_song_album FOREIGN KEY (album_id)
        REFERENCES Albums (id)
);

INSERT INTO
  Artists (name)
VALUES
  ('Tulus'),
  ('Sheila On 7'),
  ('Peterpan');


INSERT INTO
  Albums (artist_id, title, price)
VALUES
  ((select id from Artists where name='Tulus'), 'GAJAH', 35000),
  ((select id from Artists where name='Tulus'), 'MONOKROM', 35000),
  ((select id from Artists where name='Sheila On 7'), 'Sheila On 7', 40000),
  ((select id from Artists where name='Peterpan'), 'Taman Langit', 30000);

INSERT INTO
  Songs (album_id, title)
VALUES
  ((select id from Albums where title='GAJAH'), 'Sepatu'),
  ((select id from Albums where title='GAJAH'),'Jangan Cintai Aku Apa Adanya'),
  ((select id from Albums where title='MONOKROM'), 'Pamit'),
  ((select id from Albums where title='MONOKROM'), 'Manusia Kuat'),
  ((select id from Albums where title='MONOKROM'), 'Ruang Sendiri'),
  ((select id from Albums where title='MONOKROM'), 'Tukar Jiwa'),
  ((select id from Albums where title='Sheila On 7'), 'Kita'),
  ((select id from Albums where title='Sheila On 7'), 'Anugerah Terindah Yang Pernah Kumiliki'),
  ((select id from Albums where title='Sheila On 7'), 'Dan'),
  ((select id from Albums where title='Sheila On 7'), 'J.A.P'),
  ((select id from Albums where title='Taman Langit'), 'Sahabat'),
  ((select id from Albums where title='Taman Langit'), 'Aku dan Bintang'),
  ((select id from Albums where title='Taman Langit'), 'Yang Terdalam'),
  ((select id from Albums where title='Taman Langit'), 'Topeng');
