# Docker / scp deploy前の生成物掃除

Write-Host "Cleaning build artifacts..."

# portfolioルート取得
$root = Resolve-Path (Join-Path $PSScriptRoot "..\..")

$targets = @(
    # .NET build artifacts
    "$root\backend\PrivateApi\bin",
    "$root\backend\PrivateApi\obj",

    "$root\backend\PublicApi\bin",
    "$root\backend\PublicApi\obj",

    "$root\backend\CSLib\bin",
    "$root\backend\CSLib\obj",

    "$root\backend\Repository\bin",
    "$root\backend\Repository\obj",

    "$root\batch\AutoBattle\bin",
    "$root\batch\AutoBattle\obj",

    # React build artifacts
    "$root\frontend\dist"
)

foreach ($target in $targets) {
    if (Test-Path $target) {
        Write-Host "Remove: $target"
        Remove-Item $target -Recurse -Force
    }
}

# Zone.Identifier 削除
Write-Host "Removing fake Zone.Identifier files..."

Get-ChildItem -Path $root -Recurse -File |
    Where-Object { $_.Name -like "*.Zone.Identifier" -or $_.Name -like "*Zone.Identifier" } |
    ForEach-Object {
        Write-Host "Remove: $($_.FullName)"
        Remove-Item $_.FullName -Force
    }
Write-Host "Cleaning completed."