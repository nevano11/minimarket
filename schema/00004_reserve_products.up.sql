CREATE OR REPLACE FUNCTION reserve_products(codes text[])
    RETURNS BOOLEAN
    LANGUAGE plpgsql AS
$$
DECLARE
    is_available BOOLEAN;
BEGIN
    CREATE TEMP TABLE temp_available_products ON COMMIT DROP AS
    WITH required_products_ids AS (SELECT p.id
                                   FROM product p
                                   WHERE p.code = ANY (codes)),
         products_on_storages AS (SELECT DISTINCT sp.product_id,
                                                  sp.id                                     storage_product_id,
                                                  sp.count - COALESCE(r.reserve_size, 0) AS available_count
                                  FROM storage_product sp
                                           JOIN storage s ON sp.storage_id = s.id AND s.is_available = TRUE
                                           LEFT JOIN (SELECt storage_product_id, sum(reserve_size) reserve_size
                                                      FROM reservation
                                                      WHERE is_active = true
                                                      GROUP BY storage_product_id) r ON r.storage_product_id = sp.id),
         max_available AS (SELECT pos.product_id, MAX(pos.available_count) AS max_available
                           FROM required_products_ids rpi2
                                    JOIN products_on_storages pos
                                         ON pos.available_count > 0 AND rpi2.id = pos.product_id
                           GROUP BY pos.product_id)
    SELECT pos2.product_id, min(pos2.storage_product_id) storage_product_id
    FROM max_available ma
             JOIN products_on_storages pos2
                  ON pos2.product_id = ma.product_id AND pos2.available_count = ma.max_available AND
                     ma.max_available > 0
    GROUP BY pos2.product_id;

    SELECT COUNT(*) = array_length(codes, 1) INTO is_available FROM temp_available_products;

    IF is_available = false THEN
        return is_available;
    end if;

    INSERT INTO reservation (request, storage_product_id, reserve_size)
    SELECT null, tap.storage_product_id, 1
    FROM temp_available_products tap;

    RETURN true;
END
$$;