package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Identifiable represents the Identifiable schema from the OpenAPI specification
type Identifiable struct {
	TypeField string `json:"_type"`
}

// Intangible represents the Intangible schema from the OpenAPI specification
type Intangible struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
}

// QueryContext represents the QueryContext schema from the OpenAPI specification
type QueryContext struct {
	Alterationdisplayquery string `json:"alterationDisplayQuery,omitempty"` // AlteredQuery that is formatted for display purpose. The query string in the AlterationDisplayQuery can be html-escaped and can contain hit-highlighting characters
	Alterationoverridequery string `json:"alterationOverrideQuery,omitempty"` // The query string to use to force Bing to use the original string. For example, if the query string is "saling downwind", the override query string will be "+saling downwind". Remember to encode the query string which results in "%2Bsaling+downwind". This field is included only if the original query string contains a spelling mistake.
	Alteredquery string `json:"alteredQuery,omitempty"` // The query string used by Bing to perform the query. Bing uses the altered query string if the original query string contained spelling mistakes. For example, if the query string is "saling downwind", the altered query string will be "sailing downwind". This field is included only if the original query string contains a spelling mistake.
	Askuserforlocation bool `json:"askUserForLocation,omitempty"` // A Boolean value that indicates whether Bing requires the user's location to provide accurate results. If you specified the user's location by using the X-MSEdge-ClientIP and X-Search-Location headers, you can ignore this field. For location aware queries, such as "today's weather" or "restaurants near me" that need the user's location to provide accurate results, this field is set to true. For location aware queries that include the location (for example, "Seattle weather"), this field is set to false. This field is also set to false for queries that are not location aware, such as "best sellers".
	Istransactional bool `json:"isTransactional,omitempty"`
	Originalquery string `json:"originalQuery"` // The query string as specified in the request.
	TypeField string `json:"_type"`
	Adultintent bool `json:"adultIntent,omitempty"` // A Boolean value that indicates whether the specified query has adult intent. The value is true if the query has adult intent; otherwise, false.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Subcode string `json:"subCode,omitempty"` // The error code that further helps to identify the error.
	Value string `json:"value,omitempty"` // The parameter's value in the request that was not valid.
	TypeField string `json:"_type"`
	Code string `json:"code"` // The error code that identifies the category of error.
	Message string `json:"message"` // A description of the error.
	Moredetails string `json:"moreDetails,omitempty"` // A description that provides additional information about the error.
	Parameter string `json:"parameter,omitempty"` // The parameter in the request that caused the error.
}

// EntitiesEntityPresentationInfo represents the EntitiesEntityPresentationInfo schema from the OpenAPI specification
type EntitiesEntityPresentationInfo struct {
	Entitytypedisplayhint string `json:"entityTypeDisplayHint,omitempty"` // A display version of the entity hint. For example, if entityTypeHints is Artist, this field may be set to American Singer.
	Entitytypehints []string `json:"entityTypeHints,omitempty"` // A list of hints that indicate the entity's type. The list could contain a single hint such as Movie or a list of hints such as Place, LocalBusiness, Restaurant. Each successive hint in the array narrows the entity's type.
	Query string `json:"query,omitempty"`
	TypeField string `json:"_type"`
	Entityscenario string `json:"entityScenario"` // The supported scenario.
	Entitysubtypehints []string `json:"entitySubTypeHints,omitempty"`
}

// SearchResponse represents the SearchResponse schema from the OpenAPI specification
type SearchResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
}

// Thing represents the Thing schema from the OpenAPI specification
type Thing struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
}

// PostalAddress represents the PostalAddress schema from the OpenAPI specification
type PostalAddress struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
}

// Response represents the Response schema from the OpenAPI specification
type Response struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// CreativeWork represents the CreativeWork schema from the OpenAPI specification
type CreativeWork struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
}

// Places represents the Places schema from the OpenAPI specification
type Places struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
}

// ResponseBase represents the ResponseBase schema from the OpenAPI specification
type ResponseBase struct {
	TypeField string `json:"_type"`
}

// SearchResultsAnswer represents the SearchResultsAnswer schema from the OpenAPI specification
type SearchResultsAnswer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
}

// StructuredValue represents the StructuredValue schema from the OpenAPI specification
type StructuredValue struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
}

// Action represents the Action schema from the OpenAPI specification
type Action struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Copyrightholder Thing `json:"copyrightHolder,omitempty"` // Defines a thing.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Mainentity Thing `json:"mainEntity,omitempty"` // Defines a thing.
	Text string `json:"text,omitempty"` // Text content of this creative work
	Commentcount int `json:"commentCount,omitempty"`
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	About []Thing `json:"about,omitempty"` // For internal use only.
	Copyrightyear int `json:"copyrightYear,omitempty"`
	Disclaimer string `json:"disclaimer,omitempty"`
	Discussionurl string `json:"discussionUrl,omitempty"`
	Isaccessibleforfree bool `json:"isAccessibleForFree,omitempty"`
	Creator Thing `json:"creator,omitempty"` // Defines a thing.
	Mentions []Thing `json:"mentions,omitempty"` // For internal use only.
	Genre []string `json:"genre,omitempty"`
	Headline string `json:"headLine,omitempty"`
}

// Answer represents the Answer schema from the OpenAPI specification
type Answer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
}

// GeoCoordinates represents the GeoCoordinates schema from the OpenAPI specification
type GeoCoordinates struct {
	Elevation float64 `json:"elevation,omitempty"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TypeField string `json:"_type"`
}

// SearchAction represents the SearchAction schema from the OpenAPI specification
type SearchAction struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Mainentity Thing `json:"mainEntity,omitempty"` // Defines a thing.
	Text string `json:"text,omitempty"` // Text content of this creative work
	Commentcount int `json:"commentCount,omitempty"`
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	About []Thing `json:"about,omitempty"` // For internal use only.
	Copyrightyear int `json:"copyrightYear,omitempty"`
	Disclaimer string `json:"disclaimer,omitempty"`
	Discussionurl string `json:"discussionUrl,omitempty"`
	Isaccessibleforfree bool `json:"isAccessibleForFree,omitempty"`
	Creator Thing `json:"creator,omitempty"` // Defines a thing.
	Mentions []Thing `json:"mentions,omitempty"` // For internal use only.
	Genre []string `json:"genre,omitempty"`
	Headline string `json:"headLine,omitempty"`
	Copyrightholder Thing `json:"copyrightHolder,omitempty"` // Defines a thing.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Serviceurl string `json:"serviceUrl,omitempty"` // Use this URL to get additional data to determine how to take the appropriate action. For example, the serviceUrl might return JSON along with an image URL.
	Displayname string `json:"displayName,omitempty"` // A display name for the action.
	Istopaction bool `json:"isTopAction,omitempty"` // A Boolean representing whether this result is the top action.
	Location []Place `json:"location,omitempty"`
	Result []Thing `json:"result,omitempty"` // The result produced in the action.
}

// Place represents the Place schema from the OpenAPI specification
type Place struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Immediateaction []Action `json:"immediateAction,omitempty"`
	Potentialaction []Action `json:"potentialAction,omitempty"`
	Preferredclickthroughurl string `json:"preferredClickthroughUrl,omitempty"`
	Readlink string `json:"readLink,omitempty"` // The URL that returns this resource.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL to Bing's search result for this item.
	Adaptivecard string `json:"adaptiveCard,omitempty"`
	Entitypresentationinfo EntitiesEntityPresentationInfo `json:"entityPresentationInfo,omitempty"` // Defines additional information about an entity such as type hints.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
}
