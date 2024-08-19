# ascii-art-web-dockerize



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/almadhoob/ascii-art-web-dockerize">
    <img src="static/logo.png" alt="Logo" width="128" height="128">
  </a>

<h3 align="center">ASCII Art Web Dockerize</h3>

  <p align="center">
    A simple text manipulation web app.
    <br />
    <a href="https://github.com/almadhoob/ascii-art-web-dockerize"><strong>Go to the repo »</strong></a>
    <br />
    <br />
    <a href="https://github.com/almadhoob/ascii-art-web-dockerize/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/almadhoob/ascii-art-web-dockerize/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#authors">Authors</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<!-- <div align="center"><img src="images/screenshot.png" alt="Screenshot"></div> -->
<!-- <br /> -->

The ascii-art-web-dockerize is a program that is written in Go language using its standard library for the backend (net/http) and  the frontend (html/template) to be excuted as a web server with a web page. It gets two inputs before a submit button, then puts an output in the graphic representation of ASCII code. The first input is a text box that accepts any combination of English letters, Arabic numerals, whitespaces, newlines, and regular special characters only. The second input is a radio button which is the banner name that must be: standard, shadow, xor thinkertoy.

<p align="right">(<a href="#ascii-art-web-dockerize">back to top</a>)</p>



### Built With

* [Go programming language](https://go.dev/doc/)
* [Go HTML templates](https://pkg.go.dev/html/template/)
* [Cascading Style Sheets](https://developer.mozilla.org/en-US/docs/Web/CSS/)


<p align="right">(<a href="#ascii-art-web-dockerize">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

You only need a web browser (such as Mozilla Firefox) in addition to an operating system (Linux, macOS, or Windows) along with the following software:
* Docker Desktop, which is [available from this link](https://docs.docker.com/desktop/).

### Installation

1. Clone the repo
    ```sh
    git clone https://github.com/almadhoob/ascii-art-web-dockerize.git
    cd ascii-art-web-dockerize
    ```

2. Build the image
    ```sh
    docker image build --file Dockerfile --tag docker-ascii .
    ```

3. Run the container
    ```sh
    docker container run --publish 8080:8080 -d docker-ascii
    ```

<p align="right">(<a href="#ascii-art-web-dockerize">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Open the web page through: [http://127.0.0.1:8080](http://127.0.0.1:8080)

<p align="right">(<a href="#ascii-art-web-dockerize">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Add appropriate HTTP status codes.
- [x] Stylise the input form using CSS.
- [x] Containerise everything in Docker.
- [ ] ...

See the [open issues](https://github.com/almadhoob/ascii-art-web-dockerize/issues) for a full list of proposed features (and known bugs).

<p align="right">(<a href="#ascii-art-web-dockerize">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".

1. Fork the Project
2. Create your Feature Branch (git checkout -b feature/AmazingFeature)
3. Commit your Changes (git commit -m 'Added some AmazingFeature')
4. Push to the Branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

Don't forget to give the project a star! Thanks again!

<p align="right">(<a href="#ascii-art-web-dockerize">back to top</a>)</p>

<!-- AUTHORS -->
## Authors

- Ayoob Aloribi — [LinkedIn](https://bh.linkedin.com/in/ayoob-aloribi/)
- Ahmed Almadhoob — [LinkedIn](https://bh.linkedin.com/in/almadhoob/)

<p align="right">(<a href="#ascii-art-web-dockerize">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* Yaman Al-Masri — [LinkedIn](https://bh.linkedin.com/in/yaman-al-masri-1b2108244/)

<p align="right">(<a href="#ascii-art-web-dockerize">back to top</a>)</p>



<!-- LICENSE -->
## License

This is an [MIT-Licensed](./LICENSE) project which is created by its authors for [Reboot01](https://reboot01.com/).

<p align="right">(<a href="#ascii-art-web-dockerize">back to top</a>)</p>


