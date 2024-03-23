INSERT INTO storage (name, code)
VALUES ('Склад 1', '1')
     , ('Склад 2', '2')
     , ('Склад 3', '3');

INSERT INTO storage (name, code, is_available)
VALUES ('Склад 4', '4', false),
       ('Склад пустой', '5', true);

INSERT INTO product (name, size, code)
VALUES ('Футболка пивозавр', 2, '1'),
       ('Футболка CJ', 2, '2'),
       ('Куртка зеленая', 6, '3'),
       ('Куртка синяя', 6, '4'),
       ('Куртка красная', 5, '5'),
       ('Машинка УАЗ патриот', 1, '6'),
       ('Машинка полицейская', 1, '7'),
       ('Машинка скорая', 1, '8');

INSERT INTO storage_product (product_id, storage_id, count)
VALUES (1, 1, 5),
       (1, 2, 5),
       (2, 2, 5),
       (2, 3, 5),
       (2, 4, 500),
       (3, 1, 5),
       (4, 4, 500),
       (4, 2, 5),
       (5, 4, 500),
       (5, 1, 5),
       (6, 1, 5),
       (6, 2, 5),
       (6, 3, 5),
       (6, 4, 500),
       (7, 1, 5),
       (7, 2, 5),
       (7, 4, 500),
       (8, 2, 5),
       (8, 3, 5);

INSERT INTO reservation (request, storage_product_id, reserve_size)
VALUES ('5c342f39-b007-486c-aafa-8d4a68097657', 16, 1);