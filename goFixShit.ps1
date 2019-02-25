# 11/02/19
# Made by BredeFK

Write-Output "`nRunning go mod verify`n===================="
#go mod init
go mod verify

Write-Output "`nJump to project destination`n==========================="

$start = Get-Location
Set-Location $env:GOPATH\src\github.com\JohanAanesen\NTNU-Bachelor-Management-System-For-CS-Assignments

Write-Output "`nGo Vet`n========"
$result = go vet ./...
If (-not$result)
{
    Write-Output "Pass 100%"
}
Else
{
    Write-Output $result
}

Write-Output "`nGo fmt`n========"
$result = go fmt ./...
If (-not$result)
{
    Write-Output "Pass 100%"
}
Else
{
    Write-Output $result
}

Write-Output "`nGo Lint`n========"
$result = golint ./...
If (-not$result)
{
    Write-Output "Pass 100%"
}
Else
{
    Write-Output $result
}

Write-Output "`nGo Cyclo`n========"
$result = gocyclo ./...
If (-not$result)
{
    Write-Output "Pass 100%"
}
Else
{
    Write-Output $result
}

Write-Output "`nGo test`n========"
go test -cover ./controller/

Write-Output ""

Set-Location $start
