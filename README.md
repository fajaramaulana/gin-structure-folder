# Gin Clean Architecture

Clean Architecture with [Gin Web Framework](https://gin-gonic.com/)


## Usage

Setup environment variables
- Create env file
  - Key **MYSQL**
  - Key **GIN_MODE**
  - Key **PORT**
  - Key **JWT_SECRET_KEY**

```zsh
MYSQL="[YOUR_DATABASE_USERNAME]:[YOUR_DATABASE_PASS]@tcp(127.0.0.1:[YOUR MYSQL PORT])/[YOUR TABLE]?charset=utf8mb4&parseTime=True&loc=Local"
GIN_MODE=debug
PORT=":[YOUR PORT]"

JWT_SECRET_KEY= [YOUR JWT SECRET KEY]
```
## Package
- [**GIN**](https://gin-gonic.com/)
- [**GORM**](https://gorm.io/index.html)
- [**GODOTENV**](https://github.com/joho/godotenv)
- [**SLUG**](github.com/gosimple/slug)
- [**MYSQL**](github.com/go-sql-driver/mysql)
- [**validator/v10**](github.com/go-playground/validator/v10)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)