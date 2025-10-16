# Get-ProjectTree.ps1
# Generates tree listings for selected project subfolders,
# excluding common bulky directories like node_modules.

# Patterns to exclude from the tree output
$exclude = '\\node_modules\\|\\\.pnpm\\|\\\.next\\|\\dist\\|\\build\\|\\coverage\\|\\\.turbo\\|\\\.git\\'

# Output file
$outFile = "combined_tree.txt"

# Ensure file starts clean
if (Test-Path $outFile) {
    Remove-Item $outFile
}

# Function to run tree and append results with a header
function Add-TreeSection($path) {
    "`n=== $path ===`n" | Add-Content $outFile
    cmd /c "tree $path /f" | Select-String -NotMatch $exclude | Add-Content $outFile
}

# Run for each folder
Add-TreeSection "apps\api"
Add-TreeSection "apps\worker"
Add-TreeSection "infra"
Add-TreeSection "pkg"
#Add-TreeSection "."

Write-Host "Tree output written to $outFile"
