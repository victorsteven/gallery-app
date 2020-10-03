package utils


type Image struct {
	ID        uint64     `json:"id"`
	ImgUrl    string     `json:"imgUrl"`
}


//Images are retrieved from cloudinary

func GetImages() []Image {

	return []Image{
		{
			1,
			"https://res.cloudinary.com/chikodi/image/upload/v1601634127/bart-christiaanse-7QFmFdOpdFs-unsplash.jpg",
		},
		{
			2,
			"https://res.cloudinary.com/chikodi/image/upload/v1601634106/lina-castaneda-HcmdstM9IFw-unsplash.jpg",
		},
		{
			3,
			"https://res.cloudinary.com/chikodi/image/upload/v1601634070/jocke-wulcan-NMwgHV1xdHU-unsplash.jpg",
		},
		{
			4,
			"https://res.cloudinary.com/chikodi/image/upload/v1601634066/bailey-mahon-aK3qEYH_nO0-unsplash.jpg",
		},
		{
			5,
			"https://res.cloudinary.com/chikodi/image/upload/v1601634057/karolis-puidokas-3Ruy7rRNevY-unsplash.jpg",
		},
		{
			6,
			"https://res.cloudinary.com/chikodi/image/upload/v1601634049/jake-colling-9O-l0p38gPw-unsplash.jpg",
		},
		{
			7,
			"https://res.cloudinary.com/chikodi/image/upload/v1601633944/jc-gellidon-5SdGN6k8zpQ-unsplash.jpg",
		},
		{
			8,
			"https://res.cloudinary.com/chikodi/image/upload/v1601633704/jaber-ahmed-SIRrK_oox2M-unsplash.jpg",
		},
	}
}
