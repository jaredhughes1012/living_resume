
$env:POSTGRES_CONNSTR = "host=localhost port=54321 user=postgres password=postgres dbname=postgres sslmode=disable"
$env:JWT_HMAC_SECRET = "aabbccddeeffgg1122334455667799"
$env:IAM_MIGRATIONS_DIR = "file:$(Get-Location)\scripts\migrations\iam"
$env:REDIS_CONNSTR = "localhost:6379"
$env:CLIENT_URL = "http://localhost:5173"