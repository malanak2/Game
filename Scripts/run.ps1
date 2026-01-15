Remove-Item -LiteralPath "build/Resources" -Force -Recurse
Copy-Item -Path "Resources" -Destination "build/Resources" -Recurse
go build -o build/app.exe .
./build/app.exe