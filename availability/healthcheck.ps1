$response = Invoke-WebRequest -Uri 'https://www.fitcalcs.life'
$statusCode = $response.StatusCode
$statusDescription = $response.StatusDescription
$content = $response.Content

Write-Host " * * * Status Code: $statusCode * * * "
Write-Host " * * * Status Description: $statusDescription * * * "
Write-Host " * * * Content: $content * * * "
