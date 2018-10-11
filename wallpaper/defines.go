package wallpaper


var	pixabayAPIUrl = "https://pixabay.com/api/"

// Data : Importing Definition of the Response 
type Data struct {
	Total int `json:"total"`
	TotalHits int `json:"totalHits"`
	Hits []Hits `json:"hits"`
}

// Hits : Response of each hit
type Hits struct {
	ID int	`json:id`
	WebFormatHeight int `json:webformatHeight`
	WebFormatWidth int `json:webformatWidth`
	Likes int `json:likes`
	ImageWidth int `json:imageWidth`
	UserID int `json:user_id`
	Views int `json:views`
	Comments int  `json:comments`
	PageURL string  `json:pageURL`
	ImageHeight int `json:imageHeight`
	WebFormatURL string  `json:webformatURL`
	Type string  `json:type`
	PreviewHeight int  `json:previewHeight`
	Tags string  `json:tags`
	Downloads int  `json:downloads`
	User string  `json:user`
	Favorites int  `json:favorites`
	ImageSize int32 `json:imageSize`
	PreviewWidth int  `json:previewWidth`
	UserImageURL string `json:UserImageURL`
	PreviewURL string `json:previewURL`
	LargeImageURL string `json:largeImageURL`
	FullHDURL string `json:fullHDURL`	
}