package tilda_go

import "time"

type (
	Project struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"descr"`
	}

	ProjectInfo struct {
		ID                 string    `json:"id"`
		UserID             string    `json:"userid"`
		Date               time.Time `json:"date"`
		Title              string    `json:"title"`
		Description        string    `json:"descr"`
		Img                string    `json:"img"`
		Sort               string    `json:"sort"`
		Alias              string    `json:"alias"`
		IndexpageID        string    `json:"indexpageid"`
		HeaderpageID       string    `json:"headerpageid"`
		FooterpageID       string    `json:"footerpageid"`
		HeadlineFont       string    `json:"headlinefont"`
		TextFont           string    `json:"textfont"`
		HeadlineColor      string    `json:"headlinecolor"`
		TextColor          string    `json:"textcolor"`
		LinkColor          string    `json:"linkcolor"`
		LinkFontWeight     string    `json:"linkfontweight"`
		LinkLineColor      string    `json:"linklinecolor"`
		LinkLineHeight     string    `json:"linklineheight"`
		LineColor          string    `json:"linecolor"`
		BgColor            string    `json:"bgcolor"`
		GoogleAnalyticsID  string    `json:"googleanalyticsid"`
		GoogleTmID         string    `json:"googletmid"`
		CustomDomain       string    `json:"customdomain"`
		URL                string    `json:"url"`
		IsExample          string    `json:"isexample"`
		TextFontSize       string    `json:"textfontsize"`
		TextFontWeight     string    `json:"textfontweight"`
		HeadlineFontWeight string    `json:"headlinefontweight"`
		NoSearch           string    `json:"nosearch"`
		YandexMetrikaID    string    `json:"yandexmetrikaid"`
		ExportImgPath      string    `json:"export_imgpath"`
		ExportCssPath      string    `json:"export_csspath"`
		ExportJsPath       string    `json:"export_jspath"`
		ExportBasePath     string    `json:"export_basepath"`
		ViewLogin          string    `json:"viewlogin"`
		ViewPassword       string    `json:"viewpassword"`
		ViewIPs            string    `json:"viewips"`
		Copyright          string    `json:"copyright"`
		Headcode           string    `json:"headcode"`
		UserPayment        string    `json:"userpayment"`
		FormsKey           string    `json:"formskey"`
		InfoType           string    `json:"info_type"`
		InfoTags           string    `json:"info_tags"`
		Page404ID          string    `json:"page404id"`
		MyfontsJSON        string    `json:"myfonts_json"`
		IsEmail            string    `json:"is_email"`
		Kind               string    `json:"kind"`
		Blocked            string    `json:"blocked"`
		Trash              string    `json:"trash"`
		CntFolders         string    `json:"cnt_folders"`
		CntCollabs         string    `json:"cnt_collabs"`
		Collabs            string    `json:"collabs"`
		DesignerIDn        string    `json:"designeridn"`
		Changed            string    `json:"changed"`
		Images             []Image   `json:"images"`
	}

	ExportedProject struct {
		ID             string    `json:"id"`
		ProjectID      string    `json:"projectid"`
		Date           time.Time `json:"date"`
		Title          string    `json:"title"`
		Description    string    `json:"descr"`
		Img            string    `json:"img"`
		Sort           int       `json:"sort,string"`
		Published      int       `json:"published,string"`
		FeatureImg     string    `json:"featureimg"`
		Alias          string    `json:"alias"`
		Filename       string    `json:"filename"`
		ExportJsPath   string    `json:"export_jspath"`
		ExportCssPath  string    `json:"export_csspath"`
		ExportImgPath  string    `json:"export_imgpath"`
		ExportBasePath string    `json:"export_basepath"`
		ProjectAlias   string    `json:"project_alias"`
		PageAlias      string    `json:"page_alias"`
		ProjectDomain  string    `json:"project_domain"`
		HTML           string    `json:"html"`
		Images         []Image   `json:"images"`
		Js             []JS      `json:"js"`
		Css            []CSS     `json:"css"`
	}

	PageInfo struct {
		Page
		HTML string `json:"html"`
	}

	Image struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	JS struct {
		From  string   `json:"from"`
		To    string   `json:"to"`
		Attrs []string `json:"attrs"`
	}

	CSS struct {
		From string `json:"from"`
		To   string `json:"to"`
	}

	Page struct {
		ID          string   `json:"id"`
		ProjectID   string   `json:"projectid"`
		Title       string   `json:"title"`
		Description string   `json:"descr"`
		Img         string   `json:"img"`
		FeatureImg  string   `json:"featureimg"`
		Alias       string   `json:"alias"`
		Date        DateTime `json:"date"`
		Sort        int      `json:"sort,string"`
		Published   int      `json:"published,string"`
		HTML        string   `json:"html"`
		Filename    string   `json:"filename"`
		JS          []string `json:"js"`
		CSS         []string `json:"css"`
	}

	PageFull struct {
		ID          string   `json:"id"`
		ProjectID   string   `json:"projectid"`
		Title       string   `json:"title"`
		Description string   `json:"descr"`
		Img         string   `json:"img"`
		FeatureImg  string   `json:"featureimg"`
		Alias       string   `json:"alias"`
		Date        DateTime `json:"date"`
		Sort        int      `json:"sort,string"`
		Published   int      `json:"published,string"`
		HTML        string   `json:"html"`
		Filename    string   `json:"filename"`
	}

	PageExport struct {
		ID             string   `json:"id"`
		ProjectID      string   `json:"projectid"`
		Date           DateTime `json:"date"`
		Title          string   `json:"title"`
		Description    string   `json:"descr"`
		Img            string   `json:"img"`
		Sort           int      `json:"sort,string"`
		Published      int      `json:"published,string"`
		FeatureImg     string   `json:"featureimg"`
		Alias          string   `json:"alias"`
		Filename       string   `json:"filename"`
		ExportJSPath   string   `json:"export_jspath"`
		ExportCSSPath  string   `json:"export_csspath"`
		ExportImgPath  string   `json:"export_imgpath"`
		ExportBasePath string   `json:"export_basepath"`
		ProjectAlias   string   `json:"project_alias"`
		PageAlias      string   `json:"page_alias"`
		ProjectDomain  string   `json:"project_domain"`
		HTML           string   `json:"html"`
		Images         []Image  `json:"images"`
		JS             []JS     `json:"js"`
		CSS            []CSS    `json:"css"`
	}
)
