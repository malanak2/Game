#Remove-Item -LiteralPath "build/Resources" -Force -Recurse
#Copy-Item -Path "Resources" -Destination "build/Resources" -Recurse
xcopy "Resources" "build\Resources" /e /h /c /i /d /y
