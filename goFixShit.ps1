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
$env:PORT="8080"
$env:SQLDB="root:@tcp(127.0.0.1:3306)/cs53"
$env:SESSION_KEY="52 67 166 253 96 202 151 106 65 44 177 84 130 1 208 172 233 228 151 112 132 236 225 112 168 222 202 121 102 43 41 151 54 129 105 1 233 5 77 68 207 10 251 15 252 134 240 64 171 237 177 154 209 203 62 3 116 138 74 175 97 177 16 156"
go test -cover ./internal/handlers/

Write-Output ""

# Go back to where user started
Set-Location $start
