."docker/functions.ps1"
."docker/docker-install.ps1"

$dockerFile = "docker/dockerfile"
$varFileName = "version_setup_variables.sh"
$envFileName = ".env"

$varMap = Read-Value-From-File -file $varFileName
Write-Host "ARG param map:"
$varMap.GetEnumerator() | Sort-Object Name | Format-Table Name, Value -AutoSize

$envMap = Read-Value-From-File -file $envFileName
Write-Host "ENV param map:"
$envMap.GetEnumerator() | Sort-Object Name | Format-Table Name, Value -AutoSize

Replace-Values -dockerFile $dockerFile -file $varFileName -map $varMap
Replace-Values -dockerFile $dockerFile -file $envFileName -map $envMap

docker buildx build -t translate_server_image -f $dockerFile .
