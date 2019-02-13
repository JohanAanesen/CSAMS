# 11/02/19
# Made by BredeFK

Write-Output "`nJump to project destination`n==========================="

# Get current location for later use
$start = Get-Location
Set-Location $env:GOPATH\src\github.com\JohanAanesen\NTNU-Bachelor-Management-System-For-CS-Assignments

# Run Go Vet
Write-Output "`nGo Vet`n========"
$result = go vet ./...
If (-not $result) {
    Write-Output "Pass 100%"
} Else {
    Write-Output $result
}

# Run Go fmt
Write-Output "`nGo fmt`n========"
$result = go fmt ./...
If (-not $result) {
    Write-Output "Pass 100%"
} Else {
    Write-Output $result
}

# Run Go Lint
Write-Output "`nGo Lint`n========"
$result = golint ./...
If (-not $result) {
    Write-Output "Pass 100%"
} Else {
    Write-Output $result
}

# Run Go Cyclo
Write-Output "`nGo Cyclo`n========"
$result = gocyclo ./...
If (-not $result) {
    Write-Output "Pass 100%"
} Else {
    Write-Output $result
}

# Run Go test
Write-Output "`nGo test`n========"
go test -cover ./controller/

Write-Output ""

# Go back to where user started
Set-Location $start
