# 11/02/19
# Made by BredeFK

Write-Output "`nRunning go mod verify`n===================="
#go mod init
go mod verify

Write-Output "`nJump to project destination`n==========================="

$start = Get-Location
Set-Location $env:GOPATH\src\github.com\JohanAanesen\NTNU-Bachelor-Management-System-For-CS-Assignments

Write-Output "`nGo Vet`n========"
go vet ./webservice/...
go vet ./peerservice/...


Write-Output "`nGo fmt`n========"
go fmt ./webservice/...
go fmt ./peerservice/...

Write-Output "`nGo Lint`n========"
golint ./...


Write-Output "`nGo Cyclo`n========"
gocyclo ./webservice/...
gocyclo ./peerservice/...


Write-Output "`nGo test`n========"
go test ./webservice/... -cover
go test ./peerservice/... -cover

Write-Output ""

Set-Location $start
