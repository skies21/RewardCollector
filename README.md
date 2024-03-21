# Reward Collector
## Launch
To get started with the Reward Collector project, follow these steps:

1. Make sure you have Go installed on your system.
2. Navigate to the project directory in your terminal.
3. Run the following command to tidy up and fetch dependencies:
   ```
    go mod tidy
   ```
4. After dependencies are fetched, run the following command to launch the application:
   ```
   go run main.go
   ```
5. To work, you will need a data .json file in the format:
   ```
   [{"amount":1,"gameType":"ARCADE", "rarity":"COMMON","reward":"1 coin"}]
   ```
The application will create a window with a table that can be filtered by rarity. 



