# 11/02/19
# Made by BredeFK

Write-Host "`nJump to project destination`n===========================" -ForegroundColor Green

$start = Get-Location
Set-Location $env:GOPATH\src\github.com\JohanAanesen\CSAMS

Write-Host "`nRunning go mod verify`n====================" -ForegroundColor Green
go mod verify

Write-Host "`nGo Vet`n========" -ForegroundColor Cyan
go vet ./webservice/...
go vet ./peerservice/...
go vet ./mailservice/...


Write-Host "`nGo fmt`n========" -ForegroundColor Cyan
go fmt ./webservice/...
go fmt ./peerservice/...
go fmt ./mailservice/...

Write-Host "`nGo Lint`n========" -ForegroundColor Cyan
golint ./...

Write-Host "`nGo Cyclo`n========" -ForegroundColor Cyan
gocyclo ./webservice/...
gocyclo ./peerservice/...
gocyclo ./mailservice/...


Write-Host "`nGo test`n========" -ForegroundColor Yellow
go test ./webservice/... -cover
go test ./peerservice/... -cover
go test ./mailservice/... -cover

Write-Host ""

Set-Location $start
