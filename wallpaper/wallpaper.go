package wallpaper

import (
	
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/reujab/wallpaper"
)

var data Data

// DownloadAndSet : Sends request to Pixabay Server and fetches the top result according to query and 
// Sets the wallpaper of the System
func DownloadAndSet(index int, pixabayAPIKey string) error {

	if data.Total != 0 {
		// Find the URL of the Image
		fileURL := data.Hits[index].FullHDURL
		log.Println("URL of the Image for Index : ", index," is : ",fileURL)
		
		log.Println("Changing the Current Wallapaper")
		background, err := wallpaper.Get()
		if err != nil {
			log.Panic("Cannot get the Current Wallpaper : ",err)
			return err
		}
		log.Println("Current wallpaper:", background)
		err = wallpaper.SetFromURL(fileURL) 
		if err!=nil {
		 _ = getTopResponse(pixabayAPIKey)
		}

		log.Println("Command Success")
			return nil
	} else {

		log.Print("Fetching latest Results")
		err := getTopResponse(pixabayAPIKey)
		if err!=nil {
			log.Println("Error while fetchng results : ",err)
		}
	}
	return nil
}

func getTopResponse(pixabayAPIKey string) error {
	log.Println("Generating the query URL")
	
	client := &http.Client{}
	req, err := http.NewRequest("GET", pixabayAPIUrl, nil)
	if err != nil {
        log.Print(err)
        return err
	}
	q := req.URL.Query()
	q.Add("key", pixabayAPIKey)
	q.Add("orientation","horizontal")
	q.Add("order","popular")
	q.Add("image_type","photo")
	q.Add("safesearch","true")
	q.Add("editors_choice","true")
	q.Add("min_width","3543")
	q.Add("min_height","2286")
    req.URL.RawQuery = q.Encode()
    req.Header.Add("Accept", "application/json")
	log.Println(req.URL.String())

	resp, err := client.Do(req)

    if err != nil {
        log.Panic("Errored when sending request to the server", err)
        return err
    }
    defer resp.Body.Close()
    respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        log.Panic("Error in reading the Body : ", err)
        return err
	}
	
    log.Println("Status of call : ",resp.Status)
	// fmt.Println(string(resp_body))
	data = Data{}
	err = json.Unmarshal(respBody , &data)
	if err != nil {
        log.Panic("Error Occured while unmarshaling the response JSON")
        return err
    }

	return nil
}