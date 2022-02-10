# Codenames Teams Generator

A script to generate teams for [Codenames](https://codenames.game/) in Go.

## Build

```bash
go build
```

## Usage

```bash
./generate Tom Dick Harry Bob # At least 3 player names are required
```

```
Tom Dick Harry Bob
Team Red
+------+-----------+-----------+
|      | SPYMASTER | OPERATIVE |
+------+-----------+-----------+
| Tom  | ✓         |           |
| Dick |           | ✓         |
+------+-----------+-----------+

Team Blue
+-------+-----------+-----------+
|       | SPYMASTER | OPERATIVE |
+-------+-----------+-----------+
| Harry | ✓         |           |
| Bob   |           | ✓         |
+-------+-----------+-----------+
```