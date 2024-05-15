<p align="center">
<img width="350" alt="image" src="https://github.com/s4s7/soglog/assets/41041296/a1acbd3a-81e0-48ef-8c5a-99a23978c9a0">
</p>

<div align="center">
<samp>

# Minimalist Logger <br> based on slog with Otel for GoogleCloud

sog stands for **S**tructured **O**tel **G**oogleCloud, inspired by [clog](https://github.com/nownabe/clog). 
Special thanks to nownabe.

</samp>
</div>


## Key Features
1. **When the log level is set to error, a stack trace will be displayed.**  
   

|                                                                                                                         |                                                                                                                                                                                  |
|-------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| <img width="400" alt="image" src="https://github.com/s4s7/soglog/assets/41041296/de36ba4d-06ef-405a-a097-c6cff5c12b3d"> | The display of stack traces can be easily controlled at initialization. <br>For example, you can configure it to hide stack traces locally, while enabling them on Google Cloud. |

2. **You can add additional information from the context (ctx) to labels.**

|                                                                                                                         |                                                                                                                                                                     |
|-------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| <img width="200" alt="image" src="https://github.com/s4s7/soglog/assets/41041296/d1fddd4f-29a3-4cc6-add1-e2eb7adcfe9c"> | Insert code that corresponds to the following type during initialization <br> ``type LabelFieldInjector func(ctx context.Context) (key, value string, found bool)`` |

3. **You can initialize slog with a one-liner at server startup, allowing you to use its default slog methods throughout your application**

```golang
slog.SetDefault(slog.New(soglog.NewCloudLoggingHandler("YourProjectID", true, nil)
```

## Note
soglog is strictly designed within the functionalities of the slog to support cloud logging. Therefore, soglog supports only 4 log levels: `Debug`, `Info`, `Warn`, and `Error`.
