package scraper

import (
	"log"
	"regexp"
	"strings"

	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"

	"github.com/kazGear/portfolio/goBatch/internal/crawler/model"
	C "github.com/kazGear/portfolio/goBatch/pkg/constants"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

// 案件構造体の構築フレームワーク
func buildJobFrame(data map[string]string, url string, logger *log.Logger) (*model.Job) {
    job  := model.Job{}
    trim := utils.TrimSpace()

	job.Title = trim(data[C.Title])
	job.Url = trim(url)
	job.Description = data[C.Description]

	return &job
}

var whitespaceRegex = regexp.MustCompile(`\s+`)

// 文字列をスキル検索用に正規化する
func normalizeForSkillSearch(str string) string {
	normalized := norm.NFKC.String(str) // Unicode正規化（全角英数字・互換文字対策）
	normalized = width.Narrow.String(normalized)
	normalized = strings.ToLower(normalized)
	normalized = strings.ReplaceAll(normalized, "\r\n", "\n") // 改行コード統一
	normalized = strings.TrimSpace(normalized)
	normalized = whitespaceRegex.ReplaceAllString(normalized, " ") // 連続する空白・改行・タブを1スペースへv

	return normalized
}

type Skill struct {
    Name     string
    Category string
    Keywords []string // 大文字・小文字等で区分けする必要はない（検索対象が正規化済の前提）
    Patterns []*regexp.Regexp // 短いキーワードの誤検出防止用
}

var skillsLanguageDictionary = []*Skill{
    {
        Name:     "HTML",
        Category: C.Language,
        Keywords: []string{
            "html",
            "html5",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "CSS",
        Category: C.Language,
        Keywords: []string{
            "css",
            "css3",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "JavaScript",
        Category: C.Language,
        Keywords: []string{
            "javascript",
            "js",
            "java script",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bjs\b`),
        },
    },
    {
        Name:     "TypeScript",
        Category: C.Language,
        Keywords: []string{
            "typescript",
            "ts",
            "type script",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bts\b`),
        },
    },
    {
        Name:     "Java",
        Category: C.Language,
        Keywords: []string{},
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bjava\b`),
        },
    },
    {
        Name:     "C#",
        Category: C.Language,
        Keywords: []string{
            "C#",
            "Ｃ＃",
            "csharp",
            "c sharp",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Python",
        Category: C.Language,
        Keywords: []string{
            "python",
            "python3",
            "python 3",
            "python2",
            "python 2",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Go",
        Category: C.Language,
        Keywords: []string{
            "go言語",
            "golang",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bgo\b`),
        },
    },
    {
        Name:     "C",
        Category: C.Language,
        Keywords: []string{
            "c言語",
            "ansi c",
            "ansi-c",
            "c language",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bc\b`),
        },
    },
    {
        Name:     "C++",
        Category: C.Language,
        Keywords: []string{
            "C++",
            "Ｃ＋＋",
            "cpp",
            "c plus plus",
            "c plusplus",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "PHP",
        Category: C.Language,
        Keywords: []string{
            "php",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Ruby",
        Category: C.Language,
        Keywords: []string{
            "ruby",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Kotlin",
        Category: C.Language,
        Keywords: []string{
            "kotlin",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Swift",
        Category: C.Language,
        Keywords: []string{
            "swift",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Objective-C",
        Category: C.Language,
        Keywords: []string{
            "objective-c",
            "objective c",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "VB.NET",
        Category: C.Language,
        Keywords: []string{
            "vb.net",
            "vb .net",
            "vbnet",
            "vb net",
            "visualbasic.NET",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Visual Basic",
        Category: C.Language,
        Keywords: []string{
            "visual basic",
            "visualbasic",
            "vb6",
            "vb 6",
            "visual basic 6",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "VBA",
        Category: C.Language,
        Keywords: []string{
            "vba",
            "visual basic for applications",
            "excel vba",
            "マクロ",
            "ﾏｸﾛ",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "COBOL",
        Category: C.Language,
        Keywords: []string{
            "cobol",
            "cobol言語",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "FORTRAN",
        Category: C.Language,
        Keywords: []string{
            "fortran",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Perl",
        Category: C.Language,
        Keywords: []string{
            "perl",
            "perl言語",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Scala",
        Category: C.Language,
        Keywords: []string{
            "scala",
            "scala言語",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "R",
        Category: C.Language,
        Keywords: []string{
            "r言語",
            "r language",
            "r-language",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\br\b`),
        },
    },
    {
        Name:     "Dart",
        Category: C.Language,
        Keywords: []string{
            "dart",
            "dart言語",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Rust",
        Category: C.Language,
        Keywords: []string{
            "rust",
            "rust言語",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "SQL",
        Category: C.Language,
        Keywords: []string{},
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bsql\b`),
        },
    },
    {
        Name:     "PowerShell",
        Category: C.Language,
        Keywords: []string{
            "powershell",
            "power shell",
            "ps1",
            ".ps1",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Delphi",
        Category: C.Language,
        Keywords: []string{
            "delphi",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "ABAP",
        Category: C.Language,
        Keywords: []string{
            "abap",
            "sap",
            "sap abap",
            "sap-abap",
            "abap言語",
        },
        Patterns: []*regexp.Regexp{},
    },
}

var skillsFrameworkLibraryDictionary = []*Skill{
    {
        Name:     "Spring (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "spring",
            "spring framework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Spring Boot (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "spring boot",
            "springboot",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Spring MVC (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "spring mvc",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Spring Security (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "spring security",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Spring Batch (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "spring batch",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Jakarta EE (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "jakarta ee",
            "java ee",
            "jee",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Struts (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "struts",
            "apache struts",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Hibernate (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "hibernate",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "MyBatis (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "mybatis",
            "ibatis",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Quarkus (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "quarkus",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Micronaut (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "micronaut",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "ASP.NET Core (C#)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "asp.net core",
            "aspnet core",
            "asp net core",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "ASP.NET MVC (C#)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "asp.net mvc",
            "aspnet mvc",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     ".NET Framework (C#)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            ".net framework",
            "dotnet framework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     ".NET (C#)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            ".net",
            "dotnet",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Entity Framework (C#)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "entity framework",
            "entityframework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Entity Framework Core (C#)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "entity framework core",
            "ef core",
            "efcore",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Blazor (C#)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "blazor",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "React (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "react",
            "react.js",
            "reactjs",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Next.js (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "next.js",
            "nextjs",
            "next js",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Vue.js (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "vue.js",
            "vuejs",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bvue\b`),
        },
    },
    {
        Name:     "Nuxt.js (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "nuxt.js",
            "nuxtjs",
            "nuxt js",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Angular (TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "angular",
            "angular.js",
            "angularjs",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "NestJS (TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "nestjs",
            "nest.js",
            "nest js",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Express.js (JavaScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "express.js",
            "expressjs",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bexpress\b`),
        },
    },
    {
        Name:     "jQuery (JavaScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "jquery",
            "jquery.js",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Laravel (PHP)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "laravel",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Symfony (PHP)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "symfony",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "CakePHP (PHP)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "cakephp",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "CodeIgniter (PHP)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "codeigniter",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "FuelPHP (PHP)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "fuelphp",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Phalcon (PHP)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "phalcon",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Slim Framework (PHP)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "slim framework",
            "slimphp",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Yii (PHP)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "yii framework",
            "yii2",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Django (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "django",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Flask (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "flask framework",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bflask\b`),
        },
    },
    {
        Name:     "FastAPI (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "fastapi",
            "fast api",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Tornado (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "tornado",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Bottle (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "bottle framework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Pyramid (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "pyramid framework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Streamlit (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "streamlit",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Ruby on Rails (Ruby)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "ruby on rails",
            "rails framework",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\brails\b`),
        },
    },
    {
        Name:     "Sinatra (Ruby)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "sinatra framework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Hanami (Ruby)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "hanami",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Gin (Go)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "gin framework",
            "gin-gonic",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Echo (Go)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "echo framework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Fiber (Go)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "fiber framework",
            "gofiber",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Beego (Go)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "beego",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Revel (Go)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "revel framework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Flutter (Dart)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "flutter",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "SwiftUI (Swift)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "swiftui",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "UIKit (Swift/Objective-C)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "uikit",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Jetpack Compose (Kotlin/Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "jetpack compose",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Android Jetpack (Kotlin/Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "android jetpack",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "TensorFlow (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "tensorflow",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "PyTorch (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "pytorch",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Keras (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "keras",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "scikit-learn (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "scikit-learn",
            "scikit learn",
            "sklearn",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "LangChain (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "langchain",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "LlamaIndex (Python)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "llamaindex",
            "llama index",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Redux (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "redux",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "React Router (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "react router",
            "react-router",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "React Hook Form (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "react hook form",
            "react-hook-form",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Material UI (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "material ui",
            "mui",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\mui\b`),
        },
    },
    {
        Name:     "Bootstrap (JavaScript/CSS)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "bootstrap",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Tailwind CSS (JavaScript/CSS)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "tailwind css",
            "tailwindcss",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "jQuery UI (JavaScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "jquery ui",
            "jquery ui",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Spring Cloud (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "spring cloud",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Apache Camel (Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "apache camel",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Play Framework (Scala/Java)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "play framework",
            "playframework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Ktor (Kotlin)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "ktor",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Actix Web (Rust)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "actix web",
            "actix-web",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Axum (Rust)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "axum",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Rocket (Rust)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "rocket framework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Warp (Rust)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "warp framework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Electron (JavaScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "electron",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Hono (TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "hono",
            "honojs",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Fastify (JavaScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "fastify",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Koa.js (JavaScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "koa.js",
            "koajs",
            "koa js",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Svelte (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "svelte",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "SvelteKit (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "sveltekit",
            "svelte kit",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "SolidJS (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "solidjs",
            "solid js",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Alpine.js (JavaScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "alpine.js",
            "alpinejs",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Ember.js (JavaScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "ember.js",
            "emberjs",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Backbone.js (JavaScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "backbone.js",
            "backbonejs",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Quasar Framework (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "quasar framework",
            "quasar",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Ionic (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "ionic framework",
            "ionic",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Capacitor (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "capacitor",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Cordova (JavaScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "cordova",
            "apache cordova",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Remix (JavaScript/TypeScript)",
        Category: C.FrameworkLibrary,
        Keywords: []string{
            "remix",
            "remix.run",
        },
        Patterns: []*regexp.Regexp{},
    },
}

var skillsDatabaseDictionary = []*Skill{
    {
        Name:     "PostgreSQL",
        Category: C.Database,
        Keywords: []string{
            "postgresql",
            "postgres",
            "postgre sql",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "MySQL",
        Category: C.Database,
        Keywords: []string{
            "mysql",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "MariaDB",
        Category: C.Database,
        Keywords: []string{
            "mariadb",
            "maria db",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Oracle Database",
        Category: C.Database,
        Keywords: []string{
            "oracle database",
            "oracle db",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\boracle\b`),
        },
    },
    {
        Name:     "SQL Server",
        Category: C.Database,
        Keywords: []string{
            "sql server",
            "microsoft sql server",
            "mssql",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "SQLite",
        Category: C.Database,
        Keywords: []string{
            "sqlite",
            "sqlite3",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Microsoft Access",
        Category: C.Database,
        Keywords: []string{
            "microsoft access",
            "ms access",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "DB2",
        Category: C.Database,
        Keywords: []string{
            "ibm db2",
            "db2 database",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bdb2\b`),
        },
    },
    {
        Name:     "PostgreSQL Compatible Aurora",
        Category: C.Database,
        Keywords: []string{
            "aurora postgresql",
            "amazon aurora postgresql",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon Aurora MySQL",
        Category: C.Database,
        Keywords: []string{
            "aurora mysql",
            "amazon aurora mysql",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon DynamoDB",
        Category: C.Database,
        Keywords: []string{
            "dynamodb",
            "dynamo db",
            "amazon dynamodb",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "MongoDB",
        Category: C.Database,
        Keywords: []string{
            "mongodb",
            "mongo db",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Redis",
        Category: C.Database,
        Keywords: []string{
            "redis",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Apache Cassandra",
        Category: C.Database,
        Keywords: []string{
            "apache cassandra",
            "cassandra",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "CouchDB",
        Category: C.Database,
        Keywords: []string{
            "couchdb",
            "couch db",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Neo4j",
        Category: C.Database,
        Keywords: []string{
            "neo4j",
            "neo 4j",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Elasticsearch",
        Category: C.Database,
        Keywords: []string{
            "elasticsearch",
            "elastic search",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "OpenSearch",
        Category: C.Database,
        Keywords: []string{
            "opensearch",
            "open search",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "InfluxDB",
        Category: C.Database,
        Keywords: []string{
            "influxdb",
            "influx db",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "TimescaleDB",
        Category: C.Database,
        Keywords: []string{
            "timescaledb",
            "timescale db",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Firebird",
        Category: C.Database,
        Keywords: []string{
            "firebird database",
            "firebird sql",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "H2 Database",
        Category: C.Database,
        Keywords: []string{
            "h2 database",
            "h2 database engine",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bh2\b`),
        },
    },
    {
        Name:     "HSQLDB",
        Category: C.Database,
        Keywords: []string{
            "hsqldb",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "CockroachDB",
        Category: C.Database,
        Keywords: []string{
            "cockroachdb",
            "cockroach db",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "TiDB",
        Category: C.Database,
        Keywords: []string{
            "tidb",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Snowflake",
        Category: C.Database,
        Keywords: []string{
            "snowflake",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "BigQuery",
        Category: C.Database,
        Keywords: []string{
            "bigquery",
            "big query",
            "google bigquery",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon Redshift",
        Category: C.Database,
        Keywords: []string{
            "redshift",
            "amazon redshift",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Databricks",
        Category: C.Database,
        Keywords: []string{
            "databricks",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Azure Cosmos DB",
        Category: C.Database,
        Keywords: []string{
            "cosmos db",
            "azure cosmos db",
        },
        Patterns: []*regexp.Regexp{},
    },
}

var skillsCloudDictionary = []*Skill{
    {
        Name:     "Amazon Web Services (AWS)",
        Category: C.Cloud,
        Keywords: []string{
            "amazon web services",
            "aws",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "AWS EC2",
        Category: C.Cloud,
        Keywords: []string{
            "aws ec2",
            "amazon ec2",
            "elastic compute cloud",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "AWS S3",
        Category: C.Cloud,
        Keywords: []string{
            "aws s3",
            "amazon s3",
            "simple storage service",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "AWS Lambda",
        Category: C.Cloud,
        Keywords: []string{
            "aws lambda",
            "amazon lambda",
            "lambda function",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon RDS",
        Category: C.Cloud,
        Keywords: []string{
            "amazon rds",
            "aws rds",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon ECS",
        Category: C.Cloud,
        Keywords: []string{
            "amazon ecs",
            "aws ecs",
            "elastic container service",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon EKS",
        Category: C.Cloud,
        Keywords: []string{
            "amazon eks",
            "aws eks",
            "elastic kubernetes service",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon CloudFront",
        Category: C.Cloud,
        Keywords: []string{
            "cloudfront",
            "amazon cloudfront",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon Route 53",
        Category: C.Cloud,
        Keywords: []string{
            "route 53",
            "route53",
            "amazon route 53",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon VPC",
        Category: C.Cloud,
        Keywords: []string{
            "amazon vpc",
            "aws vpc",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon API Gateway",
        Category: C.Cloud,
        Keywords: []string{
            "api gateway",
            "amazon api gateway",
            "aws api gateway",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon CloudWatch",
        Category: C.Cloud,
        Keywords: []string{
            "cloudwatch",
            "amazon cloudwatch",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon SNS",
        Category: C.Cloud,
        Keywords: []string{
            "amazon sns",
            "aws sns",
            "simple notification service",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Amazon SQS",
        Category: C.Cloud,
        Keywords: []string{
            "amazon sqs",
            "aws sqs",
            "simple queue service",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Google Cloud Platform (GCP)",
        Category: C.Cloud,
        Keywords: []string{
            "google cloud platform",
            "google cloud",
            "gcp",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Google Compute Engine (GCE)",
        Category: C.Cloud,
        Keywords: []string{
            "google compute engine",
            "compute engine",
            "gce",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Google Cloud Storage",
        Category: C.Cloud,
        Keywords: []string{
            "google cloud storage",
            "gcs",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Google Cloud Functions",
        Category: C.Cloud,
        Keywords: []string{
            "google cloud functions",
            "cloud functions",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Google Kubernetes Engine (GKE)",
        Category: C.Cloud,
        Keywords: []string{
            "google kubernetes engine",
            "gke",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Microsoft Azure",
        Category: C.Cloud,
        Keywords: []string{
            "microsoft azure",
            "azure",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Azure Virtual Machines",
        Category: C.Cloud,
        Keywords: []string{
            "azure vm",
            "azure virtual machines",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Azure App Service",
        Category: C.Cloud,
        Keywords: []string{
            "azure app service",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Azure Functions",
        Category: C.Cloud,
        Keywords: []string{
            "azure functions",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Azure Kubernetes Service (AKS)",
        Category: C.Cloud,
        Keywords: []string{
            "azure kubernetes service",
            "aks",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Azure SQL Database",
        Category: C.Cloud,
        Keywords: []string{
            "azure sql database",
            "azure sql",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Cloudflare",
        Category: C.Cloud,
        Keywords: []string{
            "cloudflare",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Vercel",
        Category: C.Cloud,
        Keywords: []string{
            "vercel",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Netlify",
        Category: C.Cloud,
        Keywords: []string{
            "netlify",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Heroku",
        Category: C.Cloud,
        Keywords: []string{
            "heroku",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Firebase",
        Category: C.Cloud,
        Keywords: []string{
            "firebase",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Supabase",
        Category: C.Cloud,
        Keywords: []string{
            "supabase",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "DigitalOcean",
        Category: C.Cloud,
        Keywords: []string{
            "digitalocean",
            "digital ocean",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Oracle Cloud Infrastructure (OCI)",
        Category: C.Cloud,
        Keywords: []string{
            "oracle cloud infrastructure",
            "oci",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "IBM Cloud",
        Category: C.Cloud,
        Keywords: []string{
            "ibm cloud",
        },
        Patterns: []*regexp.Regexp{},
    },
}

var skillsInfrastructureDictionary = []*Skill{
    {
        Name:     "Docker",
        Category: C.Infrastructure,
        Keywords: []string{
            "docker",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Docker Compose",
        Category: C.Infrastructure,
        Keywords: []string{
            "docker compose",
            "docker-compose",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Kubernetes",
        Category: C.Infrastructure,
        Keywords: []string{
            "kubernetes",
            "k8s",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bk8s\b`),
        },
    },
    {
        Name:     "Amazon ECS",
        Category: C.Infrastructure,
        Keywords: []string{
            "ecs",
            "elastic container service",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\baws ecs\b`),
            regexp.MustCompile(`\becs\b`),
        },
    },
    {
        Name:     "Amazon EKS",
        Category: C.Infrastructure,
        Keywords: []string{
            "eks",
            "elastic kubernetes service",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\baws eks\b`),
            regexp.MustCompile(`\beks\b`),
        },
    },
    {
        Name:     "Linux",
        Category: C.Infrastructure,
        Keywords: []string{
            "linux",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Ubuntu",
        Category: C.Infrastructure,
        Keywords: []string{
            "ubuntu",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "CentOS",
        Category: C.Infrastructure,
        Keywords: []string{
            "centos",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Red Hat Enterprise Linux",
        Category: C.Infrastructure,
        Keywords: []string{
            "red hat enterprise linux",
            "rhel",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Nginx",
        Category: C.Infrastructure,
        Keywords: []string{
            "nginx",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Apache HTTP Server",
        Category: C.Infrastructure,
        Keywords: []string{
            "apache http server",
            "apache web server",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bapache\b`),
        },
    },
    {
        Name:     "IIS",
        Category: C.Infrastructure,
        Keywords: []string{
            "internet information services",
            "iis",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Tomcat",
        Category: C.Infrastructure,
        Keywords: []string{
            "apache tomcat",
            "tomcat",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Terraform",
        Category: C.Infrastructure,
        Keywords: []string{
            "terraform",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Ansible",
        Category: C.Infrastructure,
        Keywords: []string{
            "ansible",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Chef",
        Category: C.Infrastructure,
        Keywords: []string{
            "chef",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Puppet",
        Category: C.Infrastructure,
        Keywords: []string{
            "puppet",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Jenkins",
        Category: C.Infrastructure,
        Keywords: []string{
            "jenkins",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "GitHub Actions",
        Category: C.Infrastructure,
        Keywords: []string{
            "github actions",
            "githubaction",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "GitLab CI/CD",
        Category: C.Infrastructure,
        Keywords: []string{
            "gitlab ci",
            "gitlab cicd",
            "gitlab ci/cd",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "GitHub",
        Category: C.Infrastructure,
        Keywords: []string{
            "github",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "CircleCI",
        Category: C.Infrastructure,
        Keywords: []string{
            "circleci",
            "circle ci",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Azure DevOps",
        Category: C.Infrastructure,
        Keywords: []string{
            "azure devops",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Git",
        Category: C.Infrastructure,
        Keywords: []string{},
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bgit\b`),
        },
    },
    {
        Name:     "Subversion (SVN)",
        Category: C.Infrastructure,
        Keywords: []string{
            "subversion",
            "svn",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Apache Kafka",
        Category: C.Infrastructure,
        Keywords: []string{
            "apache kafka",
            "kafka",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "RabbitMQ",
        Category: C.Infrastructure,
        Keywords: []string{
            "rabbitmq",
            "rabbit mq",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Prometheus",
        Category: C.Infrastructure,
        Keywords: []string{
            "prometheus",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Grafana",
        Category: C.Infrastructure,
        Keywords: []string{
            "grafana",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Datadog",
        Category: C.Infrastructure,
        Keywords: []string{
            "datadog",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "New Relic",
        Category: C.Infrastructure,
        Keywords: []string{
            "new relic",
            "newrelic",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "ELK Stack",
        Category: C.Infrastructure,
        Keywords: []string{
            "elk stack",
            "elastic stack",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Logstash",
        Category: C.Infrastructure,
        Keywords: []string{
            "logstash",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Fluentd",
        Category: C.Infrastructure,
        Keywords: []string{
            "fluentd",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Vagrant",
        Category: C.Infrastructure,
        Keywords: []string{
            "vagrant",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "VirtualBox",
        Category: C.Infrastructure,
        Keywords: []string{
            "virtualbox",
            "virtual box",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "VMware",
        Category: C.Infrastructure,
        Keywords: []string{
            "vmware",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Hyper-V",
        Category: C.Infrastructure,
        Keywords: []string{
            "hyper-v",
            "hyper v",
        },
        Patterns: []*regexp.Regexp{},
    },
}

var skillsToolDictionary = []*Skill{
    {
        Name:     "Visual Studio Code",
        Category: C.Tool,
        Keywords: []string{
            "visual studio code",
            "visual studio code editor",
            "vscode",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Visual Studio",
        Category: C.Tool,
        Keywords: []string{
            "visual studio",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "JetBrains IntelliJ IDEA",
        Category: C.Tool,
        Keywords: []string{
            "intellij idea",
            "intellij",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "JetBrains Rider",
        Category: C.Tool,
        Keywords: []string{
            "rider",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "JetBrains PyCharm",
        Category: C.Tool,
        Keywords: []string{
            "pycharm",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "JetBrains WebStorm",
        Category: C.Tool,
        Keywords: []string{
            "webstorm",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Eclipse",
        Category: C.Tool,
        Keywords: []string{
            "eclipse ide",
            "eclipse",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Apache NetBeans",
        Category: C.Tool,
        Keywords: []string{
            "netbeans",
            "apache netbeans",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Xcode",
        Category: C.Tool,
        Keywords: []string{
            "xcode",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Android Studio",
        Category: C.Tool,
        Keywords: []string{
            "android studio",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Postman",
        Category: C.Tool,
        Keywords: []string{
            "postman",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Swagger",
        Category: C.Tool,
        Keywords: []string{
            "swagger",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "OpenAPI",
        Category: C.Tool,
        Keywords: []string{
            "openapi",
            "open api",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Figma",
        Category: C.Tool,
        Keywords: []string{
            "figma",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Adobe XD",
        Category: C.Tool,
        Keywords: []string{
            "adobe xd",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Jira",
        Category: C.Tool,
        Keywords: []string{
            "jira",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Confluence",
        Category: C.Tool,
        Keywords: []string{
            "confluence",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Backlog",
        Category: C.Tool,
        Keywords: []string{
            "backlog",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Redmine",
        Category: C.Tool,
        Keywords: []string{
            "redmine",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Slack",
        Category: C.Tool,
        Keywords: []string{
            "slack",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Microsoft Teams",
        Category: C.Tool,
        Keywords: []string{
            "microsoft teams",
            "teams",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Notion",
        Category: C.Tool,
        Keywords: []string{
            "notion",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Selenium",
        Category: C.Tool,
        Keywords: []string{
            "selenium",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Playwright",
        Category: C.Tool,
        Keywords: []string{
            "playwright",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Cypress",
        Category: C.Tool,
        Keywords: []string{
            "cypress",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "JMeter",
        Category: C.Tool,
        Keywords: []string{
            "jmeter",
            "apache jmeter",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "SonarQube",
        Category: C.Tool,
        Keywords: []string{
            "sonarqube",
            "sonar qube",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Sentry",
        Category: C.Tool,
        Keywords: []string{
            "sentry",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Vim",
        Category: C.Tool,
        Keywords: []string{
            "vim",
            "vi editor",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Neovim",
        Category: C.Tool,
        Keywords: []string{
            "neovim",
            "neo vim",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Nano",
        Category: C.Tool,
        Keywords: []string{
            "nano",
            "nano editor",
            "gnu nano",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Emacs",
        Category: C.Tool,
        Keywords: []string{
            "emacs",
        },
        Patterns: []*regexp.Regexp{},
    },
}

var skillsTestDictionary = []*Skill{
    {
        Name:     "JUnit (Java)",
        Category: C.Test,
        Keywords: []string{
            "junit",
            "junit5",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "NUnit (C#)",
        Category: C.Test,
        Keywords: []string{
            "nunit",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "MSTest (C#)",
        Category: C.Test,
        Keywords: []string{
            "mstest",
            "microsoft unit testing framework",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "xUnit.net (C#)",
        Category: C.Test,
        Keywords: []string{
            "xunit.net",
            "xunit",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "pytest (Python)",
        Category: C.Test,
        Keywords: []string{
            "pytest",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "unittest (Python)",
        Category: C.Test,
        Keywords: []string{
            "python unittest",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "RSpec (Ruby)",
        Category: C.Test,
        Keywords: []string{
            "rspec",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Minitest (Ruby)",
        Category: C.Test,
        Keywords: []string{
            "minitest",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Go testing (Go)",
        Category: C.Test,
        Keywords: []string{
            "go testing",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\btesting\.go\b`),
        },
    },
    {
        Name:     "Testify (Go)",
        Category: C.Test,
        Keywords: []string{
            "testify",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Jest (JavaScript/TypeScript)",
        Category: C.Test,
        Keywords: []string{
            "jest",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Vitest (JavaScript/TypeScript)",
        Category: C.Test,
        Keywords: []string{
            "vitest",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Mocha (JavaScript)",
        Category: C.Test,
        Keywords: []string{
            "mocha",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Chai (JavaScript)",
        Category: C.Test,
        Keywords: []string{
            "chai",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Jasmine (JavaScript)",
        Category: C.Test,
        Keywords: []string{
            "jasmine",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Cypress (JavaScript/TypeScript)",
        Category: C.Test,
        Keywords: []string{
            "cypress",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Playwright (JavaScript/TypeScript)",
        Category: C.Test,
        Keywords: []string{
            "playwright",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Selenium",
        Category: C.Test,
        Keywords: []string{
            "selenium",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Appium",
        Category: C.Test,
        Keywords: []string{
            "appium",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "JMeter",
        Category: C.Test,
        Keywords: []string{
            "jmeter",
            "apache jmeter",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Gatling",
        Category: C.Test,
        Keywords: []string{
            "gatling",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Locust",
        Category: C.Test,
        Keywords: []string{
            "locust",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Postman",
        Category: C.Test,
        Keywords: []string{
            "postman",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Newman",
        Category: C.Test,
        Keywords: []string{
            "newman",
            "postman newman",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "RSpec Capybara (Ruby)",
        Category: C.Test,
        Keywords: []string{
            "capybara",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Cucumber",
        Category: C.Test,
        Keywords: []string{
            "cucumber",
            "cucumber test",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Gherkin",
        Category: C.Test,
        Keywords: []string{
            "gherkin",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Mockito (Java)",
        Category: C.Test,
        Keywords: []string{
            "mockito",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "JMockit (Java)",
        Category: C.Test,
        Keywords: []string{
            "jmockit",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "WireMock",
        Category: C.Test,
        Keywords: []string{
            "wiremock",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "MockK (Kotlin)",
        Category: C.Test,
        Keywords: []string{
            "mockk",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "PHPUnit (PHP)",
        Category: C.Test,
        Keywords: []string{
            "phpunit",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Codeception (PHP)",
        Category: C.Test,
        Keywords: []string{
            "codeception",
        },
        Patterns: []*regexp.Regexp{},
    },
}

var skillsArchitectureDictionary = []*Skill{
    {
        Name:     "MVC Architecture",
        Category: C.Architecture,
        Keywords: []string{},
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bmvc\b`),
        },
    },
    {
        Name:     "MVVM Architecture",
        Category: C.Architecture,
        Keywords: []string{},
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bmvvm\b`),
        },
    },
    {
        Name:     "MVP Architecture",
        Category: C.Architecture,
        Keywords: []string{},
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bmvp\b`),
        },
    },
    {
        Name:     "Layered Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "layered architecture",
            "layer architecture",
            "n tier architecture",
            "n-tier architecture",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Three-Tier Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "three tier architecture",
            "3 tier architecture",
            "three-tier architecture",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Client Server Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "client server architecture",
            "client server",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Monolithic Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "monolithic architecture",
            "monolith architecture",
            "monolithic application",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Microservices Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "microservices",
            "microservice architecture",
            "micro services",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Service Oriented Architecture (SOA)",
        Category: C.Architecture,
        Keywords: []string{
            "service oriented architecture",
            "service-oriented architecture",
            "soa architecture",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Serverless Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "serverless architecture",
            "serverless",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Event Driven Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "event driven architecture",
            "event-driven architecture",
            "event driven",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Message Driven Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "message driven architecture",
            "message-driven architecture",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "REST Architecture",
        Category: C.Architecture,
        Keywords: []string{},
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\brestful\b`),
            regexp.MustCompile(`\brest api\b`),
        },
    },
    {
        Name:     "GraphQL Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "graphql",
            "graphql api",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "gRPC Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "grpc",
            "grpc api",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Hexagonal Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "hexagonal architecture",
            "ports and adapters",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Clean Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "clean architecture",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Domain Driven Design (DDD)",
        Category: C.Architecture,
        Keywords: []string{
            "domain driven design",
            "domain-driven design",
            "ddd",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "CQRS Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "cqrs",
            "command query responsibility segregation",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Event Sourcing",
        Category: C.Architecture,
        Keywords: []string{
            "event sourcing",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "API Gateway Pattern",
        Category: C.Architecture,
        Keywords: []string{
            "api gateway pattern",
            "api gateway architecture",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "BFF (Backend For Frontend)",
        Category: C.Architecture,
        Keywords: []string{
            "backend for frontend",
            "backend-for-frontend",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bbff\b`),
        },
    },
    {
        Name:     "MVC Model2 Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "model2",
            "model 2",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Reactive Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "reactive architecture",
            "reactive programming",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Cloud Native Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "cloud native architecture",
            "cloud native",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "12 Factor App",
        Category: C.Architecture,
        Keywords: []string{
            "12 factor app",
            "twelve factor app",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Multi Tenant Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "multi tenant architecture",
            "multi tenant",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Distributed System Architecture",
        Category: C.Architecture,
        Keywords: []string{
            "distributed system",
            "distributed architecture",
        },
        Patterns: []*regexp.Regexp{},
    },
}

var skillsMethodologyDictionary = []*Skill{
    {
        Name:     "Agile Development",
        Category: C.Methodology,
        Keywords: []string{
            "アジャイル",
            "agile",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Scrum",
        Category: C.Methodology,
        Keywords: []string{
            "スクラム",
            "scrum",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Kanban",
        Category: C.Methodology,
        Keywords: []string{},
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bkanban\b`),
        },
    },
    {
        Name:     "Extreme Programming (XP)",
        Category: C.Methodology,
        Keywords: []string{
            "extreme",
            "extreme programming",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bxp\b`),
        },
    },
    {
        Name:     "Waterfall Model",
        Category: C.Methodology,
        Keywords: []string{
            "ウォーターフォール",
            "waterfall",
            "waterfall model",
            "waterfall development",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "V Model",
        Category: C.Methodology,
        Keywords: []string{
            "v model",
            "v-model",
            "v字モデル",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "DevOps",
        Category: C.Methodology,
        Keywords: []string{
            "devops",
            "dev ops",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "CI/CD",
        Category: C.Methodology,
        Keywords: []string{
            "ci cd",
            "ci/cd",
            "continuous integration",
            "continuous delivery",
            "continuous deployment",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Test Driven Development (TDD)",
        Category: C.Methodology,
        Keywords: []string{
            "テスト駆動",
            "test driven development",
            "test-driven development",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\btdd\b`),
        },
    },
    {
        Name:     "Behavior Driven Development (BDD)",
        Category: C.Methodology,
        Keywords: []string{
            "ふるまい駆動",
            "振る舞い駆動",
            "振舞い駆動",
            "behavior driven development",
            "behaviour driven development",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bbdd\b`),
        },
    },
    {
        Name:     "Domain Driven Design (DDD)",
        Category: C.Methodology,
        Keywords: []string{
            "ドメイン駆動",
            "domain driven design",
            "domain-driven design",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bddd\b`),
        },
    },
    {
        Name:     "Pair Programming",
        Category: C.Methodology,
        Keywords: []string{
            "ペアプロ",
            "ペアープロ",
            "pair programming",
            "pair programing",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Mob Programming",
        Category: C.Methodology,
        Keywords: []string{
            "モブプロ",
            "mob programming",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Code Review",
        Category: C.Methodology,
        Keywords: []string{
            "レビュ",
            "コードレビュ",
            "code review",
            "peer review",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Pull Request Development",
        Category: C.Methodology,
        Keywords: []string{
            "プルリク",
            "pull request",
            "pull requests",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Shift Left Testing",
        Category: C.Methodology,
        Keywords: []string{
            "shift left testing",
            "shift-left testing",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "SRE (Site Reliability Engineering)",
        Category: C.Methodology,
        Keywords: []string{
            "site reliability engineering",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bsre\b`),
        },
    },
    {
        Name:     "ITIL",
        Category: C.Methodology,
        Keywords: []string{
            "itil",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "RUP (Rational Unified Process)",
        Category: C.Methodology,
        Keywords: []string{
            "rational unified process",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\brup\b`),
        },
    },
    {
        Name:     "Lean Development",
        Category: C.Methodology,
        Keywords: []string{
            "lean development",
            "lean software development",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "Rapid Application Development (RAD)",
        Category: C.Methodology,
        Keywords: []string{
            "rapid application development",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\brad\b`),
        },
    },
}

var skillsRoleDictionary = []*Skill{
    {
        Name:     "要件定義",
        Category: C.Role,
        Keywords: []string{
            "requirements definition",
            "requirement definition",
            "要件定義",
            "要件整理",
            "要求分析",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "基本設計",
        Category: C.Role,
        Keywords: []string{
            "basic design",
            "基本設計",
            "外部設計",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "詳細設計",
        Category: C.Role,
        Keywords: []string{
            "detailed design",
            "detail design",
            "詳細設計",
            "内部設計",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "アーキテクト",
        Category: C.Role,
        Keywords: []string{
            "architect",
            "architecture design",
            "アーキテクト",
            "アーキテクチャ設計",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "バックエンドエンジニア",
        Category: C.Role,
        Keywords: []string{
            "backend engineer",
            "backend developer",
            "バックエンド",
            "サーバーサイド",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "フロントエンドエンジニア",
        Category: C.Role,
        Keywords: []string{
            "frontend engineer",
            "frontend developer",
            "front end engineer",
            "フロントエンド",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "フルスタックエンジニア",
        Category: C.Role,
        Keywords: []string{
            "full stack engineer",
            "fullstack engineer",
            "フルスタック",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "インフラエンジニア",
        Category: C.Role,
        Keywords: []string{
            "infrastructure engineer",
            "infra engineer",
            "インフラ",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "クラウドエンジニア",
        Category: C.Role,
        Keywords: []string{
            "cloud engineer",
            "クラウド",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "データエンジニア",
        Category: C.Role,
        Keywords: []string{
            "data engineer",
            "データエンジニア",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "データサイエンティスト",
        Category: C.Role,
        Keywords: []string{
            "data scientist",
            "データサイエンティスト",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "AIエンジニア",
        Category: C.Role,
        Keywords: []string{
            "ai engineer",
            "aiエンジニア",
            "人工知能エンジニア",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "テスター",
        Category: C.Role,
        Keywords: []string{
            "tester",
            "test engineer",
            "テスター",
            "テストエンジニア",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "QAエンジニア",
        Category: C.Role,
        Keywords: []string{
            "qa engineer",
            "quality assurance engineer",
            "qaエンジニア",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bqa\b`),
        },
    },
    {
        Name:     "プロジェクトマネージャー",
        Category: C.Role,
        Keywords: []string{
            "project manager",
            "プロジェクトマネージャー",
            "マネージャー",
            "マネジメント",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bpm\b`),
        },
    },
    {
        Name:     "プロジェクトリーダー",
        Category: C.Role,
        Keywords: []string{
            "project leader",
            "project lead",
            "プロジェクトリーダー",
            "リーダー",
        },
        Patterns: []*regexp.Regexp{
            regexp.MustCompile(`\bpl\b`),
        },
    },
    {
        Name:     "テックリード",
        Category: C.Role,
        Keywords: []string{
            "tech lead",
            "technical lead",
            "テックリード",
            "技術リー",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "スクラムマスター",
        Category: C.Role,
        Keywords: []string{
            "scrum master",
            "スクラムマスタ",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "開発者",
        Category: C.Role,
        Keywords: []string{
            "developer",
            "programmer",
            "開発",
            "プログラマー",
            "プログラミング",
            "製造",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "運用",
        Category: C.Role,
        Keywords: []string{
            "operator",
            "operation engineer",
            "運用",
            "運用エンジニア",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "保守",
        Category: C.Role,
        Keywords: []string{
            "maintenance engineer",
            "保守",
            "保守エンジニア",
        },
        Patterns: []*regexp.Regexp{},
    },
    {
        Name:     "障害対応",
        Category: C.Role,
        Keywords: []string{
            "incident response",
            "troubleshooting",
            "障害対応",
            "障害調査",
            "障害",
        },
        Patterns: []*regexp.Regexp{},
    },
}