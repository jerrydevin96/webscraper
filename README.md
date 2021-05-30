## Webscraper

A simple web scraper written in go using colly  (https://github.com/gocolly/colly) as the scraper framework. The scraper also uses regex for pattern matching and net/http for checking link accessibility.

### Features:

- HTML version detection
- Web page title extraction
- Heading level detection (H1-H6)
- Webpage link extraction and classification
- Webpage link accessibility checks

#### Feature Description:

##### HTML Version Detection:

Fetches and returns the HTML version of the web page. Regex is used to pattern match the HTML doctype of the page to identify the HTML version.

##### Web Page Title Extraction:

The title of the web page is identified and extracted using colly framework. 

##### Heading Level Detection (H1-H6):

The headings in the web page are identified and classified as per levels (H1 - H6) using colly framework.

##### Webpage Link Extraction and Classification:

The links in the page are identified and extracted using colly framework. The extracted links are parsed and classified as internal and external links, by matching link's domain against the input URL's domain.

##### Webpage Link Accessibility Checks:

The links extracted from the webpage are checked for accessibility using net/http. The link checker is multi threaded to improve performance. 

#### Deployment and Testing:

##### Deployment:

The application is packaged into a docker image and is available as a open image in docker hub. To pull and  run the image, use the below command:

```shell
docker run -d -p 8080:8080 jerrydevin96/webscraper:latest
```

The above command will pull and run the latest stable version (v1.0.0) of the webscraper on port 8080. Pre requisite here is that docker needs to be installed in the host and port 8080 should be free and accessible. To deploy the latest version, check for the latest image in the releases section of the repo.

##### Testing:

The application exposes one API endpoint. To test the API, use any tool like postman or Insomnia.

**API URL: ** http://<hostname/ip>:8080/v0/PageDetails

###### End Point Specs:

**End Point: ** http://<host:port>/v0/PageDetails

**End Point Method: ** POST

**Input Data Structure:**

```json
{
	"url":"<url of webpage to analyze>"
}
```

**Output Data Structure (sample): **

```json
{
  "htmlVersion": "HTML 5",
  "pageTitle": "Docker Hub",
  "h1Length": 0,
  "h2Length": 0,
  "h3Length": 0,
  "h4Length": 0,
  "h5Length": 0,
  "h6Length": 0,
  "internalLinks": 0,
  "externalLinks": 0,
  "inaccessibleLinks": 0,
  "additionalInfo": {
    "internalLinks": [],
    "externalLinks": [],
    "inAccessibleLinks": []
  }
}
```

