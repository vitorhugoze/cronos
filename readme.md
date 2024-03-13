<h1 align="center">Cronos <img src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=for-the-badge&logo=go&logoColor=white"  style="position: absolute; right: 0px; margin-top: 7px"></h1>

## Description

**Backup Scheduler** is a comprehensive backup management solution built with Cobra and Viper. It consists of a client-line tool for creating the schedules and a service for executing the backups and deletions.

### Features

- **Backup Scheduling**: Schedule a backup for a specific date or time interval.
- **Multiple Schedules**: Have multiple schedules at a time for multiple folders.
- **Backup Deletion**: Add a backup deletion period to automatically manage old backups.

## Usage

#### Some examples to break it down for you

Will backup the content of the C:/source and copy it into the C:/dest path just once at the defined time.

```
cronos add C:/source C:/dest -t=2024-03-11T08:35
```

Will backup the content of the C:/source and copy it into the C:/dest path every 7 day.

```
cronos add C:/source C:/dest -i=07-00-00
```

Will backup the content of the C:/source and copy it into the C:/dest path every 7 day and delete it afther 3 days

```
cronos add C:/source C:/dest -i=07-00-00 -d=03-00-00
```

## Installation

#### Windows

- **Pre requisites**: Golang and Nssm

1- Navigate to the cronos-cli path and build it to create an executable:

```
go build -o C:\folder\
```

2- Add the path on wich you built to the PATH enviroment variable, so the cli starts with windows:

```
setx /M PATH "%PATH%;C:\folder\"
```

3- Set the enviroment variable for the scheadules config file path:

```
setx /M CRONOS_CONFIG_PATH C:\folder\
```

4- Navigate to cronos-service and build it to create an executable:

```
go build -o C:\folder\
```

5- Create a service using the binary you built for the service:

```
nssm install "Cronos Service" C:\folder\main.exe
```

6- Start the service:

```
nssm start "Cronos Service"
```
