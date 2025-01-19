# Ascii Art FS

This project is an enhancement of the previous ASCII Art project, now featuring support for selecting and using different ASCII art templates. It utilizes the Go file system (fs) API for banner/template management, while adhering to Go's best practices.

## Repository

[https://learn.zone01oujda.ma/git/abaid/ascii-art-fs](https://learn.zone01oujda.ma/git/abaid/ascii-art-fs)

## Objectives

- Extend functionality to support multiple templates.
- Handle banner/template selection as the second argument.
- Maintain compatibility with the single argument format for backward support.
- Ensure graceful error handling and clear usage instructions for unsupported formats or invalid inputs.

## Features

- Accepts a string and a banner/template name as arguments.
- Displays ASCII art output using the selected banner.
- Includes support for optional templates if implemented.
- Provides clear usage instructions for invalid inputs.

## Usage

The program follows the format:

```bash
go run . [STRING] [BANNER]
```

### Examples

#### Standard Template
```bash
$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
```

#### Shadow Template
```bash
$ go run . "Hello There!" shadow | cat -e
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
```

#### Thinkertoy Template
```bash
$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $
```

## Error Handling

Invalid usage will return the following message:
```bash
Usage: go run . [STRING] [BANNER]
```

### Example of Invalid Usage
```bash
$ go run . "Hello There!"
Usage: go run . [STRING] [BANNER]
```

## Development Requirements

- The project must be written in Go.
- Only standard Go packages are allowed.
- The code should follow Go's best practices.
- Unit testing is highly recommended.

## Learning Goals

- Deepen understanding of the Go file system (fs) API.
- Practice advanced string and data manipulation in Go.
- Explore efficient error handling and robust input validation.

## License
This project is open-source and follows the [MIT License](LICENSE).
