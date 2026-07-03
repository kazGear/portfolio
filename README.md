# Guitar data API （説明書）

ギター情報を検索・取得するためのREST APIです。

メーカー、シリーズ、カラー、ボディ材、価格帯などの条件で絞り込み検索ができます。
ページネーションやソートにも対応しています。

特徴 -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

メーカー・シリーズ・カラーなど複数条件検索  
部分一致検索（name / series）  
価格帯検索  
ソート機能  
ページネーション対応  
総件数取得  
検索条件なしでも一覧取得可能  

エンドポイント -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

GET /api/v1/guitars

クエリパラメータ -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Parameter,Type,Description

makerCd,int,メーカーコード  
name,string,ギター名（部分一致検索）  
series,string,シリーズ名（部分一致検索）  
colorCd,int,カラーコード  
bodyMaterialTopCd,int,ボディトップ材  
bodyMaterialBackCd,int,ボディバック材  
minPrice,int,最低価格  
maxPrice,int,最高価格  
sort,string,maker / name / price  
order,string,ASC / DESC（必須）  
page,int,ページ番号（必須）  
pageSize,int,1ページの件数（必須）  

使用例 -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

GET /api/v1/guitars?makerCd=1&series=Strat&page=1&pageSize=25

レスポンス例 -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

{  
  "totalCount": 283,  
  "page": 1,  
  "pageSize": 25,  
  "totalPages": 12,  
  "hasPrev": false,  
  "hasNext": true,  
  "guitars": [  
    {  
      "maker": "Fender",  
      "name": "American Professional II Stratocaster",  
      ...  
    },  
    {  
        ...  
    },  
  ]  
}  

検索仕様 -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

name, series は部分一致検索を行います。
検索条件を指定しない場合は、全件取得します。

存在しないページを指定した場合はエラーではなく、空のギター配列を返します。