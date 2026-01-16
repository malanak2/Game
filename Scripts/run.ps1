Remove-Item -LiteralPath "build/Resources" -Force -Recurse
Copy-Item -Path "Resources" -Destination "build/Resources" -Recurse
go build -o build/tapp.exe .
./build/tapp.exe