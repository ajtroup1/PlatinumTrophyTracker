-- 20240804_add_platforms.down.sql

DELETE FROM platforms WHERE name IN (
  'PlayStation', 'PlayStation 2', 'PlayStation 3', 'PlayStation 4', 'PlayStation 5',
  'Xbox', 'Xbox 360', 'Xbox One', 'Xbox Series X',
  'Switch', 'Wii', 'Wii U', 'GameCube',
  'PC'
);
