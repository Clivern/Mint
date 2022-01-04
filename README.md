<img src="/assets/logo.jpeg" width="110" />
<br/><br/>
<a href="https://gitlab.com/clivern/mint/-/releases">
  <img src="https://img.shields.io/badge/Version-1.0.1-red.svg" />
</a>
<a href="https://gitlab.com/clivern/mint/-/blob/main/LICENSE">
  <img src="https://img.shields.io/badge/LICENSE-MIT-green.svg" />
</a>


## Mint

A URL Lookup Service in Golang. It uses [echo framework](https://echo.labstack.com/) and MySQL as a database for malware URLs.

## Documentation

### Run Locally

- Download the latest [mint binary](https://gitlab.com/Clivern/mint/-/releases).
- Create a config file from [config.dist.yml](https://gitlab.com/Clivern/mint/-/blob/main/config.dist.yml).
- Run the web service like the following

```zsh
$ ./mint server -c /path/to/config.dist.yml
```
- Use the API as explained in [service API section](https://gitlab.com/Clivern/mint/-/tree/main#service-api)


### Service API

You can start using the service API as explained below.

1. Define the service URL.

```zsh
$ export URL=http://0.0.0.0:8000
```

2. Add `unsafe URL` to the database.

```zsh
$ curl -X POST \
  -H "X-API-Key: 6c68b836-6f8e-465e-b59f-89c1db53afca" \
  -d '{"hostname": "example.com", "port": 443, "uri":"/tutorial/getting-started-with-golang?ref=12342"}' \
  "$URL/api/v1/url"
```

3. Call the service to get `200` success (URL exists), meaning it is a `unsafe URL`.

```zsh
$ curl -X GET \
  "$URL/urlinfo/1/example.com:443/tutorial/getting-started-with-golang?ref=12342" -v
```

4. Call the service to get `404 not found`, meaning it is a `safe URL`.

```zsh
$ curl -X GET \
  "$URL/urlinfo/1/example.com:443/tutorial/getting-started-with-python" -v
```


### Development & Testing

You have to install [golang locally](https://go.dev/dl/) and then use the following command:

```zsh
# Install revive locally since it is needed for sanity checks
$ make install_revive
$ make run
```

To run unit test cases locally or format the code:

```zsh
# Install revive linter
$ make install_revive

# Run all sanity checks
$ make ci

# To auto format the code
$ make format

# For all available commands
$ make
```


### Deploy on Ubuntu

- Provision a new server with `terraform`. Please follow the [steps documented here](deployment/ubuntu/provision/README.md) for `digitalocean` otherwise you will have to provision the server manually.
- To run the project with docker, please follow the [steps documented here](deployment/ubuntu/docker/README.md).
- To run the project on a virtual server with ansible, please follow the [steps documented here](deployment/ubuntu/ansible/README.md).
- To run the project on a virtual server manually, please follow the [steps documented here](deployment/ubuntu/manually/README.md).


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Mint is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our Gitlab project](https://gitlab.com/Clivern/mint/-/releases) for changelogs for each release version of Mint. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://gitlab.com/clivern/mint/-/issues


## Contributing

Please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2022, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Mint** is authored and maintained by [@Clivern](http://github.com/Clivern).
