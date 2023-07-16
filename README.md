# gitfolder: ```github folders downloader tool```

![go-github release (latest SemVer)](https://img.shields.io/github/v/release/Niceblueman/gitfolder?sort=semver)

**`Gitfolder`** is a tool that allows users to download a folder from a GitHub repository using a command-line interface. The tool is built on top of the [go-github](https://pkg.go.dev/github.com/google/go-github/v53/github) library, a Go client library for accessing the GitHub API v3.

## Usage

To use the Go Commandline Publisher, follow these steps:

1. **Download the Binary**: You can download the pre-built binaries for different operating systems from the [Releases](https://github.com/Niceblueman/gitfolder/releases) page.

2. **Command Syntax**: The basic command syntax is as follows:

   ```bash
   ./gitfolder <GitHub_Folder_Link>
   ```

   Replace `<GitHub_Folder_Link>` with the URL of the GitHub folder you want to download.

3. **Example**: For example, to download a folder from the following GitHub URL:

   Run the command:

   ```bash
   ./gitfolder https://github.com/username/repository/tree/master/folder/[targeted_folder]
   ```

   The tool will create a folder named `mains` and download the files inside that folder.

## Contributions

We welcome contributions to the Go Commandline Publisher! If you find any issues or have ideas for improvements, feel free to create an issue or submit a pull request. Please follow the guidelines in [`CONTRIBUTING.md`](CONTRIBUTING.md) for contributing.

## Contributors

We appreciate the efforts of all the contributors who have helped make the Go Commandline Publisher better. Special thanks to the following contributors:

- John Doe (@johndoe)
- Jane Smith (@janesmith)

## Build Instructions

To build the Go Commandline Publisher, the following steps are required:

1. Clone the repository:

   ```bash
   git clone https://github.com/Niceblueman/gitfolder.git
   ```

2. Change directory to the repository:

   ```bash
   cd gitfolder
   ```

3. Build the executable:

   ```bash
   go build -v -ldflags "-s -w" -o bin/gitfolder .
   ```

4. Run the executable:

   ```bash
   ./bin/gitfolder <GitHub_Folder_Link>
   ```

## Supported Platforms

The Go Commandline Publisher is supported on the following platforms:

- Linux
- Windows
- macOS

## License

The Gitfolder is distributed under the BSD-style license. See the [LICENSE](./LICENSE) file for details.