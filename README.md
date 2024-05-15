# soglog : A Minimalist Logger based on slog with Google Cloud Support

sog stands for **S**tructured **O**tel **G**oogleCloud, inspired by [clog](https://github.com/nownabe/clog). 
Special thanks to nownabe.

## Key Features
1. **When the log level is set to error, a stack trace will be displayed.**  
   The display of stack traces can be easily controlled at initialization. For example, you can configure it to hide stack traces locally, while enabling them on Google Cloud.
2. **You can add additional information from the context (ctx) to labels.**
3. **You can initialize slog with a one-liner at server startup, allowing you to use its default slog methods throughout your application**

## Note
soglog is strictly designed within the functionalities of the slog to support cloud logging. Therefore, soglog supports only 4 log levels: `Debug`, `Info`, `Warn`, and `Error`.