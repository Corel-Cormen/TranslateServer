if (Get-Command docker -ErrorAction SilentlyContinue) {
    Write-Host "Docker is already installed."
} else {
    $title = "Docker is not installed."
    $message = "Would you like install the newest version?"
    $options = [System.Management.Automation.Host.ChoiceDescription[]]@(
        "&No", "&Yes"
    )
    $default = 0

    $result = $host.ui.PromptForChoice($title, $message, $options, $default)

    if ($result -eq 0) {
        Write-Host "Script is begining installation..."
        winget install -e --id Docker.DockerDesktop
        Write-Host "Docker installation command executed"
    } else {
        Write-Host "Installation rejected."
        exit 1
    }
}
