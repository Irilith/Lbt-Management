# Lbt-Management

_**Lbt**_ stands for Library BitTorrent, _**Lbt**_ also may stands for Liberty. This tool was developed for personal use and may not function as expected for other person (you).

## What it do

The tool will rename parent folder where you put your ebook
for example, when you download a torrent, with folder (subfolder in qbit) it will be similar to this: `Title [Publisher] [Uploader]`, usually i have to rename to `Title` only to keep the directory organization, but, im lazy to do it everytime so, i make this
its will turn `Title [Publisher] [Uploader]` to `Title`, but the file will remain the Publisher and Uploader

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
Copy the `.env.example` to bin and rename it to `.env` and configure it then
`./bin/LbtM`
