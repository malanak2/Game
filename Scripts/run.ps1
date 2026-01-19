Remove-Item -LiteralPath "build/Resources" -Force -Recurse
Remove-Item -LiteralPath "build/app.exe" -Force
Copy-Item -Path "Resources" -Destination "build/Resources" -Recurse
go build -o build/app.exe .
./build/app.exe