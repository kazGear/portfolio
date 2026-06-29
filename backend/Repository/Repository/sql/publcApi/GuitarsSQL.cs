namespace Repository.Repository.sql.publicApi;

/// <summary>
/// SQL文格納クラス
/// </summary>
public static class GuitarsSQL
{
    public static string GetGuitars()
    {
        string SQL = @"
            SELECT
                   maker              AS Maker,
                   name               AS Name,
                   color              AS Color,
                   color_cd           AS ColorCd,
                   body_finish        AS BodyFinish,
                   body_material      AS BodyMaterial,
                   body_material_top  AS BodyMaterialTop,
                   body_material_back AS BodyMaterialBack,
                   bridge             AS Bridge,
                   controls           AS Controls,
                   comment            AS Comment,
                   fingerboard        AS Fingerboard,
                   fret_count         AS FretCount,
                   inlays             AS Inlays,
                   joint              AS Joint,
                   neck_material      AS NeckMaterial,
                   pickups            AS Pickups,
                   price              AS Price,
                   scale_length_mm    AS ScaleLengthMm,
                   series             AS Series,
                   src                AS Src,
                   weight             AS Weight
              FROM
                   t_guitars
             WHERE
                   TRUE
    /* 動的検索条件
               AND maker = :maker
               AND name ilike '%' || :name || '%'
               AND series ilike '%' || :series || '%'
               AND color_cd = :color_cd
               AND body_material_top = :body_material_top_cd
               AND body_material_back = :body_material_back_cd
               AND price >= :min_price
               AND price <= :max_price */
      /* 動的ソート
          ORDER BY
                   'columnName' ASC or DESC */
             LIMIT
                   :page_size
            OFFSET
                   (:page - 1) * :page_size -- ページネーション
                 ;
        ";
        return SQL;
    }
}
