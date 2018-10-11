# Wallpaper Control


Now you can control your System Wallpaper, Always see the latest wallpapers on your screen.
Steps for you to use : 
  - Go to [Pixabay](https://pixabay.com/) and login.
  - Generate your own API Key.
  - Use the key as command line flag : -pixabay-key \*\*API_KEY**

Steps to build:
```sh
$ go get ***
$ go build -o wc
$ wc -pixabay-key **API_KEY**