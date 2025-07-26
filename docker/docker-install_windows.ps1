if (Get-Command docker -ErrorAction SilentlyContinue) {
    Write-Host "Docker is already installed."
} else {
    Write-Host "Docker is not installed. Installing Docker..."

    winget install -e --id Docker.DockerDesktop

    Write-Host "Docker installation command executed."
}
