# 11/02/19
# Made by BredeFK

Write-Output "`nChecking for go gets`n===================="

if(!(Test-Path -Path "$env:GOPATH\src\github.com\go-sql-driver\mysql")) {
    go get -u github.com/go-sql-driver/mysql
}

if(!(Test-Path -Path "$env:GOPATH\src\github.com\gorilla\sessions")){
    go get github.com/gorilla/sessions
}

if(!(Test-Path -Path "$env:GOPATH\src\github.com\gorilla\handlers")){
    go get github.com/gorilla/handlers
}

if(!(Test-Path -Path "$env:GOPATH\src\github.com\gorilla\mux")){
    go get -u github.com/gorilla/mux
}

if(!(Test-Path -Path "$env:GOPATH\src\github.com\gorilla\securecookie")){
    go get github.com/gorilla/securecookie
}

if(!(Test-Path -Path "$env:GOPATH\src\github.com\rs\xid")){
    go get github.com/rs/xid
}

if(!(Test-Path -Path "$env:GOPATH\src\golang.org\x\crypto\bcrypt")){
    go get golang.org/x/crypto/bcrypt
}


Write-Output "`nJump to project destination`n==========================="

$start = Get-Location
Set-Location $env:GOPATH\src\github.com\JohanAanesen\NTNU-Bachelor-Management-System-For-CS-Assignments

Write-Output "`nGo Vet`n========"
$result = go vet ./...
If (-not $result) {
    Write-Output "Pass 100%"
} Else {
    Write-Output $result
}

Write-Output "`nGo fmt`n========"
$result = go fmt ./...
If (-not $result) {
    Write-Output "Pass 100%"
} Else {
    Write-Output $result
}

Write-Output "`nGo Lint`n========"
$result = golint ./...
If (-not $result) {
    Write-Output "Pass 100%"
} Else {
    Write-Output $result
}

Write-Output "`nGo Cyclo`n========"
$result = gocyclo ./...
If (-not $result) {
    Write-Output "Pass 100%"
} Else {
    Write-Output $result
}

Write-Output "`nGo test`n========"
go test -cover ./controller/

Write-Output ""

Set-Location $start
