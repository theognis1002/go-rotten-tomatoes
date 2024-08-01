# Rotten Tomatoes Web Crawler

This project is a web crawler built in Go that scrapes movie ratings and details from Rotten Tomatoes and AMC Theatres. The crawler retrieves movie data and stores it in a SQLite database.

## Getting Started

### Prerequisites

1. **Go**: Ensure you have Go installed. You can download it from [the official Go website](https://golang.org/dl/).

2. **GORM**: Go ORM for interacting with the SQLite database.

3. **GoDotEnv**: A Go library for loading environment variables from a `.env` file.

4. **goquery**: A Go library for parsing HTML.

5. **SQLite**: The database used for storing movie data.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/rotten-tomatoes-crawler.git
   cd rotten-tomatoes-crawler
   ```

2. Install Go dependencies:

   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the root directory and add your Twilio credentials:

   ```env
   TWILIO_ACCOUNT_SID=your_twilio_account_sid
   TWILIO_AUTH_TOKEN=your_twilio_auth_token
   ```

### Usage

The project comes with a Makefile to simplify common tasks.

#### Build the Go Binary

- To build the Go binary, run:

  ```bash
  make build
  This will create an executable named rotten-tomatoes.
  ```

#### Run the Web Crawler

- To run the web crawler, use:

  ```bash
  make run
  ```

This will execute the rotten-tomatoes binary, which will scrape the Rotten Tomatoes website for movie ratings and store them in the database.

#### Clean Up Build Files

- To remove the generated binary and clean up, use:

  ```bash
  make clean
  ```

#### Default Target

The default target in the Makefile is to build the binary. Simply running make will invoke the build process.

### Contributing

Feel free to submit issues, pull requests, or suggestions to improve the crawler.
