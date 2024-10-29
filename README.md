
# Back-end Coding Challenge - Go API

This project is a back-end API service built with **Go** that provides user and action data handling. The API reads from in-memory JSON data files and exposes several endpoints to retrieve user information, action counts, action probabilities, and referral indexes.

## Project Structure
- **/users.json**: Contains the user data.
- **/actions.json**: Contains the action data.

## API Endpoints

### 1. Fetch a User by ID
- **Endpoint**: `/user/{id}`
- **Method**: GET
- **Description**: Retrieves a user based on the user ID.
- **Example**:
  ```bash
  curl http://localhost:8080/user/1
  ```

### 2. Get Total Number of Actions for a User
- **Endpoint**: `/user/{id}/actions/count`
- **Method**: GET
- **Description**: Returns the total number of actions a user has performed.
- **Example**:
  ```bash
  curl http://localhost:8080/user/1/actions/count
  ```

### 3. Get Next Action Probabilities
- **Endpoint**: `/action/{type}/next`
- **Method**: GET
- **Description**: Retrieves the probability breakdown of possible next actions based on an action type.
- **Example**:
  ```bash
  curl http://localhost:8080/action/REFER_USER/next
  ```

### 4. Get Referral Index for All Users
- **Endpoint**: `/users/referral-index`
- **Method**: GET
- **Description**: Returns the referral index of all users. The referral index is the total number of unique users invited directly or indirectly by each user.
- **Example**:
  ```bash
  curl http://localhost:8080/users/referral-index
  ```

## Installation and Setup

### Prerequisites
- Go 1.19 or higher
- Git

### Steps to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/burakalpbilkay/backend-coding-challenge.git
   cd backend-coding-challenge
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Add the JSON data files (**users.json** and **actions.json**) in the project directory.

4. Run the API server:
   ```bash
   go run main.go
   ```

5. The API will run at `http://localhost:8080`.

## Testing the API

You can test the API using **curl** or **Postman**.

Example for fetching a user by ID:
```bash
curl http://localhost:8080/user/1
curl http://localhost:8080/user/1/actions/count
curl http://localhost:8080/action/REFER_USER/next
curl http://localhost:8080/users/referral-index

```
