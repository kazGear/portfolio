import { readFile, writeFile } from "node:fs/promises";
import { fileURLToPath } from "node:url";

const frontendDirectory = fileURLToPath(new URL("..", import.meta.url));
const appFile = new URL("../src/App.jsx", import.meta.url);
const sitemapFile = new URL("../dist/sitemap.xml", import.meta.url);

// ログイン必須・補助画面は検索結果に出さない。
const nonIndexableRoutes = new Set([
    "/IndexPage",
    "/LoginPage",
    "/ShopPage",
    "/BattlePage",
    "/BattleResultPage",
    "/UserPage",
    "/EditPage",
    "/ErrorPage"
]);

const siteUrl = (process.env.SITEMAP_SITE_URL ?? "https://kazapp-trial.com").replace(/\/$/, "");
const appSource = await readFile(appFile, "utf8");
const routes = [...appSource.matchAll(/<Route\b[^>]*\bpath=\{["']([^"']+)["']\}/g)]
    .map(([, path]) => path)
    .filter((path) => !nonIndexableRoutes.has(path));

if (routes.length === 0) {
    throw new Error("sitemap に含める公開ルートが App.jsx から見つかりません。");
}

const escapeXml = (value) => value
    .replaceAll("&", "&amp;")
    .replaceAll("'", "&apos;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;");

const urls = [...new Set(routes)].map((path) => {
    const location = path === "/" ? `${siteUrl}/` : `${siteUrl}${path}`;
    return `  <url><loc>${escapeXml(location)}</loc></url>`;
});

const sitemap = [
    '<?xml version="1.0" encoding="UTF-8"?>',
    '<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">',
    ...urls,
    "</urlset>",
    ""
].join("\n");

await writeFile(sitemapFile, sitemap, "utf8");
console.log(`sitemap.xml を生成しました: ${routes.length} 件 (${frontendDirectory})`);
