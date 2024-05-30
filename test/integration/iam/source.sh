
export POSTGRES_CONNSTR="host=localhost port=54321 user=postgres password=postgres dbname=postgres sslmode=disable"
export JWT_HMAC_SECRET="aabbccddeeffgg1122334455667799"
export IAM_MIGRATIONS_DIR="file:$(Get-Location)\scripts\migrations\iam"
export REDIS_CONNSTR="localhost:6379"
export CLIENT_URL="http://localhost:5173"