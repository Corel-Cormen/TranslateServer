function Resolve-Value {
    param (
        [string]$value,
        [hashtable]$map
    )

    $maxDepth = 10
    $pattern = '\$\{([^}]+)\}'

    for ($i = 0; $i -lt $maxDepth; $i++) {
        if ($value -notmatch $pattern) {
            break
        }

        $value = [regex]::Replace($value, $pattern, {
            param($match)
            $varName = $match.Groups[1].Value
            if ($map.ContainsKey($varName)) {
                return $map[$varName]
            } else {
                Write-Error "Value for $varName not exist"
                exit 1
            }
        })
    }

    return $value
}

function Read-Value-From-File {
    param (
        [string]$file
    )

    $fileLines = Get-Content $file

    $valueMap = @{}
    foreach ($line in $fileLines) {
        if ($line -match '^\s*([^=]+)=(.*)$') {
            $key = $matches[1].Trim()
            $val = $matches[2].Trim()
            $valueMap[$key] = $val
        }
    }

    $keys = $valueMap.Keys | ForEach-Object { $_ }
    foreach ($key in $keys) {
        $resolved = Resolve-Value -value $valueMap[$key] -map $valueMap
        $valueMap[$key] = $resolved
    }

    return $valueMap
}

function Replace-Values {
    param (
        [string]$dockerFile,
        [string]$file,
        [hashtable]$map
    )

    $valueText = ""
    foreach ($key in $map.Keys) {
        $valueText += "ARG $key=$($map[$key])`n"
    }

    $fileText = Get-Content $dockerFile -Raw
    $escapedFile = [regex]::Escape($file)
    $pattern = "(?s)(# Start source $file).*?(# End source $file)"
    $replaceContent = [regex]::Replace(
        $fileText,
        $pattern,
        "`$1`n$valueText`$2"
    )
    $replaceContent = $replaceContent.TrimEnd("`r", "`n")

    Set-Content -Path $dockerFile -Value $replaceContent
}
