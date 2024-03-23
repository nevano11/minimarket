CREATE OR REPLACE FUNCTION free_products(codes text[])
    RETURNS BOOLEAN
    LANGUAGE plpgsql AS
$$
DECLARE
    is_all_products_reserved BOOLEAN;
BEGIN
    CREATE TEMP TABLE temp_reservations_by_products ON COMMIT DROP AS
    WITH required_products_ids AS (SELECT p.id
                                   FROM product p
                                   WHERE p.code = ANY (codes))
    SELECT sp.product_id, min(r.id) reservation_id
    FROM reservation r
             JOIN public.storage_product sp on sp.id = r.storage_product_id
        AND r.is_active = true
        AND product_id IN (select * FROM required_products_ids)
        AND reserve_size = 1
    GROUP BY sp.product_id;

    SELECT COUNT(*) = array_length(codes, 1) INTO is_all_products_reserved FROM temp_reservations_by_products;

    IF is_all_products_reserved = false THEN
        return is_all_products_reserved;
    end if;

    UPDATE reservation r
    SET is_active = false
    FROM temp_reservations_by_products trbp
        WHERE trbp.reservation_id = r.id;

    UPDATE storage_product sp
    SET count = count - 1
    FROM temp_reservations_by_products trbp2
    JOIN reservation res ON trbp2.reservation_id = res.id
    WHERE sp.id = res.storage_product_id;

    RETURN true;
END
$$;