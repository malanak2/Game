Remove-Item -LiteralPath "build/Resources" -Force -Recurse
Copy-Item -Path "Resources" -Destination "build/Resources" -Recurse