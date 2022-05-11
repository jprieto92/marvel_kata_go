# marvel_kata_go
The objective of this project will be to develop an application with an API that fetches information from the Marvel API of the comics that are on sale at the current date, but only the ones that are going to be published this week. The first iteration will be a CLI tool , and the only information we will need to display would be:

- Name of the comic
- Date of publication
- Price
- Link to buy it

The procedure will be, using TDD and test doubles in Go, build a CLI application that shows the comics, first without contacting the API.
Once it is working, the project shall use the actual API to retrieve the data.

The API we can use for this project is:
https://gateway.marvel.com/v1/public/comics?dateDescriptor=thisWeek&limit=100&ts=987&apikey=97f295907072a970c5df30d73d1f3816&hash=abfa1c1d42a73a5eab042242335d805d

The documentation of the API can be found here:
https://developer.marvel.com/docs

As last step, for offline code challenge, the application should expose this data through an API and should be deployable in Kubernetes. The K8s manifests / Helm chart can be saved in the same repository, under a ./deploy folder.

## Solution
### CLI version
In this version you only need to build and run the this tool (cmd/main.go). This tool retrieves the comics published this week until now.

### Microservice version
In this version you have all the configuration files to deploy this tool as a k8s microservice.

This microservice waits your request to answer to you with the comics published this week until now.

