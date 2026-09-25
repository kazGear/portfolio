import { mkdir, readFile, writeFile } from "node:fs/promises";
import { fileURLToPath } from "node:url";
import pg from "pg";
import dotenv from "dotenv";
import path, { dirname } from "node:path";

const { Client } = pg;

const scriptDirectory = path.dirname(fileURLToPath(import.meta.url));

dotenv.config({
    path: path.resolve(scriptDirectory, "../../.env.prod")
});

const frontendDirectory = fileURLToPath(new URL("..", import.meta.url));
const appFile = new URL("../src/App.jsx", import.meta.url);
const sitemapFile = fileURLToPath(
    new URL("../dist/sitemap.xml", import.meta.url)
);

// ログイン必須・補助画面は検索結果に出さない。
const nonIndexableRoutes = new Set([
    "/IndexPage",
    "/LoginPage",
    "/ShopPage",
    "/BattlePage",
    "/BattleResultPage",
    "/UserPage",
    "/EditPage",
    "/ErrorPage",
    "/GuitarGalleryPage/:makerCd/:name/:color"
]);

const siteUrl = (
    process.env.SITEMAP_SITE_URL ?? "https://kazapp-trial.com"
).replace(/\/$/, "");

// ==============================
// 固定ルートを取得
// ==============================

const appSource = await readFile(appFile, "utf8");

const routes = [
    ...appSource.matchAll(
        /<Route\b[^>]*\bpath=\{["']([^"']+)["']\}/g
    )
]
    .map(([, path]) => path)
    .filter((path) => !nonIndexableRoutes.has(path));

if (routes.length === 0) {
    throw new Error(
        "sitemap に含める公開ルートが App.jsx から見つかりません。"
    );
}

// ==============================
// DBからギター情報を取得
// ==============================

const dbClient = new Client({
    host: process.env.DB_HOST,
    port: Number(process.env.DB_PORT),
    database: process.env.DB_NAME,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
});

await dbClient.connect();

try {
    const result = await dbClient.query(`
        SELECT
               maker,
               name,
               color
          FROM
               t_guitars
      ORDER BY
               maker,
               name,
               color
             ;
    `);

    // DB → 動的URL
    const guitarRoutes = result.rows.map((guitar) => {
        const makerCd = encodeURIComponent(guitar.maker);
        const name    = encodeURIComponent(guitar.name);
        const color   = encodeURIComponent(guitar.color);

        return `/GuitarGalleryPage/${makerCd}/${name}/${color}`;
    });

    // 固定ルート + ギター詳細ルート
    routes.push(...guitarRoutes);

    console.log(`DBからギター ${result.rows.length} 件を取得しました。`);
} finally {
    await dbClient.end();
}

// ==============================
// XML生成
// ==============================

const escapeXml = (value) => value
    .replaceAll("&", "&amp;")
    .replaceAll("'", "&apos;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;");

const urls = [...new Set(routes)].map((path) => {
    const location =
        path === "/"
            ? `${siteUrl}/`
            : `${siteUrl}${path}`;

    return `  <url><loc>${escapeXml(location)}</loc></url>`;
});


const sitemap = [
    '<?xml version="1.0" encoding="UTF-8"?>',
    '<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">',
    ...urls,
    "</urlset>",
    ""
].join("\n");

await mkdir(dirname(sitemapFile), { recursive: true });
await writeFile(sitemapFile, sitemap, "utf8");

console.log(
    `sitemap.xml を生成しました: ${urls.length} 件 (${frontendDirectory})`
);