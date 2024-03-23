create or replace function select_products_by_storage(
    storage_code text
)
    returns table
            (
                id    int,
                name  text,
                size  int,
                code  text,
                count numeric
            )
    LANGUAGE plpgsql
as
$$
begin
    return query
        SELECT p.id, p.name, p.size, p.code, sum(cnts.count) AS count
        FROM product p
                 JOIN (SELECT sp.product_id, sp.count - coalesce(res_sum, 0) count
                       FROM storage_product sp
                                JOIN public.storage s on s.id = sp.storage_id AND s.code = storage_code
                                LEFT JOIN (select storage_product_id, sum(reserve_size) res_sum
                                           from reservation
                                           where is_active = true
                                           group by storage_product_id) r ON sp.id = r.storage_product_id) cnts
                      ON cnts.product_id = p.id
        GROUP BY p.id
        HAVING sum(cnts.count) > 0;

end
$$;