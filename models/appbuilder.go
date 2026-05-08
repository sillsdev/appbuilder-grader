package models

import "encoding/xml"

type AppDef struct {
	XMLName        xml.Name `xml:"app-definition"`
	Type           string   `xml:"type,attr"`
	ProgramVersion string   `xml:"program-version,attr"`

	ProjectName        string `xml:"project-name"`
	ProjectDescription string `xml:"project-description"`

	AppName         string     `xml:"app-name"`
	ApkFilename     string     `xml:"apk-filename"`
	Package         string     `xml:"package"`
	Version         Version    `xml:"version"`
	MultipleApks    string     `xml:"multiple-apks"`
	AndroidSdk      AndroidSdk `xml:"android-sdk"`
	InstallLocation string     `xml:"install-location"`

	IpaFilename       string     `xml:"ipa-filename"`
	IpaVersion        IpaVersion `xml:"ipa-version"`
	IpaAssetFilename  string     `xml:"ipa-asset-filename"`
	IpaAppType        string     `xml:"ipa-app-type"`
	IpaContainerAudio string     `xml:"ipa-container-audio"`

	PwaManifest PwaManifest `xml:"pwa-manifest"`
	Signing     Signing     `xml:"signing"`
	Resigning   Resigning   `xml:"resigning"`
	Publishing  Publishing  `xml:"publishing"`

	About        About        `xml:"about"`
	AudioSources AudioSources `xml:"audio-sources"`
	Videos       Videos       `xml:"videos"`
	DeepLinking  DeepLinking  `xml:"deep-linking"`
	Expiry       Expiry       `xml:"expiry"`

	Features    Features    `xml:"features"`
	Firebase    Firebase    `xml:"firebase"`
	Fonts       Fonts       `xml:"fonts"`
	ColorScheme string      `xml:"color-scheme"`
	ColorThemes ColorThemes `xml:"color-themes"`
	Colors      []Color     `xml:"colors"`
	Styles      Styles      `xml:"styles"`
	Traits      Traits      `xml:"traits"`

	Books               []Books             `xml:"books"`
	Plans               Plans               `xml:"plans"`
	Layouts             Layouts             `xml:"layouts"`
	Assistant           Assistant           `xml:"assistant"`
	MenuItems           []MenuItems         `xml:"menu-items"`
	Security            Security            `xml:"security"`
	InterfaceLanguages  InterfaceLanguages  `xml:"interface-languages"`
	TranslationMappings TranslationMappings `xml:"translation-mappings"`

	Images       []Images     `xml:"images"`
	AdaptiveIcon AdaptiveIcon `xml:"adaptive-icon"`
}

type Version struct {
	Code string `xml:"code,attr"`
	Name string `xml:"name,attr"`
}

type AndroidSdk struct {
	Min string `xml:"min,attr"`
}

type IpaVersion struct {
	Build string `xml:"build,attr"`
	Name  string `xml:"name,attr"`
}

type PwaManifest struct {
	Name      string  `xml:"name,attr"`
	ShortName string  `xml:"short-name,attr"`
	PwaText   PwaText `xml:"pwa-text"`
	PwaSubDir string  `xml:"pwa-sub-directory"`
}

type PwaText struct {
	Lang      string `xml:"lang,attr"`
	Direction string `xml:"direction,attr"`
}

type Signing struct {
	Keystore      string `xml:"keystore"`
	KeystorePass  string `xml:"keystore-password"`
	Alias         string `xml:"alias"`
	AliasPassword string `xml:"alias-password"`
}

type Resigning struct {
	SigningIdentity     SigningIdentity `xml:"signing-identity"`
	ProvisioningProfile string          `xml:"provisioning-profile"`
}

type SigningIdentity struct {
	Hash  string `xml:"hash,attr"`
	Value string `xml:",chardata"`
}

type Publishing struct {
	Mode           string         `xml:"mode,attr"`
	ProjectUrl     string         `xml:"project-url"`
	GooglePlay     GooglePlay     `xml:"google-play"`
	SyncOptions    SyncOptions    `xml:"sync-options"`
	ScriptureEarth ScriptureEarth `xml:"scripture-earth"`
}

type GooglePlay struct {
	Verify string `xml:"verify,attr"`
}

type SyncOptions struct {
	Audio string `xml:"audio,attr"`
	Video string `xml:"video,attr"`
}

type ScriptureEarth struct {
	Notify string `xml:"notify,attr"`
}

type About struct {
	Enabled  string `xml:"enabled,attr"`
	Filename string `xml:"filename"`
}

type AudioSources struct {
	AudioSource []AudioSource `xml:"audio-source"`
}

type AudioSource struct {
	Id            string `xml:"id,attr"`
	Type          string `xml:"type,attr"`
	Default       string `xml:"default,attr"`
	Name          string `xml:"name"`
	AccessMethods string `xml:"access-methods"`
	Folder        string `xml:"folder"`
	Address       string `xml:"address"`
	Key           string `xml:"key"`
	DamID         string `xml:"dam-id"`
}

type Videos struct {
	Video []Video `xml:"video"`
}

type Video struct {
	Id        string    `xml:"id,attr"`
	Width     string    `xml:"width,attr"`
	Height    string    `xml:"height,attr"`
	Title     string    `xml:"title"`
	Thumbnail string    `xml:"thumbnail"`
	OnlineUrl string    `xml:"online-url"`
	Placement Placement `xml:"placement"`
}

type Placement struct {
	Pos string `xml:"pos,attr"`
	Ref string `xml:"ref,attr"`
}

type DeepLinking struct {
	Enabled string `xml:"enabled,attr"`
	Uri     Uri    `xml:"uri"`
}

type Uri struct {
	Type  string `xml:"type,attr"`
	Value string `xml:"value,attr"`
}

type Expiry struct {
	Type     string `xml:"type,attr"`
	Filename string `xml:"filename"`
}

type Features struct {
	Type    string    `xml:"type,attr"`
	Feature []Feature `xml:"feature"`
}

type Feature struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Firebase struct {
	Features Features `xml:"features"`
}

type Fonts struct {
	FontHandling FontHandling `xml:"font-handling"`
	Font         []Font       `xml:"font"`
}

type FontHandling struct {
	Viewer Viewer `xml:"viewer"`
}

type Viewer struct {
	Type string `xml:"type,attr"`
}

type Font struct {
	Family      string      `xml:"family,attr"`
	FontName    string      `xml:"font-name"`
	DisplayName string      `xml:"display-name"`
	Filename    Filename    `xml:"filename"`
	StyleDecl   []StyleDecl `xml:"style-decl"`
}

type Filename struct {
	Format string `xml:"format,attr"`
	Value  string `xml:",chardata"`
}

type StyleDecl struct {
	Property string `xml:"property,attr"`
	Value    string `xml:"value,attr"`
}

type ColorThemes struct {
	ColorTheme []ColorTheme `xml:"color-theme"`
}

type ColorTheme struct {
	Name    string `xml:"name,attr"`
	Enabled string `xml:"enabled,attr"`
	Default string `xml:"default,attr"`
}

type Color struct {
	Type         string         `xml:"type,attr"`
	Name         string         `xml:"name,attr"`
	ColorMapping []ColorMapping `xml:"color-mapping"`
}

type ColorMapping struct {
	Theme string `xml:"theme,attr"`
	Value string `xml:"value,attr"`
}

type Styles struct {
}

type Traits struct {
	Trait []Trait `xml:"trait"`
}

type Trait struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Books struct {
	Id                   string        `xml:"id,attr"`
	BookCollectionName   string        `xml:"book-collection-name"`
	BookCollectionAbbrev string        `xml:"book-collection-abbrev"`
	BookCollectionDesc   string        `xml:"book-collection-description"`
	Features             Features      `xml:"features"`
	Metadata             Metadata      `xml:"metadata"`
	StylesInfo           StylesInfo    `xml:"styles-info"`
	FontChoice           FontChoice    `xml:"font-choice"`
	BookOrder            BookOrder     `xml:"book-order"`
	WritingSystem        WritingSystem `xml:"writing-system"`
	Book                 []Book        `xml:"book"`
}

type Metadata struct {
	Meta []Meta `xml:"meta"`
}

type Meta struct {
	Name    string `xml:"name,attr"`
	Content string `xml:"content,attr"`
}

type StylesInfo struct {
	TextFont         TextFont         `xml:"text-font"`
	TextSize         TextSize         `xml:"text-size"`
	LineHeight       LineHeight       `xml:"line-height"`
	TextDirection    TextDirection    `xml:"text-direction"`
	NumeralSystem    NumeralSystem    `xml:"numeral-system"`
	VerseNumberStyle VerseNumberStyle `xml:"verse-number-style"`
	Styles           Styles           `xml:"styles"`
}

type TextFont struct {
	Family string `xml:"family,attr"`
}

type TextSize struct {
	Value string `xml:"value,attr"`
}

type LineHeight struct {
	Value string `xml:"value,attr"`
}

type TextDirection struct {
	Value string `xml:"value,attr"`
}

type NumeralSystem struct {
	Value string `xml:"value,attr"`
}

type VerseNumberStyle struct {
	Value string `xml:"value,attr"`
}

type FontChoice struct {
	Type             string   `xml:"type,attr"`
	FontChoiceFamily []string `xml:"font-choice-family"`
}

type BookOrder struct {
	CanonName string `xml:"canon-name"`
}

type WritingSystem struct {
	Code         string       `xml:"code,attr"`
	DisplayNames DisplayNames `xml:"display-names"`
}

type DisplayNames struct {
}

type Book struct {
	Id         string     `xml:"id,attr"`
	Name       string     `xml:"name"`
	Abbrev     string     `xml:"abbrev"`
	Group      string     `xml:"group"`
	SubGroup   string     `xml:"sub-group"`
	FontChoice FontChoice `xml:"font-choice"`
	Filename   string     `xml:"filename"`
	Source     string     `xml:"source"`
	Audio      []Audio    `xml:"audio"`
	Features   Features   `xml:"features"`
	StylesInfo StylesInfo `xml:"styles-info"`
	BookTabs   BookTabs   `xml:"book-tabs"`
}

type Audio struct {
	Chapter        string        `xml:"chapter,attr"`
	Filename       AudioFilename `xml:"filename"`
	TimingFilename string        `xml:"timing-filename"`
}

type AudioFilename struct {
	Src   string `xml:"src,attr"`
	Len   string `xml:"len,attr"`
	Size  string `xml:"size,attr"`
	Value string `xml:",chardata"`
}

type BookTabs struct {
	MainType string `xml:"main-type,attr"`
}

type Plans struct {
	Features Features `xml:"features"`
	Plan     []Plan   `xml:"plan"`
}

type Plan struct {
	Id       string    `xml:"id,attr"`
	Days     string    `xml:"days,attr"`
	Title    string    `xml:"title"`
	Filename string    `xml:"filename"`
	Image    PlanImage `xml:"image"`
}

type PlanImage struct {
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Value  string `xml:",chardata"`
}

type Layouts struct {
	Default string   `xml:"default,attr"`
	Layout  []Layout `xml:"layout"`
}

type Layout struct {
	Mode             string             `xml:"mode,attr"`
	Enabled          string             `xml:"enabled,attr"`
	LayoutCollection []LayoutCollection `xml:"layout-collection"`
	Features         Features           `xml:"features"`
}

type LayoutCollection struct {
	Id string `xml:"id,attr"`
}

type Assistant struct {
	Enabled   string      `xml:"enabled,attr"`
	Providers Providers   `xml:"providers"`
	Heading   Translation `xml:"heading"`
	Footer    Translation `xml:"footer"`
	Features  Features    `xml:"features"`
}

type Providers struct {
	Provider []Provider `xml:"provider"`
}

type Provider struct {
	Service string `xml:"service,attr"`
	Enabled string `xml:"enabled,attr"`
}

type Translation struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type MenuItems struct {
	Type     string     `xml:"type,attr"`
	MenuItem []MenuItem `xml:"menu-item"`
}

type MenuItem struct {
	Type   string `xml:"type,attr"`
	Title  string `xml:"title"`
	Link   string `xml:"link"`
	LinkId string `xml:"link-id"`
	Images Images `xml:"images"`
}

type Security struct {
	Mode     string           `xml:"mode,attr"`
	Features SecurityFeatures `xml:"features"`
	Pin      string           `xml:"pin"`
}

type SecurityFeatures struct {
	Feature []Feature `xml:"feature"`
}

type InterfaceLanguages struct {
	UseSystemLanguage string         `xml:"use-system-language,attr"`
	WritingSystems    WritingSystems `xml:"writing-systems"`
}

type WritingSystems struct {
	WritingSystem []WritingSystemItem `xml:"writing-system"`
}

type WritingSystemItem struct {
	Code string `xml:"code,attr"`
	Type string `xml:"type,attr"`
}

type TranslationMappings struct {
	DefaultLang        string               `xml:"default-lang,attr"`
	TranslationMapping []TranslationMapping `xml:"translation-mapping"`
}

type TranslationMapping struct {
	Id          string            `xml:"id,attr"`
	Translation []TranslationItem `xml:"translation"`
}

type TranslationItem struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

type Images struct {
	Type  string      `xml:"type,attr"`
	Image []ImageItem `xml:"image"`
}

type ImageItem struct {
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Value  string `xml:",chardata"`
}

type AdaptiveIcon struct {
	Foreground Foreground `xml:"foreground"`
	Background Background `xml:"background"`
}

type Foreground struct {
	Inset string `xml:"inset,attr"`
}

type Background struct {
}
