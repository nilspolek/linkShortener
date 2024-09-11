# Link Shortener

Link Shortener is a simple service that uses SQLite for storing and managing shortened URLs. This Go-based application provides endpoints to shorten URLs and resolve them back to their original destination.

## Features

Shorten URLs: Convert long URLs into shorter, manageable links.
Resolve Shortened Links: Redirect shortened links to their original URLs.
Delete Shortened Links: Remove a shortened link and ensure it no longer redirects.
## Prerequisites

Go (Go 1.18 or later recommended)
## Start the application

```bash
git clone https://github.com/nilspolek/linkShortener.git
cd linkshortener
make run
```

## Usage

### Access a Link

```http
GET http://localhost:8080/<short>
```
This request redirects to the original URL where short is the shortened link.

### Add a URL (POST request)
```json
{
  "destination": "https://example.com"
}
```
returns
```json
{
  "short": "shortened"
}
```
### Delete a shortened URL (DELETE request):
```json
{
  "destination": "https://example.com"
  "short": ["<shortlink1>, <shortlink2>, ..."]
}
```
## Testing

The project includes tests to verify functionality. Run the tests with:

```bash
make test
```
The tests cover:

GET Request: Verify redirection from shortened link to the original URL.
POST and DELETE Requests: Test creating, resolving, and deleting shortened links.
Troubleshooting

Server Not Starting: Ensure no other services are running on port 8080.
Database Issues: Check logs for errors and ensure the database file has the correct permissions.

## Contributing

### To contribute to this project:

Fork the repository
Create a feature branch:
```bash
git checkout -b feature-branch
git commit -am 'Add new feature'
git push origin feature-branch
```
Open a Pull Request
## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE)
