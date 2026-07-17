namespace Repository.Repository.sql.publicApi;

/// <summary>
/// SQL文格納クラス
/// </summary>
public static class GuitarsSQL
{
    public static string GetGuitars(string conditions, string sort)
    {
        string SQL = @$"
            SELECT
                   guitars.maker              AS Maker,
                   maker.name                 AS MakerName,
                   guitars.name               AS Name,
                   guitars.color              AS Color,
                   guitars.color_cd           AS ColorCd,
                   guitars.body_finish        AS BodyFinish,
                   guitars.body_material      AS BodyMaterial,
                   guitars.body_material_top  AS BodyMaterialTop,
                   guitars.body_material_back AS BodyMaterialBack,
                   guitars.bridge             AS Bridge,
                   guitars.controls           AS Controls,
                   guitars.comment            AS Comment,
                   guitars.fingerboard        AS Fingerboard,
                   fingerboard.name           AS FingerboardName,
                   guitars.fret_count         AS FretCount,
                   guitars.inlays             AS Inlays,
                   guitars.joint              AS Joint,
                   guitars.neck_material      AS NeckMaterial,
                   neck.name                  AS NeckMaterialName,
                   guitars.pickups            AS Pickups,
                   guitars.price              AS Price,
                   guitars.scale_length_mm    AS ScaleLengthMm,
                   guitars.series             AS Series,
                   guitars.src                AS Src,
                   guitars.weight             AS Weight,
                   substring(guitars.updated::text, 1, 16) AS Updated
              FROM
                   t_guitars AS guitars
        INNER JOIN
                   m_code AS maker
                ON guitars.maker = maker.VALUE
               AND maker.code_id = 'guitar_makers'
        INNER JOIN
                   m_code AS neck
                ON guitars.neck_material = neck.VALUE
               AND neck.code_id          = 'guitar_woods'
        INNER JOIN
                   m_code AS fingerboard
                ON guitars.fingerboard = fingerboard.VALUE
               AND fingerboard.code_id = 'guitar_woods'

             WHERE
                   TRUE
                   {conditions}
        /* 動的検索
                   AND maker = @maker
                   AND name ilike '%' || @name || '%'
                   AND series ilike '%' || @series || '%'
                   AND color_cd = @color_cd
                   AND body_material_top = @body_material_top_cd
                   AND body_material_back = @body_material_back_cd
                   AND price >= @min_price
                   AND price <= @max_price */

                   {sort}
      /* 動的ソート 
          ORDER BY
                   'columnName' ASC or DESC */
             LIMIT
                   @page_size
            OFFSET
                   (@page - 1) * @page_size -- ページネーション
                 ;
        ";
        return SQL;
    }

    public static string GetTotalCount(string conditions)
    {
        string SQL = @$"
            SELECT
                   count(*)
              FROM
                   t_guitars AS guitars
             WHERE
                   TRUE
                   {conditions}
        /* 動的検索
                   AND maker = @maker
                   AND name ilike '%' || @name || '%'
                   AND series ilike '%' || @series || '%'
                   AND color_cd = @color_cd
                   AND body_material_top = @body_material_top_cd
                   AND body_material_back = @body_material_back_cd
                   AND price >= @min_price
                   AND price <= @max_price */
        ";
        return SQL;
    }
}
