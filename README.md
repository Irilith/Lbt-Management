# Lbt-Management

_**Lbt**_ stands for Library BitTorrent, _**Lbt**_ also may stands for Liberty. This tool was developed for personal use and may not function as expected for other person (you).

## Support

| Software     |
|---------------|
| qBittorrent   |

## Build Tools
- [Go](https://golang.org/dl/): The Go programming language.
- [Make](https://www.gnu.org/software/make/): A build automation tool. (not required, you can just do `go build path/`)

## Install and Build
  ```bash
  git clone github.com/Irilith/Lbt-Management.git
  cd Lbt-Management
  make
  ```

## Configuration

1. **Enable WebUI in qBittorrent**: Ensure that the WebUI is enabled in qBittorrent settings.
2. **Update Configuration**: Modify the `.env` file with the following details:
   - **QBIT_USERNAME**: Your qBittorrent WebUI username.
   - **QBIT_PASSWORD**: Your qBittorrent WebUI password.
   - **PORT**: The port number for WebUI access.
3. **Qbittorrent Download Settings** (recommended):
- **Enable Torrent Content Layout**: Set to "Create subfolder".
- **Default Torrent Management**: Set to "Automatic".


## Usage
`./bin/LbtM`
