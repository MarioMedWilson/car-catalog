# Car Catalog

This project a golang API that provides endpoints for managing Cars catalogs and there types. 

### Installation
1. Clone project
```bash
git clone https://github.com/MarioMedWilson/car-catalog.git
```

2. Change into the project directory
```bash
cd car-catalog
```

3. Install pkg
```bash
go mod download
```

4. Build Project
```bash
go build -o main
```

5. Run Projcet
```bash
./main
```

### Docker option
```bash
docker-compose up
```
Note that if it didn't from first time terminate and rerun


## API Endpoints

#### Car endpoints
| Method   | URL                                      | Description                              |
| -------- | ---------------------------------------- | ---------------------------------------- |
| `GET`    | `/car`                                   | Retrieve all cars.                       |
| `POST`   | `/car`                                   | Create a new car.                        |
| `GET`    | `/car/:id`                               | Retrieves a specific car by ID.          |
| `PUT`    | `/car/:id`                               | Updates a specific car by ID.            |
| `DELETE` | `/car/:id`                               | Deletes a specific car by ID.            |
| `GET`    | `/car/?filter={...}`                     | Retrieve cars with filter.               |

#### Car type endpoints

| Method   | URL                                      | Description                              |
| -------- | ---------------------------------------- | ---------------------------------------- |
| `GET`    | `/car-type`                              | Retrieve all cars types.                 |
| `POST`   | `/car-type`                              | Create a new car type.                   |
| `GET`    | `/car-type/:id`                          | Retrieves a car type by ID.              |
| `PUT`    | `/car-type/:id`                          | Updates a car type by ID.                |
| `DELETE` | `/car-type/:id`                          | Deletes a car type by ID.                |

### ERD Database
![Pic](/erd.png)
