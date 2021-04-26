SchoolDay
=========

# **Installation**
#### 0. Golang install
- **Windows**
    
    - Click the link below to download and install the Golang installer. \
    [**go1.15.3.windows-amd64.msi**](https://golang.org/dl/go1.15.3.windows-amd64.msi)

    or
    - Download and install the latest version from the official Golang website. \
    [**golang.org**](https://golang.org)

- **Ubuntu** **/** **Debian**
    ```zsh
    # Latest version
    sudo add-apt-repository ppa:ubuntu-lxc/lxd-stable
    sudo apt-get update
    sudo apt-get -y upgrade
    sudo apt-get install golang
    ```

- **Mac OS**
    ```zsh
    brew install golang
    ```

#### Authentication information setting

- **Linux** **/** **Unix**
    ```zsh
    export BOT_TOKEN="<Your discord bot token>"
    export CSAT_DATE="<The day of the K-SAT>"
    export SCHOOL_INFO_KEY="<Your API key>"
    export SCHOOL_SCHEDULE_KEY="<Your API key>"
    export ELS_TIME_TABLE_KEY="<Your API key>"
    export MIS_TIME_TABLE_KEY="<Your API key>"
    export HIS_TIME_TABLE_KEY="<Your API key>"
    export MEAL_SERVICE_DIET_INFO_KEY="<Your API key>"
    export MySQL_CREDENTIAL_USERNAME="<Your MySQL uid>"
    export MySQL_CREDENTIAL_PASSWORD="<Your MySQL pwd>"
    export MySQL_SERVER_URL="<Your MySQL server url>"
    export MySQL_DB_NAME="<Your MySQL DB name>"
    ```

- **Microsoft Windows 10 with Powershell**
    ```powershell
    $Env:BOT_TOKEN="<Your discord bot token>"
    $Env:CSAT_DATE="<The day of the K-SAT>"
    $Env:SCHOOL_INFO_KEY="<Your API key>"
    $Env:SCHOOL_SCHEDULE_KEY="<Your API key>"
    $Env:ELS_TIME_TABLE_KEY="<Your API key>"
    $Env:MIS_TIME_TABLE_KEY="<Your API key>"
    $Env:HIS_TIME_TABLE_KEY="<Your API key>"
    $Env:MEAL_SERVICE_DIET_INFO_KEY="<Your API key>"
    $Env:MySQL_CREDENTIAL_USERNAME="<Your MySQL uid>"
    $Env:MySQL_CREDENTIAL_PASSWORD="<Your MySQL pwd>"
    $Env:MySQL_SERVER_URL="<Your MySQL server url>"
    $Env:MySQL_DB_NAME="<Your MySQL DB name>"
    ```
