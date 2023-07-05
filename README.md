## AddressTrail

<!-- ![AddressTrail Logo](web/static/img/trail.png) -->
<img src="web/static/img/trail.png" width="850" height="350">

AddressTrail is a web application that keeps track of user addresses and provides a unique ID to each user. When a user needs to provide their address to a company, they simply provide the unique ID, and the company can fetch the address via an API call and save it in their database. When the user updates their address on the webapp, it triggers API calls to the companies' callback URLs with the updated address associated with the user. The main idea is to allow users to update their address once, and the application will update the address given by the user to every organization. This gives users the flexibility to move without worrying about updating their address everywhere.

### Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)

## Installation

This webapp is built using Docker, Golang, MongoDB, and Redis. To install and run the application, follow these steps:

1. Install [Docker](https://www.docker.com/get-started) on your system.
2. Clone the repository: `git clone https://github.com/anujshah3/AddressTrail.git`
3. Navigate to the project directory: `cd AddressTrail`
4. Create a `.env` file in the project directory with the following content:

```
GOOGLE_CLIENT_ID=`Get your client ID from google`
GOOGLE_CLIENT_SECRET=`Get your client Secret from google`
AUTH_CALLBACK_URL="http://localhost:8080/auth/google/callback"
RANDOM_STRING=`Any Random String`
SESSION_SECRET=`32 length random string`
MONGO_URI="mongodb://admin:admin@mongodb:27017"
```

5. Run the Docker container: `docker-compose up`

The application should now be running on your local machine.

## Usage

1. Open a web browser and navigate to `http://localhost:8080`.
2. Log in using your Google account.
3. Add or update your address in the webapp.
4. Share your unique ID with companies to provide your address.

When you update your address in the webapp, it will automatically update the address for all associated organizations.

## Contributing

If you would like to contribute to this project, please fork and create a new pull request with a new branch with descriptive name. I appreciate your contributions and will review your changes as soon as possible.

## Support and Contact

If you encounter any issues or have questions, please open an issue on the [Issues](https://github.com/anujshah3/AddressTrail/issues) and I will do my best to address your concerns.
